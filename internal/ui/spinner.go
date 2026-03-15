package ui

import (
	"time"

	"github.com/briandowns/spinner"
)

type Spinner struct {
	s *spinner.Spinner
}

func New(suffix string) *Spinner {
	s := spinner.New(spinner.CharSets[14], 120*time.Millisecond)
	s.Suffix = " " + suffix

	return &Spinner{s: s}
}

func (s *Spinner) Start() {
	s.s.Start()
}

func (s *Spinner) Stop() {
	s.s.Stop()
}
