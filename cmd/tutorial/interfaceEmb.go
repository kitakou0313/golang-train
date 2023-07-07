package main

type A interface {
	doInA()
}

type B interface {
	A
	doInB()
}

type IB struct {
}

func (ia *IB) doInA() {

}

func (ia *IB) doInB() {

}

type As struct {
}

type Bs struct {
	As
}

func main() {
	// interfaceの埋め込み
	// 埋め込まれたinterfaceの型の変数に代入できる
	// 複数interfaceの同時実装に近そう
	var ib B = &IB{}
	var ia A = &IB{}

	ib.doInA()
	ib.doInB()

	ia.doInA()
	// ia.doInB()はエラー

	// structの埋め込み
	var sb Bs = Bs{}
	// structの場合は埋め込まれた側の型の変数に代入できない（継承とは異なる）
	var sa As = Bs{}
}
