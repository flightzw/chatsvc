package query

import (
	"gorm.io/gen/field"
)

type userQuery struct {
	ID          field.Int32
	Username    field.String // 用户名
	Password    field.String // 密码
	Nickname    field.String // 昵称
	AvatarURL   field.String // 头像url
	Status      field.Int32  // 状态: 1正常，2封禁
	LastLoginAt field.Time   // 最后上线时间
	LastLoginIP field.String // 最后上线ip
	CreatedAt   field.Time   // 创建时间
	UpdatedAt   field.Time   // 更新时间
	DeletedAt   field.Field  // 删除时间
}

func NewUserQuery() *userQuery {
	tableName := ""
	return &userQuery{
		ID:          field.NewInt32(tableName, "id"),
		Username:    field.NewString(tableName, "username"),
		Password:    field.NewString(tableName, "password"),
		Nickname:    field.NewString(tableName, "nickname"),
		AvatarURL:   field.NewString(tableName, "avatar_url"),
		Status:      field.NewInt32(tableName, "status"),
		LastLoginAt: field.NewTime(tableName, "last_login_at"),
		LastLoginIP: field.NewString(tableName, "last_login_ip"),
		CreatedAt:   field.NewTime(tableName, "created_at"),
		UpdatedAt:   field.NewTime(tableName, "updated_at"),
		DeletedAt:   field.NewField(tableName, "deleted_at"),
	}
}

type friendQuery struct {
	ID              field.Int32
	UserID          field.Int32  // 用户id
	FriendID        field.Int32  // 好友id
	FriendNickname  field.String // 昵称
	FriendAvatarURL field.String // 头像url
	Remark          field.String // 备注
	CreatedAt       field.Time   // 创建时间
}

func NewFriendQuery() *friendQuery {
	tableName := ""
	return &friendQuery{
		ID:              field.NewInt32(tableName, "id"),
		UserID:          field.NewInt32(tableName, "user_id"),
		FriendID:        field.NewInt32(tableName, "friend_id"),
		FriendNickname:  field.NewString(tableName, "friend_nickname"),
		FriendAvatarURL: field.NewString(tableName, "friend_avatar_url"),
		Remark:          field.NewString(tableName, "remark"),
		CreatedAt:       field.NewTime(tableName, "created_at"),
	}
}

type privateMessageQuery struct {
	ID        field.Int32
	SendID    field.Int32  // 用户uid
	RecvID    field.Int32  // 好友uid
	Content   field.String // 发送内容
	Type      field.Int32  // 消息类型
	Status    field.Int32  // 状态 0:未送达 1:已送达 2:撤回 3:已读
	CreatedAt field.Time   // 创建时间
}

func NewPrivateMessageQuery() *privateMessageQuery {
	tableName := ""
	return &privateMessageQuery{
		ID:        field.NewInt32(tableName, "id"),
		SendID:    field.NewInt32(tableName, "send_id"),
		RecvID:    field.NewInt32(tableName, "recv_id"),
		Content:   field.NewString(tableName, "content"),
		Type:      field.NewInt32(tableName, "type"),
		Status:    field.NewInt32(tableName, "status"),
		CreatedAt: field.NewTime(tableName, "created_at"),
	}
}

type sensitiveWordQuery struct {
	ID        field.Int32
	Content   field.String // 敏感词内容
	Enabled   field.Int32  // 是否启用 1:是 0:否
	CreatedAt field.Time   // 创建时间
	UpdatedAt field.Time   // 更新时间
}

func NewSensitiveWordQuery() *sensitiveWordQuery {
	tableName := ""
	return &sensitiveWordQuery{
		ID:        field.NewInt32(tableName, "id"),
		Content:   field.NewString(tableName, "content"),
		Enabled:   field.NewInt32(tableName, "enabled"),
		CreatedAt: field.NewTime(tableName, "created_at"),
		UpdatedAt: field.NewTime(tableName, "updated_at"),
	}
}
