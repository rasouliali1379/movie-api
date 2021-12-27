package config

type (
	Config struct {
		Jwt Jwt `yaml:"jwt"`
		Api Api `yaml:"api"`
	}
	Jwt struct {
		Secret string `yaml:"secret"`
	}

	Api struct {
		Port string `yaml:"port"`
	}
)
