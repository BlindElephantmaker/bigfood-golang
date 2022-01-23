package postgres

type Config struct {
	host     string
	port     string
	username string
	password string
	dbName   string
	sslMode  string
}

func NewConfig(host, port, username, password, dbName, sslMode string) *Config {
	return &Config{
		host:     host,
		port:     port,
		username: username,
		password: password,
		dbName:   dbName,
		sslMode:  sslMode,
	}
}
