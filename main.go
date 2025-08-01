package main

import (
	"flag"
	"strconv"
)

var (
	buildVersion = "0.0.1" // buildVersion 构建版本号
	gitBranch    = "dev"   // gitBranch git 分支
	gitHash      = "debug" // gitHash git 提交点哈希值
	release      string    // release 发布模式 true/false
	buildTime    string    // buildTime 构建时间戳
)

// configDir 自定义配置目录
var configDir = flag.String("conf", "./configs", "config directory, eg: -conf /configs/")

func getBuildRelease() bool {
	v, _ := strconv.ParseBool(release)
	return v
}

func main() {
	flag.Parse()
}
