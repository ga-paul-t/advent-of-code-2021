package aoc

import "io/ioutil"

func ReadInputFile() string {
	data, _ := ioutil.ReadFile("input.txt")
	return string(data)
}
