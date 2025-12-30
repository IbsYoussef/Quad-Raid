package quad_test

import (
	"bytes"
	"io"
	"os"
)

// Helper function to capture output of what gets printed to stdout
// This helper function is shared across all test files in quad_test package
func captureOutput(f func()) string {
	// Step 1: Save the original stdout
	oldStdout := os.Stdout

	// Step 2: Create a pipe (r = read end, w = write end)
	r, w, _ := os.Pipe()

	// Step 3: Redirect stdout to the write end of the pipe
	os.Stdout = w

	// Step 4: Run the funciton (it will print to our pipe instead of terminal)
	f()

	// Step 5: Close the write end and restore original stdout
	w.Close()
	os.Stdout = oldStdout

	// Step 6: Read everything from the pipe into a buffer
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Step 7: Return the captured output as a string
	return buf.String()
}
