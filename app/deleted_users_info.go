package app

import (
	"database/sql"
	"log"
)

func CopyUsersToDeletedUsers(db *sql.DB) {
	// Copy users to DeletedUsers table with deleted set to FALSE and update the row if the user already exists
	copyQuery := `INSERT INTO DeletedUsers (
		id_student, deleted, delete_date, lastname, firstname, username, email, avatar, birth_date, bio, website, github, xp, rang_rank_, school_year, creation_date, update_date
	)
	SELECT 
		id_student, FALSE, NULL, lastname, firstname, username, email, avatar, birth_date, bio, website, github, xp, rang_rank_, school_year, creation_date, update_date
	FROM users
	ON CONFLICT (id_student) DO UPDATE SET
		deleted = EXCLUDED.deleted,
		delete_date = EXCLUDED.delete_date,
		lastname = EXCLUDED.lastname,
		firstname = EXCLUDED.firstname,
		username = EXCLUDED.username,
		email = EXCLUDED.email,
		avatar = EXCLUDED.avatar,
		birth_date = EXCLUDED.birth_date,
		bio = EXCLUDED.bio,
		website = EXCLUDED.website,
		github = EXCLUDED.github,
		xp = EXCLUDED.xp,
		rang_rank_ = EXCLUDED.rang_rank_,
		school_year = EXCLUDED.school_year,
		creation_date = EXCLUDED.creation_date,
		update_date = EXCLUDED.update_date;`

	if _, err := db.Exec(copyQuery); err != nil {
		log.Fatal(err)
	}
}
