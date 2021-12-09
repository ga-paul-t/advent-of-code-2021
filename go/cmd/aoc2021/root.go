package main

import (
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc"
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day01a"
	"github.com/spf13/cobra"
)

const (
	benchNum = 100
)

var (
	puzzles = []aoc.Runner{
		day01a.Puzzle{},
	}
)

func newRootCmd(args []string, out io.Writer) (*cobra.Command, error) {
	var bench bool

	cmd := &cobra.Command{
		Use:          "aoc-2021",
		Short:        "Advent of Code 2021 Solutions",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			if bench {
				benchmarkPuzzles(out)
			} else {
				runPuzzles(out)
			}
		},
	}

	f := cmd.Flags()
	f.BoolVarP(&bench, "benchmark", "b", false, "run in benchmarking mode")

	return cmd, nil
}

func runPuzzles(out io.Writer) {
	fmt.Fprintln(out, "🎄 Advent of Code 2021")
	fmt.Fprintln(out)

	for _, puz := range puzzles {
		now := time.Now()
		res := puz.Run()
		dur := time.Since(now)
		fmt.Fprintf(out, "🧩 Puzzle %s: %d [time taken: %s]\n", puz, res, dur)
	}
}

func benchmarkPuzzles(out io.Writer) {
	fmt.Fprintln(out, "🎄 Advent of Code 2021 - Benchmark")
	fmt.Fprintln(out)

	btimes := make([]time.Duration, benchNum)

	for _, puz := range puzzles {
		res := 0

		// Run each puzzle the required benchmark sample rate, collecting each duration
		for b := 0; b < benchNum; b++ {
			now := time.Now()
			res = puz.Run()
			btimes[b] = time.Since(now)
		}
		sort.Slice(btimes, func(i, j int) bool {
			return btimes[i] < btimes[j]
		})

		fmt.Fprintf(out, "🧩 Puzzle %s: %d [time taken: %s]\n", puz, res, btimes[0])
	}
}
