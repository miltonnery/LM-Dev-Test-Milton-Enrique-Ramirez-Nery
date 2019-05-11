package config

type Configurations struct {
	Server Server `yaml: "server"`
	Uris   []Uri  `yaml: "uris"`
}
