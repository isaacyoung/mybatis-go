
// cd src/isaac
// go run main.go

package main

import (
	"fmt"
	"mybatis-go/content"
	"mybatis-go/code"
)

func main() {

	// config
	con, err := content.Build("config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}

	con.ClearTarget()

	// code
	code.Model(con)
	code.Dao(con)
	code.Serv(con)
	code.ServImpl(con)
	code.Xml(con)

	// copy
	if con.IsOver() {
		con.CopyToProject()
	}
}
