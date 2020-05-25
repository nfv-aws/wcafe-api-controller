package config

import (
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Endpoint string
		User     string
		Password string
		Name     string
	}
	SQS struct {
		Region           string
		Pets_Queue_Url   string
		Stores_Queue_Url string
	}
	LOG struct {
		File_path string
	}
}

var C Config

func Configure() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("$GOPATH/src/github.com/nfv-aws/wcafe-api-controller/config")

	// 環境変数 export WCAFE_XXXで設定値を上書きできるように設定
	// ex) Database.Password ->  export WCAFE_DB_PASSWORD
	viper.SetEnvPrefix("wcafe")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// conf読み取り
	if err := viper.ReadInConfig(); err != nil {
		log.Println("config file read error: ", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Println("config file Unmarshal error: ", err)
		os.Exit(1)
	}
}
