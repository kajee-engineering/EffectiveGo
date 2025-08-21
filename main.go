package main

import (
	// "MyGolang/basics"
	"MyGolang/effective"
	"MyGolang/effective/udon"
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

	// テストコミット
}

var (
	// コマンドライン引数を定義
	FlagStr = flag.String("string", "default", "文字列フラグ")
	FlagInt = flag.Int("int", -1, "数値フラグ")
)
