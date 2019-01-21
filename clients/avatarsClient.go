package clients

import (
	"fmt"
	"github.com/simple-music/gateway/common"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/utils"
	"github.com/valyala/fasthttp"
	"net/http"
)

const (
	AvatarNotFoundErrMessage = "avatar not found"
)

type AvatarsClient struct {
	client      *utils.RestClient
	notFoundErr *errs.Error
}

func NewAvatarsClient() *AvatarsClient {
	return &AvatarsClient{
		client: utils.NewRestClient(
			utils.RestClientConfig{
				ServiceName: "avatars-service",
			},
			utils.RestClientComponents{
				Logger:          common.Logger,
				DiscoveryClient: common.DiscoveryClient,
			},
		),
		notFoundErr: errs.NewError(errs.NotFound, AvatarNotFoundErrMessage),
	}
}

func (c *AvatarsClient) AddAvatar(user string, content []byte) *errs.Error {
	req := fasthttp.AcquireRequest()
	path := fmt.Sprintf("/avatars/%s", user)

	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("image/jpeg")
	req.AppendBody(content)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return c.wrapErr(resp)
	}

	return nil
}

func (c *AvatarsClient) GetAvatar(user string) ([]byte, *errs.Error) {
	req := fasthttp.AcquireRequest()
	path := fmt.Sprintf("/avatars/%s", user)

	req.Header.SetMethod(http.MethodGet)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, c.notFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return nil, c.wrapErr(resp)
	}

	return resp.Body(), nil
}

func (c *AvatarsClient) DeleteAvatar(user string) *errs.Error {
	panic("not implemented") //TODO
}

func (c *AvatarsClient) wrapErr(resp *fasthttp.Response) *errs.Error {
	err := fmt.Errorf("unexpected response: %v", resp)
	return errs.NewServiceError(err)
}
