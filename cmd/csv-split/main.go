package main

import (
	"fmt"
	"os"
	"github.com/harkaitz/go-dbt1"
)
func main() {
	var arg, val string
	var csv dbt1.CSV
	for _, arg = range os.Args[1:] {
		csv = dbt1.CSV(arg)
		for _, val = range csv.Getl() {
			fmt.Println(val)
		}
	}
}
