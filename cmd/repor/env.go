// DB_NAME: repor_test
// RETHINKDB_URL: localhost:28015

package main

import (
	"os"
)

// SetupEnv is a local, ignored file for setting environment variables.
func SetupEnv() {
	os.Setenv("DB_NAME", "repor_test")
	os.Setenv("RETHINKDB_URL", "localhost:28015")
}
