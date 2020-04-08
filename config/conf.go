package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// 構造体定義
type Config struct {
	User     string `json:"user"`
	Pass     string `json:"pass"`
	Endpoint string `json:"endpoint"`
	Dbname   string `json:"dbname"`
}

func ReadDbConfig() *Config {

	jsonString, err := ioutil.ReadFile("config/db_conf.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// config.jsonの中身を出力
	// そのままだとbyteデータになるのでstringに変換
	log.Println(string(jsonString))

	// 設定変数用意
	c := new(Config)

	// 設定
	err = json.Unmarshal(jsonString, c)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	return c

}
