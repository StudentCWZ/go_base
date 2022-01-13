package main

import (
	"fmt"
	"reflect"
)

// reflect demo

func reflectType(x interface{}) {
	// 我是不知道别人调用我这个函数的时候会传进来什么类型的变量
	// 1. 方式一：通过类型断言
	// 2. 方法二：借助反射
	obj := reflect.TypeOf(x)
	fmt.Println(obj, obj.Name(), obj.Kind())
	fmt.Printf("%T\n", obj)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	fmt.Printf("%v, %T\n", v, v)
	k := v.Kind() // 拿到值对应的类型
	fmt.Println(k)
	// 如何得到一个传入时候的变量
	switch k {
	case reflect.Float32:
		// 把反射渠道的值转换成一个 int32 类型的变量。
		ret := float32(v.Float())
		fmt.Printf("%v %T\n", ret, ret)
		fmt.Println(ret)
	case reflect.Int32:
		ret := int32(v.Int())
		fmt.Printf("%v %T\n", ret, ret)
		fmt.Println(ret)
	}

}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	// Elem() 用来根据指针取对应的值
	k := v.Elem().Kind()
	switch k {
	case reflect.Int32:
		v.Elem().SetInt(100)
	case reflect.Float32:
		v.Elem().SetFloat(3.21)
	}
}

type Cat struct{}

type Dog struct{}

func main() {
	var a float32 = 1.234
	reflectType(a)
	var b int8 = 1
	reflectType(b)
	// 结构体
	var c Cat
	reflectType(c)
	var d Dog
	reflectType(d)
	// slice
	var e []int
	reflectType(e)
	var f []string
	reflectType(f)

	// valueOf
	var aa int32 = 100
	reflectValue(aa)
	var bb float32 = 1.234
	reflectValue(bb)

	// set value
	var aaa int32 = 10
	reflectSetValue(&aaa)
	fmt.Println(aaa)

	/*
		IsNil() 常被用于判断指针是否为空；IsValid() 常被用于判定返回值是否有效
	*/
	// *int 类型指针
	var x *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(x).IsNil())
	// nil 值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	y := struct{}{}
	// 尝试从结构体中查找 abc 字段
	fmt.Println("不存在的结构体成员:", reflect.ValueOf(y).FieldByName("abc").IsValid())
	// 尝试从结构体中查找 abc 方法
	fmt.Println("不存在的结构体方法:", reflect.ValueOf(y).MethodByName("abc").IsValid())
	// map
	z := map[string]int{}
	// 尝试从 map 中查找一个不存在的键
	fmt.Println("map 中不存在的键:", reflect.ValueOf(z).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}
