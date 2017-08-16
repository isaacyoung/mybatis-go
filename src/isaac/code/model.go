// generate entity.java
package code

import (
	"isaac/content"
	"os"
	"fmt"
)

func Model(con *content.Content) {
	os.MkdirAll(con.GetModelPath(), os.ModeDir)

	for _, t := range con.Tables {
		className := content.GetModelName(t.Table.Name)

		file, err := os.Create(con.GetModelPath() + "/" + className +  ".java")
		if err != nil {
			fmt.Println(err)
			return
		}

		file.WriteString(con.Config.Pkg.Model + ";\n\n")

		file.WriteString("/**\n")
		file.WriteString(" * " + t.Table.Comment + "\n")
		file.WriteString(" */\n")
		file.WriteString("public class " + className + " {\n")

		for _, c := range t.Columns {
			field := content.GetFieldName(c.Name, false)
			file.WriteString("    /**\n")
			file.WriteString("     * " + c.Comment + "\n")
			file.WriteString("     */\n")
			file.WriteString("    private String " + field + ";\n\n")
		}

		for _, c := range t.Columns {
			field := content.GetFieldName(c.Name, false)
			file.WriteString("    public String " + content.GetFieldGetter(c.Name) + "() {\n")
			file.WriteString("        return this." + field + ";\n")
			file.WriteString("    }\n\n")

			file.WriteString("    public void " + content.GetFieldSetter(c.Name) + "(string _" + field + ") {\n")
			file.WriteString("        this." + field + " = _" + field + ";\n")
			file.WriteString("    }\n\n")
		}

		file.WriteString("}")
		file.Close()
	}

}
