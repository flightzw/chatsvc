package stringx

import (
	"fmt"
	"testing"

	"github.com/vcaesar/cedar"
)

func Test_Trie(t *testing.T) {
	words := []string{"高清视频", "高清 CV", "东京冷", "东京热"}

	d := cedar.New()
	for idx, w := range words {
		d.Insert([]byte(w), idx)
	}

	fmt.Println(d.Jump([]byte("东京冷吗"), 0))
	fmt.Println(d.Find([]byte("东京冷吗"), 0))

	fmt.Println(d.PrefixMatch([]byte("东京冷吗"), 0))
	fmt.Println(d.ExactMatch([]byte("东京冷吗")))
}
