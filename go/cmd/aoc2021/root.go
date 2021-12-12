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
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day03b"
	"github.com/spf13/cobra"
)

const (
	benchNum = 200
)

var (
	puzzles = map[int][]aoc.Runner{
		1: {day01a.Puzzle{}, day01b.Puzzle{}},
		2: {day02a.Puzzle{}, day02b.Puzzle{}},
		3: {day03a.Puzzle{}, day03b.Puzzle{}},
	}
)

type options struct {
	Bench  bool
	Puzzle int
	Times  int
	PartA  bool
	PartB  bool
}

func newRootCmd(args []string, out io.Writer) (*cobra.Command, error) {
	opts := options{}

	cmd := &cobra.Command{
		Use:          "aoc-2021",
		Short:        "Advent of Code 2021 Solutions",
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			toRun := make([]aoc.Runner, 0, len(puzzles)*2)
			if opts.Puzzle > 0 {
				pz, ok := puzzles[opts.Puzzle]
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

			if opts.Bench {
				benchmarkPuzzles(out, toRun, opts)
			} else {
				runPuzzles(out, toRun)
			}

			return nil
		},
	}
	f := cmd.Flags()
	f.BoolVarP(&opts.Bench, "benchmark", "b", false, "run in benchmarking mode")
	f.IntVarP(&opts.Times, "times", "t", benchNum, "number of puzzle executions during benchmark")
	f.IntVarP(&opts.Puzzle, "puzzle", "p", 0, "run a specific puzzle only")
	f.BoolVarP(&opts.PartA, "part-a", "a", false, "run part A of puzzles only")
	f.BoolVarP(&opts.PartB, "part-b", "b", false, "run part B of puzzles only")

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

func benchmarkPuzzles(out io.Writer, pzs []aoc.Runner, opts options) {
	fmt.Fprintf(out, "🎄 Advent of Code 2021 - Benchmark [executions: %d]\n", opts.Times)
	fmt.Fprintln(out)

	exec := make([]time.Duration, opts.Times)

	for _, pz := range pzs {
		res := 0

		// Run each puzzle the required benchmark sample rate, collecting each duration
		for b := 0; b < opts.Times; b++ {
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
