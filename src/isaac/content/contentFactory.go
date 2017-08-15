// config and tables info
package content

import (
	"fmt"
	"isaac/jdbc"
	"isaac/config"
)

type TableInfo struct {
	Table jdbc.Table
	Columns []jdbc.Column
}

type Content struct {
	Config *config.Config
	Tables []TableInfo
}

func Build(path string) (*Content, error) {

	// config
	con, err := config.GetConfig("config.ini")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	content := new(Content)
	content.Config = con

	// jdbc
	tableInfos, err := getTables(con, content)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	content.Tables = tableInfos

	return content, nil
}

func getTables(config *config.Config, content *Content) ([]TableInfo, error) {
	tables, err := jdbc.GetTables(config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	columns, err := jdbc.GetColumns(config, config.Jdbc.Table)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	tableInfos := []TableInfo{}

	for _, t := range tables {
		tableInfo := TableInfo{}
		tableInfo.Table = t
		columnList := []jdbc.Column{}
		for _, c := range columns {
			if t.Name == c.TableName {
				columnList = append(columnList, c)
			}
		}
		tableInfo.Columns = columnList
		tableInfos = append(tableInfos, tableInfo)
	}
	return tableInfos, nil
}