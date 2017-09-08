// generate Mapper.xml
package code

import (
	"isaac/content"
	"os"
	"fmt"
)

func Xml(con *content.Content) {
	os.MkdirAll(con.GetXmlPath(), os.ModeDir)

	for _, t := range con.Tables {
		file, err := os.Create(con.GetXmlPath() + "/" + t.GetDaoName() + ".xml")
		if err != nil {
			fmt.Println(err)
			return
		}

		file.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
		file.WriteString(`<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd">` + "\n")
		file.WriteString(`<mapper namespace="` + con.Config.Pkg.Dao + `.` + t.GetDaoName() + `">` + "\n")

		// BaseResultMap
		file.WriteString(`  <resultMap id="BaseResultMap" type="` + con.Config.Pkg.Model + `.` + t.ModelName + `">` + "\n")
		for _, c := range t.Columns {
			if c.IsKey() {
				file.WriteString(`    <id column="` + c.Column.Name + `" jdbcType="` + c.JdbcType + `" property="` + c.Field + `" />` + "\n")
			} else {
				file.WriteString(`    <result column="` + c.Column.Name + `" jdbcType="` + c.JdbcType + `" property="` + c.Field + `" />` + "\n")
			}
		}
		file.WriteString(`  </resultMap>` + "\n\n")

		// selectById
		file.WriteString(`  <select id="selectById" parameterType="java.lang.Integer" resultMap="BaseResultMap">` + "\n")
		file.WriteString(`    select * from ` + t.Table.Name + "\n")
		keys := t.GetKeys()
		for i, v := range keys {
			if i == 0 {
				file.WriteString("    where ")
			} else {
				file.WriteString(" and ")
			}
			file.WriteString(v.Column.Name + " = #{" + v.Field + ",jdbcType=" + v.JdbcType + "}")
		}
		file.WriteString("\n")
		file.WriteString(`  </select>` + "\n\n")

		// deleteById
		file.WriteString(`  <delete id="deleteById" parameterType="java.lang.Integer">` + "\n")
		file.WriteString(`    delete from ` + t.Table.Name + "\n")
		for i, v := range keys {
			if i == 0 {
				file.WriteString("    where ")
			} else {
				file.WriteString(" and ")
			}
			file.WriteString(v.Column.Name + " = #{" + v.Field + ",jdbcType=" + v.JdbcType + "}")
		}
		file.WriteString("\n")
		file.WriteString(`  </delete>` + "\n\n")

		// select
		file.WriteString(`  <select id="select" parameterType="` + con.Config.Pkg.Model + `.` + t.ModelName + `" resultMap="BaseResultMap">` + "\n")
		file.WriteString(`    select ` + "\n")
		file.WriteString(`    a.*` + "\n")
		file.WriteString(`    from ` + t.Table.Name + " a\n")
		file.WriteString(`    <trim prefix="where" prefixOverrides="and|or">` + "\n")
		for i, c := range t.Columns {
			if c.IsString() {
				file.WriteString(`      <if test="` + c.Field + ` != null and ` + c.Field + ` !=  ''">` + "\n")
			} else if c.IsDate() {
				file.WriteString(`      <if test="` + c.Field + ` != null ">` + "\n")
			} else {
				file.WriteString(`      <if test="` + c.Field + ` != null and ` + c.Field + ` !=  -1 ">` + "\n")
			}
			if i == 0 {
				file.WriteString(`        `)
			} else {
				file.WriteString(`        and `)
			}
			file.WriteString(`a.` + c.Column.Name + " = #{" + c.Field + ",jdbcType=" + c.JdbcType + "}" + "\n")
			file.WriteString(`      </if>` + "\n")
		}

		str := `      <if test="sqlCondition != null and sqlCondition.size > 0">
        and
        <foreach collection="sqlCondition" index="index" item="condition" separator=" and ">
          ${condition}
        </foreach>
      </if>`
		file.WriteString(str + "\n")
		file.WriteString(`    </trim>` + "\n")
		file.WriteString(`  </select>` + "\n\n")

		// insert
		file.WriteString(`  <insert id="insert" keyProperty="id" parameterType="` + con.Config.Pkg.Model + `.` + t.ModelName + `" useGeneratedKeys="true">` + "\n")
		sql := getInsertSql(t.Columns)
		file.WriteString(`    insert into ` + t.Table.Name + " (" + sql[0] + ")\n")
		file.WriteString(`    values` + "\n")
		file.WriteString(`    <foreach collection="list" index="index" item="item" separator=",">` + "\n")
		file.WriteString("       (" + sql[1] + ")\n")
		file.WriteString(`    </foreach>` + "\n")
		file.WriteString(`  </insert>` + "\n\n")

		// update
		file.WriteString(`  <update id="update" parameterType="` + con.Config.Pkg.Model + `.` + t.ModelName + `">` + "\n")
		file.WriteString(`    <foreach close="" collection="list" index="index" item="item" open="" separator=";">` + "\n")
		file.WriteString(`      update ` + t.Table.Name + "\n")
		file.WriteString(`      <set>` + "\n")
		for i, c := range t.Columns {
			if c.IsKey() {
				continue
			}
			if c.IsString() {
				file.WriteString(`        <if test="item.` + c.Field + ` != null and item.` + c.Field + ` !=  ''">` + "\n")
			} else if c.IsDate() {
				file.WriteString(`        <if test="item.` + c.Field + ` != null ">` + "\n")
			} else {
				file.WriteString(`        <if test="item.` + c.Field + ` != null and item.` + c.Field + ` !=  -1 ">` + "\n")
			}

			file.WriteString(`          ` + c.Column.Name + " = #{item." + c.Field + ",jdbcType=" + c.JdbcType + "}")
			if i != len(t.Columns) -1 {
				file.WriteString(",")
			}
			file.WriteString("\n")
			file.WriteString(`        </if>` + "\n")
		}
		file.WriteString(`      </set>` + "\n")
		for i, v := range keys {
			if i == 0 {
				file.WriteString("      where ")
			} else {
				file.WriteString(" and ")
			}
			file.WriteString(v.Column.Name + " = #{item." + v.Field + ",jdbcType=" + v.JdbcType + "}")
		}
		file.WriteString("\n")

		file.WriteString(`    </foreach>` + "\n")
		file.WriteString(`  </update>` + "\n\n")

		// delete
		file.WriteString(`  <delete id="delete" parameterType="` + con.Config.Pkg.Model + `.` + t.ModelName + `">` + "\n")
		file.WriteString(`    delete from ` + t.Table.Name + "\n")
		file.WriteString(`    <trim prefix="where" prefixOverrides="and|or">` + "\n")
		for i, c := range t.Columns {
			if c.IsString() {
				file.WriteString(`      <if test="` + c.Field + ` != null and ` + c.Field + ` !=  ''">` + "\n")
			} else if c.IsDate() {
				file.WriteString(`      <if test="` + c.Field + ` != null ">` + "\n")
			} else {
				file.WriteString(`      <if test="` + c.Field + ` != null and ` + c.Field + ` !=  -1 ">` + "\n")
			}

			if i == 0 {
				file.WriteString(`        `)
			} else {
				file.WriteString(`        and `)
			}
			file.WriteString(c.Column.Name + " = #{" + c.Field + ",jdbcType=" + c.JdbcType + "}\n")

			file.WriteString(`      </if>` + "\n")
		}

		file.WriteString(`    </trim>` + "\n")
		file.WriteString(`  </delete>` + "\n\n")

		file.WriteString(`</mapper>`)
		file.Close()
	}
}

func getInsertSql(columns []content.ColumnInfo) []string {
	result := []string{}
	list := ""
	values := ""
	temp := ""
	for i, v := range columns {
		if len(temp) > 80 {
			list += "\n      "
			values += "\n        "
			temp = ""
		}
		list += v.Column.Name
		values += "#{item." + v.Field + ",jdbcType=" + v.JdbcType + "}"
		temp += "#{item." + v.Field + ",jdbcType=" + v.JdbcType + "}"

		if i != len(columns)-1 {
			list += ", "
			values += ", "
			temp += ", "
		}
	}
	result = append(result, list)
	result = append(result, values)
	return result
}
