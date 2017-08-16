
// cd src/isaac
// go run main.go

package main

import (
	"fmt"
	"isaac/content"
	"isaac/code"
)

func main() {

	// config
	con, err := content.Build("config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	con.ClearTarget()

	// TODO code
	code.Model(con)
	code.Dao(con)
}
