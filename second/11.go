package main

func main() {

}

// 1.声明两个函数 查看 return 和 defer 执行的时机
// 1.1 第一个函数中 声明一个变量 a， defer一个内部函数，函数中执行 a++，最后由外层函数返回 a
// 1.2 第一个函数接受一个变量 a，defer一个内部函数，函数中执行 a++，最后由外层函数返回 a
// 1.3 查看上方执行不同

// 2.使用defer panic抛出异常 提示：要使用 recover() 参考 22.go