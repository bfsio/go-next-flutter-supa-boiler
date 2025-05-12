
package main

import (
	"fmt"
	"os"
)

func main() {
	checklist := []string{
		"1. Is per-tenant business logic better served via external services or plugins?",
		"2. Should we extract and version design tokens in a shared library?",
		"3. Should migrations and seeds be dry-runnable for CI tests?",
		"4. Should we abstract DB access with sqlc, gorm, or go full PGX?",
		"5. Do we need a plugin system for per-tenant logic overrides?",
	}

	fmt.Println("Pet Rock Development Checklist")
	fmt.Println("----------------------------")

	for _, item := range checklist {
		fmt.Println(item)
	}

	// Provide an exit code if desired
	os.Exit(0)
}
