package basics

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// いくつかのデータをまとめて塊としてあつけるのが構造体
// 構造体も複合型
// ブレースの中にフィールド（メンバーとなる変数）を列挙する
type Book struct {
	// 先頭を大文字にしなければ外部のパッケージから利用できない
	// 中に関数を定義できない
	// jsonタグを定義しておくとこの定義に従って、構造体のフィールドをJSONに書き出したり、JSONの情報をフィールドにマッピングできる。
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	ReleasedAt time.Time `json:"released_at"`
	ISBN       string    `json:"isbn"`
}

// この構造体は型であり、メモリ上に存在するものではない。
// メモリ上に「インスタンス」を作らなければデータを保存したりできない。

func DemoStruct() {

	// 構造体の利用(フィールドはすべてゼロ値に初期化)
	var b Book
	fmt.Println(b)

	// フィールドを初期化しながらインスタンス作成
	b2 := Book{
		Title: "Twisted Network Programing Essentials",
	}
	fmt.Println(b2)

	// フィールドを初期化しながらインスタンス作成
	// 変数にはポインタを格納
	b3 := &Book{
		Title: "Learn to Golang",
	}
	fmt.Println(b3)

	f, err := os.Open("book.json")
	if err != nil {
		log.Fatal("file open error:", err)
	}
	d := json.NewDecoder(f)
	d.Decode(&b)
	fmt.Println(b)

	// GoはJavaと異なり、リフレクションを使って動的に拡張することを多用することは稀である。
	// 数少ない用途が、このタグを使ったマッピングである。ウェブブラウザからのリクエストや、JSONなどの構造化ファイル、データベースとのマッピングなどに活用する。
	// リフレクションを使って動的にstructの情報を取得する
}
