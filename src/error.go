package src

import (
	"fmt"
	"os"
)

func HandleInputError(Len uint64, numLen uint64, charLen uint64) {
	if numLen > Len || charLen > Len {
		fmt.Println("The values of -n and -c must be less than the value of -l.")
		os.Exit(1)
	}
}
