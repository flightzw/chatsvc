package openai

// 交互角色
type RoleType string

const (
	RoleTypeSystem    RoleType = "system"    // 系统
	RoleTypeUser      RoleType = "user"      // 用户
	RoleTypeAssistant RoleType = "assistant" // AI助手
)

// 模型类别
type ModelType string

const (
	// 讯飞
	ModelTypeXFYunLite    ModelType = "lite"
	ModelTypeXFYunUltraV4 ModelType = "4.0Ultra"
	// DeepSeek
	ModelTypeDeepSeekChat ModelType = "deepseek-chat"
)

type ReasonType string

const (
	ReasonTypeStop            ReasonType = "stop"
	ReasonTypeLength          ReasonType = "length"
	ReasonTypeContentFilter   ReasonType = "content_filter"
	ReasonTypeResInsufficient ReasonType = "insufficient_system_resource"
)
