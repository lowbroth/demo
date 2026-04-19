package main

import (
		"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	engine := gin.Default()
	viper.New()
	viper.SetConfigName("configs/demo.yml")
	serviceConfig := &ServerConfig{}
	viper.Unmarshal(serviceConfig)

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success", "version": 1.0, "name": serviceConfig.ServiceName})
				fmt.Println("好的")
	})
	engine.Run(":8080")
}

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
