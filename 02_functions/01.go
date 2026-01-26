package main

import "fmt"

func greet(){
	fmt.Println("hello")
}

func taro(){
	var taro string = "太郎"
	var old int = 23
	fmt.Printf("%s is %d years old.\n", taro, old)
}

func ReturnFunction () int {
	return 20
}

func calc_num (a, b int) (int, int ){
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {
	greet()
	taro()


	var return_number int = ReturnFunction() 
	fmt.Println(return_number)

	var a int = 5
	var b int = 2
	var sum, diff int = calc_num(a,b)
	fmt.Println(sum)
	fmt.Println(diff)
}