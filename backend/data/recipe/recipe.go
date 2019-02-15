package recipe

import (
	"database/sql"
	"github.com/alx-t/cookbook/backend/models"
)

func ReadOne(db *sql.DB, id int) (models.Recipe, error) {
	var rec models.Recipe
	row := db.QueryRow("SELECT id, user_id, title, description, created_at, updated_at FROM recipes WHERE id=$1 ORDER BY id", id)
	return rec, row.Scan(&rec.Id, &rec.UserId, &rec.Title, &rec.Description, &rec.CreatedAt, &rec.UpdatedAt)
}

func Read(db *sql.DB, str string) ([]models.Recipe, error) {
	var rows *sql.Rows
	var err error

	if str != "" {
		rows, err = db.Query("SELECT * FROM recipes WHERE name LIKE $1 ORDER BY id",
			"%"+str+"%")
	} else {
		rows, err = db.Query("SELECT id, user_id, title, description, created_at, updated_at FROM recipes ORDER BY id")
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rs = make([]models.Recipe, 0)
	var rec models.Recipe
	for rows.Next() {
		if err = rows.Scan(&rec.Id, &rec.UserId, &rec.Title, &rec.Description, &rec.CreatedAt, &rec.UpdatedAt); err != nil {
			return nil, err
		}
		rs = append(rs, rec)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return rs, nil
}

func Insert(db *sql.DB, rcp *models.Recipe) (sql.Result, error) {
	return db.Exec("INSERT INTO recipes (user_id, title, description) VALUES (default, $1, $2, $3)",
		rcp.UserId, rcp.Title, rcp.Description)
}

func Remove(db *sql.DB, id int) (sql.Result, error) {
	return db.Exec("DELETE FROM recipes WHERE id=$1", id)
}

func Update(db *sql.DB, id int, rcp *models.Recipe) (sql.Result, error) {
	return db.Exec("UPDATE recipes SET user_id = $1, title = $2, description = $3, updated_at = now() WHERE id=$4",
		rcp.UserId, rcp.Title, rcp.Description)
}
