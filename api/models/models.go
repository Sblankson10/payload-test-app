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
		fmt.Println("I can not insert")
		return 0, fmt.Errorf("add provider: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("add provider: %v", err)
	}

	return id, nil
}

func GetProvidersByDepositor(depositorID int64, db *sql.DB) ([]entities.Provider, error) {
	var providers []entities.Provider

	result, err := db.Query("SELECT deposits_id, provider_ref, created, modified FROM payloads WHERE deposits_id = ?", depositorID)
	if err != nil {
		return nil, fmt.Errorf("GetProvider %v: %v", depositorID, err)
	}
	defer result.Close()

	// loop through rows
	for result.Next() {
		var provider entities.Provider
		if err := result.Scan(&provider.DepositsId, &provider.ProviderRef, &provider.Created, &provider.Modified); err != nil {
			return nil, fmt.Errorf("GetProvider %v: %v", depositorID, err)
		}
		providers = append(providers, provider)
		if err := result.Err(); err != nil {
			return nil, fmt.Errorf("GetProvider %v: %v", depositorID, err)
		}
	}

	return providers, nil
}

func GetProviders(db *sql.DB) ([]entities.Provider, error) {
	var providers []entities.Provider

	result, err := db.Query("SELECT deposits_id, provider_ref, created, modified FROM payloads")
	if err != nil {
		return nil, fmt.Errorf("GetProvider: %v", err)
	}
	defer result.Close()

	// loop through rows
	for result.Next() {
		var provider entities.Provider
		if err := result.Scan(&provider.DepositsId, &provider.ProviderRef, &provider.Created, &provider.Modified); err != nil {
			return nil, fmt.Errorf("GetProvider: %v", err)
		}
		providers = append(providers, provider)
		if err := result.Err(); err != nil {
			return nil, fmt.Errorf("GetProvider: %v", err)
		}
	}

	return providers, nil
}
