package basics // 同じディレクトリ内のパッケージ名はすべて同じである必要がある

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog/log"
)

// 固定長の文字列を返すだけのシンプルなWebサービス
func DemoWebApp() {
	// ①
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
		log.Info().Msg("receive hello world request")
	})

	// ②
	fmt.Println("Listening on port 8080")

	// ③
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// ④
		fmt.Fprintf(os.Stderr, "")
		io.WriteString(os.Stderr, err.Error())
		os.Exit(1)
	}

	// ① まず、URLごとの処理を登録する。ここでは/heeloというパスに文字列を返す処理を登録している。
	// ② Webサービスがどこで起動しているかわからないとテストするのに不便なので画面にテキストを表示する。
	// ③ ここではhttpパッケージのListenAndServe()関数を使ってサーバを起動する。
	// ④ もし、オープンしようとしているポートが占有されている場合は即座にエラーが返ってきてくる。このエラーを表示してプログラムをクローズさせる、

	// http://localhost:8080/helloにアクセスするとテキストが表示される。

	// これにJSON形式のリクエストやレスポンスを読み書きしたり、データベース機能が入ったり、リクエストから情報を取って書き込んだりすれば、実用的なアプリケーションになる。

	// まとめ
	// シニアなメンバーが開発したアプリケーションに、多少の機能追加を行うのに必要な知識に絞って紹介してきた
	// (終わり)
}
