package aoc

import "fmt"

type Runner interface {
	fmt.Stringer
	Run() int
}
