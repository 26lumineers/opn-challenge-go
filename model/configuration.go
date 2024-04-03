package model

type Configuration struct{
	App	AppConfiguration
}
type AppConfiguration struct{
	PublicKey string `mapstructure:"pub_key"`
	SecretKey string `mapstructure:"secret_key"`
}