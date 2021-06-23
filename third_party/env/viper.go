package env

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init(env string) {
	viper.SetConfigName(env)
	viper.SetConfigType("json")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./configs") // optionally look for config in the working directory
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
}
