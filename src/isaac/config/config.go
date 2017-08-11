package config

import "github.com/go-ini/ini"

type Jdbc struct {
	Driver   string `ini:"driver"`
	Url      string `ini:"url"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
	Table    string `ini:"table"`
}

type Pkg struct {
	Model string `ini:"model"`
	Dao   string `ini:"dao"`
	Serv  string `ini:"serv"`
	Xml   string `ini:"xml"`
}

type Out struct {
	Target string `ini:"target"`
}

type Proj struct {
	Target string `ini:"target"`
	Over   string `ini:"over"`
}

type Config struct {
	Jdbc *Jdbc
	Pkg  *Pkg
	Out  *Out
	Proj *Proj
}

func GetContent(path string) (*Config, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, err
	}
	j := new(Jdbc)
	err = cfg.Section("jdbc").MapTo(j)
	p := new(Pkg)
	err = cfg.Section("pkg").MapTo(p)
	o := new(Out)
	err = cfg.Section("out").MapTo(o)
	pj := new(Proj)
	err = cfg.Section("proj").MapTo(pj)
	c := Config{j, p, o, pj}
	return &c, err
}
