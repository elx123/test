package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)
sadjhfkjhfkjds
// Person 结构体，包含多个字段
type Person struct {
	Name    string
	Age     int
	Address string
}

// transformPerson 转换 Person 结构体以便比较
func transformPerson(p Person) (transformedPerson struct {
	Name string
	Age  int
}) {
	// 在这里，我们只关注 Name 和 Age 字段
	return struct {
		Name string
		Age  int
	}{
		Name: p.Name,
		Age:  p.Age,
	}
}

func main() {
	// 创建两个 Person 实例
	p1 := Person{Name: "Alice", Age: 30, Address: "123 Elm St"}
	p2 := Person{Name: "Alice", Age: 30, Address: "456 Oak St"}

	// 使用 cmp 比较两个 Person 实例，但只关注 Name 和 Age 字段
	equal := cmp.Equal(p1, p2, cmp.Transformer("TransformPerson", transformPerson))

	// 输出比较结果
	fmt.Printf("两个人是否相同: %v\n", equal)
}
