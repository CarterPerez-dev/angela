// ©AngelaMos | 2026
// spinner.go

package ui

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var frames = []string{
	"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏",
}

// Spinner displays an animated braille spinner in the terminal.
type Spinner struct {
	msg    string
	done   chan struct{}
	wg     sync.WaitGroup
	mu     sync.Mutex
	active bool
}

// NewSpinner creates a spinner with the given message.
func NewSpinner(msg string) *Spinner {
	return &Spinner{msg: msg}
}

// Start begins the spinner animation in a goroutine.
func (s *Spinner) Start() {
	s.mu.Lock()
	if s.active {
		s.mu.Unlock()
		return
	}
	s.active = true
	s.done = make(chan struct{})
	s.wg.Add(1)
	s.mu.Unlock()

	go s.run()
}

// Stop halts the spinner and clears the line.
func (s *Spinner) Stop() {
	s.mu.Lock()
	if !s.active {
		s.mu.Unlock()
		return
	}
	s.active = false
	close(s.done)
	s.mu.Unlock()
	s.wg.Wait()
}

func (s *Spinner) run() {
	defer s.wg.Done()
	fmt.Print("\033[?25l")

	ticker := time.NewTicker(80 * time.Millisecond)
	defer ticker.Stop()

	idx := 0
	for {
		select {
		case <-s.done:
			clearLine()
			fmt.Print("\033[?25h")
			return
		case <-ticker.C:
			frame := frames[idx%len(frames)]
			fmt.Printf(
				"\r  %s %s",
				CyanBold(frame),
				HiMagenta(s.msg),
			)
			idx++
		}
	}
}

func clearLine() {
	fmt.Print("\r" + strings.Repeat(" ", 80) + "\r")
}
