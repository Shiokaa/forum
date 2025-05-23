package main

import (
	"fmt"
	"forum/src/configs"
	"forum/src/controllers"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	configs.EnvInit() // Chargement des variables d'environnements

	// Lancement de la base de donnée
	db, dbErr := configs.DbInit()
	if dbErr != nil {
		log.Fatalf(" Initialisation de la DB impossible : %v ", dbErr)
	}

	defer db.Close() // Fermeture de la base de donnée une fois toute les données récupérées

	// Récupération des templates
	temp, tempErr := template.ParseGlob("./templates/*.html")
	if tempErr != nil {
		log.Fatalf(" Récupération des templates impossible : %v", tempErr)
	}

	inscriptionController := controllers.InscriptionControllerInit(temp) // Initialisation du template inscription

	router := mux.NewRouter() // Initialisation du router

	// Routage des différents controllers
	inscriptionController.InsciptionRouter(router)

	// Mise en place du serveur sur le port 3000
	serveErr := http.ListenAndServe(":3000", router)
	if serveErr != nil {
		log.Fatalf("Erreur lancement serveur - %v", serveErr)
	}
	fmt.Println("Serveur lancé : http://localhost:3000")
}
