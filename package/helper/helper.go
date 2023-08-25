// A package to provide common functionality over the application
package helper

import (
	"fmt"
	"os"
	"strings"
)

// function to check if file exists
func DoesFileExist(fileName string) bool {
	_, error := os.Stat(fileName)

	// check if error is "file not exists"
	if os.IsNotExist(error) {
		fmt.Printf("%v file does not exist\n", fileName)
		return false
	} else {
		fmt.Printf("%v file exist\n", fileName)
		return true
	}
	return false
}

//strip bearer form the jwt authentication token
func StripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 6 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}
