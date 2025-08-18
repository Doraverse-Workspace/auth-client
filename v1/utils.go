package v1

import "fmt"

// Errorf prints an error message if an error occurs
// err is the error
// message is the message to print
// isDebug is a boolean to check if the error should be printed
func Errorf(err error, message string, isDebug bool) {
	if err != nil {
		if isDebug {
			fmt.Println("❌ ", message, err)
		}
	}
}
