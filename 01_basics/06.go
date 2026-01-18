package main

import (
	"fmt"
)

/*
【06_arrays_slices.go の練習問題】
以下の3つの問題を解いてください。
*/

// ========== 問題1 ==========
/*
【問題1】配列とスライスの基本操作
配列とスライスの違いを理解するプログラムを作成してください。

【要件】
1. 5つの整数を格納する配列を作成し、値を代入
2. 同じ値を持つスライスを作成
3. 配列とスライスの長さと容量を出力
4. 配列の一部をスライスで切り出す
5. スライスに要素を追加（append）

【出力例】
========== 配列 ==========
配列: [10 20 30 40 50]
長さ: 5

========== スライス ==========
スライス: [10 20 30 40 50]
長さ: 5, 容量: 5

========== スライス操作 ==========
部分スライス [1:4]: [20 30 40]
要素追加後: [10 20 30 40 50 60 70]
長さ: 7, 容量: 10

【学習ポイント】
- 配列とスライスの宣言
- len()とcap()の違い
- スライス演算子の使い方
- appendの動作
*/

func problem1() {
	// 配列とスライスの違い

	// len　スライスの現在の要素数を返す
	// cap スライスが確保しているメモリの最大数
	// append(s, val) s(すらいす)の末尾にvalを更新して、返す
	// make([]型, len, cap)長さと容量を指定してスライスを作成する
	// copy(dest, src) スライスの内容をスライスにコピーする
	slice1 := []int{1,2,3,4,5}
	array1 := [8]int{1,2,3,4,5}
	fmt.Printf("len: %d, cap: %d, そのまま: %v\n",len(slice1), cap(slice1), slice1)
	fmt.Printf("len: %d, cap: %d, そのまま: %v\n",len(array1), cap(array1), array1)
	fmt.Printf("部分スライス%v\n", slice1[1:])
	slice1 = append(slice1, 2,2)
	// 容量が足りなくなると二倍のメモリ領域を確保する。
	fmt.Printf("len: %d, cap: %d, そのまま: %v\n",len(slice1), cap(slice1), slice1)
}

// ========== 問題2 ==========
/*
【問題2】成績管理システム
スライスを使って、学生の成績を管理するプログラムを作成してください。

【要件】
1. 学生名のスライスを作成（5人分）
2. 各学生の点数のスライスを作成
3. 全学生の平均点を計算
4. 最高点と最低点を見つける
5. 80点以上の学生をフィルタリング
6. 結果を整形して出力

【出力例】
========== 成績一覧 ==========
太郎: 85点
花子: 92点
次郎: 78点
美咲: 88点
健太: 95点
---
平均点: 87.6点
最高点: 95点 (健太)
最低点: 78点 (次郎)
---
========== 優秀者（80点以上）==========
- 太郎: 85点
- 花子: 92点
- 美咲: 88点
- 健太: 95点

【学習ポイント】
- スライスのループ処理
- rangeの使い方
- スライスのフィルタリング
- 複数のスライスの連携
*/

func problem2() {
	// names := []string{"太郎","花子","次郎","美咲","健太"}
	// scores := []int{85,92,78,88,95}
	// var minScore int = scores[0] 
	// var minName string = names[0]
	// var maxScore int = scores[0]
	// var maxName string = names[0]
	// var total float64 = 0
	// for i := range names {
	// 	fmt.Printf("%s: %d点\n",names[i], scores[i])
	// 	if minScore > scores[i] {
	// 		minScore = scores[i]
	// 		minName = names[i]
	// 	}
	// 	if maxScore < scores[i] {
	// 		maxScore = scores[i]
	// 		maxName = names[i]
	// 	}
	// 	total += float64(scores[i])
	// }
	// var average float64 = total / float64(len(names))
	// fmt.Println("---")
	// fmt.Printf("平均点: %f\n",average)
	// fmt.Printf("最高点: %d (%s)\n",maxScore, maxName)
	// fmt.Printf("最低点: %d (%s)\n",minScore, minName)
	// fmt.Println("---")

	// var border int = 80
	// fmt.Println("---成績優秀者---")
	// for i := range scores {
	// 	if scores[i] >= border {
	// 		fmt.Printf("%s: %d点\n", names[i], scores[i])
	// 	}
	// }

	type Student struct {
		Name string
		Score int
	}

	var students = []Student{
		{Name: "太郎", Score: 85},
		{Name: "花子", Score: 92},
		{Name: "次郎", Score: 78},
		{Name: "美咲", Score: 88},
		{Name: "健太", Score: 95},
	}

	best := students[0]
	worst := students[0]
	total := 0.0
	for _, s := range students {
	fmt.Printf("%s: %d点\n", s.Name, s.Score)
	total += float64(s.Score)

	if s.Score > best.Score {
		best = s
	}
	if s.Score < worst.Score {
		worst = s
	}
}

// 5. 結果の出力
fmt.Printf("---\n平均点: %.1f点\n", total/float64(len(students)))
fmt.Printf("最高点: %d点 (%s)\n", best.Score, best.Name)
fmt.Printf("最低点: %d点 (%s)\n---", worst.Score, worst.Name)

// 6. フィルタリング（優秀者）
fmt.Println("\n========== 優秀者（80点以上）==========")
for _, s := range students {
	if s.Score >= 80 {
		fmt.Printf("- %s: %d点\n", s.Name, s.Score)
	}
}
} 

// ========== 問題3 ==========
/*
【問題3】動的配列の操作
スライスの挿入、削除、結合などの高度な操作を実装してください。

【要件】
1. 初期スライス: [1, 2, 3, 4, 5]
2. インデックス2に10を挿入 → [1, 2, 10, 3, 4, 5]
3. インデックス4の要素を削除 → [1, 2, 10, 3, 5]
4. 別のスライス [6, 7, 8] を結合
5. スライスをコピーして、元のスライスと独立していることを確認
6. 各操作後のスライスの状態を出力

【出力例】
========== スライス操作 ==========
初期状態: [1 2 3 4 5]
長さ: 5, 容量: 5

インデックス2に10を挿入:
結果: [1 2 10 3 4 5]
長さ: 6, 容量: 10

インデックス4の要素(4)を削除:
結果: [1 2 10 3 5]
長さ: 5, 容量: 10

[6 7 8]を結合:
結果: [1 2 10 3 5 6 7 8]
長さ: 8, 容量: 10

========== コピーの確認 ==========
元のスライス: [1 2 10 3 5 6 7 8]
コピー: [1 2 10 3 5 6 7 8]
コピーの先頭を変更: [999 2 10 3 5 6 7 8]
元のスライス: [1 2 10 3 5 6 7 8] (変化なし)

【学習ポイント】
- スライスへの挿入
- スライスからの削除
- スライスの結合
- copy()の使い方
- スライスの参照の理解
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
go run 06.go

【学習ポイント】
1. 配列とスライスの違い
2. len()とcap()の理解
3. スライス演算子 [start:end]
4. append()の使い方
5. スライスの挿入・削除
6. copy()とスライスの参照
7. rangeを使ったループ
*/
