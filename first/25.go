package main

import "fmt"

// 1.声明1个usb接口，接口中包含 start stop方法
type Usb interface {
	start()
	stop()
}

// 2.声明1个Phone结构体
type Phone struct {
	Name string
}

// 3.实现 Usb 接口中的所有方法
func (thisPhone Phone) start() {
	fmt.Printf("start %v\n", thisPhone.Name)
}
func (thisPhone Phone) stop() {
	fmt.Printf("stop %v\n", thisPhone.Name)
}

func main() {
	// 4.实例化结构体
	var phone1 Usb = Phone{
		Name: "华为",
	}
	// 5.实现Usb接口，调用方法
	phone1.start()
	phone1.stop()
}
