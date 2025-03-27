package openai

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

type XFYunClient struct {
	client *resty.Client
}

const (
	xfyunURL string = "https://spark-api-open.xf-yun.com/v1"
)

func NewXFYunClient(apiKey string, debug bool) *XFYunClient {
	return &XFYunClient{
		client: resty.New().SetBaseURL(xfyunURL).
			SetAuthToken(apiKey).SetDebug(debug),
	}
}

type XFYunChatRequest struct {
	*ChatCompletionRequest

	User string `json:"user"`
}

type XFYunChatReply struct {
	ChatCompletionReply

	Code    int    `json:"code"`
	Message string `json:"message"`
	SID     string `json:"sid"`
}

func (reply *XFYunChatReply) Error() string {
	switch reply.Code {
	case 10007:
		return fmt.Sprintf("%d-流量受限，问题尚未处理完成，无法处理下一个请求", reply.Code)
	case 10013, 10019:
		return fmt.Sprintf("%d-输入内容审核不通过，涉嫌违规，请重新调整输入内容", reply.Code)
	case 10014:
		return fmt.Sprintf("%d-输出内容涉及敏感信息，生成终止", reply.Code)
	case 10907:
		return fmt.Sprintf("%d-会话长度超过上限，需精简输入内容", reply.Code)
	case 11200, 11201, 11202, 11203:
		return fmt.Sprintf("%d-授权错误，权限不足或触发流量限制", reply.Code)
	default:
		return fmt.Sprintf("%d-服务内部异常，功能暂不可用", reply.Code)
	}
}

func (c *XFYunClient) ChatCompletion(ctx context.Context, req *ChatCompletionRequest, opts *ChatOptions) (*ChatCompletionReply, error) {
	resp, err := c.client.R().SetBody(&XFYunChatRequest{
		User:                  fmt.Sprintf("user_%d", opts.UserID),
		ChatCompletionRequest: req,
	}).SetError(&ErrorReply{}).SetResult(&XFYunChatReply{}).Post(_pathChatCompletion)
	if err != nil {
		return nil, errors.Wrap(err, "http request failed")
	}
	if resp.StatusCode() >= 400 {
		return nil, resp.Error().(*ErrorReply)
	}
	res, ok := resp.Result().(*XFYunChatReply)
	if !ok {
		return nil, errors.Errorf("http response assert failed: %v", resp.Result())
	}
	if res.Code != 0 {
		return nil, res
	}
	return &res.ChatCompletionReply, nil
}
