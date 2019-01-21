package clients

import (
	"fmt"
	"github.com/simple-music/gateway/common"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/utils"
	"github.com/valyala/fasthttp"
)

type AvatarsClient struct {
	client *utils.RestClient
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
	}
}

func (c *AvatarsClient) AddAvatar(user string, content []byte) *errs.Error {
	req := fasthttp.AcquireRequest()
	path := fmt.Sprintf("/avatars/%s", user)

	req.AppendBody(content)
	req.Header.SetContentType("image/jpeg")

	resp, err := c.client.PerformRequest(req, path)
	if err != nil {
		return err
	}

	fmt.Println(resp) //TODO

	return nil
}

func (c *AvatarsClient) GetAvatar(user string) ([]byte, *errs.Error) {
	panic("not implemented") //TODO
}

func (c *AvatarsClient) DeleteAvatar(user string) *errs.Error {
	panic("not implemented") //TODO
}
