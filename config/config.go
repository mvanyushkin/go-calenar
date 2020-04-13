package config

type Config struct {
	HttpListen       string `config:"httplisten"`
	LogFile          string `config:"logfile"`
	LogLevel         string `config:"loglevel"`
	ConnectionString string `config:"connectionstring"`
	RabbitMQ         string `config:"rabbitmq"`
}
