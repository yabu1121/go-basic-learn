package main 

import "fmt"

func main(){
	// 自動で0で初期化
	var arr1 [5]int;
	fmt.Printf("%v", arr1);

	arr1[0] = 1;
	fmt.Printf("%v", arr1);

	// 値指定で初期化
	arr2 := [5]int{1,2,3,4,5}
	fmt.Printf("%v", arr2);
	
	// 一部初期化宣言
	arr3 :=[5]int{1,2}
	fmt.Printf("%v", arr3);
	
	// len()で配列の長さを求められる。
	arr4 := [...]int{10,20,40}
	fmt.Printf("%v, %d\n", arr3, len(arr4));

	for i, v := range arr2{
		fmt.Printf("arr2[%d] = %d\n", i, v)
	}
	

	// スライスというのは可変長で変形できる
	// len()で長さを、cap()で容量を求められる

	var slice1 []int;
	fmt.Printf("s:%v l:%d c:%d\n", slice1, len(slice1), cap(slice1));
	// メソッドでいじるんじゃなくて再代入をする、
	slice1 = append(slice1, 1);
	fmt.Printf("s:%v l:%d c:%d\n", slice1, len(slice1), cap(slice1));
	// len(slice1) 長さを特定
	// slice1[len(slice1)-1] 要素にアクセス
	slice2 := []int{1,2,3,4,5};
	fmt.Printf("s:%v l:%d c:%d\n", slice2, len(slice2), cap(slice2));

	// スライスをmakeで作成できるようになる。
	slice3 := make([]int, 3);
	fmt.Printf("s:%v l:%d c:%d\n", slice3, len(slice3), cap(slice3));

	slice4 := make([]int, 3,5);
	fmt.Printf("s:%v l:%d c:%d\n", slice4, len(slice4), cap(slice4));
}