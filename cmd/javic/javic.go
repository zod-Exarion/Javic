package javic

import (
	"javic/test"
	"testing"
)

func Javic(fileName string) {
	// Read the QBASIC File
	// content, err := os.ReadFile(fileName)
	// if err != nil {
	// 	fmt.Printf("Error reading file: %v\n", err)
	// 	os.Exit(1)
	// }

	test.TestNextToken(&testing.T{})
}
