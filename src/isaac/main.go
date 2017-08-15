
// cd src/isaac
// go run main.go

package main

import (
	"fmt"
	"isaac/content"
)

func main() {

	// config
	con, err := content.Build("config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	if con.Config.Jdbc.Driver != "com.mysql.jdbc.Driver" {
		fmt.Println("not get config values")
		return
	}

	con.ClearTarget()

	// TODO code


}
