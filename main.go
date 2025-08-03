package main

import (
	// "MyGolang/basics"
	"MyGolang/effective"
	"fmt"
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
}
