package progressbar

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"golang.org/x/term"
)

type ProgressBar struct {
	total        int
	current      int
	customString string
	faint        func(a ...interface{}) string
	hiRed        func(a ...interface{}) string
}

// NewProgressBar initializes a new progress bar.
func NewProgressBar(total int, customString string) *ProgressBar {
	if total <= 0 {
		total = 50 // default total length if invalid input
	}

	faint := color.New(color.Faint).SprintFunc()
	hiRed := color.RGB(249, 38, 114).SprintFunc()

	return &ProgressBar{
		total:        total,
		customString: customString,
		faint:        faint,
		hiRed:        hiRed,
	}
}

// display renders the progress bar to the terminal.
func (pb *ProgressBar) display() {
	// Get terminal width
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		width = 80 // Default to 80 columns if unable to get terminal size
	}

	// Reserve space for custom string and padding
	progressBarWidth := width - len(pb.customString) - 10 // 10 for percentage and extra padding
	if progressBarWidth < 20 {
		progressBarWidth = 20 // Ensure a minimum width for the progress bar
	}

	// Calculate filled and unfilled parts
	filledLength := pb.current * progressBarWidth / pb.total
	progressBar := ""
	for i := 0; i < progressBarWidth; i++ {
		if i < filledLength {
			progressBar += pb.hiRed("━")
		} else {
			progressBar += pb.faint("━")
		}
	}

	// Calculate raw percentage
	percentage := float64(pb.current) / float64(pb.total) * 100

	// Print the progress bar with the custom string
	fmt.Printf("\r%s %s %.2f%%", pb.customString, progressBar, percentage)
}

// UpdateProgress updates the progress bar's state and redraws it.
func (pb *ProgressBar) UpdateProgress(percentage float64, customString ...string) {
	if percentage < 0 || percentage > 100 {
		return // Invalid percentage
	}

	// Update progress bar state
	pb.current = int(float64(pb.total) * percentage / 100)

	// Update custom string if provided
	if len(customString) > 0 {
		pb.customString = customString[0]
	}

	// Redraw the progress bar
	pb.display()
}

// Finish clears the progress bar and moves to a new line.
func (pb *ProgressBar) Finish() {
	fmt.Println() // Move to the next line
}
