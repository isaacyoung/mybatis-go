package main

import (
	 "isaac/config"
	"fmt"
	"isaac/jdbc"
)

func main() {

	// config
	con, err := config.GetContent("config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	if con.Jdbc.Driver != "com.mysql.jdbc.Driver" {
		fmt.Println("not get config values")
		return
	}

	// TODO jdbc
	tables, err := jdbc.GetTables(con)
	if err != nil {
		fmt.Println(err)
		return 
	}

	for _, t := range tables {
		fmt.Println(t.Name + " " + t.Comment)
	}

	columns, err := jdbc.GetColumns(con, con.Jdbc.Table)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, t := range columns {
		fmt.Println(t.Name + " " + t.Comment)
	}

	// TODO code


}
