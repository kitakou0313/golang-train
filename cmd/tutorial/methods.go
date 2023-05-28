package main

import "fmt"

// 値レシーバとポイントレシーバの違いを見る

type Type1 struct {
	val int
}

func (t1 Type1) methodValReciever(arg int) {
	t1.val = arg
}

func (t1 *Type1) methodPointerReciver(arg int) {
	t1.val = arg
}

// Golangでは上記のレシーバは内部的にはレシーバを引数に持つ関数として定義される
// Type1.methodValReciever(t1, arg)
// Pythonもこんな感じだった記憶がある（self）

// なので値レシーバには値渡しでレシーバのオブジェクトが渡されるため，fieldを変更してもレシーバには反映されない

// メソッドの定義はポインタ型と値型で別々だが，呼び出し時はコンパイラがいい感じに変換してくれるらしい

func main() {

	type1 := Type1{
		val: 1,
	}

	type1.methodValReciever(10)
	fmt.Println(type1)

	type1.methodPointerReciver(20)
	fmt.Println(type1)

}
