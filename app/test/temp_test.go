package main

import "testing"

// AssertEqual 测试预期值和实际值是否匹配
func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	t.Helper() // 标记为测试辅助函数
	switch expected := expected.(type) {
	case string:
		if expected != actual {
			t.Errorf("Error:\nexpected: %s\nactual: %s", expected, actual)
		}
	default:
		t.Errorf("Unsupported type")
	}
}

func TestExample(t *testing.T) {
	expected := "Hello, world!"
	actual := "Helldo, world!"
	AssertEqual(t, expected, actual) // 调用测试辅助函数
}
