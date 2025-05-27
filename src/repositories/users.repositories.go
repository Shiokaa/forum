package repositories

import (
	"database/sql"
	"fmt"
	"forum/src/models"
)

// Structure permettant l'injection de la base de donnée
type UsersRepositories struct {
	db *sql.DB
}

// Fonction pour initialiser le repositorie de user avec l'injection de la base de donnée
func UsersRepositoriesInit(db *sql.DB) *UsersRepositories {
	return &UsersRepositories{db}
}

// Fonction permettant d'initialiser la création de l'utilisateur en récupérant les colonnes de la base de donnée
func (r *UsersRepositories) CreateUser(user models.Users) (int, error) {
	query := "INSERT INTO `users`(`role_id`, `name`, `email`, `password`) VALUES (?,?,?,?);" // Query pour insérer des valeurs dans une table

	// Utilisation de la query en remplaçant les valeurs par celles à injecter
	sqlResult, sqlErr := r.db.Exec(query,
		user.Role_id,
		user.Name,
		user.Email,
		user.Password,
	)
	if sqlErr != nil {
		return -1, fmt.Errorf(" Erreur ajout utilisateur - Erreur : \n\t %s", sqlErr.Error())
	}

	// Récupération du dernier ID, cela permet de savoir si on a bien ajouté un utilisateur ou non
	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf(" Erreur ajout utilisateur - Erreur récupération identifiant : \n\t %s", idErr.Error())
	}

	return int(id), nil
}
