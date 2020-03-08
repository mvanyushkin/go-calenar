package main

type Config struct {
	HttpListen string `config:"httplisten"`
	LogFile    string `config:"logfile"`
	LogLevel   string `config:"loglevel"`
}
