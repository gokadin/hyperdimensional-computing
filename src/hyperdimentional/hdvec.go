package hyperdimentional

import (
	"fmt"
	"math"
	"math/rand"
)

type Vector []float64

func New(size int) Vector {
	vector := make([]float64, size)

    for index := range vector {
		random := rand.Intn(100)
		if random >= 50 {
            vector[index] = 1.0
		} else {
			vector[index] = -1.0
		}
	}

    return vector
}

func Rotate(v Vector) Vector {
	length := len(v)
	result := make([]float64, length)

	for i := 0; i < length; i++ {
		if i == length - 1 {
			result[i] = v[0]
			break
		}

		result[i] = v[i + 1]
	}

	return result
}

func Multiply(v1, v2 Vector) Vector {
	length := len(v1)
	result := make([]float64, length)

    for i := 0; i < length; i++ {
		result[i] = v1[i] * v2[i]
	}

    return result
}

func Add(v1, v2 Vector) Vector {
	length := len(v1)
	result := make([]float64, length)

    for i := 0; i < length; i++ {
		result[i] = v1[i] + v2[i]
	}

	return result
}

func Cosine(v1, v2 Vector) float64 {
    dot := Dot(v1, v2)

    magnitudeProduct := v1.Magnitude() * v2.Magnitude()

    return dot / magnitudeProduct
}

func Dot(v1, v2 Vector) float64 {
	length := len(v1)
	result := 0.0
    for i := 0; i < length; i++ {
        result += v1[i] * v2[i]
	}

	return result
}

func (v Vector) Magnitude() float64 {
	result := 0.0

	for _, value := range v {
        result += value * value
	}

	return math.Sqrt(result)
}

func (v Vector) Print() {
	for _, value := range v {
		fmt.Println(value)
	}
}
