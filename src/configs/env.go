package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvInit() {
	// Charger le fichier .env
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		log.Fatalf(" Lancement du projet impossible. Erreur lors du chargement du fichier .env : %v", errEnv.Error())
	}
}

func GetEnvWithDefault(key string, defaultValue string) string {
	// Méthode pour récupérer les variables d'environnements avec une valeur par défaut
	envVar, envVarCheck := os.LookupEnv(key)
	if !envVarCheck {
		return defaultValue
	}

	return envVar
}
