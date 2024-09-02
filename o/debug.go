package o

import "fmt"

// output to the screen during testing only
func DebugOutput(debugText any) {
	fmt.Println(debugText)
}

// will add debug during testing and will add to log
func DebugLogAndEmail(debugText any, addToLog bool, sendToEmail string) {
	fmt.Println()
}
