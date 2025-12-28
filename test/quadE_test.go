package quad_test

import (
	"quad/quads"
	"testing"
)

// Test: Standard 5x3 rectangle
func TestQuadE_5x3(t *testing.T) {
	expected := "ABBBC\n" +
		"B   B\n" +
		"CBBBA\n"

	actual := captureOutput(func() {
		quads.QuadE(5, 3)
	})

	if actual != expected {
		t.Errorf("QuadE(5, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single character (1x1)
func TestQuadE_1x1(t *testing.T) {
	expected := "A\n"

	actual := captureOutput(func() {
		quads.QuadE(1, 1)
	})

	if actual != expected {
		t.Errorf("QuadE(1, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single row (5x1)
func TestQuadE_5x1(t *testing.T) {
	expected := "ABBBC\n"

	actual := captureOutput(func() {
		quads.QuadE(5, 1)
	})

	if actual != expected {
		t.Errorf("QuadE(5, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Single column (1x5)
func TestQuadE_1x5(t *testing.T) {
	expected := "A\n" +
		"B\n" +
		"B\n" +
		"B\n" +
		"C\n"

	actual := captureOutput(func() {
		quads.QuadE(1, 5)
	})

	if actual != expected {
		t.Errorf("QuadE(1, 5) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Tall and narrow (1x3)
func TestQuadE_1x3(t *testing.T) {
	expected := "A\n" +
		"B\n" +
		"C\n"

	actual := captureOutput(func() {
		quads.QuadE(1, 3)
	})

	if actual != expected {
		t.Errorf("QuadE(1, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Wide and short (3x1)
func TestQuadE_3x1(t *testing.T) {
	expected := "ABC\n"

	actual := captureOutput(func() {
		quads.QuadE(3, 1)
	})

	if actual != expected {
		t.Errorf("QuadE(3, 1) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Minimum valid rectangle (2x2)
func TestQuadE_2x2(t *testing.T) {
	expected := "AC\n" +
		"CA\n"

	actual := captureOutput(func() {
		quads.QuadE(2, 2)
	})

	if actual != expected {
		t.Errorf("QuadE(2, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Small square (3x3)
func TestQuadE_3x3(t *testing.T) {
	expected := "ABC\n" +
		"B B\n" +
		"CBA\n"

	actual := captureOutput(func() {
		quads.QuadE(3, 3)
	})

	if actual != expected {
		t.Errorf("QuadE(3, 3) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Medium rectangle (7x4)
func TestQuadE_7x4(t *testing.T) {
	expected := "ABBBBBC\n" +
		"B     B\n" +
		"B     B\n" +
		"CBBBBBA\n"

	actual := captureOutput(func() {
		quads.QuadE(7, 4)
	})

	if actual != expected {
		t.Errorf("QuadE(7, 4) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Larger rectangle (10x6)
func TestQuadE_10x6(t *testing.T) {
	expected := "ABBBBBBBBC\n" +
		"B        B\n" +
		"B        B\n" +
		"B        B\n" +
		"B        B\n" +
		"CBBBBBBBBA\n"

	actual := captureOutput(func() {
		quads.QuadE(10, 6)
	})

	if actual != expected {
		t.Errorf("QuadE(10, 6) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very wide (15x2)
func TestQuadE_15x2(t *testing.T) {
	expected := "ABBBBBBBBBBBBBC\n" +
		"CBBBBBBBBBBBBBA\n"

	actual := captureOutput(func() {
		quads.QuadE(15, 2)
	})

	if actual != expected {
		t.Errorf("QuadE(15, 2) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Very tall (2x10)
func TestQuadE_2x10(t *testing.T) {
	expected := "AC\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"BB\n" +
		"CA\n"

	actual := captureOutput(func() {
		quads.QuadE(2, 10)
	})

	if actual != expected {
		t.Errorf("QuadE(2, 10) FAILED")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero width
func TestQuadE_0x5(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadE(0, 5)
	})

	if actual != expected {
		t.Errorf("QuadE(0, 5) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - zero height
func TestQuadE_5x0(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadE(5, 0)
	})

	if actual != expected {
		t.Errorf("QuadE(5, 0) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative width
func TestQuadE_Negative5x3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadE(-5, 3)
	})

	if actual != expected {
		t.Errorf("QuadE(-5, 3) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}

// Test: Invalid - negative height
func TestQuadE_5xNegative3(t *testing.T) {
	expected := "Invalid input parameters\n"

	actual := captureOutput(func() {
		quads.QuadE(5, -3)
	})

	if actual != expected {
		t.Errorf("QuadE(5, -3) should print invalid message")
		t.Logf("Expected:\n%s", expected)
		t.Logf("Got:\n%s", actual)
	}
}
