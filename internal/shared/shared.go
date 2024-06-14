package shared

import (
	"math"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func LoadEnv() error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	envPath := "./.env"
	if strings.Contains(exePath, "debug") || strings.Contains(exePath, "tmp") {
		log.Info("with Air")
		// envPath = "../../.env"  		DEBUG without Air
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return err
	}

	return nil
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
