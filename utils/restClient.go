package utils

//go:generate easyjson

import (
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/logs"
	"github.com/valyala/fasthttp"
	"net/http"
	"sync"
)

type RestClientConfig struct {
	ServiceName     string
	ServiceID       string
	ServicePassword string
}

type RestClientComponents struct {
	Logger          *logs.Logger
	DiscoveryClient *DiscoveryClient
}

type RestClient struct {
	client      fasthttp.Client
	credentials serviceCredentials

	authToken      serviceAuthToken
	authTokenMutex *sync.RWMutex

	config     RestClientConfig
	components RestClientComponents
}

func NewRestClient(cf RestClientConfig, cp RestClientComponents) *RestClient {
	srvCred := serviceCredentials{
		ServiceID:       cf.ServiceID,
		ServicePassword: cf.ServicePassword,
	}
	srvCred.JSON, _ = easyjson.Marshal(&srvCred)

	return &RestClient{
		config:     cf,
		components: cp,

		client: fasthttp.Client{
			Name: cf.ServiceName + "-client",
		},
		credentials: srvCred,

		authTokenMutex: &sync.RWMutex{},
	}
}

func (c *RestClient) PerformRequest(req *fasthttp.Request, path string) (*fasthttp.Response, *errs.Error) {
	address, err := c.components.DiscoveryClient.ResolveName(c.config.ServiceName)
	if err != nil {
		return nil, err
	}

	req.SetRequestURI(fmt.Sprintf("http://%s%s", address, path))

	resp, err := c.performReq(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusUnauthorized {
		if err = c.performAuthReq(address); err != nil {
			return nil, err
		}

		resp, err = c.performReq(req)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func WrapUnexpectedResponse(resp *fasthttp.Response) *errs.Error {
	err := fmt.Errorf("unexpected response: %v", resp)
	return errs.NewServiceError(err)
}

func (c *RestClient) performReq(req *fasthttp.Request) (*fasthttp.Response, *errs.Error) {
	c.authTokenMutex.RLock()
	req.Header.Set("X-Gateway-Token", c.authToken.AuthToken)
	c.authTokenMutex.RUnlock()

	resp := fasthttp.AcquireResponse()
	err := c.client.Do(req, resp)
	if err != nil {
		return nil, errs.NewServiceError(err)
	}

	return resp, nil
}

func (c *RestClient) performAuthReq(address string) *errs.Error {
	req := fasthttp.AcquireRequest()

	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetBody(c.credentials.JSON)

	req.SetRequestURI(fmt.Sprintf("http://%s/auth", address))

	resp, err := c.performReq(req)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return WrapUnexpectedResponse(resp)
	}

	c.authTokenMutex.Lock()
	defer c.authTokenMutex.Unlock()

	jsonErr := easyjson.Unmarshal(resp.Body(), &c.authToken)
	if jsonErr != nil {
		return errs.NewServiceError(jsonErr)
	}

	return nil
}

//easyjson:json
type serviceCredentials struct {
	ServiceID       string `json:"serviceId"`
	ServicePassword string `json:"servicePassword"`
	JSON            []byte `json:"-"`
}

//easyjson:json
type serviceAuthToken struct {
	AuthToken string `json:"authToken"`
}
