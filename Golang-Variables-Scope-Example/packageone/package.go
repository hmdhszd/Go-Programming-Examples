package packageone

import "fmt"

var PublicVar = "I'm a PublicVar in packageone"
var privateVar = "I'm a PrivateVar (on exported) in packageone"

func notExportedFunction() {
	fmt.Println("This is notExportedFunction")
}

func ExportedFunction() {
	fmt.Println("This is ExportedFunction")
}
