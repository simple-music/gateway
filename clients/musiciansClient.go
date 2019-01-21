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
	MusicianNotFoundErrMessage           = "user not found"
	MusicianAttributeDuplicateErrMessage = "user attribute duplicate"
)

type MusiciansClient struct {
	client *utils.RestClient

	notFoundErr *errs.Error
	conflictErr *errs.Error
}

func NewMusiciansClient() *MusiciansClient {
	return &MusiciansClient{
		client: utils.NewRestClient(
			utils.RestClientConfig{
				ServiceName:     "musicians-service",
				ServiceID:       config.MusiciansServiceID,
				ServicePassword: config.MusiciansServicePassword,
			},
			utils.RestClientComponents{
				Logger:          common.Logger,
				DiscoveryClient: common.DiscoveryClient,
			},
		),
		notFoundErr: errs.NewError(errs.NotFound, MusicianNotFoundErrMessage),
		conflictErr: errs.NewError(errs.Conflict, MusicianAttributeDuplicateErrMessage),
	}
}

func (c *MusiciansClient) AddMusician(newMusician *models.NewMusician, musician *models.Musician) *errs.Error {
	req := fasthttp.AcquireRequest()
	path := "/musicians"

	req.Header.SetMethod(http.MethodPost)
	req.Header.SetContentType("application/json")

	content, _ := easyjson.Marshal(newMusician)
	req.AppendBody(content)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusConflict {
		return c.conflictErr
	} else if resp.StatusCode() != http.StatusCreated {
		return utils.WrapUnexpectedResponse(resp)
	}

	if err := easyjson.Unmarshal(resp.Body(), musician); err != nil {
		return errs.NewServiceError(err)
	}

	return nil
}

func (c *MusiciansClient) FindMusician(nickname string, musician *models.Musician) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodGet)

	path := fmt.Sprintf("/musicians?nickname=%s", nickname)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return c.notFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return utils.WrapUnexpectedResponse(resp)
	}

	if err := easyjson.Unmarshal(resp.Body(), musician); err != nil {
		return errs.NewServiceError(err)
	}

	return nil
}

func (c *MusiciansClient) GetMusician(id string, musician *models.Musician) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodGet)

	path := fmt.Sprintf("/musicians/%s", id)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return c.notFoundErr
	} else if resp.StatusCode() != http.StatusOK {
		return utils.WrapUnexpectedResponse(resp)
	}

	if err := easyjson.Unmarshal(resp.Body(), musician); err != nil {
		return errs.NewServiceError(err)
	}

	return nil
}

func (c *MusiciansClient) UpdateMusician(id string, update *models.MusicianUpdate) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodPatch)
	req.Header.SetContentType("application/json")

	content, _ := easyjson.Marshal(update)
	req.AppendBody(content)

	path := fmt.Sprintf("/musicians/%s", id)

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	if resp.StatusCode() == http.StatusNotFound {
		return c.notFoundErr
	} else if resp.StatusCode() == http.StatusConflict {
		return c.conflictErr
	} else if resp.StatusCode() != http.StatusOK && resp.StatusCode() != http.StatusNotModified {
		return utils.WrapUnexpectedResponse(resp)
	}

	return nil
}

func (c *MusiciansClient) DeleteMusician(id string) *errs.Error {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(http.MethodDelete)

	path := fmt.Sprintf("/musicians/%s", id)

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
