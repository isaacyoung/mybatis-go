// generate ServiceImpl.java
package code

import (
	"mybatis-go/content"
	"os"
	"fmt"
	"sort"
)

func ServImpl(con *content.Content) {
	os.MkdirAll(con.GetServiceImplPath(), os.ModeDir)

	for _, t := range con.Tables {
		file, err := os.Create(con.GetServiceImplPath() + "/" + t.GetServiceImplName() + ".java")
		if err != nil {
			fmt.Println(err)
			return
		}

		file.WriteString("package " + con.Config.Pkg.Serv + ".impl;\n\n")

		imports := []string{}
		imports = append(imports, "com.cdsq.manage.base.BaseServiceImp")
		imports = append(imports, "org.springframework.stereotype.Service")
		imports = append(imports, con.Config.Pkg.Model+"."+t.ModelName)
		imports = append(imports, con.Config.Pkg.Dao+"."+t.GetDaoName())
		imports = append(imports, con.Config.Pkg.Serv+"."+t.GetServiceName())
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
		file.WriteString("@Service\n")
		file.WriteString("public class " + t.GetServiceImplName() + " extends BaseServiceImp<" + t.ModelName + ", " + t.GetDaoName() + "> implements " + t.GetServiceName() + " {\n")

		file.WriteString("}")
		file.Close()
	}

}
