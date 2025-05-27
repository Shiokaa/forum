package main

import (
	"forum/src/configs"
	"forum/src/controllers"
	"forum/src/services"
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
	templates, tempErr := template.ParseGlob("./templates/*.html")
	if tempErr != nil {
		log.Fatalf(" Récupération des templates impossible : %v", tempErr)
	}

	usersServices := services.UsersServicesInit(db) // Initialisation du service user
	topicServices := services.TopicsServicesInit(db)

	inscriptionController := controllers.InscriptionControllerInit(templates, usersServices) // Initialisation du controller inscription
	accueilController := controllers.AccueilControllerInit(templates, topicServices)         // Initialisation du controller accueil

	router := mux.NewRouter() // Initialisation du router

	// Routage des différents controllers
	inscriptionController.InsciptionRouter(router)
	accueilController.AccueilRouter(router)

	// Sert les fichiers static
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets"))))

	// Mise en place du serveur sur le port 8080
	log.Println("Démarrage du serveur sur http://localhost:8080 ...")
	serveErr := http.ListenAndServe(":8080", router)
	if serveErr != nil {
		log.Fatalf("Erreur lancement serveur - %v", serveErr)
	}
}
