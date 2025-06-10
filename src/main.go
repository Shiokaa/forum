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

	configs.EnvInit()              // Chargement des variables d'environnements
	store := configs.SessionInit() // Chargement de la session

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

	usersServices := services.UsersServicesInit(db)       // Initialisation du service user
	topicServices := services.TopicsServicesInit(db)      // Initialisation du service topic
	messagesServices := services.MessagesServicesInit(db) // Initialisation du service message

	inscriptionController := controllers.InscriptionControllerInit(templates, usersServices, store) // Initialisation du controller inscription
	accueilController := controllers.AccueilControllerInit(templates, topicServices, store)         // Initialisation du controller accueil
	topicController := controllers.TopicControllerInit(templates, topicServices, store)             // Initialisation du controller topic
	messageController := controllers.MessageControllerInit(templates, messagesServices, store)      // Initialisation du controller message
	connexionController := controllers.ConnexionControllerInit(templates, usersServices, store)     // Initialisation du controller connexion
	errorController := controllers.ErrorControllerInit(templates)                                   // Initialisation du controller d'erreur

	router := mux.NewRouter() // Initialisation du router

	// Routage des différents controllers
	inscriptionController.InsciptionRouter(router)
	accueilController.AccueilRouter(router)
	topicController.TopicRouteur(router)
	messageController.MessageRouter(router)
	connexionController.ConnexionRouter(router)
	errorController.ErrorRouter(router)

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./assets")))) // Sert les fichiers static

	// Mise en place du serveur sur le port 8080
	log.Println("Démarrage du serveur sur http://localhost:8080 ...")
	serveErr := http.ListenAndServe(":8080", router)
	if serveErr != nil {
		log.Fatalf("Erreur lancement serveur - %v", serveErr)
	}
}
