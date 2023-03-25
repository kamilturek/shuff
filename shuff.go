package shuff

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const version = "0.2.0"

func RunCLI() {
	var showHelp bool
	var showVersion bool

	flag.BoolVar(&showHelp, "help", showHelp, "show help message")
	flag.BoolVar(&showHelp, "h", showHelp, "show help message (shorthand)")
	flag.BoolVar(&showVersion, "version", showVersion, "show version")
	flag.BoolVar(&showVersion, "v", showVersion, "show version (shorthand)")
	flag.Parse()

	if showHelp {
		fmt.Fprintf(os.Stderr, "Usage: %s [FILE]\n", os.Args[0])
		os.Exit(0)
	}

	if showVersion {
		fmt.Fprintln(os.Stderr, version)
		os.Exit(0)
	}

	s, err := NewShuffler(
		WithInputFromArgs(os.Args[1:]),
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := s.Shuffle(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

type shuffler struct {
	input  io.Reader
	output io.Writer
	seed   int64
}

type option func(*shuffler) error

func NewShuffler(opts ...option) (*shuffler, error) {
	s := &shuffler{
		input:  os.Stdin,
		output: os.Stdout,
		seed:   time.Now().UnixNano(),
	}

	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func WithInput(input io.Reader) option {
	return func(s *shuffler) error {
		if input == nil {
			return errors.New("nil input reader")
		}

		s.input = input

		return nil
	}
}

func WithInputFromArgs(args []string) option {
	return func(s *shuffler) error {
		if len(args) == 0 {
			return nil
		}

		f, err := os.Open(args[0])
		if err != nil {
			return err
		}

		s.input = f

		return nil
	}
}

func WithOutput(output io.Writer) option {
	return func(s *shuffler) error {
		if output == nil {
			return errors.New("nil output writer")
		}

		s.output = output

		return nil
	}
}

func WithSeed(seed int64) option {
	return func(s *shuffler) error {
		s.seed = seed

		return nil
	}
}

func (s *shuffler) Shuffle() error {
	scanner := bufio.NewScanner(s.input)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	rand.Seed(s.seed)
	rand.Shuffle(len(lines), func(i, j int) {
		lines[i], lines[j] = lines[j], lines[i]
	})

	w := bufio.NewWriter(s.output)

	for _, line := range lines {
		w.WriteString(line + "\n")
	}

	w.Flush()

	return nil
}
