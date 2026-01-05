package main
import "fmt"

func main(){
	// マップというデータ構造。
	// おそらくキーバリューをペアにして保持するデータ構造。
	// 簡単にmake関数で宣言をして作れる。
	// var map[string]int[]内はキーの値、その外はバリューの値の型を宣言する
	var m1 map[string]int	
	fmt.Printf("m1: %v (nil: %t)\n", m1, m1 == nil)

	m2 := make(map[string]int)
	fmt.Printf("m2: %v (nil: %t)\n", m2, m2 == nil)

	m3 := map[string]int{
		"太郎": 25,
		"花子": 23,
		"次郎": 30,
	}
	fmt.Printf("m3: %v\n", m3)

	ages := make(map[string]int)

	ages["吉田"] = 2
	ages["畑村"] = 14
	
	fmt.Printf("%v", ages)
	ages["畑村"] = 121
	fmt.Printf("%v", ages)
	fmt.Printf("%v", ages["太郎"])
	fmt.Printf("%v", ages["吉田"])
	age, exists := ages["太郎"]
	fmt.Print(age, exists)
	age, exists = ages["吉田"]
	fmt.Print(age, exists)
}