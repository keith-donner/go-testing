package util

import (
	"fmt"
	"regexp"
)

// output to the screen during testing only
//
// ________________________________________________________________
//
// debugText: can be of any variable type
func DebugOutput(debugText any) {
	fmt.Println(debugText)
}

// will add debug during testing and will add to log
//
// ________________________________________________________________
//
// debugText: can be of any variable type
//
// addToLog: if true, this debugText will be added to the log
//
// sendToEmail: will send this debug text to the email address(es). Separate with a ;
// func DebugLogAndEmail(debugText any, addToLog bool, sendToEmail string) {
// 	fmt.Println(debugText)

// }

func RegexGroups(pattern, matchString string) (paramsMap map[string]string) {

	var compRegEx = regexp.MustCompile(pattern)
	match := compRegEx.FindStringSubmatch(matchString)

	paramsMap = make(map[string]string)
	for i, name := range compRegEx.SubexpNames() {
		if i > 0 && i <= len(match) {
			paramsMap[name] = match[i]
		}
	}
	return paramsMap
}
