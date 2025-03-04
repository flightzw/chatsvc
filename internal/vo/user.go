package vo

type UserVO struct {
	ID        int32  `json:"id"`
	Username  string `json:"username"`   // 用户名
	AvatarURL string `json:"avatar_url"` // 头像url
	Nickname  string `json:"nickname"`   // 昵称
	Gender    int32  `json:"gender"`     // 性别：0未知，1男，2女
	Signature string `json:"signature"`  // 个性签名
	IsOnline  bool   `json:"is_online"`  // 是否在线
}
