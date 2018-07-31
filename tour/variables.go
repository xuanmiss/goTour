package main

import "fmt"

/**
 *Created By miss
 *
 *At 2018/7/27
 */
const flag = true //常量生命使用const

var c, java, python int
var cc, javaa, pythonn = "c or c++", "java or kt", "python or ruby" //包级别变量，全局可用
func doVariables() {
	myvar := "hello miss" //1、func 内的变量可以用短变量声明 :=  2、函数级别变量，仅在当前函数可用
	fmt.Println(myvar)
	fmt.Println(cc)

}

func typeChange() {
	s := 18 * 2
	var num = string(s) //转为 s 对应的ASCII 编码对应字符
	fmt.Println(s, num)
}

func main() {

	doVariables()
	typeChange()
	//fmt.Println(c,java,python)
	//fmt.Println(cc,"\n",javaa,"\n",pythonn)
}
