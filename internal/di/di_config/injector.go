package di_config

func InitializeConfig(envFilePath string) (*Config, error) {
	config, err := NewConfig(envFilePath)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
