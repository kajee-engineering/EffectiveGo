package udon

import (
	"fmt"
	"time"
)

// GoにはJavaのオーバーロードやPythonのキーワード引数といったオプション引数がない
// 以下は、うどん屋のオプションを題材にオプション引数を実現する方法

type Portion int

const (
	Regular Portion = iota
	Small
	Large
)

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

// めんの量、油揚げ、海老天の有無でインスタンス作成
func AllParam(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

// 上のコードは全てのパラメタを渡さなければならないが、これを以下に使いやすくするかが本章のゴール

// 海老天2本入りの大盛りうどん
var tempuraUdon = AllParam(Regular, false, 2)

// 次のコードは、よく利用されるバリエーションを関数として提供しています。
func New2th(p Portion) *Udon {
	return &Udon{
		men:      p, // ただしバリエーションの組み合わせが爆発するのを抑えるため、量に関するオプションは引数に残している
		aburaage: false,
		ebiten:   0,
	}
}

func NewKitsune(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: true,
		ebiten:   0,
	}
}

func NewTempura(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   3,
	}
}

// これ以外はよく使う用途別にラッパー関数を複数用意する方法がある。例えばos.Create()がこれに該当する。

// 構造体を利用したオプション引数
type Option struct {
	Men      Portion
	Aburaage bool
	Ebiten   uint
}

func New3th(opt Option) *Udon {
	// ゼロ値に対するデフォルト値処理は関数/メソッド内部で行う
	// 朝食時間は海老天1本無料
	if opt.Ebiten == 0 && time.Now().Hour() < 10 {
		opt.Ebiten = 1
	}
	return &Udon{
		men:      opt.Men,
		aburaage: opt.Aburaage,
		ebiten:   opt.Ebiten,
	}
}

// ビルダーを利用したオプション引数
type fluentOpt struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func Kake3th(p Portion) *fluentOpt {
	// デフォルトはコンストラクタ関数で設定
	// 必須オプションはここに付与可能
	return &fluentOpt{
		men:      p,
		aburaage: false,
		ebiten:   1,
	}
}

func (o *fluentOpt) Aburaage() *fluentOpt {
	o.aburaage = true
	return o
}

func (o *fluentOpt) Ebiten(n uint) *fluentOpt {
	o.ebiten = n
	return o
}

func (o *fluentOpt) Order() *Udon {
	return &Udon{
		men:      o.men,
		aburaage: o.aburaage,
		ebiten:   o.ebiten,
	}
}

func UseFluentInterfasce() {
	// 複数のオブジェクトの生成パターンを分離して段階的に構成するデザインパターン(=ビルダー)
	oomoriKitune := Kake3th(Large).Aburaage().Order()
	fmt.Println(oomoriKitune)
}

// Functional Optionパターンを使ったオプション引数
type OptFunc func(r *Udon) // 関数型を定義

func New4th(opts ...OptFunc) *Udon {
	r := &Udon{}
	for _, opt := range opts {
		opt(r) // ループ1回目は opt(r) == func(r *Udon) { r.men = p } を実行
	}
	return r
}

func OptMen(p Portion) OptFunc {
	return func(r *Udon) { r.men = p }
}

func OptAburaage() OptFunc {
	return func(r *Udon) { r.aburaage = true }
}

func OptEbiten(u uint) OptFunc {
	return func(r *Udon) { r.ebiten = u }
}

func UseFuncOption() {
	// 引数を評価してからNew4th()呼び出す
	// New4th(
	//        func(r *Udon) { r.men = p },
	//        func(r *Udon) { r.aburaage = true } ,
	//        func(r *Udon) { r.ebiten = u },
	//       )
	tokuseiUdon := New4th(
		OptMen(Large),
		OptAburaage(),
		OptEbiten(3),
	)
	fmt.Println(tokuseiUdon)
}
