package basics

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
) // パッケージ名はすべて小文字でつける

// 関数定義
// 同じ型が続く場合は、後ろにあるものが前にあるものとして扱われる。
// 関数の中に関数は定義できない。
func calc(x, y int) int {
	return x + y
}

// calcAge 与えられた引数が古い間は返り値をインクリメントする
// 返り値に名前をつける。複数になる場合は、かっこでくくる。
func calcAge(y int, m time.Month, d int) (age int, err error) {

	b := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
	fmt.Println(b)
	n := time.Now()
	if b.After(n) {
		err = errors.New("誕生日が未来です")
		return
	}
	for {
		b = time.Date(y+age+1, m, d, 0, 0, 0, 0, time.Local)
		if b.After(n) {
			return
		}
		age++
	}

}

func Demo() {
	fmt.Println("Hello, World!")

	// Goは型を名前の後ろに定義する。省略可能性が高いものほど後ろになる。
	// var num1 int = 123

	// 右辺で型が決まる場合は型名は不要
	// 関数外（パッケージレベル）でも利用できる
	// var num2 = 123

	// 変数宣言と代入を同時に行う短縮記法
	// ただし関数内部でのみ利用可能
	// num3 := 123

	// Goは宣言だけして未使用な変数がコードに残っているとエラーになるためコメントアウトしている。

	x := 1 // 関数ローカルな変数または引数で、宣言場所と利用場所が近く、説明がなくてもいいものは短く。1文字などの場合もあり。エラーを返す変数はerr。コンテキストはctx。
	if true {
		// コードブロックの内側で変数 x を再宣言できる
		// 外側の x には影響を及ぼさない
		x := 2         // シャドーイングされた値(よく見るとGolandが外側の x と違う色で表示しており、視覚的に区別してくれている。)
		fmt.Println(x) // CamelCaseまたはcamelCaseとする。先頭が大文字なのでパッケージ外からも使えるパブリックな要素。なお先頭が小文字の場合はプライベートな要素である。
	}
	fmt.Println(x) // ここの位置では必ず1になる

	// 空行なしで関数や型の宣言の直前につけると、
	// それらの要素の説明として利用される
	// func CommentedFunc() {
	// }

	// var i int = 123
	// 数値同士の変換はかっこでくくり型を前置する
	// var f float64 = float64(i)
	// 64ビットOSで64ビットのintと、int64の明示的な変換が必要
	// var i64 int64 = int64(i)
	// boolへの変換は比較演算子を使う
	// var b bool = i != 0

	// 文字列との変換はstrconvパッケージを利用
	in := 12345
	fmt.Println(in)
	// strconvの数値入力はint64, uint64, float64なので
	// それ以外の変数を使うときは型変換が必要
	s := strconv.FormatInt(int64(in), 10) // 10進数
	fmt.Println(s)                        // "12345"

	// Parse系はエラー変換失敗時にエラーを返す
	// 成功時のerrはnil
	f, err := strconv.ParseFloat("12.3456", 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

	/*
		ポインター
		変数を作ってデータを格納すると、その変数はOSからもらったメモリのどこかに保存される。
		メモリにはアドレスがあります。そのアドレスを扱う機能がポインターである。
		Goのポインターは次のように使う。
		- 変数のポインター型には*（アスタリスク）を前置する
		- 既存の変数のポインターを取り出すには&を利用する
	*/
	var i int = 10 // ポインターの参照先となる普通の変数
	var p *int     // ポインターを格納する変数（デフォルトはnil）。この場合、ポインタの参照先がint型なので*intになる。stringであれば*stringになる。
	// fmt.Println(*p) // nilの参照先を取り出すと "panic: runtime error: invalid memory address or nil pointer dereference" が発生する
	p = &i // pにはiのアドレスが入る
	fmt.Println(p)

	/*
		ゼロ値
		変数宣言をしたが、まだ値を設定していない変数はゼロ値で初期化される。
		int 0
		bool false
		string 空文字列
		ポインター nil
		インタフェース nil

		他の言語の中には未初期化の変数はバグの原因になったり、セキュリティホールの原因となるものもある（未初期化変数利用とバッファオーバーラン（確保した以上のメモリにアクセスしてしまう不具合））
		Goはどちらも行えないようになっている。
	*/

	// Goは配列よりスライスを多用する。
	// Goの配列は固定長。一方、スライスは不定。
	// あまり使わない配列
	var nums [3]int = [3]int{1, 2, 3}
	fmt.Println(nums)
	fmt.Println(nums[0], nums[1], nums[2])
	fmt.Println(len(nums))

	// スライスの変数宣言
	var nums1 []int
	fmt.Println(nums1)
	fmt.Println(nums1 == nil) // ゼロ値はnilなのでtrue

	num2 := []int{1, 2, 3}
	fmt.Println(num2)
	num3 := nums[0:2] // 配列からスライス作成
	fmt.Println(num3) // [1 2]
	num4 := num3[1:3] // [2 3] 元の配列numsを参照している
	fmt.Println(num4)

	// 配列と同じようにブラケットで要素取得可能
	// 範囲外アクセスはパニック。セキュリティホールの原因になり得るため、安全性の高い言語であるGoはエラーとして扱う。
	fmt.Println(num2[1]) // 2
	//fmt.Println(num2[3]) // panic: runtime error: index out of range [3] with length 3

	// 要素の割り当ても可能
	num2[0] = 100
	fmt.Println(num2)

	// 長さも取得可能
	fmt.Println(len(num2)) // 3

	// スライスに要素を追加
	// 再代入が必要
	num2 = append(num2, 4)
	fmt.Println(num2) // [100 2 3 4]
	// 再代入しない場合
	fmt.Println(append(num2, 50)) // [100 2 3 4 50] の出力になるが...
	// append関数に渡す引数は追加されない。
	fmt.Println(num2) // [100 2 3 4]

	// 辞書、ハッシュなどとも呼ばれるデータ構造はGoではmapと呼ぶ
	// スライスと同様に複合型である
	// 数字がキーで値が文字列のマップ
	// HTTPステータスを格納
	hs := map[int]string{
		200: "OK",
		404: "Not found",
	}
	fmt.Println(hs)
	fmt.Println(hs[200])

	// makeで作る
	author := make(map[string][]string)
	fmt.Println(author == nil) // 空のmapが存在するのでfalse

	// ブラケットで要素アクセス
	author["Go"] = []string{"Robert Griesemer", "Rob Pike", "Ken Thompson"}
	fmt.Println(author)

	// データ取得
	status := hs[200]
	fmt.Println(status) // "OK"

	// 存在しない要素にアクセスするとゼロ値
	fmt.Println(hs[0]) // string型のゼロ値は空文字

	// あるかどうかの情報も一緒に取得
	status, exist := hs[304]
	// status = ""
	// exist = false
	fmt.Println(status, exist)

	status, exist = hs[200]
	fmt.Println(status, exist)

	// if文/if else/elseの基本的な書き方
	statusCode := 200
	if statusCode == 200 {
		fmt.Println("no error")
	} else if statusCode < 500 {
		fmt.Println("client error")
	} else {
		fmt.Println("server error")
	}

	// データ取得とチェックを同時に行う
	if result, ok := hs[300]; ok {
		fmt.Println("result is :", result)
	} else {
		fmt.Println("result is :", result)
		fmt.Println("ok is :", ok)
	}

	// スライスやマップの各要素に対してループ
	scketches := []string{"Dead Parrot", "Killer joke", "Spanish Inquisition", "Spam"}
	for i, s := range scketches {
		fmt.Println(i, s)
	}

	// 1つ変数だけ書けばインデックスのみを受け取れる
	for i := range scketches {
		fmt.Println(i)
	}

	// ブランク識別子でインデックスを読み飛ばして値だけを受け取れる
	for _, s := range scketches {
		fmt.Println(s)
	}

	fmt.Println("---------------------")
	for _, s := range scketches {
		if strings.HasPrefix(s, "K") {
			continue
		}

		if strings.HasSuffix(s, "n") {
			break
		}
		fmt.Println(s)
	}

	// 回り続けるループ
	counter := 0
	for counter < 10 {
		fmt.Println("ブール値がtrueの間は回り続けるループ")
		counter += 1
	}

	//end := time.Now().Add(time.Second)
	//for {
	//	fmt.Println("breakやreturnで抜けないと終わらないループ")
	//	if end.Before(time.Now()) {
	//		break
	//	}
	//}

	// 伝統的なforループ
	for i := 0; i < 10; i++ {
		fmt.Println(i, "回目")
	}

	// switch文
	w := []string{"running", "run", "walk", "stop", "stop", "sleep"}
	for _, s := range w {
		switch s {
		case "running":
			fallthrough
		case "run":
			fmt.Println("run")
			fallthrough
		case "r":
			fmt.Println("r")
		case "stop":
			fmt.Println("stop")
		default:
			fmt.Println("other...")
		}
	}

	// 名前付き関数はトップレベルでしか作れないが、無名関数は関数の中でも作れる。
	mf := func(x, y int) int {
		return x + y
	}

	// 無名関数の呼び出し
	fmt.Printf("%d\n", mf(10, 20)) // "30"

	// 無名関数として定義したものを渡せる
	doCalc(10, 20, mf)
	// 引数内で無名関数を定義できる
	doCalc(10, 20, func(a, b int) int {
		return 1
	})

	age, err := calcAge(2000, time.April, 1)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(age, "years old")

	DemoDefer()
	DemoWebApp()
	DemoStruct()
}

func doCalc(x, y int, f func(int, int) int) {
	fmt.Println(f(1, 1))
	fmt.Println(x, y)
}
