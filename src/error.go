package src

import (
	"fmt"
	"os"
)

// HandleInputError handles the error of input parameters.
func HandleInputError(totalLength uint64, numLength uint64, charLength uint64) {
	if numLength > totalLength || charLength > totalLength {
		fmt.Println("The values of -n and -c must be less than the value of -l.")
		os.Exit(1)
	}
}
