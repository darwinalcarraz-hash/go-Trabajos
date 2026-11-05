package main

import (
	"fmt"
	"time"
)

// 1. Estructura base Medicine con campos privados
type Medicine struct {
	name            string
	manufacturer    string
	manufactureDate time.Time
	shelfLife       int // en meses
}

// Constructor para Medicine (para que Tablet y Syrup lo utilicen como pide la instrucción 4)
func NewMedicine(name, manufacturer string, manufactureDate time.Time, shelfLife int) (Medicine, error) {
	m := Medicine{}
	if err := m.SetName(name); err != nil {
		return Medicine{}, err
	}
	if err := m.SetManufacturer(manufacturer); err != nil {
		return Medicine{}, err
	}
	if err := m.SetManufactureDate(manufactureDate); err != nil {
		return Medicine{}, err
	}
	if err := m.SetShelfLife(shelfLife); err != nil {
		return Medicine{}, err
	}
	return m, nil
}

// 2. Getters y Setters para cada campo de Medicine con las validaciones apropiadas
func (m *Medicine) GetName() string {
	return m.name
}

func (m *Medicine) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("el nombre no puede estar vacío")
	}
	m.name = name
	return nil
}

func (m *Medicine) GetManufacturer() string {
	return m.manufacturer
}

func (m *Medicine) SetManufacturer(manufacturer string) error {
	if manufacturer == "" {
		return fmt.Errorf("el fabricante no puede estar vacío")
	}
	m.manufacturer = manufacturer
	return nil
}

func (m *Medicine) GetManufactureDate() time.Time {
	return m.manufactureDate
}

func (m *Medicine) SetManufactureDate(manufactureDate time.Time) error {
	if manufactureDate.IsZero() {
		return fmt.Errorf("la fecha de fabricación debe ser una fecha válida")
	}
	m.manufactureDate = manufactureDate
	return nil
}

func (m *Medicine) GetShelfLife() int {
	return m.shelfLife
}

func (m *Medicine) SetShelfLife(shelfLife int) error {
	if shelfLife <= 0 {
		return fmt.Errorf("la vida útil debe ser un número positivo")
	}
	m.shelfLife = shelfLife
	return nil
}

// 6. Método en Medicine que calcule la fecha de caducidad en base a la fecha de fabricación y la vida útil
func (m *Medicine) CalculateExpirationDate() time.Time {
	return m.manufactureDate.AddDate(0, m.shelfLife, 0)
}

// 3. Dos estructuras que "heredan" de Medicine: Tablet y Syrup
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

// Getters y Setters adicionales o específicos si se requieren, 
// o se pueden actualizar mediante los métodos de Medicine y los campos directos de las hijas.

// 4. Constructores para Tablet y Syrup que inicialicen correctamente los valores y usen el constructor de Medicine
func NewTablet(name, manufacturer string, manufactureDate time.Time, shelfLife int, dosePerTablet float64, isPrescription bool) (Tablet, error) {
	if dosePerTablet <= 0 {
		return Tablet{}, fmt.Errorf("dosePerTablet debe ser mayor a 0")
	}

	// Usando el constructor de Medicine
	med, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return Tablet{}, err
	}

	return Tablet{
		Medicine:       med,
		dosePerTablet:  dosePerTablet,
		isPrescription: isPrescription,
	}, nil
}

func NewSyrup(name, manufacturer string, manufactureDate time.Time, shelfLife int, volume float64, flavor string) (Syrup, error) {
	if volume <= 0 {
		return Syrup{}, fmt.Errorf("volume debe ser mayor a 0")
	}

	// Usando el constructor de Medicine
	med, err := NewMedicine(name, manufacturer, manufactureDate, shelfLife)
	if err != nil {
		return Syrup{}, err
	}

	return Syrup{
		Medicine: med,
		volume:   volume,
		flavor:   flavor,
	}, nil
}

// 5. Define métodos para Tablet y Syrup que muestren todos los detalles del medicamento
func (t Tablet) ShowDetails() {
	fmt.Printf("[Tablet] Nombre: %s | Fabricante: %s | F. Fabricación: %s | Vida Útil: %d meses | Dosis: %.2f mg | Receta: %t\n",
		t.GetName(), t.GetManufacturer(), t.GetManufactureDate().Format("2006-01-02"), t.GetShelfLife(), t.dosePerTablet, t.isPrescription)
}

func (s Syrup) ShowDetails() {
	fmt.Printf("[Syrup] Nombre: %s | Fabricante: %s | F. Fabricación: %s | Vida Útil: %d meses | Volumen: %.2f ml | Sabor: %s\n",
		s.GetName(), s.GetManufacturer(), s.GetManufactureDate().Format("2006-01-02"), s.GetShelfLife(), s.volume, s.flavor)
}

// 7. Función main
func main() {
	// Fechas válidas de prueba
	date1, _ := time.Parse("2006-01-02", "2024-05-10")
	date2, _ := time.Parse("2006-01-02", "2024-01-15")

	// Creación de instancias
	t1, _ := NewTablet("Paracetamol", "PharmaCorp", date1, 24, 500.0, false)
	t2, _ := NewTablet("Ibuprofeno", "MediLab", date2, 12, 400.0, true)

	s1, _ := NewSyrup("Jarabe Tos", "HealthLab", date1, 18, 120.5, "Cereza")
	s2, _ := NewSyrup("Amoxicilina", "BioPharma", date2, 6, 90.0, "Fresa")

	// - Crea dos listas de inventario que contengan instancias de Tablet y Syrup cada una.
	inventario1 := []any{t1, s1}
	inventario2 := []any{t2, s2}

	// - Muestra los detalles de todos los medicamentos en el inventario.
	// - Calcula y muestra la fecha de caducidad de cada medicamento.
	fmt.Println("=== INVENTARIO 1 ===")
	for _, item := range inventario1 {
		switch v := item.(type) {
		case Tablet:
			v.ShowDetails()
			fmt.Printf("   -> Fecha de Caducidad: %s\n", v.CalculateExpirationDate().Format("2006-01-02"))
		case Syrup:
			v.ShowDetails()
			fmt.Printf("   -> Fecha de Caducidad: %s\n", v.CalculateExpirationDate().Format("2006-01-02"))
		}
	}

	fmt.Println("\n=== INVENTARIO 2 ===")
	for _, item := range inventario2 {
		switch v := item.(type) {
		case Tablet:
			v.ShowDetails()
			fmt.Printf("   -> Fecha de Caducidad: %s\n", v.CalculateExpirationDate().Format("2006-01-02"))
		case Syrup:
			v.ShowDetails()
			fmt.Printf("   -> Fecha de Caducidad: %s\n", v.CalculateExpirationDate().Format("2006-01-02"))
		}
	}

	// - Actualiza algún atributo de un medicamento usando un setter y muestra el cambio usando un getter.
	fmt.Println("\n=== PRUEBA DE SETTER Y GETTER ===")
	fmt.Printf("Nombre original (t1): %s\n", t1.GetName())

	err := t1.SetName("Paracetamol 500mg (Actualizado)")
	if err == nil {
		fmt.Printf("Nombre actualizado (t1 usando getter): %s\n", t1.GetName())
	}
}