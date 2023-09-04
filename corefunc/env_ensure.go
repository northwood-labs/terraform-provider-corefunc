package corefunc

import (
	"errors"
	"os"
)

/*
EnvEnsure ensures that a given environment variable is set to a non-empty value.
If the environment variable is unset or if it is set to an empty string,
EnvEnsure will respond with an error.

----

* name (string): The name of the environment variable to check.
*/
func EnvEnsure(name string) error {
	if os.Getenv(name) == "" {
		return errors.New("environment variable " + name + " is not defined") // lint:allow_errorf
	}

	return nil
}
