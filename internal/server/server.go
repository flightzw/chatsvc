package server

import (
	"regexp"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	wsserver "github.com/flightzw/chatsvc/internal/ws/server"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, wsserver.InitSessionHub)

type replaceRule struct {
	pattern *regexp.Regexp
	replace string
}

func makeLogFilter() func(level log.Level, keyvals ...interface{}) bool {
	rules := []*replaceRule{
		{pattern: regexp.MustCompile(`"password":"[a-zA-Z0-9!@#$%^&*]+"`), replace: `"password":"******"`},
	}
	return func(level log.Level, keyvals ...interface{}) bool {
		for keyIdx := 0; keyIdx < len(keyvals); keyIdx += 2 {
			if keyvals[keyIdx] == "args" {
				value := keyvals[keyIdx+1].(string)
				for _, rule := range rules {
					if rule.pattern.MatchString(value) {
						keyvals[keyIdx+1] = rule.pattern.ReplaceAllString(value, rule.replace)
					}
				}
			}
		}
		return false
	}
}
