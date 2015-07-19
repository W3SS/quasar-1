package quasar

type Config struct {
	Name           string `json:"name"`
	Version        string `json:"version"`
	UUID           string `json:"uuid"`
	Description    string `json:"description"`
	ServiceChannel string `json:"service_channel"`
	RobotChannel   string `json:"robot_channel"`
	TenyksChannel  string `json:"tenyks_channel"`
	Redis          RedisConfig
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}

func GetConfig() *Config {
	return &Config{}
}
