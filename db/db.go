package db

import "math/rand"

type Length float64

var db = map[string]Length{
	"finger":         .09,
	"hockey stick":   .81,
	"tree":           20,
	"M16 5.56 Rifle": .986,
	"lava lamp":      0.3048,
}

func Get(key string) Length {
	return db[key]
}

func Keys() []string {
	keys := make([]string, 0, len(db))
	for k := range db {
		keys = append(keys, k)
	}
	return keys
}

func RandKey() string {
	randIndex := rand.Intn(len(db))
	return Keys()[randIndex]
}

func RandValue() Length {
	return Get(RandKey())
}
