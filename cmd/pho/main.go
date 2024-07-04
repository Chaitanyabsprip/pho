package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/Chaitanyabsprip/pho/internal/display"
	"github.com/Chaitanyabsprip/pho/internal/pho"
)

func run(ctx context.Context, args []string) (int, error) {
	_ = ctx
	config, err := ParseArgs(args)
	if err != nil {
		return 1, err
	}
	analysis, err := pho.Analyze(config.Filepaths, config.Recurse)
	if err != nil {
		return 1, err
	}
	plotables := make([]display.Plotable, 0)
	for i := range analysis {
		plotables = append(plotables, analysis[i])
	}
	display.Display(config.DisplayMethod, plotables)
	open(fmt.Sprint("./", config.DisplayMethod, ".html"))
	return 0, nil
}

func main() {
	ctx := context.Background()
	if code, err := run(ctx, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(code)
	}
}

func open(path string) {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", path}
	case "windows":
		args = []string{"cmd", "/c", "start", path}
	default:
		args = []string{"xdg-open", path}
	}
	cmd := exec.Command(args[0], args[1:]...)
	err := cmd.Run()
	if err != nil {
		log.Printf("openinbrowser: %v %v\n", err, path)
	}
}
