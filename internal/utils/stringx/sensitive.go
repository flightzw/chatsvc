package stringx

import (
	"sync"

	"github.com/importcjj/sensitive"
)

type Filter struct {
	lock  sync.RWMutex
	inner *sensitive.Filter
}

func NewFilter() *Filter {
	return &Filter{
		lock:  sync.RWMutex{},
		inner: sensitive.New(),
	}
}

// 添加敏感词
func (f *Filter) AddWord(words ...string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.inner.AddWord(words...)
}

func (f *Filter) DelWord(words ...string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.inner.DelWord(words...)
}

// 更新去噪模式
func (f *Filter) UpdateNoisePattern(pattern string) {
	f.lock.Lock()
	defer f.lock.Unlock()
	f.inner.UpdateNoisePattern(pattern)
}

// 替换敏感词
func (f *Filter) Replace(text string, repl rune) string {
	f.lock.RLock()
	defer f.lock.RUnlock()
	return f.inner.Replace(text, repl)
}
