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
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day04a"
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day04b"
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day05a"
	"github.com/ga-paul-t/advent-of-code-2021/internal/aoc/day05b"
	"github.com/spf13/cobra"
)

const (
	benchNum = 200
)

// slice [][]aoc.Runner
// puzzle - 1 when selecting

var (
	puzzles = [][]aoc.Runner{
		{day01a.Puzzle{}, day01b.Puzzle{}},
		{day02a.Puzzle{}, day02b.Puzzle{}},
		{day03a.Puzzle{}, day03b.Puzzle{}},
		{day04a.Puzzle{}, day04b.Puzzle{}},
		{day05a.Puzzle{}, day05b.Puzzle{}},
	}
)

type options struct {
	Bench  bool
	Puzzle int
	Times  int
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
				if opts.Puzzle > len(puzzles) {
					return errors.New("no puzzle found")
				}

				toRun = append(toRun, puzzles[opts.Puzzle-1]...)

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

	return cmd, nil
}

func runPuzzles(out io.Writer, pzs []aoc.Runner) {
	fmt.Fprintln(out, "🎄 Advent of Code 2021")
	fmt.Fprintln(out)

	// Capture all durations for displaying total elapsed time
	elapsed := make([]time.Duration, len(pzs))

	for _, pz := range pzs {
		now := time.Now()
		res := pz.Run()
		dur := time.Since(now)
		fmt.Fprintf(out, "🧩 Puzzle %s: %-14d [time taken: %s]\n", pz, res, dur)

		elapsed = append(elapsed, dur)
	}

	fmt.Fprintln(out)
	fmt.Fprintf(out, "Total elapsed time: [%s]\n", sumElapsedTime(elapsed))
}

func sumElapsedTime(dur []time.Duration) time.Duration {
	var t time.Duration
	for _, d := range dur {
		t += d
	}

	return t
}

func benchmarkPuzzles(out io.Writer, pzs []aoc.Runner, opts options) {
	fmt.Fprintf(out, "🎄 Advent of Code 2021 - Benchmark [executions: %d]\n", opts.Times)
	fmt.Fprintln(out)

	// Capture all durations for displaying total elapsed time
	elapsed := make([]time.Duration, len(pzs))

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

		fmt.Fprintf(out, "🧩 Puzzle %s: %-14d [time taken: %s]\n", pz, res, exec[0])
		elapsed = append(elapsed, exec[0])
	}

	fmt.Fprintln(out)
	fmt.Fprintf(out, "Total elapsed time: [%s]\n", sumElapsedTime(elapsed))
}
