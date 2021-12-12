package main

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"time"

	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc"
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day01a"
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day01b"
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day02a"
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day02b"
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03a"
	"github.com/spf13/cobra"
)

const (
	benchNum = 200
)

var (
	puzzles = map[int][]aoc.Runner{
		1: {day01a.Puzzle{}, day01b.Puzzle{}},
		2: {day02a.Puzzle{}, day02b.Puzzle{}},
		3: {day03a.Puzzle{}},
	}

	times int
)

func newRootCmd(args []string, out io.Writer) (*cobra.Command, error) {
	var bench bool
	var puzzle int

	cmd := &cobra.Command{
		Use:          "aoc-2021",
		Short:        "Advent of Code 2021 Solutions",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			toRun := make([]aoc.Runner, 0, len(puzzles)*2)
			if puzzle > 0 {
				pz, ok := puzzles[puzzle]
				if ok {
					toRun = append(toRun, pz...)
				} else {
					return errors.New("no puzzle found")
				}
			} else {
				for _, pz := range puzzles {
					toRun = append(toRun, pz...)
				}
			}

			if bench {
				benchmarkPuzzles(out, toRun)
			} else {
				runPuzzles(out, toRun)
			}

			return nil
		},
	}
	f := cmd.Flags()
	f.BoolVarP(&bench, "benchmark", "b", false, "run in benchmarking mode")
	f.IntVarP(&times, "times", "t", benchNum, "number of puzzle executions during benchmark")
	f.IntVarP(&puzzle, "puzzle", "p", 0, "run a specific puzzle only")

	return cmd, nil
}

func runPuzzles(out io.Writer, pzs []aoc.Runner) {
	fmt.Fprintln(out, "🎄 Advent of Code 2021")
	fmt.Fprintln(out)

	for _, pz := range pzs {
		now := time.Now()
		res := pz.Run()
		dur := time.Since(now)
		fmt.Fprintf(out, "🧩 Puzzle %s: %d [time taken: %s]\n", pz, res, dur)
	}
}

func benchmarkPuzzles(out io.Writer, pzs []aoc.Runner) {
	fmt.Fprintf(out, "🎄 Advent of Code 2021 - Benchmark [executions: %d]\n", times)
	fmt.Fprintln(out)

	exec := make([]time.Duration, times)

	for _, pz := range pzs {
		res := 0

		// Run each puzzle the required benchmark sample rate, collecting each duration
		for b := 0; b < times; b++ {
			now := time.Now()
			res = pz.Run()
			exec[b] = time.Since(now)
		}
		sort.Slice(exec, func(i, j int) bool {
			return exec[i] < exec[j]
		})

		fmt.Fprintf(out, "🧩 Puzzle %s: %d [time taken: %s]\n", pz, res, exec[0])
	}
}
