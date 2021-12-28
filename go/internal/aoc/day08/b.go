package day08

import (
	"math"
	"strings"
)

type PuzzleB struct{}

func (p PuzzleB) String() string {
	return "08b"
}

func (p PuzzleB) Run() int {
	lines := strings.Split(input, "\n")

	t := 0
	for _, line := range lines {
		parts := strings.Split(line, " | ")

		sigs := strings.Fields(parts[0])
		digs := strings.Fields(parts[1])

		// Digits 1, 4, 7 and 8 have a unique number of segments. All other digits have either
		// 5 [2, 3 and 5] or 6 [0, 6, and 9] segments. All digits based on the illustration contain
		// parts of either digit 1 or 4. By generating a lookup table, it should be possibly to decode
		// the other digits

		// This is probably not very efficient, but we need to traverse the signals to identify digits
		// 1 and 4, if both exist. Then traverse the digits to decode them
		one, four := "", ""
		for i := range sigs {
			switch len(sigs[i]) {
			case 2:
				one = sigs[i]
			case 4:
				four = sigs[i]
			}
		}

		// Utility function to count the number of segments from digits 1 and 4 the exist
		// in the given digit. Return a tuple for pattern matching
		countSegments := func(in string) (int, int) {
			oc, fc := 0, 0
			for i := range in {
				oc += strings.Count(one, string(in[i]))
				fc += strings.Count(four, string(in[i]))
			}

			return oc, fc
		}

		num := [4]int{}
		for i := range digs {
			switch len(digs[i]) {
			case 2:
				num[i] = 1
			case 3:
				num[i] = 7
			case 4:
				num[i] = 4
			case 7:
				num[i] = 8
			case 5:
				oc, fc := countSegments(digs[i])

				// Mappings:
				// 3 > 4: [3 parts] - 1: [2 parts]
				// 5 > 4: [3 parts] - 1: [*]
				// 2 > 4: [*]       - 1: [*]

				if fc == 3 {
					if oc == 2 {
						num[i] = 3
					} else {
						num[i] = 5
					}
				} else {
					num[i] = 2
				}
			case 6:
				oc, fc := countSegments(digs[i])

				// Mappings:
				// 9 :> 4: [4 parts] - 1: [*]
				// 0 :> 4: [3 parts] - 1: [2 parts]
				// 6 :> 4: [3 parts] - 1: [*]

				if fc == 4 {
					num[i] = 9
				} else if fc == 3 {
					if oc == 2 {
						num[i] = 0
					} else {
						num[i] = 6
					}
				}
			}
		}

		// Add decoded number to total ([0]thousands,[1]hundreds,[2]tens,[3]units)
		for i := range num {
			t += num[i] * int(math.Pow10(3-i))
		}
	}

	return t
}
