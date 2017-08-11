package jdbc

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"isaac/config"
	"fmt"
)

type Connect struct {
	db *sql.DB
}

func (c *Connect) connect(config *config.Config) error {
	var err error
	c.db, err = sql.Open("mysql",
		config.Jdbc.UserName+":"+config.Jdbc.Password+"@/"+config.GetDataBaseFromUrl())
	if err != nil {
		return err
	}
	return nil
}

func (c *Connect) close() error {
	return c.db.Close()
}

type Table struct {
	Name    string
	Comment string
}

func GetTables(config *config.Config) ([]Table, error) {
	if config.Jdbc.Table != "" && config.Jdbc.Table != "%" {
		return getTable(config)
	}
	str := "SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.`TABLES` WHERE TABLE_SCHEMA=? AND TABLE_TYPE='BASE TABLE'"
	return getTables(config, str, config.GetDataBaseFromUrl(), "")
}

func getTable(config *config.Config) ([]Table, error) {
	str := "SELECT TABLE_NAME,TABLE_COMMENT FROM information_schema.`TABLES` WHERE TABLE_SCHEMA=? AND TABLE_TYPE='BASE TABLE' AND TABLE_NAME=?"
	return getTables(config, str, config.GetDataBaseFromUrl(), config.Jdbc.Table)
}

func getTables(config *config.Config, q string, v string, v2 string) ([]Table, error) {
	connect := new(Connect)
	err := connect.connect(config)
	defer connect.close()
	if err != nil {
		return nil, err
	}

	var rows *sql.Rows

	if v2 != "" {
		rows, err = connect.db.Query(q, v, v2)
	} else {
		rows, err = connect.db.Query(q, v)
	}

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := []Table{}

	for rows.Next() {
		var (
			name    string
			comment string
		)
		rows.Scan(&name, &comment)
		result = append(result, Table{name, comment})
	}
	return result, nil
}

type Column struct {
	TableName string
	Name      string
	Key       bool
	Type      string
	Scale     int
	Length    int64
	Comment   string
}

func GetColumns(config *config.Config) ([]Column, error) {
	str := "SELECT TABLE_NAME,COLUMN_NAME,COLUMN_KEY,DATA_TYPE,IFNULL(NUMERIC_SCALE,0) AS NUMERIC_SCALE,IFNULL(CHARACTER_MAXIMUM_LENGTH,0) AS CHARACTER_MAXIMUM_LENGTH,COLUMN_COMMENT FROM information_schema.`COLUMNS` WHERE TABLE_SCHEMA=? AND TABLE_NAME=?"
	return getColumns(config, str, config.GetDataBaseFromUrl(),config.Jdbc.Table)
}

func getColumns(config *config.Config, q string, v string, v2 string) ([]Column, error) {
	connect := new(Connect)
	err := connect.connect(config)
	defer connect.close()
	if err != nil {
		return nil, err
	}
	fmt.Println(q)
	fmt.Println(v)

	rows, err := connect.db.Query(q, v, v2)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	result := []Column{}

	for rows.Next() {
		var (
			tableName string
			name      string
			key       string
			tp        string
			scale     int
			length    int64
			comment   string
		)
		err := rows.Scan(&tableName, &name, &key, &tp, &scale, &length, &comment)
		if err != nil {
			return nil, err
		}
		result = append(result, Column{tableName, name, key == "PRI", tp, scale, length, comment})
	}
	return result, nil
}
