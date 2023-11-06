package config

type listen struct {
	port string `json:"port"`
}

type mongo struct {
	dsn string `json:"dsn"`
	db  string `json:"db"`
}
type config struct {
	listen listen `json:"listen"`
	mongo  mongo  `json:"mongo"`
}

func Getconfig() {

}
