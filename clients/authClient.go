package clients

import (
	"fmt"
	"github.com/mailru/easyjson"
	"github.com/simple-music/gateway/common"
	"github.com/simple-music/gateway/config"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/models"
	"github.com/simple-music/gateway/utils"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

const (
	CredentialsNotFoundErrMessage  = "credentials not found"
	RefreshTokenNotFoundErrMessage = "refresh token not found"
)

type AuthClient struct {
	client *utils.RestClient

	credentialsNotFoundErr  *errs.Error
	refreshTokenNotFoundErr *errs.Error
}

func NewAuthClient() *AuthClient {
	return &AuthClient{
		client: utils.NewRestClient(
			utils.RestClientConfig{
				ServiceName:     "auth-service",
				ServiceID:       config.AuthServiceID,
				ServicePassword: config.AuthServicePassword,
			},
			utils.RestClientComponents{
				Logger:          common.Logger,
				DiscoveryClient: common.DiscoveryClient,
			},
		),
		credentialsNotFoundErr:  errs.NewError(errs.NotFound, CredentialsNotFoundErrMessage),
		refreshTokenNotFoundErr: errs.NewError(errs.NotFound, RefreshTokenNotFoundErrMessage),
	}
}

func (c *AuthClient) AddCredentials(credentials *models.NewCredentials) *errs.Error {
	req := fasthttp.AcquireRequest()

	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")

	content, _ := easyjson.Marshal(credentials)
	req.SetBody(content)

	resp, err := c.client.PerformRequest(req, "/credentials")
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return utils.WrapUnexpectedResponse(resp)
	}

	return nil
}

func (c *AuthClient) DeleteCredentials(user string) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodDelete)

	path := fmt.Sprintf("/credentials/%s", user)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return c.credentialsNotFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return utils.WrapUnexpectedResponse(resp)
	}

	return nil
}

func (c *AuthClient) GetAuthCode(authCode *models.AuthCode) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodGet)

	path := fmt.Sprintf("/clients?clientId=%s&clientSecret=%s",
		authCode.ClientID, authCode.ClientSecret,
	)

	log.Println(path)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return c.credentialsNotFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return utils.WrapUnexpectedResponse(resp)
	}

	if err := easyjson.Unmarshal(resp.Body(), authCode); err != nil {
		return errs.NewServiceError(err)
	}

	return nil
}

func (c *AuthClient) StartSession(credentials *models.AuthCredentials) ([]byte, *errs.Error) {
	req := fasthttp.AcquireRequest()

	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")

	content, _ := easyjson.Marshal(credentials)
	req.SetBody(content)

	resp, err := c.client.PerformRequest(req, "/sessions")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, c.credentialsNotFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return nil, utils.WrapUnexpectedResponse(resp)
	}

	return resp.Body(), nil
}

func (c *AuthClient) RefreshSession(refreshToken string) ([]byte, *errs.Error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodPatch)

	path := fmt.Sprintf("/sessions?refreshToken=%s", refreshToken)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, c.refreshTokenNotFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return nil, utils.WrapUnexpectedResponse(resp)
	}

	return resp.Body(), nil
}

func (c *AuthClient) DeleteSession(refreshToken string) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodDelete)

	path := fmt.Sprintf("/sessions?refreshToken=%s", refreshToken)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return utils.WrapUnexpectedResponse(resp)
	}

	return nil
}
