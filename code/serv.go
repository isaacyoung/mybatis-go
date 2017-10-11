// generate Service.java
package code

import (
	"mybatis-go/content"
	"os"
	"fmt"
	"sort"
)

func Serv(con *content.Content) {
	os.MkdirAll(con.GetServicePath(), os.ModeDir)

	for _, t := range con.Tables {
		file, err := os.Create(con.GetServicePath() + "/" + t.GetServiceName() + ".java")
		if err != nil {
			fmt.Println(err)
			return
		}

		file.WriteString("package " + con.Config.Pkg.Serv + ";\n\n")

		imports := []string{}
		imports = append(imports, "com.cdsq.manage.base.BaseService")
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
		file.WriteString("public interface " + t.GetServiceName() + " extends BaseService<" + t.ModelName + "> {\n")

		file.WriteString("}")
		file.Close()
	}

}
