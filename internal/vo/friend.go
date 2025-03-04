package vo

type FriendVO struct {
	ID        int32  `json:"id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
	AvatarUrl string `json:"avatar_url"`
	Gender    int32  `json:"gender"`    // 性别：0未知，1男，2女
	Signature string `json:"signature"` // 个性签名
	Remark    string `json:"remark"`
	IsOnline  bool   `json:"is_online"`
}
