package main

import "fmt"

/*
【02_variables.go の練習問題】
以下の3つの問題を解いてください。
*/

// ========== 問題1 ==========
/*
【問題1】温度変換プログラム
摂氏温度を華氏温度に変換するプログラムを作成してください。

【要件】
1. 摂氏温度を変数celsiusに格納（例: 25.0）
2. 華氏温度 = 摂氏温度 × 9/5 + 32 の式で計算
3. 両方の温度を出力

【出力例】
摂氏: 25.0°C
華氏: 77.0°F

【使用する概念】
- 変数宣言（var または :=）
- 型推論
- 算術演算
- fmt.Printf()
*/

func problem1() {
	// ここにコードを書いてください
	var celsius float32 = 25.0
	var fahrenheit float32 = celsius * 9 / 5 + 32
	var cmsg string = fmt.Sprintf("%.1f°C", celsius)
	fmt.Println(cmsg)
	var fmsg string = fmt.Sprintf("%.1f°F", fahrenheit)
	fmt.Println(fmsg)
}

// ========== 問題2 ==========
/*
【問題2】ユーザー情報の管理
以下の情報を適切な変数に格納し、出力してください。

【格納する情報】
- ユーザーID: 12345 (int型)
- ユーザー名: "田中花子" (string型)
- メールアドレス: "hanako@example.com" (string型)
- 年齢: 28 (int型)
- プレミアム会員: true (bool型)
- ポイント残高: 1500.50 (float64型)

【出力例】
========== ユーザー情報 ==========
ID: 12345
名前: 田中花子
メール: hanako@example.com
年齢: 28歳
プレミアム会員: true
ポイント残高: 1500.50pt
==================================

【使用する概念】
- 複数の変数宣言
- 異なる型の変数
- ゼロ値の理解
*/

type Person struct {
	Id int
	Name string
	Email string
	Age int
	Is_Premium bool
	Point float64
}

// func problem2() {
// 	// ここにコードを書いてください
// 	var user_id int = 12345
// 	var user_name string = "田中花子"
// 	var user_email string = "hanako@example.com"
// 	var user_age int = 28
// 	var user_is_premium bool = true
// 	var user_point float64 = 1500.50

// 	fmt.Printf("ID: %d\n", user_id)
// 	fmt.Printf("名前: %s\n", user_name)
// 	fmt.Printf("メール: %s\n", user_email)
// 	fmt.Printf("年齢: %d歳\n", user_age)
// 	fmt.Printf("プレミアム会員: %v\n", user_is_premium)
// 	fmt.Printf("Iポイント残高 %.2fpt\n", user_point)
// }

func problem2() {
	// ここにコードを書いてください
	// var p Person = new(Person)
  // p := new(Person)
	// p.Id = 12345
	// p.Name = "田中花子"
	// p.Email = "hanako@example.com"
	// p.Age = 28
	// p.Is_Premium = true
	// p.Point = 1500.50

	p := Person {
		Id: 12345,
		Name: "田中花子",
		Email: "hanako@example.com",
		Age: 28,
		Is_Premium: true,
		Point: 1500.50,
	}

	fmt.Printf("ID: %d\n", p.Id)
	fmt.Printf("名前: %s\n", p.Name)
	fmt.Printf("メール: %s\n", p.Email)
	fmt.Printf("年齢: %d歳\n", p.Age)
	fmt.Printf("プレミアム会員: %v\n", p.Is_Premium)
	fmt.Printf("Iポイント残高 %.2fpt\n", p.Point)
}

// ========== 問題3 ==========
/*
【問題3】定数を使った計算
円周率(π)を定数として定義し、円の面積と円周を計算してください。

【要件】
1. 円周率を定数Piとして定義（3.14159）
2. 半径を変数radiusに格納（例: 5.0）
3. 面積 = π × 半径²
4. 円周 = 2 × π × 半径
5. 結果を出力

【出力例】
半径: 5.00cm
円周: 31.42cm
面積: 78.54cm²

【使用する概念】
- 定数の宣言（const）
- 変数と定数の違い
- 算術演算
- 型変換（必要に応じて）
*/

func problem3() {
	// きほんてきにGoではfloatは64, constには型を指定しない
	const Pi = 3.1415926535
	var radius float64 = 5.0
	var area float64 = Pi * radius * radius
	var circumference float64 = 2 * Pi * radius
	fmt.Printf("半径: %.2fcm\n",radius)
	fmt.Printf("円周: %.2fcm\n",circumference)
	fmt.Printf("面積: %.2f㎠\n",area)
}

func main() {
	fmt.Println("========== 問題1 ==========")
	problem1()

	fmt.Println("\n========== 問題2 ==========")
	problem2()

	fmt.Println("\n========== 問題3 ==========")
	problem3()
}

/*
【実行方法】
go run 02.go

【学習ポイント】
1. var、:=、constの使い分け
2. 型推論と明示的な型宣言
3. 異なる型の変数の扱い
4. 定数の使用場面
5. ゼロ値の理解
*/
