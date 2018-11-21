package godoc_test

import (
	"fmt"

	"github.com/gmailzj/godoc"
)

// 此注释将会被展示在页面上
// 此函数将被展示在OverView区域
func Example() {
	fmt.Println("Hello OverView")

	// Output:
	// Hello OverView
}
func ExampleGetUserName() {
	username := godoc.GetUserName()
}

// 此函数将被展示在Function区域
// 展示big标签
func ExampleGetUserName_big() {
	fmt.Println("Hello Banana")

	// Output:
	// Hello Banana
}
