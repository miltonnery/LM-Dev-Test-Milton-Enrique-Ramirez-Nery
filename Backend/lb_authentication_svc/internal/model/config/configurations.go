package model

type Configurations struct {
	Server      Server     `yaml: "server"`
	Datasource  Datasource `yaml: "datasource"`
}
