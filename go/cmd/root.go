package main

import (
	"fmt"
	"io"
	"time"

	"github.com/ga-paul-t/aoc2021/internal/aoc"
	"github.com/ga-paul-t/aoc2021/internal/aoc/day01a"
	"github.com/spf13/cobra"
)

var (
	puzzles = []aoc.Runner{
		day01a.Puzzle{},
	}
)

func newRootCmd(args []string, out io.Writer) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:          "aoc-2021",
		Short:        "Advent of Code 2021 Solutions",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			runPuzzles(out)
		},
	}

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
