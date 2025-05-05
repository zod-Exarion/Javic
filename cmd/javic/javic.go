package javic

import (
	"fmt"
	"javic/core/tokenizer"
	"os"
)

func Javic(fileName string) {
	// Read the QBASIC File
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(content))
	fmt.Println(tokenizer.LookupIdent("cls"))
}
