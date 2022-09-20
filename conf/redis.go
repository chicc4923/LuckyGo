package conf

type RdsConfig struct {
	Host      string
	Port      int
	User      string
	Pwd       string
	isRunning bool
}

var RdsConfigList = []RdsConfig{
	{
		Host:      "127.0.0.1",
		Port:      3306,
		User:      "root",
		Pwd:       "123456",
		isRunning: true,
	},
}

var RdsMaster RdsConfig = RdsConfigList[0]
