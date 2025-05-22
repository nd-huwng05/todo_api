package yml

import "os"

type YamlConfig struct {
	configFile  string
	serviceName string
	data        map[string]interface{}
	rootPath    string
}

func (y *YamlConfig) Initial(serviceName string, configFile string, rootPath string) {
	y.serviceName = serviceName
	y.rootPath = rootPath
	y.configFile = configFile

	// Initialize empty data
	y.data = make(map[string]interface{})

	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		if os.Getenv("ENV") == "test" || os.Getenv("ENV") == "" {

		}
	}
}
