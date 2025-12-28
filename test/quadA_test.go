package quad_test

import (
	"quad/quads"
	"testing"
)

// Test: Standard 5x3 Rectangle
func TestQuadA_5x3(t *testing.T) {
	expected := "o---o\n|   |\no---o\n"

	actual := captureOutput(func() {
		quads.QuadA(5, 3)
	})

	if actual != expected {
		t.Errorf("QuadA(5, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single character (1x1)
func TestQuadA_1x1(t *testing.T) {
	expected := "o\n"

	actual := captureOutput(func() {
		quads.QuadA(1, 1)
	})

	if actual != expected {
		t.Errorf("QuadA(1, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single row (5x1)
func TestQuadA_5x1(t *testing.T) {
	expected := "o---o\n"

	actual := captureOutput(func() {
		quads.QuadA(5, 1)
	})

	if actual != expected {
		t.Errorf("QuadA(5, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single column (1x5)
func TestQuadA_1x5(t *testing.T) {
	expected := "o\n|\n|\n|\no\n"

	actual := captureOutput(func() {
		quads.QuadA(1, 5)
	})

	if actual != expected {
		t.Errorf("QuadA(1, 5) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Tall and narrow (1x3)
func TestQuadA_1x3(t *testing.T) {
	expected := "o\n|\no\n"

	actual := captureOutput(func() {
		quads.QuadA(1, 3)
	})

	if actual != expected {
		t.Errorf("QuadA(1, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Wide and short (3x1)
func TestQuadA_3x1(t *testing.T) {
	expected := "o-o\n"

	actual := captureOutput(func() {
		quads.QuadA(3, 1)
	})

	if actual != expected {
		t.Errorf("QuadA(3, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Minimum valid rectangle (2x2)
func TestQuadA_2x2(t *testing.T) {
	expected := "oo\noo\n"

	actual := captureOutput(func() {
		quads.QuadA(2, 2)
	})

	if actual != expected {
		t.Errorf("QuadA(2, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Slightly bigger (3x3)
func TestQuadA_3x3(t *testing.T) {
	expected := "o-o\n| |\no-o\n"

	actual := captureOutput(func() {
		quads.QuadA(3, 3)
	})

	if actual != expected {
		t.Errorf("QuadA(3, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Medium rectangle (7x4)
func TestQuadA_7x4(t *testing.T) {
	expected := "o-----o\n|     |\n|     |\no-----o\n"

	actual := captureOutput(func() {
		quads.QuadA(7, 4)
	})

	if actual != expected {
		t.Errorf("QuadA(7, 4) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Larger rectangle (10x6)
func TestQuadA_10x6(t *testing.T) {
	expected := "o--------o\n" +
		"|        |\n" +
		"|        |\n" +
		"|        |\n" +
		"|        |\n" +
		"o--------o\n"

	actual := captureOutput(func() {
		quads.QuadA(10, 6)
	})

	if actual != expected {
		t.Errorf("QuadA(10, 6) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very wide (15x2)
func TestQuadA_15x2(t *testing.T) {
	expected := "o-------------o\no-------------o\n"

	actual := captureOutput(func() {
		quads.QuadA(15, 2)
	})

	if actual != expected {
		t.Errorf("QuadA(15, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very tall (2x10)
func TestQuadA_2x10(t *testing.T) {
	expected := "oo\n" +
		"||\n" +
		"||\n" +
		"||\n" +
		"||\n" +
		"||\n" +
		"||\n" +
		"||\n" +
		"||\n" +
		"oo\n"

	actual := captureOutput(func() {
		quads.QuadA(2, 10)
	})

	if actual != expected {
		t.Errorf("QuadA(2, 10) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero width
func TestQuadA_0x5(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadA(0, 5)
	})

	if actual != expected {
		t.Errorf("QuadA(0, 5) should print nothing")
		t.Logf("Expected: (empty)")
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero height
func TestQuadA_5x0(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadA(5, 0)
	})

	if actual != expected {
		t.Errorf("QuadA(5, 0) should print nothing")
		t.Logf("Expected: (empty)")
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - both zero
func TestQuadA_0x0(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadA(0, 0)
	})

	if actual != expected {
		t.Errorf("QuadA(0, 0) should print nothing")
		t.Logf("Expected: (empty)")
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative width
func TestQuadA_Negative5x3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadA(-5, 3)
	})

	if actual != expected {
		t.Errorf("QuadA(-5, 3) should print nothing")
		t.Logf("Expected: (empty)")
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative height
func TestQuadA_5xNegative3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadA(5, -3)
	})

	if actual != expected {
		t.Errorf("QuadA(5, -3) should print nothing")
		t.Logf("Expected: (empty)")
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - both negative
func TestQuadA_NegativeNegative(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadA(-5, -3)
	})

	if actual != expected {
		t.Errorf("QuadA(-5, -3) should print nothing")
		t.Logf("Expected: (empty)")
		t.Logf("Got:\n%s", actual)
	}
}
