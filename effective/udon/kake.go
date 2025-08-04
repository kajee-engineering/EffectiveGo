package udon

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
func NewUdon(p Portion, aburaage bool, ebiten uint) *Udon {
	return &Udon{
		men:      p,
		aburaage: aburaage,
		ebiten:   ebiten,
	}
}

// 海老天2本入りの大盛りうどん
var tempuraUdon = NewUdon(Regular, false, 2)

// 次のコードは、よく利用されるバリエーションを関数として提供しています。
// ただしバリエーションの組み合わせが爆発するのを抑えるため、量に関するオプションは引数に残している。
func NewKakeUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   0,
	}
}

func NewKitsuneUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: true,
		ebiten:   0,
	}
}

func NewTempuraUdon(p Portion) *Udon {
	return &Udon{
		men:      p,
		aburaage: false,
		ebiten:   3,
	}
}

// これ以外はよく使う用途別にラッパー関数を複数用意する方法がある。例えばos.Create()がこれに該当する。
