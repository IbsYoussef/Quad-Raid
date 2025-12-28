package quad_test

import (
	"quad/quads"
	"testing"
)

// Test: Standard 5x3 rectangle
func TestQuadD_5x3(t *testing.T) {
	expected := "ABBBC\n" +
		"B   B\n" +
		"ABBBC\n"

	actual := captureOutput(func() {
		quads.QuadD(5, 3)
	})

	if actual != expected {
		t.Errorf("QuadD(5, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single character (1x1)
func TestQuadD_1x1(t *testing.T) {
	expected := "A\n"

	actual := captureOutput(func() {
		quads.QuadD(1, 1)
	})

	if actual != expected {
		t.Errorf("QuadD(1, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single row (5x1)
func TestQuadD_5x1(t *testing.T) {
	expected := "ABBBC\n"

	actual := captureOutput(func() {
		quads.QuadD(5, 1)
	})

	if actual != expected {
		t.Errorf("QuadD(5, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single column (1x5)
func TestQuadD_1x5(t *testing.T) {
	expected := "A\n" +
		"B\n" +
		"B\n" +
		"B\n" +
		"A\n"

	actual := captureOutput(func() {
		quads.QuadD(1, 5)
	})

	if actual != expected {
		t.Errorf("QuadD(1, 5) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Tall and narrow (1x3)
func TestQuadD_1x3(t *testing.T) {
	expected := "A\n" +
		"B\n" +
		"A\n"

	actual := captureOutput(func() {
		quads.QuadD(1, 3)
	})

	if actual != expected {
		t.Errorf("QuadD(1, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Wide and short (3x1)
func TestQuadD_3x1(t *testing.T) {
	expected := "ABC\n"

	actual := captureOutput(func() {
		quads.QuadD(3, 1)
	})

	if actual != expected {
		t.Errorf("QuadD(3, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Minimum valid rectangle (2x2)
func TestQuadD_2x2(t *testing.T) {
	expected := "AC\n" +
		"AC\n"

	actual := captureOutput(func() {
		quads.QuadD(2, 2)
	})

	if actual != expected {
		t.Errorf("QuadD(2, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Small square (3x3)
func TestQuadD_3x3(t *testing.T) {
	expected := "ABC\n" +
		"B B\n" +
		"ABC\n"

	actual := captureOutput(func() {
		quads.QuadD(3, 3)
	})

	if actual != expected {
		t.Errorf("QuadD(3, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Medium rectangle (7x4)
func TestQuadD_7x4(t *testing.T) {
	expected := "ABBBBBC\n" +
		"B     B\n" +
		"B     B\n" +
		"ABBBBBC\n"

	actual := captureOutput(func() {
		quads.QuadD(7, 4)
	})

	if actual != expected {
		t.Errorf("QuadD(7, 4) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Larger rectangle (10x6)
func TestQuadD_10x6(t *testing.T) {
	expected := "ABBBBBBBBC\n" +
		"B        B\n" +
		"B        B\n" +
		"B        B\n" +
		"B        B\n" +
		"ABBBBBBBBC\n"

	actual := captureOutput(func() {
		quads.QuadD(10, 6)
	})

	if actual != expected {
		t.Errorf("QuadD(10, 6) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very wide (15x2)
func TestQuadD_15x2(t *testing.T) {
	expected := "ABBBBBBBBBBBBBC\n" +
		"ABBBBBBBBBBBBBC\n"

	actual := captureOutput(func() {
		quads.QuadD(15, 2)
	})

	if actual != expected {
		t.Errorf("QuadD(15, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very tall (2x10)
func TestQuadD_2x10(t *testing.T) {
	expected := "AC\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"AC\n"

	actual := captureOutput(func() {
		quads.QuadD(2, 10)
	})

	if actual != expected {
		t.Errorf("QuadD(2, 10) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero width
func TestQuadD_0x5(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadD(0, 5)
	})

	if actual != expected {
		t.Errorf("QuadD(0, 5) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero height
func TestQuadD_5x0(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadD(5, 0)
	})

	if actual != expected {
		t.Errorf("QuadD(5, 0) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative width
func TestQuadD_Negative5x3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadD(-5, 3)
	})

	if actual != expected {
		t.Errorf("QuadD(-5, 3) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative height
func TestQuadD_5xNegative3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadD(5, -3)
	})

	if actual != expected {
		t.Errorf("QuadD(5, -3) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}
