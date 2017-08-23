package main

import "fmt"

func main(){
	var str string
	fmt.Printf("请输入内容：")

	fmt.Scanf("%s", &str)

	fmt.Printf("输入了：%s", str)
}