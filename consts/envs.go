package consts

import (
	"fmt"
	"os"
)

func CheckEnv() error {
	envs := []string{EnvUser, EnvPass}
	for _, v := range envs {
		if len(os.Getenv(v)) == 0 {
			return fmt.Errorf("env %s not found", v)
		}
	}
	return nil
}

func GetUser() string {
	return os.Getenv(EnvUser)
}

func GetPass() string {
	return os.Getenv(EnvPass)
}
