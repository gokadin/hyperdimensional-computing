package hyperdimensional

import (
	"log"
	"math"
	"math/rand"
)

const vectorDefaultSize = 10000

type HdVec struct {
    values []float32
}

func NewRandBipolar() *HdVec {
	return NewRandBipolarOfSize(vectorDefaultSize)
}

func NewRandBipolarOfSize(size int) *HdVec {
	vec := NewEmptyBipolarOfSize(size)

	for index := range vec.values {
		random := rand.Intn(2)
		if random == 0 {
			vec.values[index] = -1.0
		} else {
			vec.values[index] = 1.0
		}
	}

	return vec
}

func NewEmptyBipolar() *HdVec {
	return NewEmptyBipolarOfSize(vectorDefaultSize)
}

func NewEmptyBipolarOfSize(size int) *HdVec {
	return &HdVec{
		values: make([]float32, size),
	}
}

func FromSlice(input []float32) *HdVec {
	return &HdVec{
		values: input,
	}
}

func CircularConvolution(v1, v2 *HdVec) *HdVec {
	if v1.Size() != v2.Size() {
		log.Fatalf("vector sizes do not match: %d and %d", v1.Size(), v2.Size())
	}

	result := NewEmptyBipolarOfSize(v1.Size())
	for j := 0; j < v1.Size(); j++ {
		for k := 0; k < v1.Size(); k++ {
			v2Index := ((j - k) % v1.Size() + v1.Size()) % v1.Size()
			result.Set(j, result.At(j) + v1.At(k) * v2.At(v2Index))
		}
	}

	return result
}

/**
	Rotate for:
	[ x_0 x_1 x_2 x_3 ... x_n ]
	Produces:
	[ x_1 x_2 x_3 ... x_n x_0 ]
 */
func Rotate(v *HdVec, count int) *HdVec {
	result := NewEmptyBipolarOfSize(v.Size())
	for i := 0; i < result.Size(); i++ {
		if i >= result.Size() - count {
			result.values[i] = v.values[count - (result.Size() - i)]
			continue
		}

		result.values[i] = v.values[i + count]
	}

	return result
}

func Multiply(vectors ...*HdVec) *HdVec {
	if len(vectors) < 2 {
		log.Fatal("you must provide at least 2 vectors")
	}

	result := NewEmptyBipolarOfSize(vectors[0].Size())
    for i := 0; i < result.Size(); i++ {
    	for j := 1; j < len(vectors); j++ {
			result.values[i] = vectors[j - 1].values[i] * vectors[j].values[i]
		}
	}

    return result
}

func Add(vectors ...*HdVec) *HdVec {
	if len(vectors) < 2 {
		log.Fatal("you must provide at least 2 vectors")
	}

	result := NewEmptyBipolarOfSize(vectors[0].Size())
	for i := 0; i < result.Size(); i++ {
		for j := 1; j < len(vectors); j++ {
			result.values[i] = vectors[j - 1].values[i] + vectors[j].values[i]
		}
	}

	return result
}

func (v *HdVec) Add(v2 *HdVec) {
    for i := 0; i < v.Size(); i++ {
		v.values[i] += v2.values[i]
	}
}

func (v *HdVec) ToBipolar() {
	for i := 0; i < v.Size(); i++ {
		if v.values[i] > 0 {
			v.values[i] = 1
		} else {
			v.values[i] = -1
		}
	}
}

func Cosine(v1, v2 *HdVec) float32 {
    dot := Dot(v1, v2)

    magnitudeProduct := v1.Magnitude() * v2.Magnitude()

    return dot / magnitudeProduct
}

func Dot(v1, v2 *HdVec) float32 {
	var result float32
    for i := 0; i < v1.Size(); i++ {
        result += v1.values[i] * v2.values[i]
	}

	return result
}

func (v *HdVec) Magnitude() float32 {
	var result float32
	for _, value := range v.values {
        result += value * value
	}

	return float32(math.Sqrt(float64(result)))
}

/**
	Involution for:
	[ x_0 x_1 x_2 x_3 x_n ... ]
	Produces:
	[ x_0 x_n ... x_3 x_2 x_1 ]
 */
func (v *HdVec) Involution() *HdVec {
	for i := 1; i < v.Size(); i++ {
		 v.Set(i, v.At(v.Size() - i))
	}
	return v
}

func (v *HdVec) Values() []float32 {
	return v.values
}

func (v *HdVec) Size() int {
	return len(v.values)
}

func (v *HdVec) At(index int) float32 {
	return v.values[index]
}

func (v *HdVec) Set(index int, value float32) {
	v.values[index] = value
}

func (v *HdVec) ScaleUp(size int) {
	scaled := make([]float32, size)
	scaleFactor := size / len(v.values)

	for i, value := range v.values {
		for j := 0; j < scaleFactor; j++ {
			scaled[i * scaleFactor + j] = value
		}
	}

	v.values = scaled
}
