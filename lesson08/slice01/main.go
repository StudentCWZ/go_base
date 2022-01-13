package main

import "fmt"

func main() {
	// var a []string
	// var b []int
	// var c = []bool{false, true}
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// 基于数组得到切片
	a := [5]int{55, 56, 57, 58, 59}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println("--------------")
	// 切片再次切片
	c := b[:]
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println("--------------")
	// make 函数构造切片
	d := make([]int, 5, 10)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println("--------------")
	// 通过 len() 函数获取切片的长度
	fmt.Println(len(d))
	// 通过 cap() 函数获取切片的容量
	fmt.Println(cap(d))
	fmt.Println("--------------")
	// nil
	var v []int     // 声明但是并没有初始化
	var n = []int{} // 声明并且初始化
	u := make([]int, 0)
	if v == nil {
		fmt.Println("v == nil")
	}
	fmt.Println(v, len(v), cap(v))
	if n == nil {
		fmt.Println("n == nil")
	}
	fmt.Println(n, len(n), cap(n))
	if u == nil {
		fmt.Println("u == nil")
	}
	fmt.Println(u, len(u), cap(u))
	// 切片赋值拷贝
	j := make([]int, 3) // [0, 0, 0]
	h := j
	h[0] = 100
	fmt.Println(j)
	fmt.Println(h)
	fmt.Println("--------------")
	// 切片的遍历
	l := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	fmt.Println("--------------")
	// for range 遍历
	for index, value := range l {
		fmt.Println(index, value)
	}
	fmt.Println("--------------")
	// 切片的扩容
	// 切片要初始化才能使用
	var m []int // 此时并没有申请内存
	// m[0] = 100 	// 报错： index out of range
	for i := 0; i < 10; i++ {
		m = append(m, i)
		fmt.Printf("%v len: %d cap: %d ptr:%p \n", m, len(m), cap(m), m)
	}
	fmt.Println("--------------")
	// 一次追加多个元素
	var p []int // 此时并没有申请内存
	p = append(p, 1, 2, 3, 4, 5)
	o := []int{12, 13, 14, 15}
	p = append(p, o...)
	fmt.Println(p)
	fmt.Println("--------------")
	// 切片的 copy
	q := []int{1, 2, 3, 4, 5}
	r := make([]int, 5)
	s := r
	copy(r, q)
	r[0] = 100
	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)
	fmt.Println("--------------")
	// 切片删除元素
	t := []string{"北京", "上海", "广州", "深圳"}
	t = append(t[0:2], t[3:]...)
	fmt.Println(t)
	fmt.Println("--------------")
	// 练习题
	var w = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		w = append(w, fmt.Sprintf("%v", i))
	}
	fmt.Println(w) // [     0 1 2 3 4 5 6 7 8 9]
	fmt.Println(len(w))
}
