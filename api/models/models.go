package models

import (
	"database/sql"
	"fmt"
	"payload-app/api/entities"
)

func Insert(provider *entities.CreateProvider, db *sql.DB) (int64, error) {
	result, err := db.Exec("INSERT INTO payloads (deposits_id, provider_ref) VALUES (?, ?)",
		provider.DepositsId, provider.ProviderRef)
	if err != nil {
		return 0, fmt.Errorf("add provider: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add provider: %v", err)
	}

	return id, nil
}
