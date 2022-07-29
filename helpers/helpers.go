package helpers

import (
	"fmt"
	"log"
	"os"
)

func GetEnvVar(envvar string, fatal bool) string {
	envVarVal, doesExist := os.LookupEnv(envvar)
	if doesExist == false {
		if fatal {
			log.Fatal(fmt.Sprintf("The Env Var %s Is Not Set !", envvar))
		}
		log.Println(fmt.Sprintf("The Env Var %s Is Not Set, using default mode!", envvar))
		return ""
	}
	return envVarVal
}
