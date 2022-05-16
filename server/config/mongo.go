package config

type Mongo struct {
	Uri string `mapstructure:"uri" json:"uri" yaml:"uri"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
}
