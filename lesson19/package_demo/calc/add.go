package calc

import (
	"fmt"
	"lesson19/package_demo/snow"
)

// 标识符首字母大写表示对外可见
// 通常不对外的标识符都是要首字母小写

// Name 定义一个测试的全局变量
var Name = "沙河小王子"

// Add 是一个计算两个 int 类型数据和的函数
func Add(x, y int) int {
	snow.Snow()
	_ = Sub(x, y) // 同一个包中的不同文件可以直接调用
	return x + y
}

// init 函数在包导入的时候自动执行
// init 函数没有参数也没有返回值
func init() {
	fmt.Println("calc.init()")
	fmt.Println(Name)
}
