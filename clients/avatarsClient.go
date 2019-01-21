package clients

import (
	"github.com/simple-music/gateway/common"
	"github.com/simple-music/gateway/errs"
	"github.com/simple-music/gateway/utils"
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
	panic("not implemented") //TODO
}

func (c *AvatarsClient) GetAvatar(user string) ([]byte, *errs.Error) {
	panic("not implemented") //TODO
}

func (c *AvatarsClient) DeleteAvatar(user string) *errs.Error {
	panic("not implemented") //TODO
}
