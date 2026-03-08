package main

import (
	"fmt"
	"time"
)

func main() {
	// sayHello("同期")
	// sayHello("同期")　// 1 2 3 1 2 3

	// go sayHello("非同期1")
	// go sayHello("非同期2")
	// go sayHello("非同期3")
	// time.Sleep(1000 * time.Millisecond) // 0 0 0 1 1 1 2 2 2, timeSleepがないと何も起こらないので自動終了してしまう。
	// ぐちゃぐちゃ

	// //無名関数 go 
	// go func () {
	// 	fmt.Println("無名関数")
	// }()
	// // この奥の()はそのまま実行するということでお間違いないでしょうか。

	// go func (message string){
	// 	fmt.Printf("%s", message)
	// }("hello")
	// time.Sleep(500 * time.Millisecond)

	// 複数の並列
	// for  i:= 1; i <= 5; i++ {
	// 	go func (n int){
	// 		fmt.Printf("%d\n", i)	
	// 	}(i)
	// 	time.Sleep(1000 * time.Millisecond)
	// }i

	// start := time.Now()

	// go task("1", 500)
	// go task("2", 500)
	// go task("3", 500)
	// time.Sleep(600 * time.Millisecond)
	// fmt.Printf("%v", time.Since(start))
	// 同期 1.50s
	// 非同期　0.15ms
}

// sayHello関数はとりあえず三回繰り返しで
// func sayHello(name string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Printf("%s: %d\n", name, i)
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func printNumber(n int) {
// 	fmt.Printf("番号: %d\n", n)
// }

func task(name string, ms int) {
	fmt.Printf("%s 開始\n", name)
	time.Sleep(time.Duration(ms) * time.Millisecond)
	fmt.Printf("%s 完了\n", name)
}

/*
【実行方法】
go run 01_goroutines.go

【重要な概念】
1. Goroutine は軽量スレッド
2. go キーワードで起動
3. main関数が終了すると全Goroutineも終了
4. 実行順序は保証されない

【Goroutine の特徴】
- 非常に軽量（数KB）
- 数千〜数万のGoroutineを起動可能
- Goランタイムが自動的にスケジューリング

【注意点】
1. main関数が終了するとGoroutineも終了
2. ループ変数のキャプチャに注意
3. Goroutineリークに注意（終了しないGoroutine）
4. 共有メモリへのアクセスは同期が必要

【ベストプラクティス】
1. Goroutineの終了を管理する
2. WaitGroupやChannelで同期
3. ループ変数は引数として渡す
4. context でキャンセル可能にする

【次のステップ】
02_channels.go でChannelを学びましょう
*/
