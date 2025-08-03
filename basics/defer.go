package basics

import (
	"fmt"
	"io"
	"os"
)

func DemoDefer() {
	fmt.Println("hello world")

	// プログラムの中ではとある処理の後に後片付けが必要な場合がある。
	// Pythonのwith句のようにさまざまな言語が「後処理を実行」のメカニズムを用意している。
	// Golangも準備と同じ場所で予約するのが1番確実だ。という原則は変わらない。距離が離れると、整合が取れていない場合の見落としが増える。

	// ファイルを開く
	f, err := os.Create("sample.txt")
	if err != nil {
		fmt.Println("err", err)
		return
	}
	// この関数のスコープを抜けたら自動でファイルをクローズ
	defer f.Close()
	io.WriteString(f, "hello world")

	// エラー処理におけるGolangの慣習は次のとおり
	// 失敗する可能性のある関数の末尾をerror型とする
	// 成功時にはnilを、失敗時にはそこに詳細なエラーを割り当てて返す
	// 関数を呼び出した後は、if err != nilというif文でエラーをチェックし、追加の情報をラップしたり、そのままreturnで呼び出し元に返す。
	// エラーハンドリングの詳細は5章で詳しく説明する
}
