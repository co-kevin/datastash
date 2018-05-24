package main

type config struct {
	Port       int    `env:"DATASTASH_PORT" envDefault:"9999"`
	EurekaHost string `env:"DATASTASH_EUREKA_HOST" envDefault:"http://localhost:8761/eureka"`
	MongoURL   string `env:"DATASTASH_MONGO_URL" envDefault:"mongodb://localhost:27017"`
}
