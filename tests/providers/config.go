package providers

type Config struct {
	Database *Database
}

type Database struct {
	Adapter  string
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Charset  string
}

func (config *Config) Register() {
	config.Database = &Database{
		Adapter:  "Mysql",
		Host:     "127.0.0.1",
		Port:     "3306",
		Username: "root",
		Password: nil,
		Dbname:   "phalcon",
		Charset:  "utf8",
	}
}

func (config *Config) GetInstance() *Config {
	return config
}
