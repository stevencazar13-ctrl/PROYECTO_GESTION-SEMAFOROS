# Sistema de Sincronización de Semáforos en Go


Proyecto de investigación aplicada para un sistema de sincronización de semáforos que optimiza el flujo vehicular utilizando Go (Golang) y patrones de concurrencia.

---

## Sobre el Proyecto

Este proyecto es un sistema para la sincronización de semáforos urbanos, desarrollado en el lenguaje de programación Go como parte de la asignatura de Programación Estructurada y Funcional.

La congestión de tráfico en ciudades como Quito genera múltiples problemas, entre ellos:

* Incremento en tiempos de viaje
* Mayor consumo de combustible 
* Aumento en emisiones contaminantes 
* Estrés en conductores y pasajeros 

Este sistema busca aplicar la programación concurrente para optimizar el flujo vehicular y mitigar estos problemas.

## Objetivos del Proyecto

[cite_start]El objetivo general es **diseñar e implementar un sistema de sincronización semafórica** usando Go para optimizar el flujo vehicular.

Los objetivos específicos incluyen:
1.  Analizar algoritmos de sincronización existentes
2.  Diseñar la arquitectura del sistema usando patrones concurrentes
3.  Implementar el núcleo del sistema en Go
4.  Validar el funcionamiento mediante simulaciones

---

## Tecnologías Utilizadas

* **Lenguaje:** Go 1.21+ 
* **Librerías Clave (Planeadas):**
    * `sync` (Para concurrencia) 
    * `time` (Para ciclos y temporizadores) 
    * `encoding/json` (Para manejo de configuraciones) 
* **Herramientas:** Git y Visual Studio Code 

---


## Ejemplo de Código Básico

El siguiente archivo `main.go` muestra la estructura fundamental de un `Semaforo` y sus métodos básicos, que sirve como punto de partida para el sistema.

```go
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


