package biz

import (
	"context"

	"github.com/flightzw/chatsvc/api/chatsvc/errno"
	"github.com/flightzw/chatsvc/internal/conf"
	"github.com/flightzw/chatsvc/internal/utils/hash"
	"github.com/flightzw/chatsvc/internal/utils/openai"
)

const (
	_redisLockKeyAIChatLimit = "aichat-limit:id:%d"
)

// 生成 hash 密码串
func generatePasswordStr(password string) string {
	salt := hash.GenerateSalt(16)
	hashPassword := hash.GenerateBcryptHashPassword(password, salt)
	return hashPassword + ":" + salt
}

// 获取AI聊天结果
func getAIChatMessage(modelType openai.ModelType, apikey string, userID int32, message string) (string, error) {
	client, err := openai.NewAIChatClient(modelType, apikey, conf.Getenv("RUN_ENV").(string) != "production")
	if err != nil {
		return "", errno.ErrorSystemInternalFailure("配置项异常，服务暂不可用").WithCause(err)
	}
	reply, err := client.ChatCompletion(context.Background(),
		&openai.ChatCompletionRequest{
			Model:    modelType,
			Messages: []*openai.Message{{Content: message, Role: openai.RoleTypeUser}}},
		&openai.ChatOptions{
			UserID: userID,
		})
	if err != nil {
		return "", err
	}
	if len(reply.Choices) < 1 {
		return "", errno.ErrorSystemInternalFailure("模型调用异常，响应数据为空")
	}
	return reply.Choices[0].Message.Content, nil
}
