package udon

import "time"

// GoにはJavaのオーバーロードやPythonのキーワード引数がない
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
func New(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

// 海老天2本入りの大盛りうどん
var tempuraUdon = New(Regular, false, 2)

// 次のコードは、よく利用されるバリエーションを関数として提供しています。
// ただしバリエーションの組み合わせが爆発するのを抑えるため、量に関するオプションは引数に残している。
func NewKake(p Portion) *Udon {
	return &Udon{
		men:      p,
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

type Option struct {
	Men      Portion
	Aburaage bool
	Ebiten   uint
}

func NewOption(opt Option) *Udon {
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
