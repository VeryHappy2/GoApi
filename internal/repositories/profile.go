package repositories

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/VeryHappy2/GoApi/internal/storage"
	"github.com/VeryHappy2/GoApi/internal/storage/sqlite"
)

type Profile struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
}

type ProfileRep struct {
	Object Profile
}

func (profile ProfileRep) Add(db *sqlite.Storage) error {
	var operation string = "internal.repositories.add"

	stmt, err := db.DB.Prepare("INSERT INTO t_profile (name, last_name) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare statement: %w", operation, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(profile.Object.Name, profile.Object.LastName)
	if err != nil {
		return fmt.Errorf("%s: failed to execute statement: %w", operation, err)
	}

	return nil
}

func (profile ProfileRep) Update(db *sqlite.Storage) error {
	var operation string = "internal.repositories.update"

	stmt, err := db.DB.Prepare("update t_profile set name = ? last_name = ? where id = ?")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare statement: %w", operation, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(profile.Object.Name, profile.Object.LastName, profile.Object.Id)
	if err != nil {
		return fmt.Errorf("%s: failed to execute statement: %w", operation, err)
	}

	return nil
}

func (profile ProfileRep) GetById(db *sqlite.Storage) (*Profile, error) {
	var operation string = "internal.repositories.getById"

	var dbProfile Profile
	query := "SELECT id, name, last_name FROM t_profile WHERE id = ?"
	err := db.DB.QueryRow(query, profile.Object.Id).Scan(&dbProfile.Id, &dbProfile.LastName, &dbProfile.Name)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, storage.ErrNotFound
		}

		return nil, fmt.Errorf("%s: execute statement: %w", operation, err)
	}

	return &dbProfile, nil
}

func (profile ProfileRep) Delete(db sqlite.Storage) error {
	var operation string = "internal.repositories.delete"

	stmt, err := db.DB.Prepare("delete from t_profile where id = ?")
	if err != nil {
		return fmt.Errorf("%s: failed to prepare statement: %w", operation, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(profile.Object.Id)
	if err != nil {
		return fmt.Errorf("%s: failed to execute statement: %w", operation, err)
	}

	return nil
}
