package entity

type AnyMap map[string]interface{}

func (m AnyMap) Assert() map[string]interface{} {
	return m
}

type LoginInfo struct {
	Token string `json:"token"`
}
