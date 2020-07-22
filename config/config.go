package config

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
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
		Users_Queue_Url  string
	}
	DynamoDB struct {
		Region string
	}
	LOG struct {
		File_path string
	}
	Conductor struct {
		Ip   string
		Port string
	}
}

var C Config

func Configure() {
	//configファイルの読み込み設定
	if os.Getenv("CONFIG_ACCESS") == "Production" {
		log.Debug().Caller().Msg("Production Config")
		viper.SetConfigName("config_production")
	} else {
		log.Debug().Caller().Msg("Local Config")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	// viper.AddConfigPath("$GOPATH/src/github.com/nfv-aws/wcafe-api-controller/config")

	// 環境変数 export WCAFE_XXXで設定値を上書きできるように設定
	// ex) Database.Password ->  export WCAFE_DB_PASSWORD
	viper.SetEnvPrefix("wcafe")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// conf読み取り
	if err := viper.ReadInConfig(); err != nil {
		log.Error().Caller().Err(err).Send()
		os.Exit(1)
	}

	if err := viper.Unmarshal(&C); err != nil {
		log.Error().Caller().Err(err).Send()
		os.Exit(1)
	}
}
