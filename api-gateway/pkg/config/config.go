package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type Config struct {
	Port         string `mapstructure:"PORT"`
	JobQuestAuth string `mapstructure:"JobQuest_Auth"`
	JobQuestJob  string `mapstructure:"JobQuest_Job"`
	ChatSvcUrl   string `mapstructure:"CHAT_SVC_URL"`
	JobQuestFollow string `mapstructure:"JobQuest_follow"`

	KafkaPort  string `mapstructure:"KAFKA_PORT"`
	KafkaTopic string `mapstructure:"KAFKA_TOPIC"`

	Email    string `mapstructure:"EMAIL"`
	Password string `mapstructure:"PASSWORD"`
}

var envs = []string{
	"PORT", "JobQuest_Auth", "JobQuest_Job", "EMAIL", "PASSWORD",
}

func LoadConfig() (Config, error) {
	var config Config

	viper.AddConfigPath("./")
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			return config, err
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil

}


type LinkedIn struct {
	LinkedInLoginConfig oauth2.Config
}

var AppConfig LinkedIn

func OauthSetup() oauth2.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	AppConfig.LinkedInLoginConfig = oauth2.Config{
		RedirectURL:  "https://oauth.pstmn.io/v1/callback",
		ClientID:     os.Getenv("LINKEDIN_CLIENT_ID"),
		ClientSecret: os.Getenv("LINKEDIN_CLIENT_SECRET"),
		Scopes: []string{"openid","profile","email"},
		Endpoint: oauth2.Endpoint{ // LinkedIn OAuth2 endpoints
			AuthURL:  "https://www.linkedin.com/oauth/v2/authorization",
			TokenURL: "https://www.linkedin.com/oauth/v2/accessToken",
		},
	}

	return AppConfig.LinkedInLoginConfig
}