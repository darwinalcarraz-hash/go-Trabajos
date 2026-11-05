package main

import (
	"errors"
	"fmt"
)

// Estructura base Medicine con campos privados  
type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate string
	shelfLife       int
}

// Getters y Setters para Medicine con validaciones
func (m Medicine) GetName() string { 
	return m.name }
func (m *Medicine) SetName(name string) error {
	if name == "" {
		return errors.New("ERROR: el nombre no puede estar vacío")
	}
	m.name = name
	return nil
}

func (m Medicine) GetManufacturer() string { 
	return m.manufacturer }
func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return errors.New("ERROR: el fabricante no puede estar vacío")
	}
	m.manufacturer = manufacturer
	return nil
}

func (m Medicine) GetManufactureDate() string { 
	return m.manufactureDate }
func (m *Medicine) SetManufactureDate(manufactureDate string) error {
	if manufactureDate == "" {
		return errors.New("ERROR: la fecha no puede estar vacía")
	}
	m.manufactureDate = manufactureDate
	return nil
}

func (m Medicine) GetShelfLife() int { 
	return m.shelfLife }
func (m *Medicine) SetShelfLife(shelfLife int) error {
	if shelfLife <= 0 {
		return errors.New("ERROR: la vida útil debe ser un número positivo")
	}
	m.shelfLife = shelfLife
	return nil
}

// Estructuras que heredan de Medicine Tablet y Syrup
type Tablet struct {
	Medicine
	dosePerTablet  float64
	isPrescription bool
}

type Syrup struct {
	Medicine
	volume float64
	flavor string
}