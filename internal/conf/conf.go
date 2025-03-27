package conf

var (
	config *Bootstrap
)

func BindConfig(data *Bootstrap) {
	config = data
}

func Getenv(key string) any {
	switch key {
	case "RUN_ENV":
		return config.Data.Env
	case "AI_CONFIG_ID":
		return config.Data.Dbconfig.AiconfigId
	default:
		return ""
	}
}
