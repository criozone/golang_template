package helper

import (
	"os"
	"path/filepath"
)

func GetExecPath() (string, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return "", err
	}

	return dir, nil
}

func Env(envKey, defaultVal string) string {
	val, ok := os.LookupEnv(envKey)
	if !ok {
		val = defaultVal
	}

	return val
}
