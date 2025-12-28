package quad_test

import (
	"quad/quads"
	"testing"
)

// Test: Standard 5x3 Rectangle
func TestQuadB_5x3(t *testing.T) {
	expected := "/***\\\n" +
		"*   *\n" +
		"\\***/\n"

	actual := captureOutput(func() {
		quads.QuadB(5, 3)
	})

	if actual != expected {
		t.Errorf("QuadB(5, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single character (1x1)
func TestQuadB_1x1(t *testing.T) {
	expected := "/\n"

	actual := captureOutput(func() {
		quads.QuadB(1, 1)
	})

	if actual != expected {
		t.Errorf("QuadB(1, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single row (5x1)
func TestQuadB_5x1(t *testing.T) {
	expected := "/***\\\n"

	actual := captureOutput(func() {
		quads.QuadB(5, 1)
	})

	if actual != expected {
		t.Errorf("QuadB(5, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single column (1x5)
func TestQuadB_1x5(t *testing.T) {
	expected := "/\n" +
		"*\n" +
		"*\n" +
		"*\n" +
		"\\\n"

	actual := captureOutput(func() {
		quads.QuadB(1, 5)
	})

	if actual != expected {
		t.Errorf("QuadB(1, 5) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Tall and narrow (1x3)
func TestQuadB_1x3(t *testing.T) {
	expected := "/\n" +
		"*\n" +
		"\\\n"

	actual := captureOutput(func() {
		quads.QuadB(1, 3)
	})

	if actual != expected {
		t.Errorf("QuadB(1, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Wide and short (3x1)
func TestQuadB_3x1(t *testing.T) {
	expected := "/*\\\n"

	actual := captureOutput(func() {
		quads.QuadB(3, 1)
	})

	if actual != expected {
		t.Errorf("QuadB(3, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Minimum valid rectangle (2x2)
func TestQuadB_2x2(t *testing.T) {
	expected := "/\\\n" +
		"\\/\n"

	actual := captureOutput(func() {
		quads.QuadB(2, 2)
	})

	if actual != expected {
		t.Errorf("QuadB(2, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Small square (3x3)
func TestQuadB_3x3(t *testing.T) {
	expected := "/*\\\n" +
		"* *\n" +
		"\\*/\n"

	actual := captureOutput(func() {
		quads.QuadB(3, 3)
	})

	if actual != expected {
		t.Errorf("QuadB(3, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Medium rectangle (7x4)
func TestQuadB_7x4(t *testing.T) {
	expected := "/*****\\\n" +
		"*     *\n" +
		"*     *\n" +
		"\\*****/\n"

	actual := captureOutput(func() {
		quads.QuadB(7, 4)
	})

	if actual != expected {
		t.Errorf("QuadB(7, 4) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Larger rectangle (10x6)
func TestQuadB_10x6(t *testing.T) {
	expected := "/********\\\n" +
		"*        *\n" +
		"*        *\n" +
		"*        *\n" +
		"*        *\n" +
		"\\********/\n"

	actual := captureOutput(func() {
		quads.QuadB(10, 6)
	})

	if actual != expected {
		t.Errorf("QuadB(10, 6) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very wide (15x2)
func TestQuadB_15x2(t *testing.T) {
	expected := "/*************\\\n" +
		"\\*************/\n"

	actual := captureOutput(func() {
		quads.QuadB(15, 2)
	})

	if actual != expected {
		t.Errorf("QuadB(15, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very tall (2x10)
func TestQuadB_2x10(t *testing.T) {
	expected := "/\\\n" +
		"**\n" +
		"**\n" +
		"**\n" +
		"**\n" +
		"**\n" +
		"**\n" +
		"**\n" +
		"**\n" +
		"\\/\n"

	actual := captureOutput(func() {
		quads.QuadB(2, 10)
	})

	if actual != expected {
		t.Errorf("QuadB(2, 10) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero width
func TestQuadB_0x5(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadB(0, 5)
	})

	if actual != expected {
		t.Errorf("QuadB(0, 5) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero height
func TestQuadB_5x0(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadB(5, 0)
	})

	if actual != expected {
		t.Errorf("QuadB(5, 0) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative width
func TestQuadB_Negative5x3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadB(-5, 3)
	})

	if actual != expected {
		t.Errorf("QuadB(-5, 3) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative height
func TestQuadB_5xNegative3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadB(5, -3)
	})

	if actual != expected {
		t.Errorf("QuadB(5, -3) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}
