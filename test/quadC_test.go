package quad_test

import (
	"quad/quads"
	"testing"
)

// Test: Standard 5x3 rectangle
func TestQuadC_5x3(t *testing.T) {
	expected := "ABBBA\n" +
		"B   B\n" +
		"CBBBC\n"

	actual := captureOutput(func() {
		quads.QuadC(5, 3)
	})

	if actual != expected {
		t.Errorf("QuadC(5, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single character (1x1)
func TestQuadC_1x1(t *testing.T) {
	expected := "A\n"

	actual := captureOutput(func() {
		quads.QuadC(1, 1)
	})

	if actual != expected {
		t.Errorf("QuadC(1, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single row (5x1)
func TestQuadC_5x1(t *testing.T) {
	expected := "ABBBA\n"

	actual := captureOutput(func() {
		quads.QuadC(5, 1)
	})

	if actual != expected {
		t.Errorf("QuadC(5, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single column (1x5)
func TestQuadC_1x5(t *testing.T) {
	expected := "A\n" +
		"B\n" +
		"B\n" +
		"B\n" +
		"C\n"

	actual := captureOutput(func() {
		quads.QuadC(1, 5)
	})

	if actual != expected {
		t.Errorf("QuadC(1, 5) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Tall and narrow (1x3)
func TestQuadC_1x3(t *testing.T) {
	expected := "A\n" +
		"B\n" +
		"C\n"

	actual := captureOutput(func() {
		quads.QuadC(1, 3)
	})

	if actual != expected {
		t.Errorf("QuadC(1, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Wide and short (3x1)
func TestQuadC_3x1(t *testing.T) {
	expected := "ABA\n"

	actual := captureOutput(func() {
		quads.QuadC(3, 1)
	})

	if actual != expected {
		t.Errorf("QuadC(3, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Minimum valid rectangle (2x2)
func TestQuadC_2x2(t *testing.T) {
	expected := "AA\n" +
		"CC\n"

	actual := captureOutput(func() {
		quads.QuadC(2, 2)
	})

	if actual != expected {
		t.Errorf("QuadC(2, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Small square (3x3)
func TestQuadC_3x3(t *testing.T) {
	expected := "ABA\n" +
		"B B\n" +
		"CBC\n"

	actual := captureOutput(func() {
		quads.QuadC(3, 3)
	})

	if actual != expected {
		t.Errorf("QuadC(3, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Medium rectangle (7x4)
func TestQuadC_7x4(t *testing.T) {
	expected := "ABBBBBA\n" +
		"B     B\n" +
		"B     B\n" +
		"CBBBBBC\n"

	actual := captureOutput(func() {
		quads.QuadC(7, 4)
	})

	if actual != expected {
		t.Errorf("QuadC(7, 4) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Larger rectangle (10x6)
func TestQuadC_10x6(t *testing.T) {
	expected := "ABBBBBBBBA\n" +
		"B        B\n" +
		"B        B\n" +
		"B        B\n" +
		"B        B\n" +
		"CBBBBBBBBC\n"

	actual := captureOutput(func() {
		quads.QuadC(10, 6)
	})

	if actual != expected {
		t.Errorf("QuadC(10, 6) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very wide (15x2)
func TestQuadC_15x2(t *testing.T) {
	expected := "ABBBBBBBBBBBBBA\n" +
		"CBBBBBBBBBBBBBC\n"

	actual := captureOutput(func() {
		quads.QuadC(15, 2)
	})

	if actual != expected {
		t.Errorf("QuadC(15, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very tall (2x10)
func TestQuadC_2x10(t *testing.T) {
	expected := "AA\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"CC\n"

	actual := captureOutput(func() {
		quads.QuadC(2, 10)
	})

	if actual != expected {
		t.Errorf("QuadC(2, 10) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero width
func TestQuadC_0x5(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadC(0, 5)
	})

	if actual != expected {
		t.Errorf("QuadC(0, 5) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero height
func TestQuadC_5x0(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadC(5, 0)
	})

	if actual != expected {
		t.Errorf("QuadC(5, 0) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative width
func TestQuadC_Negative5x3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadC(-5, 3)
	})

	if actual != expected {
		t.Errorf("QuadC(-5, 3) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative height
func TestQuadC_5xNegative3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadC(5, -3)
	})

	if actual != expected {
		t.Errorf("QuadC(5, -3) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}
