package main

import (
	// "MyGolang/basics"
	"MyGolang/effective"
	"MyGolang/effective/udon"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"strings"
	"time"
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
	fmt.Println("---------")
	flag.Parse()
	log.Println(*FlagStr) // コンソールの末尾に表示される
	log.Println(*FlagInt)
	log.Println(flag.Args())
	fmt.Println("---------")

	//
	// 1.8 メモリ起因のパフォーマンス低下を解消する
	//

	// Goでパフォーマンスに差が現れがちなポイントはスライスとマップである
	// スライスやマップのメモリ確保を高速化するトピックを紹介する

	// 使用するメモリの最大値がわかっている場合は最初から最大値を指定しましょう
	// 正確な長さがわかっている場合
	s1 := make([]int, 1000) // 最初に実際のサイズまで一緒に確保してしまう
	fmt.Println(len(s1))
	fmt.Println(cap(s1))

	// 正確な長さはわからないがキャパシティだけ増やす場合
	// メモリの再割り当てとコピー回数が減らせる
	s2 := make([]int, 0, 1000)
	fmt.Println(len(s2))
	fmt.Println(cap(s2))

	// 1.9 文字列の結合方法
	src := []string{"Black", "To", "The", "Future", "Part", "III"}
	var title string
	for i, word := range src {
		if i != 0 {
			title += " "
		}
		title += word
	}
	log.Println(title)
	// Goの文字列は不変な文字列のため文字列を追加する、または一部を取り出す場合は新しい文字列が生成されてメモリが確保される。
	// ループで1つ1つ結合すると、中間の文字列が全てメモリ上に一度は確保されるので、速度が遅くなる。

	// 大量に文字列を結合するときはstrings.Builderを使うのが良い
	// 結合後のサイズがわかっている場合はGrowを使うこともできる
	var builder strings.Builder
	builder.Grow(100) // 最大100文字以下と仮定できる場合
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	log.Println(builder.String())

	// 1.10 日時の取り扱い
	// Time, Date, DateTimeが別々で提供されている言語もあるが、Goは1つの型で全て扱う
	// 現在時刻のtime.Timeインスタンス取得
	now := time.Now()

	// 指定日時のtime.Timeインスタンス取得
	tz, _ := time.LoadLocation("America/Los_Angeles")
	future := time.Date(2015, time.October, 21, 7, 28, 0, 0, tz)

	fmt.Println(now.String())
	fmt.Println(future.Format(time.RFC3339Nano))

	// 定義済みのローカルタイムゾーン
	now = time.Date(1985, time.October, 26, 9, 0, 0, 0, tz)
	fmt.Println(now.String())

	// 定義済みのUTCローカルタイムゾーン
	past := time.Date(1955, time.November, 12, 6, 38, 0, 0, tz)
	fmt.Println(past.String())
	// タイムゾーンの情報には夏時間の開始日時など、人が年毎に決めている情報も含まれるため、OSの更新をおこなった上でOSの情報を参照するのがベスト
	// Go1.15からはタイムゾーン情報をアプリケーションにバンドルできるようになったので、アプリケーションをこまめに最新の処理系でビルドし直せる場合はこの方法でもよい。
	// _ "time/tzdata"

	// 日時とは別に時間を表す型はtime.Durationで提供されている
	// time.Durationの作成方法はいくつかある
	// 例えばtime.Time同士の差をSub()メソッドで計算
	// 例えばtime.Secondなどの既存の時間のインスタンスの積により作成

	// 5分を作成
	// Nanosecond, Millisecond, Second, Minute, Hourが定義済み
	fiveMinutes := 5 * time.Minute
	fmt.Println(fiveMinutes) // 5m0s

	// intとは形違いで直接演算できないので、即値との計算以外はtime.Durationへの明示的なキャストが必要
	// キャストがないと次のエラーが発生する
	// invalid operation: secondes * time.Second (mismatched types int and time.Duration)
	var secondes int = 10
	tenSeconds := time.Duration(secondes) * time.Second
	//tenSeconds := secondes * time.Second // キャストしない
	fmt.Println(tenSeconds)

	// Timeの演算でDuration作成
	past = time.Date(1955, time.November, 12, 6, 38, 0, 0, time.UTC)
	dur := time.Now().Sub(past)
	fmt.Println(dur)

	// Truncate()メソッドを用いて5分単位で切り詰めることもできる
	// 1時間にまとめてバッチで読み込むファイル名を取得する場合
	filepath := time.Now().Truncate(time.Hour).Format("20060102150405.json")
	fmt.Println(filepath)
	// 5分後と5分前の時刻
	fiveMinutesAfter := time.Now().Add(fiveMinutes)
	fmt.Println(fiveMinutesAfter)
	fiveMinutesBefore := time.Now().Add(-fiveMinutes) // time.Time型はAdd()かAddDate()しか持っていないのでマイナスの値をAdd()に渡す
	fmt.Println(fiveMinutesBefore)

	// 3秒停止
	fmt.Println("3秒スリープスタート")
	time.Sleep(3 * time.Second) // 現在のゴルーチンを停止する
	fmt.Println("3秒スリープ完了")

	// 10秒間待つ
	// selectを使って他の処理の完了待ちをするのに便利 @see 378ページ
	fmt.Println("10秒停止スタート")
	timer := time.NewTimer(10 * time.Second) // Timerのポインタを受け取る
	defer timer.Stop()                       // Timerのリソースを確実に解放するためにdeferを使う
	<-timer.C                                // タイマーのチャネルから通知を受け取る
	fmt.Println("10秒停止完了")

	// 定義型のレシーバにアプリケーション固有のドメインロジックを宣言することで、関心事を分離したり、凝集度を高めることができる。
	// これはコードの見通しがよくなることに加えて、ユニットテストが書きやすくなる効果もある。
	c1 := consumer{"taro", true}
	c2 := consumer{"suzuki", false}
	cs := consumers{c1, c2}
	fmt.Println(cs.activeConsumer()) // 例えばレスポンスの顧客一覧から有効な顧客を抽出する問い合わせをレシーバに移譲できる。
	// 今回であれば、cunsumers型を定義して、ロジックをレシーバとして実装する方法がおすすめ。
	// この時、戻り値もcunsmers型にする実装がおすすめ。なぜならチェーンメソッドで複数のコレクション操作を記述できるため。
	// consumers := cs.activeConsumer().expires(time.Now()).SortedByExpiredAt() // 有効なユーザから期限が切れた一覧をソートして受け取る

	// 2.3.2 値への型定義
	// SKU(Stock Keeping Unit, 在庫管理単位のこと)の例にする
	// 具体的にはT01230101は最初の5桁が商品コード(T0123)、次の2桁がサイズ、次の2桁がカラーを示す
	// これを素直に実装すると低レイヤーのコードが入り込むことになる
	// skuCD, _ := r.URL.Query()["skuCode"]
	// itemCD, sizeCD, colerCD := skuCD[0:5], skuCD[5:7], skuCD[7:9]
	// 上記のコードを呼び出し元で整合性をチェックする必要が出てくるため少々厄介

	// 2.3.3 列挙への型定義
	// 例: シーズンに対応した受発注サービス

	// 2.5 機密情報を扱うフィールドを定義して出力書式をカスタマイズ
	c := ConfidentialCustomer{
		CustomerID: 1,
		CreditCard: "4111-1111-1111-1111",
	}
	fmt.Println(c)
	fmt.Printf("%v\n", c)
	fmt.Printf("%+v\n", c)
	fmt.Printf("%#v\n", c)

	bytes, _ := json.Marshal(c)
	fmt.Println("JSON: ", string(bytes))

}

