package app

import (
	"database/sql"
	"fmt"
	"net/http"
)

// Get all subjects and all questions without answer
func GuestHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Guest layout
		subjects, err := FetchAllSubjects(db, 0)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		/* w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(subjects) */
		//Use html/template
		err = ParseAndExecuteTemplate("guest", subjects, w)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
