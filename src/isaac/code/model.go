// generate entity.java
package code

import (
	"isaac/content"
	"os"
	"fmt"
	"sort"
)

func Model(con *content.Content) {
	os.MkdirAll(con.GetModelPath(), os.ModeDir)

	for _, t := range con.Tables {
		file, err := os.Create(con.GetModelPath() + "/" + t.ModelName +  ".java")
		if err != nil {
			fmt.Println(err)
			return
		}

		file.WriteString("package " + con.Config.Pkg.Model + ";\n\n")

		imports := t.GetImportType()
		imports = append(imports, "java.util.ArrayList")
		imports = append(imports, "java.util.List")
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
		file.WriteString("public class " + t.ModelName + " {\n")

		for _, c := range t.Columns {
			file.WriteString("    /**\n")
			file.WriteString("     * " + c.Column.Comment + "\n")
			file.WriteString("     */\n")
			file.WriteString("    private " + c.ShortJavaType + " " + c.Field + ";\n\n")
		}

		file.WriteString("    /**\n")
		file.WriteString("     * 自定义条件\n")
		file.WriteString("     */\n")
		file.WriteString("    private List sqlCondition;\n\n")

		for _, c := range t.Columns {
			file.WriteString("    public " + c.ShortJavaType + " " + c.GetFieldGetter() + "() {\n")
			file.WriteString("        return this." + c.Field + ";\n")
			file.WriteString("    }\n\n")

			file.WriteString("    public void " + c.GetFieldSetter() + "(" + c.ShortJavaType + " _" + c.Field + ") {\n")
			file.WriteString("        this." + c.Field + " = _" + c.Field + ";\n")
			file.WriteString("    }\n\n")
		}

		str := `	public List getSqlCondition() {
        return sqlCondition;
    }

    public void setSqlCondition(List _sqlCondition) {
        this.sqlCondition = _sqlCondition;
    }

    public void addCondition(String condition) {
        if(sqlCondition == null || sqlCondition.isEmpty()) {
             sqlCondition = new ArrayList();
        }
        sqlCondition.add(condition);
    }
		`

		file.WriteString(str + "\n")

		file.WriteString("}")
		file.Close()
	}

}
