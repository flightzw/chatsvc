package openai

import "fmt"

type RespType string

const (
	RespText RespType = "text"
	RespJson RespType = "json_object"
)

type Message struct {
	Content string   `json:"content"`
	Role    RoleType `json:"role"`
}

type ResponseFormat struct {
	Type RespType `json:"type"`
}

type StreamOpts struct {
	IncludeUsage bool `json:"include_usage"`
}

type ChatCompletionRequest struct {
	Model    ModelType  `json:"model"`    // 使用模型
	Messages []*Message `json:"messages"` // 消息上下文
	Stream   bool       `json:"stream"`   // 流式输出控制
	// MaxTokens        int            `json:"max_tokens"`               // 模型回答使用 token 最大值
	// FrequencyPenalty float64        `json:"frequency_penalty"`        //
	// PresencePenalty  float64        `json:"presence_penalty"`         //
	// ResponseFormat   ResponseFormat `json:"response_format"`          // 格式配置
	// Stop             []string       `json:"stop,omitempty"`           // 支持 null 或字符串数组
	// StreamOptions    *StreamOpts    `json:"stream_options,omitempty"` // 流式选项（根据文档定义具体结构）
	// Temperature      float64        `json:"temperature"`              // 温度参数
	// TopP             float64        `json:"top_p"`                    // 核心采样概率
	// Tools            any            `json:"tools,omitempty"`          // 工具列表（根据文档定义具体结构）
	// ToolChoice       string         `json:"tool_choice"`              // 工具选择策略
	// Logprobs         bool           `json:"logprobs,omitempty"`       // 是否返回概率信息
	// TopLogprobs      *int           `json:"top_logprobs,omitempty"`   // 返回的最高概率数量（指针支持 null）
}

type ChoiceInfo struct {
	FinishReason ReasonType `json:"finish_reason"` // 生成结束原因
	Index        int64      `json:"index"`         // 索引
	Message      *Message   `json:"message"`       // 响应消息
}

type TokenUseInfo struct {
	CompletionTokens int64 `json:"completion_tokens"` // 响应内容消耗 token 数
	PromptTokens     int64 `json:"prompt_tokens"`     // 请求内容消耗 token 数
	TotalTokens      int64 `json:"total_tokens"`      // 总消耗 token 数
}

type ChatCompletionReply struct {
	Choices []*ChoiceInfo `json:"choices"`
	Usage   *TokenUseInfo `json:"usage"`
}

type ErrorReply struct {
	APIError *httpError `json:"error,omitempty"`
}

func (reply *ErrorReply) Error() string {
	return fmt.Sprintf("[%s] %s => %s", reply.APIError.Code, reply.APIError.Type, reply.APIError.Message)
}

type httpError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    string `json:"code"`
	Param   any    `json:"param"`
}
