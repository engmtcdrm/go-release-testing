package main

import (
	"fmt"

	"github.com/engmtcdrm/go-release-testing/app"
)

func main() {
	fmt.Printf("WELCOME TO %s PREPARE TO BE BUILT\n", app.Name)
	fmt.Printf("The current version is %s\n", app.Version)
}
