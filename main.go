package main

import (
	// "MyGolang/basics"
	"MyGolang/effective"
	"MyGolang/effective/udon"
	"fmt"
	"io"
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
	var k = udon.New(udon.Large, false, 0)
	fmt.Println(k)

	uo := udon.Option{
		Men:      udon.Large,
		Aburaage: false,
		Ebiten:   0,
	}
	var u = udon.NewOption(uo)
	fmt.Println(u)

}
