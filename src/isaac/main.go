package main

import (
	 "isaac/config"
	"fmt"
)

func main() {
	con, err := config.GetContent("config.ini")
	if err != nil {
		fmt.Println(err)
		return
	}
	if con.Jdbc.Driver != "com.mysql.jdbc.Driver" {
		fmt.Println("not get config values")
		return
	}


}
