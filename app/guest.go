package app

import (
	"database/sql"
	"html/template"
	"net/http"
)

// Get all subjects and all questions without answer
func GuestHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Guest layout
		subjects, err := FetchAllSubjects(db, 0)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		/* w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(subjects) */
		//Use html/template
		w.Header().Set("Content-Type", "text/html")
		tmpl, err := template.ParseFiles("templates/guest.html")
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, subjects)
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
	}
}
