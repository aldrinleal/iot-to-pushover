package util

import "os"

func EnvIf(args ...string) string {
	result := args[-1+len(args)]

	for _, envVarName := range args[0 : -1+len(args)] {
		if newValue, ok := os.LookupEnv(envVarName); ok {
			return newValue
		}
	}

	return result
}


func IsRunningOnLambda() bool {
	return os.Getenv("_LAMBDA_SERVER_PORT") != ""
}