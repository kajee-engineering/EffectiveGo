package effective

import (
	"bufio"
	"os"
)

func New(text string) error {
	return &errorString{text}
}

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

// 一般的なエラーハンドリングは戻り値の最後にerrorを定義する
func ReadFile(name string) ([]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		// エラーハンドリング
		return nil, err // 戻り値のerror型の値はerr
	}
	r := bufio.NewReader(f)
	l, err := r.ReadString('\n')
	if err != nil {
		f.Close()
	}
	// 一般的にガード節と言われる手法で例外を早めに処理することで見通しが良くなる

	// ここではエラーは発生していない
	return []byte(l[:len(l)-1]), nil // エラーが発生していないので戻り値であるerror型の値はnil
}
