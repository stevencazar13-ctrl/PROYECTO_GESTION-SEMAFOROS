#Sistema de Sincronización de Semáforos en Go
Este proyecto es un sistema para la sincronización de semáforos urbanos, desarrollado en el lenguaje de programación Go. Su objetivo es optimizar el flujo vehicular usando programación concurrente.

#Contexto
La congestión de tráfico genera múltiples problemas, como aumento en los tiempos de viaje, mayor consumo de combustible y estrés en los conductores. Una sincronización inteligente de semáforos puede reducir estos problemas.

#Objetivo Principal
Diseñar e implementar un sistema de sincronización semafórica usando Go para optimizar el flujo vehicular.

Tecnologías Utilizadas
Lenguaje: Go 1.21+

Librerías Principales: fmt

Herramientas: Git, Visual Studio Code

(Nota: Librerías como sync y time se planean usar a medida que el proyecto crezca, como se indica en el PDF del proyecto, pero este ejemplo básico solo usa fmt.)

Ejemplo de Código Básico
El siguiente archivo main.go muestra la estructura fundamental de un Semaforo y sus métodos básicos.

Go

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
    fmt.Printf("El semáforo en %s ahora está en %s\n", s.Area, s.Color)
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
Autores
Juan Sebastián Báez Fuentes


Steven Cazar
