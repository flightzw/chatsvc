package openai

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	deepseekURL = "https://api.deepseek.com"
)

var (
	deepseekErrCodeMap = map[int]string{
		400: "格式错误",
		401: "认证失败",
		402: "余额不足",
		422: "参数错误",
		429: "速率受限",
		500: "服务故障",
		503: "服务繁忙",
	}
)

type DeepSeekClient struct {
	client *resty.Client
}

func NewDeepSeekClient(apiKey string, debug bool) *DeepSeekClient {
	return &DeepSeekClient{
		client: resty.New().SetBaseURL(deepseekURL).
			SetAuthToken(apiKey).SetDebug(debug),
	}
}

func (c *DeepSeekClient) ChatCompletion(ctx context.Context, req *ChatCompletionRequest, opts *ChatOptions) (*ChatCompletionReply, error) {
	resp, err := c.client.R().SetBody(req).SetError(&ErrorReply{}).SetResult(&ChatCompletionReply{}).Post(_pathChatCompletion)
	if err != nil {
		return nil, errors.Wrap(err, "http request failed")
	}

	if code := resp.StatusCode(); code >= 400 {
		return nil, errors.Errorf("%s: %s", deepseekErrCodeMap[code], resp.Error())
	}
	res, ok := resp.Result().(*ChatCompletionReply)
	if !ok {
		return nil, errors.Errorf("http response assert failed: %v", resp.Result())
	}
	return res, nil
}
