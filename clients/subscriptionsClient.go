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
	"net/http"
)

const (
	UserNotFoundErrMessage = "user not found"
)

type SubscriptionsClient struct {
	client      *utils.RestClient
	notFoundErr *errs.Error
}

func NewSubscriptionsClient() *SubscriptionsClient {
	return &SubscriptionsClient{
		client: utils.NewRestClient(
			utils.RestClientConfig{
				ServiceName:     "subscriptions-service",
				ServiceID:       config.SubscriptionsServiceID,
				ServicePassword: config.SubscriptionsServicePassword,
			},
			utils.RestClientComponents{
				Logger:          common.Logger,
				DiscoveryClient: common.DiscoveryClient,
			},
		),
		notFoundErr: errs.NewError(errs.NotFound, UserNotFoundErrMessage),
	}
}

func (c *SubscriptionsClient) AddSubscription(user string, subscription string) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodPost)

	path := fmt.Sprintf("/users/%s/subscriptions/%s", user, subscription)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return utils.WrapUnexpectedResponse(resp)
	}

	return nil
}

func (c *SubscriptionsClient) CheckSubscription(user string, subscription string) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodGet)

	path := fmt.Sprintf("/users/%s/subscriptions/%s", user, subscription)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return c.notFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return utils.WrapUnexpectedResponse(resp)
	}

	return nil
}

func (c *SubscriptionsClient) DeleteSubscription(user string, subscription string) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodDelete)

	path := fmt.Sprintf("/users/%s/subscriptions/%s", user, subscription)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return c.notFoundErr
	} else if (resp.StatusCode() != http.StatusOK) && (resp.StatusCode() != http.StatusNoContent) {
		return utils.WrapUnexpectedResponse(resp)
	}

	return nil
}

func (c *SubscriptionsClient) GetStatus(user string, status *models.SubscriptionsStatus) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodGet)

	path := fmt.Sprintf("/users/%s/status", user)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return c.notFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return utils.WrapUnexpectedResponse(resp)
	}

	if err := easyjson.Unmarshal(resp.Body(), status); err != nil {
		return errs.NewServiceError(err)
	}

	return nil
}

func (c *SubscriptionsClient) GetSubscribers(user string, pageIndex, pageSize int) ([]byte, *errs.Error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodGet)

	path := fmt.Sprintf("/users/%s/subscribers?page=%d", user, pageIndex)
	if pageSize != 0 {
		path += fmt.Sprintf("&size=%d", pageSize)
	}

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, c.notFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return nil, utils.WrapUnexpectedResponse(resp)
	}

	return resp.Body(), nil
}

func (c *SubscriptionsClient) GetSubscriptions(user string, pageIndex, pageSize int) ([]byte, *errs.Error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodGet)

	path := fmt.Sprintf("/users/%s/subscriptions?page=%d", user, pageIndex)
	if pageSize != 0 {
		path += fmt.Sprintf("&size=%d", pageSize)
	}

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return nil, c.notFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return nil, utils.WrapUnexpectedResponse(resp)
	}

	return resp.Body(), nil
}

func (c *SubscriptionsClient) DeleteUser(user string) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodDelete)

	path := fmt.Sprintf("/users/%s", user)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return c.notFoundErr
	} else if resp.StatusCode() != http.StatusNoContent {
		return utils.WrapUnexpectedResponse(resp)
	}

	return nil
}
