package config

type Oauth struct {
	OursparkAppid   string `mapstructure:"ourspark-appid" json:"oursparkAppid" yaml:"ourspark-appid"`
	OursparkSecret  string `mapstructure:"ourspark-secret" json:"oursparkSecret" yaml:"ourspark-secret"`
	OursparkBaseUrl string `mapstructure:"ourspark-base-url" json:"oursparkBaseUrl" yaml:"ourspark-base-url"`
}
