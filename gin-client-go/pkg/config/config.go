package config

import (
	"gopkg.in/yaml.v3"
	"k8s.io/klog/v2"
	"os"
)

var KeyMap map[KeyName]string

type Config struct {
	Server Server
}

//1、` `表示字段的标签(tag)，tag是结构体字段的元数据，为字段附加额外的信息。
//2、标签的格式是键值对的形式，键和值之间使用冒号分隔，多个键值对之间使用空格分隔。
//3、标签通常用于描述字段的用途、序列化和反序列化、验证等。在实际开发中，
//可以使用反射机制来读取和解析结构体标签，以达到自动化处理的目的，
//比如根据标签进行 JSON 序列化和反序列化，或者进行数据验证等操作。

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// 初始化时读取配置文件
func init() {
	var config Config
	yamlFile, err := os.ReadFile("./.gin-client-go2.yaml")
	if err != nil {
		klog.Fatal(err)
		return
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		klog.Fatal(err)
		return
	}
	KeyMap = make(map[KeyName]string)
	KeyMap[ServerName] = config.Server.Name
}
