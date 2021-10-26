package goarea

import "math"

// Pi é uma constante matemática, uma relação
// entre a circuferência e seu diâmetro
const Pi = 3.1416

// Circ calcula area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula area do retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// Func privada! ("_")
func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
