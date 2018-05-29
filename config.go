package main

type config struct {
	Port               int    `env:"DATASTASH_PORT" envDefault:"9999"`
	EurekaHost         string `env:"DATASTASH_EUREKA_HOST" envDefault:"http://localhost:8761/eureka"`
	MongoURL           string `env:"DATASTASH_MONGO_URL" envDefault:"mongodb://localhost:27017"`
	MongoAuthMechanism string `env:"DATASTASH_MONGO_AUTH_MECHANISM" envDefault:"SCRAM-SHA-1"`
	MongoUsername      string `env:"DATASTASH_MONGO_USERNAME" envDefault:"admin"`
	MongoPassword      string `env:"DATASTASH_MONGO_PASSWORD" envDefault:"admin"`
	MongoAuthSource    string `env:"DATASTASH_MONGO_AUTH_SOURCE" envDefault:"admin"`
}
