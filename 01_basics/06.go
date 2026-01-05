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
	
}