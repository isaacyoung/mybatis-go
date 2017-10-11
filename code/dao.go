// generate Mapper.java
package code

import (
	"mybatis-go/content"
	"os"
	"fmt"
	"sort"
)

func Dao(con *content.Content) {
	os.MkdirAll(con.GetDaoPath(), os.ModeDir)

	for _, t := range con.Tables {
		file, err := os.Create(con.GetDaoPath() + "/" + t.GetDaoName() + ".java")
		if err != nil {
			fmt.Println(err)
			return
		}

		file.WriteString("package " + con.Config.Pkg.Dao + ";\n\n")

		imports := []string{}
		imports = append(imports, "com.cdsq.manage.base.BaseMapper")
		imports = append(imports, con.Config.Pkg.Model+"."+t.ModelName)
		sort.Strings(imports)
		if imports != nil {
			for _, v := range imports {
				file.WriteString("import " + v + ";\n")
			}
			file.WriteString("\n")
		}

		file.WriteString("/**\n")
		file.WriteString(" * " + t.Table.Comment + "\n")
		file.WriteString(" */\n")
		file.WriteString("public interface " + t.GetDaoName() + " extends BaseMapper<" + t.ModelName + "> {\n")

		file.WriteString("}")
		file.Close()
	}

}
