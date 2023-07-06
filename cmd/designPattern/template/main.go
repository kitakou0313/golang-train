package main

import "fmt"

// template pattern
/*
アルゴリズムの全体構造は変わらないが，
一部の詳細の振る舞いが異なるようなものの
実装に適したパターン

親クラス，interfaceで全体の構造を定義しておき，
その中で呼び出されるメソッドをoverrideすることで
詳細の振る舞いを変化させる
*/
func main() {
	smsOtp := &Sms{}

	o := Otp{
		iOtp: smsOtp,
	}
	o.genAndSendOTP(4)

	fmt.Println()

	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.genAndSendOTP(4)
}