// 以下の2つのinterfaceを拡張することで、例えばクレジットカードのように注意深く扱う情報のログ出力をマスキングできる
// type Stringer interface { // fmtモジュールのPrintなどで、値を出力する際に呼び出される。
//
//		String() string
//	}
//
// type GoStringer interface { // fmtモジュールのPrintなどで、値を出力する際に呼び出される。
//
//		GoString() string
//	}
type ConfidentialCustomer struct {
	CustomerID int64
	CreditCard CreditCard
}

type CreditCard string // ユーザ定義型

func (c CreditCard) String() string { // ユーザ定義型(CreditCard)にString()インターフェースを実装して、
	return "xxxx-xxx-xxxx-xxxx" // 機密情報をマスキングしている。
}
func (c CreditCard) GoString() string {
	return "xxxx-xxx-xxxx-xxxx"
}

type season int

const (
	peak   season = iota + 1 // 繁忙期
	normal                   // 通常
	off                      // 閑散期
)

// 例えば季節変動が理由で料金が異なる場合の要件はよく出てくる
func (s season) price(price float64) season {
	if s == peak {
		return s + 200
	}
	return s
}

// 少し面倒だが、次のように実装しておくと利用方法を一目で理解しやすくなる
type SKUcode string

func (c SKUcode) Invalid() bool {
	// 桁数や利用可能文字のチェックを行う
	return true
}

func (c SKUcode) ItemCD() string {
	return string(c)[0:5]
}

func (c SKUcode) SizeCD() string {
	return string(c)[5:7]
}

func (c SKUcode) ColerCD() string {
	return string(c)[7:9]
}

var (
	// コマンドライン引数を定義
	FlagStr = flag.String("string", "default", "文字列フラグ")
	FlagInt = flag.Int("int", -1, "数値フラグ")
)

// 場合によっては上記のチェーンメソッドを関数として宣言して、実装をカプセル化できる。
func (c consumers) RequiredFollows() consumers {
	return c.activeConsumer().expires(time.Now()).SortedByExpiredAt()
}

func (c consumers) expires(t time.Time) consumers {
	// 引数tで有効期限が失効するユーザに絞り込む処理を記述する
	return c
}

func (c consumers) SortedByExpiredAt() consumers {
	// 有効期限をキーに昇順でソートする
	return c
}

// 2.3.1 スライスへの型定義
type consumers []consumer

func (c consumers) activeConsumer() consumers {
	resp := make([]consumer, 0, len(c))
	for _, v := range c {
		if v.activeFlg() {
			resp = append(resp, v)
		}
	}
	return resp
}

type consumer struct {
	name     string
	isActive bool
}

func (c *consumer) activeFlg() bool {
	return c.isActive
}
