package underline

import (
	"fmt"
	"testing"
)

type test struct {
	title     string
	symbol    string
	with_text bool
}

var input = []test{
	test{"Solid / Цельное", "─", true},
	test{"Dashed / Пунктирное", "-", true},
	test{"Dotted / Из точек", ".", true},
	test{"Custom line style / Собственный символ", "+", true},
	test{"", "solid", true},
}

func Test_Underline(t *testing.T) {
	for _, d := range input {
		fmt.Println(Custom(d.title, d.symbol, d.with_text) + "\n")
	}
}

func Benchmark_Underline(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Custom("Custom line style / Собственный символ", "/", false)
	}
}
