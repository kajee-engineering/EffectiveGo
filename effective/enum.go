package effective

// 文字列として出力可能にする
//
//go:generate stringer -type=CarType
type CarType int

// 最もシンプルな列挙型相当の定数
const (
	Sedan CarType = iota + 1 // iotaの初期値は0なので、初期値と最初の要素の0を区別するために+1するのが一般的
	HatchBack
	MPV
	SUV
	Crossover
	Coupe
	Convertible
)

type CarOption uint64

const (
	GPS CarOption = 1 << iota
	AWD
	SunRoof
	HeatedSeat
	DriverAssist
)
