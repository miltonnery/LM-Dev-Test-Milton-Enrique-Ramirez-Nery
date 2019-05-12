package model

type Datasource struct {
	Postgres struct {
		Driver  string `yaml: "driver"`
		Host    string `yaml: "host"`
		Port    string `yaml: "port"`
		User    string `yaml: "user"`
		Pass    string `yaml: "pass"`
		Dbname  string `yaml: "dbname"`
		Sslmode string `yaml: "sslmode"`
		Schema  string `yaml: "schema"`
	} `yaml: "postgres"`
}
