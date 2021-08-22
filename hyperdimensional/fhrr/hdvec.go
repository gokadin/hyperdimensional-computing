package fhrr

import (
	"encoding/json"
	"math"
	"math/rand"
)

const VectorDefaultSize = 10000

type HdVec struct {
	values []float64
}

func Rand() *HdVec {
	return RandOfSize(VectorDefaultSize)
}

func RandOfSize(size int) *HdVec {
	vec := NewEmptyOfSize(size)

	for index := range vec.values {
		vec.values[index] = -math.Pi + rand.Float64()*(math.Pi+math.Pi)
	}

	return vec
}

func NewEmptyOfSize(size int) *HdVec {
	return &HdVec{
		values: make([]float64, size),
	}
}

func FromSlice(values []float64) *HdVec {
	vec := &HdVec{values}
	return vec
}

func FromHDVec(from *HdVec) *HdVec {
	vec := NewEmptyOfSize(from.Size())
	for i := 0; i < from.Size(); i++ {
		vec.values[i] = from.values[i]
	}
	return vec
}

func Bind(a, b *HdVec) *HdVec {
	result := NewEmptyOfSize(a.Size())

	for i := 0; i < a.Size(); i++ {
		angleSum := a.values[i] + b.values[i]
		angleSum = normalizeAngle(angleSum)
		result.values[i] = angleSum
	}

	return result
}

func Unbind(a, b *HdVec) *HdVec {
	result := NewEmptyOfSize(a.Size())

	for i := 0; i < a.Size(); i++ {
		angleSum := a.values[i] - b.values[i]
		angleSum = normalizeAngle(angleSum)
		result.values[i] = angleSum
	}

	return result
}

func Bundle(vectors ...*HdVec) *HdVec {
	result := NewEmptyOfSize(vectors[0].Size())
	for i := 0; i < vectors[0].Size(); i++ {
		var resultValue float64
		for _, vector := range vectors {
			resultValue += math.Pow(math.E, vector.values[i])
		}
		resultValue = math.Log(resultValue)
		result.values[i] = normalizeAngle(resultValue)
	}

	return result
}

func normalizeAngle(angle float64) float64 {
	for angle < -math.Pi {
		angle += 2 * math.Pi
	}

	for angle > math.Pi {
		angle -= 2 * math.Pi
	}

	return angle
}

func (v *HdVec) Size() int {
	return len(v.values)
}

func (v *HdVec) At(index int) float64 {
	return v.values[index]
}

func (v *HdVec) Set(index int, value float64) {
	v.values[index] = value
}

func Similarity(a, b *HdVec) float64 {
	var sum float64
	for i := 0; i < a.Size(); i++ {
		sum += math.Cos(a.values[i] - b.values[i])
	}
	return sum / float64(a.Size())
}

func Equal(a, b *HdVec) bool {
	if a.Size() != b.Size() {
		return false
	}

	for i := 0; i < len(a.values); i++ {
		if a.values[i] != b.values[i] {
			return false
		}
	}

	return true
}

func (v *HdVec) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, &v.values)
}

func (v *HdVec) MarshalJSON() ([]byte, error) {
	return json.Marshal(&v.values)
}
