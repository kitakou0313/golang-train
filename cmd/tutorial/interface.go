package main

import "fmt"

type IDoSomething interface {
	doSome()
}

type interfaceImplementsWithPointerReciever struct {
}

func (i *interfaceImplementsWithPointerReciever) doSome() {
	fmt.Println("hoge!")
}

type interfaceImplementsWithValReciever struct {
}

func (i interfaceImplementsWithValReciever) doSome() {
	fmt.Println("hoge?")
}

// Golangではメソッドがポインタ型について定義されているか値型について定義されているかでinterface型への代入の仕方が異なる

// 1つ目ではポインタレシーバについて定義されているので，interface型に代入するにはポインタにする必要がある

func useIDoSomething() {
	var hoge IDoSomething = &interfaceImplementsWithPointerReciever{}

	var huga IDoSomething = interfaceImplementsWithValReciever{}

	hoge.doSome()
	huga.doSome()
}
