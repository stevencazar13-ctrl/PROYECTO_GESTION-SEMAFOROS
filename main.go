package main

import "fmt"

// Clase Semaforo (struct)
type Semaforo struct {
	Area  string
	Color string
	Ciclo float64
}

// Método para cambiar el color
func (s *Semaforo) CambiarColor(nuevoColor string) {
	s.Color = nuevoColor
	fmt.Printf("El semáforo en %s ahora está en %s\n", s.Area, nuevoColor)
}

// Método para mostrar estado actual
func (s *Semaforo) MostrarEstado() {
	fmt.Printf("Área: %s | Color: %s | Ciclo: %.1f segundos\n",
		s.Area, s.Color, s.Ciclo)
}

// Constructor (función que crea instancia)
func NuevoSemaforo(area string, color string, ciclo float64) *Semaforo {
	return &Semaforo{
		Area:  area,
		Color: color,
		Ciclo: ciclo,
	}
}

func main() {

	semaforoCentro := Semaforo{
		Area:  "Centro Ciudad",
		Color: "Rojo",
		Ciclo: 30.0,
	}

	//constructor
	semaforoEscuela := NuevoSemaforo("Frente Escuela", "Verde", 15.0)

	semaforoCentro.MostrarEstado()
	semaforoEscuela.MostrarEstado()

	// Cambiar estado
	semaforoCentro.CambiarColor("Verde")
	semaforoCentro.MostrarEstado()

	// Modificación directa de propiedades
	semaforoEscuela.Ciclo = 20.0
	fmt.Printf("Nuevo ciclo: %.1f segundos\n", semaforoEscuela.Ciclo)
}
