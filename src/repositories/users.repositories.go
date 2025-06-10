package repositories

import (
	"database/sql"
	"errors"
	"fmt"
	"forum/src/models"

	"golang.org/x/crypto/bcrypt"
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

func (r *UsersRepositories) ConnectUser(email string, password string) (models.Users, error) {
	var user models.Users

	query := `
	SELECT user_id, password
	FROM users
	WHERE email = ?
	`

	sqlErr := r.db.QueryRow(query, email).Scan(&user.User_id, &user.Password)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Users{}, nil
		}
		return models.Users{}, fmt.Errorf(" Erreur récupération item - Erreur : \n\t %s", sqlErr.Error())
	}

	hashedErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if hashedErr != nil {
		return models.Users{}, errors.New("identifiants invalides")
	}

	return user, nil
}

func (r *UsersRepositories) GetUserById(id int) (models.Users, error) {
	var user models.Users

	query := `
	SELECT user_id, role_id, name, email, password, created_at, updated_at
	FROM users
	WHERE user_id = ?
	`

	sqlErr := r.db.QueryRow(query, id).Scan(&user.User_id, &user.Role_id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Users{}, nil
		}
		return models.Users{}, fmt.Errorf(" Erreur récupération item - Erreur : \n\t %s", sqlErr.Error())
	}

	return user, nil
}
