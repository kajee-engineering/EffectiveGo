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
type Book struct { // 構造体はメモリ上に存在しないデータのレイアウト定義
	// 先頭を大文字にしなければ外部のパッケージから利用できない
	// 中に関数を定義できない
	// jsonタグを定義しておくとこの定義に従って、構造体のフィールドをJSONに書き出したり、JSONの情報をフィールドにマッピングできる。
	Title      string    `json:"title"`
	Author     string    `json:"author"`
	Publisher  string    `json:"publisher"`
	ReleasedAt time.Time `json:"released_at"` // 他(Time)の構造体も埋め込むことができる。しかしアップキャスト、ダウンキャストができないので継承ではない。
	ISBN       string    `json:"isbn"`
	// 他の言語ではフィールドはプライベートにして読み書きはメソッド経由で行う設計があるが、Goの場合は外部から操作されると整合性が壊れてしまうもの(銀行の残高とか?)以外は禁止しない
}

// この構造体は型であり、メモリ上に存在するものではない。
// メモリ上に「インスタンス」を作らなければデータを保存したりできない。

func DemoStruct() {

	// 構造体の利用(フィールドはすべてゼロ値に初期化)
	var b Book // インスタンスを作成。インスタンスはフィールドのデータを一括で保存するメモリ領域である。オブジェクトとも呼ばれる
	fmt.Println(b)

	// フィールドを初期化しながらインスタンス作成
	b2 := Book{
		Title: "Twisted Network Programing Essentials",
	}
	fmt.Println(b2.Title) // 要素へのアクセスはピリオドを使う

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

	// 関数の中でしか使わない構造体を関数の中で定義することも可能
	type Person struct {
		Name struct { // 構造体の中に構造体を定義することも可能
			First string
			Last  string
		}
		Age int
	}

	// インスタンス化が冗長になる
	p := Person{
		Name: struct {
			First string
			Last  string
		}{First: "suzuki", Last: "taro"},
		Age: 42,
	}
	// NameとPerson分割して定義する。

	fmt.Println(p)

}
