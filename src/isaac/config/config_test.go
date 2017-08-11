package config

import "testing"

func TestGetContent(t *testing.T) {
	c := GetContent()
	if c.Jdbc.Driver != "com.mysql.jdbc.Driver" {
		t.Error("no value")
	}
}
