# Sistema de Sincronización de Semáforos en Go


Proyecto de investigación aplicada para un sistema de sincronización de semáforos que optimiza el flujo vehicular utilizando Go (Golang) y patrones de concurrencia.

---

## Sobre el Proyecto

[cite_start]Este proyecto es un sistema para la sincronización de semáforos urbanos, desarrollado en el lenguaje de programación Go como parte de la asignatura de Programación Estructurada y Funcional[cite: 3].

[cite_start]La congestión de tráfico en ciudades como Quito genera múltiples problemas[cite: 12], entre ellos:

* [cite_start]Incremento en tiempos de viaje [cite: 15]
* [cite_start]Mayor consumo de combustible [cite: 16]
* [cite_start]Aumento en emisiones contaminantes [cite: 17]
* [cite_start]Estrés en conductores y pasajeros [cite: 18]

Este sistema busca aplicar la programación concurrente para optimizar el flujo vehicular y mitigar estos problemas.

## 🎯 Objetivos del Proyecto

[cite_start]El objetivo general es **diseñar e implementar un sistema de sincronización semafórica** usando Go para optimizar el flujo vehicular[cite: 25].

Los objetivos específicos incluyen:
1.  [cite_start]Analizar algoritmos de sincronización existentes[cite: 27].
2.  [cite_start]Diseñar la arquitectura del sistema usando patrones concurrentes[cite: 28].
3.  [cite_start]Implementar el núcleo del sistema en Go[cite: 29].
4.  [cite_start]Validar el funcionamiento mediante simulaciones[cite: 30].

---

## Tecnologías Utilizadas

* [cite_start]**Lenguaje:** Go 1.21+ [cite: 52]
* **Librerías Clave (Planeadas):**
    * [cite_start]`sync` (Para concurrencia) [cite: 62]
    * [cite_start]`time` (Para ciclos y temporizadores) [cite: 63]
    * [cite_start]`encoding/json` (Para manejo de configuraciones) [cite: 66]
* [cite_start]**Herramientas:** Git y Visual Studio Code [cite: 68]

---

## 🏛️ Arquitectura del Sistema

El diseño del sistema es modular para manejar diferentes zonas de control y semáforos de forma concurrente. La arquitectura planeada se basa en el siguiente diagrama de componentes del proyecto:



[Image of the system's class diagram]


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
