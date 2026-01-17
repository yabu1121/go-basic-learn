package main

import "fmt"

/*
【04_operators.go の練習問題】
以下の3つの問題を解いてください。
*/

// ========== 問題1 ==========
/*
【問題1】電卓プログラム
2つの数値に対して、すべての算術演算を行うプログラムを作成してください。

【要件】
1. 変数 a に 17、b に 5 を代入
2. 加算、減算、乗算、除算、剰余の結果を出力
3. 複合代入演算子（+=, -=, *=, /=, %=）を使った例も示す

【出力例】
========== 基本演算 ==========
a = 17, b = 5
17 + 5 = 22
17 - 5 = 12
17 * 5 = 85
17 / 5 = 3
17 % 5 = 2
---
========== 複合代入演算 ==========
result = 100
result += 20 → 120
result -= 30 → 90
result *= 2 → 180
result /= 3 → 60
result %= 7 → 4

【学習ポイント】
- 算術演算子の使い方
- 整数除算の挙動
- 複合代入演算子
*/

func problem1() {
	var a int = 17
	var b int = 5
	result := a + b
	fmt.Printf("%d\n", result)
	result = a - b
	fmt.Printf("%d\n", result)
	result = a * b
	fmt.Printf("%d\n", result)
	result = a / b
	fmt.Printf("%d\n", result)
	result = a % b
	fmt.Printf("%d\n", result)

	result = a + b
	fmt.Printf("%d\n", result)
	result -= a - b
	fmt.Printf("%d\n", result)
	result *= a * b
	fmt.Printf("%d\n", result)
	result /= a / b
	fmt.Printf("%d\n", result)
	result %= a % b
	fmt.Printf("%d\n", result)
}

// ========== 問題2 ==========
/*
【問題2】条件判定プログラム
比較演算子と論理演算子を使って、複雑な条件を判定してください。

【要件】
1. 年齢（age）、所持金（money）、会員ステータス（isMember）の変数を用意
2. 以下の条件を判定して結果を出力：
   - 成人かどうか（age >= 20）
   - 購入可能かどうか（money >= 1000）
   - 割引対象かどうか（成人 AND 会員）
   - 特別オファー対象かどうか（会員 OR 所持金が5000円以上）

【出力例】
========== ユーザー情報 ==========
年齢: 25歳
所持金: 3000円
会員: true
---
========== 判定結果 ==========
成人: true
購入可能: true
割引対象: true (成人 AND 会員)
特別オファー: true (会員 OR 所持金>=5000)

【学習ポイント】
- 比較演算子（==, !=, <, <=, >, >=）
- 論理演算子（&&, ||, !）
- 複合条件の評価
*/

func problem2() {
	// ここにコードを書いてください

}

// ========== 問題3 ==========
/*
【問題3】ビット演算の実践
ビット演算子を使って、フラグ管理を実装してください。

【要件】
権限フラグを以下のように定義：
- READ権限: 1 (0001)
- WRITE権限: 2 (0010)
- EXECUTE権限: 4 (0100)
- DELETE権限: 8 (1000)

1. ユーザーにREADとWRITE権限を付与
2. 各権限を持っているか確認
3. EXECUTE権限を追加
4. WRITE権限を削除
5. 各ステップの権限状態を2進数で表示

【出力例】
========== 権限管理システム ==========
READ権限: 1 (0001)
WRITE権限: 2 (0010)
EXECUTE権限: 4 (0100)
DELETE権限: 8 (1000)
---
初期権限: 3 (0011) [READ, WRITE]
READ権限あり: true
WRITE権限あり: true
EXECUTE権限あり: false
---
EXECUTE権限追加後: 7 (0111)
WRITE権限削除後: 5 (0101) [READ, EXECUTE]

【ヒント】
- 権限追加: flags |= permission
- 権限削除: flags &^= permission
- 権限確認: (flags & permission) != 0
- 2進数表示: fmt.Printf("%04b", num)
*/

func problem3() {
	// ここにコードを書いてください

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
go run 04.go

【学習ポイント】
1. 算術演算子（+, -, *, /, %）
2. 比較演算子（==, !=, <, <=, >, >=）
3. 論理演算子（&&, ||, !）
4. ビット演算子（&, |, ^, &^, <<, >>）
5. 代入演算子（=, +=, -=, *=, /=, %=）
6. 実用的なビット演算の使い方
*/
