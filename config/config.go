package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/liushuochen/gotable"
	"sync"
)

// Config 配置文件
type Config struct {
	Port  string `toml:"port"`
	Mysql struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
		User string `toml:"user"`
		Pwd  string `toml:"pwd"`
		Db   string `toml:"db"`
	} `toml:"mysql"`
}

var once sync.Once

func LoadConfig() *Config {
	var config *Config
	once.Do(func() {
		conf, err := ReadConf(`./config.toml`, config)
		if err != nil {
			panic("配置文件读取失败:" + err.Error())
		}
		config = conf
	})

	go printTable(config)

	return config

}

func ReadConf(filename string, config *Config) (*Config, error) {
	if _, err := toml.DecodeFile(filename, &config); err != nil {
		return nil, err
	}
	return config, nil
}

func printTable(config *Config) {
	fmt.Println(`
__   _______   _______     __
 \ \ / /  __ \ / ____\ \   / /
  \ V /| |  | | (___  \ \_/ / 
   > < | |  | |\___ \  \   /  
  / . \| |__| |____) |  | |   
 /_/ \_\_____/|_____/   |_|`)
	table, err := gotable.Create("category", "description")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}

	table.AddRow([]string{"Web Port", config.Port})
	table.AddRow([]string{"Mysql Port", config.Mysql.Port})

	fmt.Println(table)
}
