package db

import (
	_ "embed"
	"encoding/json"
	"math/rand"
)

//go:embed "db.json"
var f []byte

type (
	Length   = float64
	Database map[string]Length
	DataPair struct {
		Key   string
		Value Length
	}
)

func get() (db map[string]Database) {
	if err := json.Unmarshal(f, &db); err != nil {
		return nil
	}
	return db
}

func Fingers() Database {
	return get()["fingers"]
}

func Units() Database {
	return get()["units"]
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
