package tools

type MysqlConfig struct {
	Username     string
	Password     string
	Path         string
	Dbname       string
	MaxIdleConns int
	MaxOpenConns int
}

func GetMysqlConfig() (m MysqlConfig) {
	return MysqlConfig{
		Username:     "root",
		Password:     "root",
		Path:         "127.0.0.1:3306",
		Dbname:       "test",
		MaxIdleConns: 200,
		MaxOpenConns: 100,
	}
}
