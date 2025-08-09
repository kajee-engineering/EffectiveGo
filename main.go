package main

import (
	// "MyGolang/basics"
	"MyGolang/effective"
	"MyGolang/effective/udon"
	"flag"
	"fmt"
	"io"
	"log"
)

func main() {

	// basics.Demo()

	var t effective.CarType
	t = effective.SUV
	fmt.Println("SUV :", t)
	fmt.Printf("愛車は%s\n", t)

	var o effective.CarOption
	o = effective.SunRoof
	if o&effective.SunRoof != 0 {
		fmt.Println("サンルーフ付き")
	}

	// error
	err := effective.New("これはエラーです")
	fmt.Println(err)
	fmt.Println(err.Error())
	fmt.Println(io.EOF) // EOFのerrors.NEW()を使って定義している

	// udon
	var k = udon.AllParam(udon.Large, false, 0)
	fmt.Println(k)

	uo := udon.Option{
		Men:      udon.Large,
		Aburaage: false,
		Ebiten:   0,
	}
	u := udon.New3th(uo)
	fmt.Println(u)

	udon.UseFluentInterfasce()
	result := udon.UseFuncOption
	fmt.Println(result)

	//
	// 1.7 プログラムを制御する引数
	//

	// コマンドライン引数
	flag.Parse()
	log.Println(*FlagStr)
	log.Println(*FlagInt)
	log.Println(flag.Args())

}

var (
	// コマンドライン引数を定義
	FlagStr = flag.String("string", "default", "文字列フラグ")
	FlagInt = flag.Int("int", -1, "数値フラグ")
)
