package configs

import (
	"database/sql"
	"fmt"
	"log"
)

func DbInit() (*sql.DB, error) {

	// Chargement de toutes les variables d'environnements
	dbName := GetEnvWithDefault("DB_NAME", "")
	dbUser := GetEnvWithDefault("DB_USER", "")
	dbHost := GetEnvWithDefault("DB_HOST", "")
	dbPort := GetEnvWithDefault("DB_PORT", "")

	// Verification de la présence des variables
	if dbName == "" || dbUser == "" || dbHost == "" || dbPort == "" {
		return nil, fmt.Errorf(" Erreur connexion base de donnée. Données manquantes")
	}

	connectionString := fmt.Sprintf("%s@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName) // Création d'une string pour l'utiliser après

	var err error // Variable error pour l'utiliser lors du .Ping()

	// Ouverture de la connexion avec la base de donnée en utilisant la string au dessus
	dbContext, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf(" Erreur lors de l'ouverture de la connexion : %v", err)
	}

	// Test de connexion avec la base de donnée
	err = dbContext.Ping()
	if err != nil {
		dbContext.Close() // Fermeture de la session si erreur
		log.Fatalf(" Erreur lors de la tentative de ping : %v", err)
	}

	return dbContext, nil
}
