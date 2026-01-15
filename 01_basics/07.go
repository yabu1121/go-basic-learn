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
	// キーバリューがmap[key]value
	fmt.Printf("m3: %v\n", m3)

	// makeで作る
	ages := make(map[string]int)

	ages["吉田"] = 2
	ages["畑村"] = 14
	
	fmt.Printf("%v", ages)
	ages["畑村"] = 121
	fmt.Printf("%v", ages)
	fmt.Printf("%v", ages["太郎"])
	fmt.Printf("%v", ages["吉田"])
	// 存在しないならageに0が入る
	age, exists := ages["太郎"]
	fmt.Print(age, exists)
	age, exists = ages["吉田"]
	fmt.Print(age, exists)

	fmt.Println("")
	fmt.Println("削除")
	fmt.Printf("%v\n", m3)
	// delete(mapの変数、key)
	delete(m3, "太郎")
	fmt.Printf("%v\n", m3)


	scores := map[string]int{
		"数学": 85,
		"英語": 92,
		"国語": 78,
		"理科": 88,
	}

		// キーと値の両方
	for subject, score := range scores {
		fmt.Printf("%s: %d点\n", subject, score)
	}

	fmt.Println(len(scores))


	// ネストマップの宣言
		students := map[string]map[string]int{
		"太郎": {
			"数学": 85,
			"英語": 90,
		},
		"花子": {
			"数学": 95,
			"英語": 88,
		},
	}
	println(students)


	// 構造体
		type Person struct {
		Name string
		Age  int
	}

	// rune = int32
	fmt.Println("\n=== 実用例：文字カウンター ===")

	text := "hello world"
	charCount := make(map[rune]int)

	for _, char := range text {
		charCount[char]++
	}

	for char, count := range charCount {
		fmt.Printf("'%c': %d回\n", char, count)
	}
}