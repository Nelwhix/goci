package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func run(project string, out io.Writer) error {
	if project == "" {
		return fmt.Errorf("Project directory is required: %w", ErrValidation)
	}

	args := []string{"build", ".", "errors"}
	cmd := exec.Command("go", args...)
	cmd.Dir = project

	if err := cmd.Run(); err != nil {
		return &stepErr{step: "go build", msg: "go build failed", cause: err}
	}

	_, err := fmt.Fprintln(out, "Go Build: SUCCESS")
	return err
}

func main() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "goci, a continuous integration and deployment tool by Nelson Isioma")
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	proj := flag.String("p", "", "Project directory")
	flag.Parse()

	if err := run(*proj, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	
}