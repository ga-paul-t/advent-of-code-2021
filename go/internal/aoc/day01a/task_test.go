package day01a

import "testing"

func Benchmark_Run(b *testing.B) {
	p := Puzzle{}

	for i := 0; i < b.N; i++ {
		p.Run()
	}
}
