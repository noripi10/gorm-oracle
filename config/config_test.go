package config

import (
	"testing"
)

func TestConfigNew(t *testing.T) {
	t.Setenv("ORACLE_DB_USER", "USER")
	t.Setenv("ORACLE_DB_PASSWORD", "PASSW0RD")
	t.Setenv("ORACLE_DB_ADDRESS", "ADDRESS")
	t.Setenv("ORACLE_DB_PORT", "1521")
	t.Setenv("ORACLE_DB_SERVICE", "SERVICE")

	conf, err := New()
	if err != nil {
		t.Fatal(err)
	}

	if conf.DBUser != "USER" {
		t.Error("DBUser set error")
	}

	if conf.DBPassword != "PASSW0RD" {
		t.Error("DBPassword set error")
	}
}
