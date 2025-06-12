package utilitaire

import (
	"net/http"
	"time"
)

func ConvertTime(created_at string, updated_at string, w http.ResponseWriter, r *http.Request) (string, string) {

	var str1 string
	var str2 string

	parsedCreated_at, err := time.Parse("2006-01-02 15:04:05", created_at)
	if err != nil {
		http.Redirect(w, r, "/erreur?code=invalid_time", http.StatusSeeOther)
		return "", ""
	}

	parsedUpdated_at, err := time.Parse("2006-01-02 15:04:05", created_at)
	if err != nil {
		http.Redirect(w, r, "/erreur?code=invalid_time", http.StatusSeeOther)
		return "", ""
	}

	str1 = parsedCreated_at.Format("02/01/2006 à 15:04")
	str2 = parsedUpdated_at.Format("02/01/2006 à 15:04")

	return str1, str2

}
