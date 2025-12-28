# Quad Raid Project

<div align="center">

![Demo](assets/demo.gif)

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

</div>

---

A Go-based ASCII art rectangle generator that produces five distinct geometric patterns using nested loops and conditional logic. Built as part of the **Piscine-Go** at [01 Founders](https://01-edu.org/).

## Table of Contents

- [About](#about)
- [Features](#features)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Pattern Examples](#pattern-examples)
- [Technical Implementation](#technical-implementation)
- [Testing](#testing)
- [Development](#development)
- [Learning Outcomes](#learning-outcomes)
- [Author](#author)
- [License](#license)

<a name="about"></a>

## 🎯 About

This project implements five different ASCII rectangle drawing functions (QuadA through QuadE), each producing unique corner and edge patterns. It demonstrates fundamental programming concepts including nested loops, conditional logic, and systematic problem-solving approaches.

### Part of 01 Founders Piscine

<div align="center">

<img src="assets/image.png" alt="01 Founders Logo" width="400"/>

</div>

This exercise is part of the **Piscine-Go** selection process at [01 Founders](https://01-edu.org/), a 3-week intensive coding bootcamp. The Piscine (French for "swimming pool") is an immersive, peer-to-peer learning experience that tests:

- **Problem-solving under pressure** - Complete exercises with time constraints
- **Self-directed learning** - Figure out solutions through documentation and experimentation
- **Collaboration** - Work with peers to understand complex problems (Raid projects are team-based)
- **Code quality** - Write clean, well-documented, testable code

The Quad Raid specifically focuses on:

- Understanding nested loop mechanics and execution flow
- Implementing position-based conditional logic
- Building systematic testing strategies
- Creating reusable, modular code structures

<a name="features"></a>

## ✨ Features

- **Five Distinct Patterns**: QuadA, QuadB, QuadC, QuadD, and QuadE - each with unique corner/edge characters
- **Nested Loop Architecture**: Demonstrates row-by-column iteration with position-aware logic
- **Input Validation**: Handles invalid dimensions (zero, negative) with error messages
- **Comprehensive Testing**: 70+ unit tests covering edge cases, various sizes, and invalid inputs
- **Clean Code Structure**: Modular design with separated concerns and detailed documentation
- **Professional Demo**: Formatted output display showcasing all patterns

<a name="project-structure"></a>

## 📁 Project Structure

```
quad/
├── main.go                    # Demo program with formatted output
├── quads/                     # Shape implementations
│   ├── quadA.go              # Pattern: o---o | o---o (corners: o, edges: -, |)
│   ├── quadB.go              # Pattern: /***\ * \***/ (corners: /, \, edges: *)
│   ├── quadC.go              # Pattern: ABBBA B CBBBC (corners: A, C, edges: B)
│   ├── quadD.go              # Pattern: ABBBC B ABBBC (corners: A, C, edges: B)
│   └── quadE.go              # Pattern: ABBBC B CBBBA (corners: A, C, edges: B)
├── test/                      # Comprehensive test suite
│   ├── helpers_test.go       # Shared test utilities (captureOutput)
│   ├── quadA_test.go         # 17 tests for QuadA
│   ├── quadB_test.go         # 17 tests for QuadB
│   ├── quadC_test.go         # 17 tests for QuadC
│   ├── quadD_test.go         # 17 tests for QuadD
│   └── quadE_test.go         # 17 tests for QuadE
├── assets/                    # Project assets
│   ├── 01foundersLogo.webp   # 01 Founders branding
│   └── demo.gif              # Terminal recording showcase
├── go.mod                     # Module definition
├── go.sum                     # Dependency checksums
├── LICENSE                    # MIT License
└── README.md                  # This file
```

### File Overview

**`main.go`**

- Demonstrates all five quad patterns with formatted output
- Includes styled banner and section headers
- Showcases 5x3 rectangles for each pattern

**`quads/`**

- Each quad implementation in its own file
- Self-contained with full documentation
- Consistent structure across all five patterns
- Comments explain corner/edge logic and algorithm flow

**`test/`**

- `helpers_test.go`: Shared `captureOutput()` function for stdout testing
- Individual test files for each quad with 17 tests each:
  - Standard cases: 5x3, 1x1, 5x1, 1x5, 1x3, 3x1
  - Edge cases: 2x2, 3x3, 7x4, 10x6, 15x2, 2x10
  - Invalid inputs: 0x5, 5x0, -5x3, 5x-3

**`assets/`**

- `01foundersLogo.webp`: 01 Founders branding and logo
- `demo.gif`: Animated terminal recording demonstrating all quadrangle patterns

<a name="installation"></a>

## 🚀 Installation

### Prerequisites

- Go 1.21.0 or later
- Git

### Setup

```bash
# Clone the repository
git clone https://github.com/IbsYoussef/Quad-Raid.git
cd Quad-Raid

# Download dependencies
go mod download

# Run the demo
go run main.go
```

<a name="usage"></a>

## 💻 Usage

### Demo Program

Run the formatted demonstration:

```bash
go run main.go
```

**Output:**

```
╔════════════════════════════════════════╗
║   QUADRANGLE PATTERN GENERATOR         ║
║        01 Founders - Raid 1            ║
╚════════════════════════════════════════╝

--- QuadA (5x3) ---
o---o
|   |
o---o

--- QuadB (5x3) ---
/***\
*   *
\***/

--- QuadC (5x3) ---
ABBBA
B   B
CBBBC

--- QuadD (5x3) ---
ABBBC
B   B
ABBBC

--- QuadE (5x3) ---
ABBBC
B   B
CBBBA

╔════════════════════════════════════════╗
║   ✓ All demonstrations complete!       ║
╚════════════════════════════════════════╝
```

### Using in Your Code

```go
package main

import (
    "quad/quads"
)

func main() {
    // Draw different patterns
    quads.QuadA(10, 5)
    quads.QuadB(7, 3)
    quads.QuadC(15, 8)
    quads.QuadD(20, 10)
    quads.QuadE(12, 6)

    // Invalid inputs print error message
    quads.QuadA(0, 5)     // "Invalid input parameters"
    quads.QuadB(-3, 4)    // "Invalid input parameters"
}
```

<a name="pattern-examples"></a>

## 🎨 Pattern Examples

### QuadA - Classic Corners

```
5x3:              1x1:       10x5:
o---o             o          o--------o
|   |                        |        |
o---o                        |        |
                             |        |
                             o--------o
```

**Pattern:** `o` at all corners, `-` on horizontal edges, `|` on vertical edges

---

### QuadB - Diagonal Corners

```
5x3:              1x1:       10x5:
/***\             /          /********\
*   *                        *        *
\***/                        *        *
                             *        *
                             \********/
```

**Pattern:** `/` at top-left & bottom-right, `\` at top-right & bottom-left, `*` on other edges

---

### QuadC - A-C Pattern

```
5x3:              1x1:       10x5:
ABBBA             A          ABBBBBBBBA
B   B                        B        B
CBBBC                        B        B
                             B        B
                             CBBBBBBBBC
```

**Pattern:** `A` on top corners, `C` on bottom corners, `B` on other edges

---

### QuadD - Left-Right A-C

```
5x3:              1x1:       10x5:
ABBBC             A          ABBBBBBBBC
B   B                        B        B
ABBBC                        B        B
                             B        B
                             ABBBBBBBBC
```

**Pattern:** `A` on left corners, `C` on right corners, `B` on other edges

---

### QuadE - Diagonal A-C (Reversed)

```
5x3:              1x1:       10x5:
ABBBC             A          ABBBBBBBBC
B   B                        B        B
CBBBA                        B        B
                             B        B
                             CBBBBBBBBA
```

**Pattern:** `A` at top-left & bottom-right, `C` at top-right & bottom-left, `B` on other edges

<a name="technical-implementation"></a>

## 🔧 Technical Implementation

### Algorithm Structure

All quads follow the same nested loop pattern:

```go
// Outer loop: iterate rows (top to bottom)
for row := 0; row < y; row++ {
    // Inner loop: iterate columns (left to right)
    for col := 0; col < x; col++ {
        // Position-based logic determines character
        if /* corner condition */ {
            // Print corner character
        } else if /* edge condition */ {
            // Print edge character
        } else {
            // Print interior (space)
        }
    }
    // Print newline after each row
    z01.PrintRune('\n')
}
```

### Key Concepts

**Position Detection:**

- `row == 0` → Top row
- `row == y-1` → Bottom row
- `col == 0` → Left column
- `col == x-1` → Right column

**Condition Priority (QuadA example):**

1. **Corners** (most specific): `(row == 0 || row == y-1) && (col == 0 || col == x-1)`
2. **Horizontal edges**: `row == 0 || row == y-1`
3. **Vertical edges**: `col == 0 || col == x-1`
4. **Interior** (catch-all): Everything else

**Pattern Variations:**

- **QuadA**: All corners same character → Single combined condition
- **QuadB/C/D/E**: Different corners → Separate conditions for each corner position

### Edge Cases Handled

- **1x1 rectangles**: Single character (corner only)
- **Single row (Nx1)**: Horizontal line with corners
- **Single column (1xN)**: Vertical line with corners
- **2x2 rectangles**: Four corners, no edges or interior
- **Invalid inputs**: Zero or negative dimensions trigger error message

<a name="testing"></a>

## 🧪 Testing

### Run All Tests

```bash
# Run complete test suite
go test -v ./test/

# Run tests for specific quad
go test -v ./test/ -run TestQuadA
go test -v ./test/ -run TestQuadB

# Run with coverage
go test -v -cover ./test/
```

### Test Coverage

**85 total tests across all quads:**

| Quad  | Tests | Coverage                                |
| ----- | ----- | --------------------------------------- |
| QuadA | 17    | Valid sizes, edge cases, invalid inputs |
| QuadB | 17    | Valid sizes, edge cases, invalid inputs |
| QuadC | 17    | Valid sizes, edge cases, invalid inputs |
| QuadD | 17    | Valid sizes, edge cases, invalid inputs |
| QuadE | 17    | Valid sizes, edge cases, invalid inputs |

**Test Categories:**

- ✅ Standard rectangles: 5x3, 7x4, 10x6
- ✅ Edge cases: 1x1, 1x5, 5x1, 1x3, 3x1, 2x2, 3x3
- ✅ Large sizes: 15x2, 2x10
- ✅ Invalid inputs: 0xN, Nx0, -NxM

### Example Test Output

```bash
$ go test -v ./test/ -run TestQuadA_5x3
=== RUN   TestQuadA_5x3
--- PASS: TestQuadA_5x3 (0.00s)
PASS
ok      quad/test       0.003s
```

<a name="development"></a>

## 🛠️ Development

### Code Quality

```bash
# Format code
go fmt ./...

# Run linter
go vet ./...

# Check for common issues
golangci-lint run
```

### Adding New Patterns

To add a new quad pattern:

1. Create `quads/quadF.go`:

```go
package quads

import (
    "fmt"
    "github.com/01-edu/z01"
)

func QuadF(x, y int) {
    // Validate input
    // Implement nested loops
    // Define corner/edge logic
}
```

2. Create `test/quadF_test.go`:

```go
package quad_test

import (
    "testing"
    "quad/quads"
)

func TestQuadF_5x3(t *testing.T) {
    // ... test implementation
}
```

3. Add to demo in `main.go`:

```go
showQuad("QuadF", 5, 3, quads.QuadF)
```

<a name="learning-outcomes"></a>

## 📚 Learning Outcomes

Through this project, I strengthened understanding of:

**Core Programming Concepts:**

- ✅ Nested loop mechanics and execution flow
- ✅ Position-based conditional logic
- ✅ Systematic problem decomposition
- ✅ Edge case identification and handling

**Go-Specific Skills:**

- ✅ Package organization and imports
- ✅ Function design and naming conventions
- ✅ Character output with `z01.PrintRune`
- ✅ Error handling and input validation

**Software Engineering Practices:**

- ✅ Unit testing with table-driven tests
- ✅ Test helper functions and DRY principles
- ✅ Code documentation and comments
- ✅ Version control and commit messages

**Problem-Solving Approach:**

- ✅ Breaking complex problems into smaller pieces
- ✅ Identifying patterns across similar problems
- ✅ Iterative refinement and debugging
- ✅ Learning from mistakes (off-by-one errors!)

<a name="author"></a>

## 👤 Author

**[IbsYoussef](https://github.com/IbsYoussef)** - Built as part of the 01 Founders Piscine-Go selection process.

<a name="license"></a>

## 📄 License

MIT License - Copyright (c) 2025 IbsYoussef

This project was created as part of the [01 Founders](https://01-edu.org/) Piscine-Go curriculum.

---

<div align="center">

**[⬆ Back to Top](#quad-raid-project)**

</div>
