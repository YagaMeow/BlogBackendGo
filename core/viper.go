package core

import (
	"blog-backend/core/internal"
	"blog-backend/global"
	"flag"
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitViper(path ...string) *viper.Viper {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "选择配置文件.")
		flag.Parse()
		if config == "" {
			if configEnv := os.Getenv(internal.ConfigEnv); configEnv == "" {
				switch gin.Mode() {
				case gin.DebugMode:
					config = internal.ConfigDebugFile
				case gin.TestMode:
					config = internal.ConfigTestFile
				case gin.ReleaseMode:
					config = internal.ConfigReleaseFile
				default:
					config = internal.ConfigDebugFile
				}
				fmt.Println("[Viper]未配置环境变量")
				fmt.Printf("[Viper]正在使用%s配置\n", gin.EnvGinMode)
				fmt.Printf("[Viper]配置路径%s\n", config)
			} else {
				config = configEnv
				fmt.Printf("[Viper]配置路径%s\n", config)
			}

		} else {
			fmt.Println("[Viper]正在使用命令行")
			fmt.Printf("[Viper]配置路径%s\n", config)
		}
	} else {
		config = path[0]
		fmt.Println("[Viper]正在通过传参配置")
		fmt.Printf("[Viper]配置路径%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	var err = v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("[Viper]读取配置失败 %s", err))
	}

	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("[Viper]配置文件已变更")
		if err = v.Unmarshal(&global.YAGAMI_CONFIG); err != nil {
			fmt.Printf("[Viper]解析配置失败 %s\n", err)
		}
	})
	if err = v.Unmarshal(&global.YAGAMI_CONFIG); err != nil {
		panic(fmt.Errorf("[Viper]解析配置失败 %s", err))
	}

	fmt.Println("[Viper]读取配置成功")

	return v
}
