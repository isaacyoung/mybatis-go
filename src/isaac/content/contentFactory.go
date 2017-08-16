// config and tables info
package content

import (
	"fmt"
	"isaac/jdbc"
	"isaac/config"
	"os"
	"strings"
)

type TableInfo struct {
	Table jdbc.Table
	Columns []jdbc.Column
}

type Content struct {
	Config *config.Config
	Tables []TableInfo
}

func (c *Content) ClearTarget() error {
	return os.RemoveAll(c.GetTargetPath())
}

func (c *Content) GetModelPath() string {
	return c.GetTargetPath() + packageToPath(c.Config.Pkg.Model)
}

var jdbcFlag = "_@$# /&"

func GetModelName(tableName string) string {
	return GetFieldName(tableName, true)
}

func GetFieldName(name string, fileUpper bool) string {
	var result string
	var nextUpperCase = false
	for _, v := range name {
		if strings.ContainsRune(jdbcFlag, v) {
			nextUpperCase = true
		} else if nextUpperCase {
			result += strings.ToUpper(string(v))
			nextUpperCase = false
		} else {
			result += string(v)
		}
	}
	if fileUpper {
		result = strings.ToUpper(result[0:1]) + result[1:]
	}
	return result
}

func GetFieldGetter(name string) string {
	return "get" + GetModelName(name)
}

func GetFieldSetter(name string) string {
	return "set" + GetModelName(name)
}

func packageToPath(pkg string) string {
	return strings.Replace(pkg, ".", "/", -1)
}

func (c *Content) GetTargetPath() string {
	return c.Config.Out.Target
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