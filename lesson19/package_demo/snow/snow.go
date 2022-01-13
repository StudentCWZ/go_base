package snow

import "fmt"

// 被 calc 包导入的一个包

func Snow() {
	fmt.Println("下雪了。。。")
}

func init() {
	fmt.Println("snow.init()")
}
