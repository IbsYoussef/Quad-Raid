package quad_test

import (
	"bytes"
	"io"
	"os"
	"quad/quads"
	"testing"
)

// Helper function to capture output of what gets printer to stdout
func captureOutput(f func()) string {
	// Step 1: Save the original stdout
	oldStdout := os.Stdout

	// Step 2: Create a pipe (r = read end, w = write end)
	r, w, _ := os.Pipe()

	// Step 3: Redirect stdout to the write end of the pipe
	os.Stdout = w

	// Step 4: Run the funciton (it will print to our pipe instead of terminal)
	f()

	// Step 5: Close the write end and restore originsl stdout
	w.Close()
	os.Stdout = oldStdout

	// Step 6: Read everything from the pipe into a buffer
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Step 7: Return the captured output as a string
	return buf.String()
}

// TestQuadA_5x3 tests QuadA with 5 width and 3 height
func TestQuadA_5x3(t *testing.T) {
	// What we expect to see
	expected := "o---o\n|   |\no---o\n"

	// Capture what QuadA actually prints
	actual := captureOutput(func() {
		quads.QuadA(5, 3)
	})

	// Compare expected vs actual
	if actual != expected {
		t.Errorf("QuadA(5, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)

		// Show character-by-character for debugging
		t.Logf("Expected bytes: %v", []byte(expected))
		t.Logf("Actual bytes:   %v", []byte(actual))
	}
}
