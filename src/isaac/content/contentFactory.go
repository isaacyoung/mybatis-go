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
	ModelName string
	Columns []ColumnInfo
}

type ColumnInfo struct {
	Column jdbc.Column
	Field string
	JdbcType string
	JavaType string
	ShortJavaType string
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
		columnList := []ColumnInfo{}
		for _, c := range columns {
			if t.Name == c.TableName {
				columnInfo := ColumnInfo{}
				columnInfo.Column = c
				columnInfo.Field = GetFieldName(c.Name, false)
				columnInfo.JdbcType = columnInfo.getJdbcType()
				columnInfo.JavaType = columnInfo.getJavaType()
				columnInfo.ShortJavaType = columnInfo.getShortJavaType()
				columnList = append(columnList, columnInfo)
			}
		}
		tableInfo.Columns = columnList
		tableInfos = append(tableInfos, tableInfo)
	}
	return tableInfos, nil
}

func (c *ColumnInfo) getJdbcType() string {
	jdbcType := strings.ToUpper(c.Column.Type)
	if jdbcType == "INT" {
		return "INTEGER"
	} else if jdbcType == "DATETIME" {
		return "TIMESTAMP"
	}
	return jdbcType
}

func (c *ColumnInfo) getJavaType() string {
	jdbcType := strings.ToUpper(c.Column.Type)
	if jdbcType == "BIT" {
		return "java.lang.Byte"
	} else if jdbcType == "TINYINT" {
		return "java.lang.Byte"
	} else if jdbcType == "BOOL" {
		return "java.lang.Boolean"
	} else if jdbcType == "BOOLEAN" {
		return "java.lang.Boolean"
	} else if jdbcType == "SMALLINT" {
		return "java.lang.Integer"
	} else if jdbcType == "MEDIUMINT" {
		return "java.lang.Integer"
	} else if jdbcType == "INT" {
		return "java.lang.Integer"
	} else if jdbcType == "INTEGER" {
		return "java.lang.Integer"
	} else if jdbcType == "BIGINT" {
		return "java.lang.Long"
	} else if jdbcType == "NUMERIC" {
		return "java.math.BigDecimal"
	} else if jdbcType == "DECIMAL" {
		return "java.math.BigDecimal"
	} else if jdbcType == "DEC" {
		return "java.lang.Double"
	} else if jdbcType == "FIXED" {
		return "java.lang.Double"
	} else if jdbcType == "FLOAT" {
		return "java.lang.Double"
	} else if jdbcType == "DOUBLE" {
		return "java.lang.Double"

	} else if jdbcType == "DATE" {
		return "java.util.Date"
	} else if jdbcType == "DATETIME" {
		return "java.util.Date"
	} else if jdbcType == "TIMESTAMP" {
		return "java.util.Date"
	} else if jdbcType == "TIME" {
		return "java.util.Date"
	} else if jdbcType == "YEAR" {
		return "java.util.Date"

	} else if jdbcType == "CHAR" {
		return "java.lang.String"
	} else if jdbcType == "VARCHAR" {
		return "java.lang.String"
	} else if jdbcType == "BINARY" {
		return "java.lang.String"
	} else if jdbcType == "VARBINARY" {
		return "java.lang.String"
	} else if jdbcType == "TINYBLOB" {
		return "java.lang.String"
	} else if jdbcType == "TINYTEXT" {
		return "java.lang.String"
	} else if jdbcType == "BLOB" {
		return "java.lang.String"
	} else if jdbcType == "TEXT" {
		return "java.lang.String"
	} else if jdbcType == "MEDIUMBLOB" {
		return "java.lang.String"
	} else if jdbcType == "MEDIUMTEXT" {
		return "java.lang.String"
	} else if jdbcType == "LONGBLOB" {
		return "java.lang.String"
	} else if jdbcType == "LONGTEXT" {
		return "java.lang.String"
	} else if jdbcType == "ENUM" {
		return "java.lang.String"
	} else if jdbcType == "SET" {
		return "java.lang.String"
	}
	return "java.lang.Object"
}

func (c *ColumnInfo) getShortJavaType() string {
	fullType := c.getJavaType()
	index := strings.LastIndex(fullType, ".")
	return fullType[index+1:]
}

func (t *TableInfo) GetImportType() []string {
	result := []string{}
	for _, v := range t.Columns {
		if strings.Index(v.JavaType, "java.lang") == -1 && !contain(result, v.JavaType) {
			result = append(result, v.JavaType)
		}
	}
	return result
}

func contain(c []string, s string) bool {
	for _, v := range c {
		if v == s {
			return true
		}
	}
	return false
}