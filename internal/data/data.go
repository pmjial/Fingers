package data

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"slices"
	"strings"
)

//go:embed seed.json
var seed []byte

type Measurements map[string]float64

type Dataset struct {
	Units        Measurements `json:"units"`
	Measurements Measurements `json:"measurements"`
}

func LoadSeed() (Dataset, error) {
	dataset, err := decode(bytes.NewReader(seed))
	if err != nil {
		return Dataset{}, fmt.Errorf("decode embedded seed: %w", err)
	}
	if err := dataset.validate(); err != nil {
		return Dataset{}, fmt.Errorf("validate embedded seed: %w", err)
	}
	return dataset, nil
}

func (measurements Measurements) Names() []string {
	names := make([]string, 0, len(measurements))
	for name := range measurements {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}

func decode(reader io.Reader) (Dataset, error) {
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()

	var dataset Dataset
	if err := decoder.Decode(&dataset); err != nil {
		return Dataset{}, err
	}

	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		if err == nil {
			return Dataset{}, fmt.Errorf("unexpected data after dataset")
		}
		return Dataset{}, err
	}
	return dataset, nil
}

func (dataset Dataset) validate() error {
	if len(dataset.Units) == 0 {
		return fmt.Errorf("must contain at least one unit")
	}
	if len(dataset.Measurements) == 0 {
		return fmt.Errorf("must contain at least one measurement")
	}
	if err := validateMeasurements("units", dataset.Units); err != nil {
		return err
	}
	return validateMeasurements("fingers", dataset.Measurements)
}

func validateMeasurements(section string, measurements Measurements) error {
	for name, value := range measurements {
		if strings.TrimSpace(name) == "" {
			return fmt.Errorf("%s contains a blank measurement name", section)
		}
		if value <= 0 || math.IsNaN(value) || math.IsInf(value, 0) {
			return fmt.Errorf("%s measurement %q must be positive and finite", section, name)
		}
	}
	return nil
}
