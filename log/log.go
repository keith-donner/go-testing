package debug

import (
	"fmt"
)

func Log(text string) {
	fmt.Print(text)
	//do somthing
}

func LogEmail(subject string, text string, to string) {
	fmt.Print(subject)
	fmt.Print(text)
	fmt.Print(to)
	//do something
}
