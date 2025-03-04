package enum

type AccountStatus int32

const (
	AccountStatusNotEnable AccountStatus = iota // 未启用
	AccountStatusNormal                         // 正常
	AccountStatusBan                            // 封禁
)

var _AccountStatusMap = map[AccountStatus]string{
	AccountStatusNotEnable: "未启用",
	AccountStatusNormal:    "正常",
	AccountStatusBan:       "已封禁",
}

func (x AccountStatus) Map() string {
	return _AccountStatusMap[x]
}
