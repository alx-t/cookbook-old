package user

import (
	"database/sql"
	"github.com/alx-t/cookbook/backend/models"
)

func ReadOne(db *sql.DB, id int) (models.User, error) {
	var rec models.User
	row := db.QueryRow("SELECT id, nick_name, email, created_at, updated_at FROM users WHERE id=$1 ORDER BY id", id)
	return rec, row.Scan(&rec.Id, &rec.NickName, &rec.Email, &rec.CreatedAt, &rec.UpdatedAt)
}
