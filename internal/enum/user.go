package enum

// 用户类型
type UserType int32

const (
	UserTypeNormal  UserType = iota + 1 // 普通用户
	UserTypeAI                          // AI
	UserTypeManager                     // 管理员
)

// 用户状态
type UserStatus int32

const (
	UserStatusNotEnable UserStatus = iota // 未启用
	UserStatusNormal                      // 正常
	UserStatusBan                         // 封禁
)

var _UserStatusMap = map[UserStatus]string{
	UserStatusNotEnable: "未启用",
	UserStatusNormal:    "正常",
	UserStatusBan:       "已封禁",
}

func (x UserStatus) Map() string {
	return _UserStatusMap[x]
}
