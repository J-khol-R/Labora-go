package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
}

type Cuadrado struct {
	Lado float64
}

type Rectangulo struct {
	Base   float64
	Altura float64
}

type Triangulo struct {
	Base   float64
	Altura float64
}

type Rombo struct {
	Diagonal1 float64
	Diagonal2 float64
}

type Trapecio struct {
	Base1  float64
	Base2  float64
	Altura float64
}

type Circulo struct {
	Radio float64
}

type Pentagono struct {
	Lado    float64
	Apotema float64
}

func (s *Cuadrado) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("la oprecion falló:", err)
		}
	}()

	if s.Lado <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}

	return s.Lado * s.Lado
}

func (s *Rectangulo) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("la oprecion falló:", err)
		}
	}()

	if s.Altura <= 0 || s.Base <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if s.Base == s.Altura {
		panic("Para que sea un rectangulo la base y la altura debe de ser distinta ")
	}
	return s.Base * s.Altura
}

func (s *Triangulo) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("la oprecion falló:", err)
		}
	}()

	if s.Altura <= 0 && s.Base <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	return (s.Base * s.Altura) / 2
}

func (s *Rombo) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("la oprecion falló:", err)
		}
	}()

	if s.Diagonal1 <= 0 && s.Diagonal2 <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if s.Diagonal1 == s.Diagonal2 {
		panic("Una diagonal debe de ser mayor a la otra, no iguales ")
	}
	return (s.Diagonal1 * s.Diagonal2) / 2
}

func (s *Trapecio) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("la oprecion falló:", err)
		}
	}()

	if s.Base1 <= 0 && s.Base2 <= 0 && s.Altura <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if s.Base1 == s.Base2 {
		panic("Para que sea un trapecio las dos bases deben de ser distintas ")
	}
	return ((s.Base1 + s.Base2) / 2) * s.Altura
}

func (s *Circulo) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("la oprecion falló:", err)
		}
	}()

	if s.Radio <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	return 3.1416 * s.Radio * s.Radio
}

func (s *Pentagono) Area() float64 {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, "ingresa de nuevo los datos:")
		}
	}()

	if s.Lado <= 0 && s.Apotema <= 0 {
		panic("Los números ingresados no pueden ser menor o igual a cero ")
	}
	if s.Lado < s.Apotema {
		panic("El lado no puede ser menor que el apotema")
	}
	return ((5 * s.Lado) * s.Apotema) / 2
}

func imprimirArea(shapes map[string]Shape) {
	for key, valor := range shapes {
		actual := valor.Area()
		if actual > 5000 {
			fmt.Printf("El area de la figura (%s) es demasiado grande: %v\n", key, actual)
		} else if actual != 0 {
			fmt.Printf("El area del %s es: %v\n", key, actual)
		}
	}
}

func main() {

	shapes := map[string]Shape{
		"Cuadrado":   &Cuadrado{Lado: 4},
		"Rectangulo": &Rectangulo{Base: 40, Altura: 40},
		"Triangulo":  &Triangulo{Base: 3, Altura: 6},
		"Rombo":      &Rombo{Diagonal1: 4, Diagonal2: 6},
		"Trapecio":   &Trapecio{Base1: 2, Base2: 4, Altura: 3},
		"Circulo":    &Circulo{Radio: 2},
		"Pentagono":  &Pentagono{Lado: 6, Apotema: 4},
	}

	imprimirArea(shapes)

}
