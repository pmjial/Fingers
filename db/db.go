package db

import (
	"encoding/json"
	"math/rand"
	"os"
)

const dbFile = "db.json"

type (
	Length   = float64
	Database map[string]Length
	DataPair struct {
		Key   string
		Value Length
	}
)

func Get() (db Database) {
	f, err := os.ReadFile(dbFile)
	if err != nil {
		return nil
	}

	if err := json.Unmarshal(f, &db); err != nil {
		return nil
	}
	return db
}

func Keys[K comparable, V any](db map[K]V) []K {
	keys := make([]K, 0, len(db))
	for k := range db {
		keys = append(keys, k)
	}
	return keys
}

func (db Database) RandKey() string {
	randIndex := rand.Intn(len(db))
	return Keys(db)[randIndex]
}

func (db Database) Commit() error {
	f, _ := json.MarshalIndent(db, "", "\t")
	return os.WriteFile(dbFile, f, 0755)
}

func (db Database) Append(data ...DataPair) {
	for _, datum := range data {
		db[datum.Key] = datum.Value
	}
}
