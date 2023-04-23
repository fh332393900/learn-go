package add

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(1, 3)
	want := 4

	if want != got {
		t.Errorf("excepted:%v, got:%v", want, got) // 测试失败输出错误提示
	}
}
