package config

import (
	_ "github.com/go-sql-driver/mysql"

	"math/rand"
)

type Config struct {
	HttpPort int    `json:"http_port"`
	FsPort   int    `json:"fs_port"`
	Mysql    string `json:"mysql"`
	Redis    struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	} `json:"redis"`
	Object4d []string `json:"object4d"`
}

//RandomElement根据传入的数组,随机返回一个数组中的元素
func (c Config)RandomElement() (string ){
	if len(c.Object4d) == 0 {
		panic("传入数组不能为空")
	}
	return c.Object4d[rand.Intn(len(c.Object4d) - 1)]
}