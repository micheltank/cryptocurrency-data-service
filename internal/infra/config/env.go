package config

import (
	"github.com/Netflix/go-env"
	"log"
)

type Environment struct {
	Port           int    `env:"PORT"`
	SoChainApiHost string `env:"SO_CHAIN_API_HOST"`
	Extras         env.EnvSet
}

var Env Environment

func init() {
	es, err := env.UnmarshalFromEnviron(&Env)
	if err != nil {
		log.Fatal(err)
	}
	// Remaining environment variables.
	Env.Extras = es
}
