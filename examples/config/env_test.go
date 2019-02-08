package config

import (
	"log"
	"testing"
	"webservice-library/config"
)

func TestEnv(t *testing.T) {
	expected := "{{env env_test.ini} postgres 10 true Pacific/Auckland}"
	env := &Env{}
	if err := env.SetPath("env_test.ini"); err != nil {
		log.Fatal(err)
	}

	if err := env.SetPrefix("env"); err != nil {
		log.Fatal(err)
	}

	if err := config.LoadConfig(env); err != nil {
		log.Fatal(err)
	}

	if env.DbEngine != "postgres" ||
		env.MaxConnection != 10 ||
		!env.Production ||
		env.TimeZone.String() != "Pacific/Auckland" {
		t.Errorf("expect=%s\nactual=%v", expected, *env)
	}
}
