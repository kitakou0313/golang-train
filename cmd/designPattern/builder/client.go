package main

import "fmt"

// 多数のfieldを持つクラスの初期化に用いられる手法

/*
Prob:多数のfieldを持つクラスのインスタンスの初期化は困難である
- サブクラスに分類する場合：サブクラスのパターンが膨大になる
- 一つのクラスに全て集約する場合：コンストラクタの引数が膨大になり，default値のものもある

Solution:直接コンストラクタを呼ばず，インスタンスの初期化を責務とするbuilderクラス内でインスタンスを生成する

膨大な数の引数を持つコンストラクタを呼ぶよりもいくつかのステップで初期化を行う方が可読性が高いため，

builderクラスは
- 複数の初期化ステップ
- 生成されたインスタンス自体を返すメソッド
で構成される．

初期化ステップ呼び出し順の固定のため，
クライアントからの呼び出しを受けつけ既定された順序で
初期化を行うdirectorクラスを定義することもある．

builderが初期化のパラメータの状態の保持，各ステップでの初期化が責務
directorは実際のインスタンス生成が責務

summary:
field数が膨大であったり初期化時の手間が煩雑なクラスのインスタンス生成に対抗する手法
頻繁に呼び出される形式をbuilderクラスでの各ステップとして定義し，
Directorから順番を保って呼び出すことで不適切な初期化を防ぐ

Q:builderのgetResult内でbuild step呼ぶのは？
A:overrideでのサブクラス実装時，呼び出し順序が変化してしまうのでNG
Product間で共通して呼び出し順序を固定したい場合はDirectorクラスを定義するべき
*/

// Client
func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	fmt.Println(normalHouse)

	director.builder = iglooBuilder
	iglooHouse := director.buildHouse()

	fmt.Println(iglooHouse)

}
