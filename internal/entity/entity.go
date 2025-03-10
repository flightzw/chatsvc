package entity

type AnyMap map[string]any

func (m AnyMap) Assert() map[string]any {
	return m
}

type LoginInfo struct {
	Token string `json:"token"`
}
