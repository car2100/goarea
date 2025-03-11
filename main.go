package goarea

import "math"

const Pi = 3.1415

// circ calcula a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcula area de retangulo
func Rect(base, altura float64) float64 {
	return base * altura
}

// não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
