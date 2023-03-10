package main

import (
	"log"
	"ncp/cmd"
	"ncp/color"
	"ncp/config"
	"ncp/output"
	"ncp/parser"
)

func main() {
	if len(config.RootPath) == 0 {
		var err error
		config.RootPath, err = cmd.CheckNginx()
		if err != nil {
			output.ErrorFatal(1004, err)
		}
	}
	// 从根配置开始解析，如果用户输入文件，解析用户文件，否则按照nginx -t解析输出文件
	content, err := parser.ConfigParser(config.RootPath)
	if err != nil {
		log.Fatal(err)
	}

	color.New(color.FgCyan).Println("> Generate results...")
	output.Nginx(content)
}
