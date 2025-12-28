package main

import (
	"fmt"

	"quad/quads"
)

func main() {
	printBanner()

	// Demonstrate all Quad outputs
	showQuad("QuadA", 5, 3, quads.QuadA)
	showQuad("QuadB", 5, 3, quads.QuadB)
	showQuad("QuadC", 5, 3, quads.QuadC)
	showQuad("QuadD", 5, 3, quads.QuadD)
	showQuad("QuadE", 5, 3, quads.QuadE)

	printFooter()
}

func printBanner() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   QUADRANGLE PATTERN GENERATOR         ║")
	fmt.Println("║        01 Founders - Raid 1            ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()
}

func showQuad(name string, width, height int, quadFunc func(int, int)) {
	fmt.Printf("--- %s (%dx%d) ---\n", name, width, height)
	quadFunc(width, height)
	fmt.Println()
}

func printFooter() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║   ✓ All demonstrations complete!       ║")
	fmt.Println("╚════════════════════════════════════════╝")
}