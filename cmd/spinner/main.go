package main

import (
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
)

func main() {
	newSpinnerWithColor()
}

func newSpinner() {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Start()
	time.Sleep(4 * time.Second) // Run for some time to simulate work
	s.Stop()
}

func newSpinnerWithColor() {
	suffixColor := color.New(color.Bold, color.FgGreen)
	sp := spinner.New(
		spinner.CharSets[14],
		time.Second*2,
		spinner.WithColor("green"),
		spinner.WithHiddenCursor(true),
		spinner.WithSuffix(suffixColor.Sprintf(" %s", "console print")))
	sp.Start()
	time.Sleep(time.Second * 10)
	sp.Stop()
}
