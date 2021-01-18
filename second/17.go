package main

import "fmt"

// 1.声明一个公有结构体、声明一个私有结构体
type Teacher struct {
	Name    string
	Age     int
	Subject string
}
type student struct {
	name      string
	age       int
	className string
}

func main() {
	// 2.实例化上方声明的结构体
	// 2.1 请使用至少3种实例化结构体的方式（要包括普通赋值和指针赋值）
	tea1 := Teacher{
		Name:    "王老师",
		Age:     30,
		Subject: "数学",
	}
	tea2 := &Teacher{
		Name:    "张老师",
		Age:     50,
		Subject: "语文",
	}
	tea3 := new(Teacher)
	tea3.Name = "赵老师"
	tea3.Age = 20
	tea3.Subject = "体育"
	fmt.Printf("%+v\n", tea1)
	fmt.Printf("%+v\n", *tea2)
	fmt.Printf("%+v\n", tea3)
	// 3.声明一个使用类型默认值实例化结构体
	stuDefault := student{}
	fmt.Printf("stuDefault = %#v\n", stuDefault)
	// 4.请使用一段代码证明结构体是值类型
	copyStu := stuDefault
	copyStu.name = "钱老师"
	fmt.Printf("copyStu = %#v\n", copyStu)
}
