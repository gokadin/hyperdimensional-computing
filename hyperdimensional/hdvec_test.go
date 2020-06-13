package hyperdimensional

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func Test_default_size(t *testing.T) {
    vector := NewRandBipolar()

    assert.Equal(t, vectorDefaultSize, vector.Size())
}

func Test_Rotate(t *testing.T) {
	vector := NewRandBipolarOfSize(10)

	rotated := Rotate(vector, 1)

	assert.Equal(t, vector.Size(), rotated.Size())
	assert.Equal(t, vector.At(1), rotated.At(0))
	assert.Equal(t, vector.At(0), rotated.At(rotated.Size() - 1))
}

func Test_Rotate_isCorrectWhenCountIsMoreThanOne(t *testing.T) {
	vector := NewRandBipolarOfSize(10)

	rotated := Rotate(vector, 3)

	assert.Equal(t, vector.Size(), rotated.Size())
	assert.Equal(t, vector.At(3), rotated.At(0))
	assert.Equal(t, vector.At(0), rotated.At(7))
	assert.Equal(t, vector.At(4), rotated.At(1))
	assert.Equal(t, vector.At(1), rotated.At(8))
	assert.Equal(t, vector.At(5), rotated.At(2))
	assert.Equal(t, vector.At(2), rotated.At(9))
}

func Test_Multiply(t *testing.T) {
	vec1 := NewRandBipolarOfSize(10)
	vec2 := NewRandBipolarOfSize(10)

	multiplied := Multiply(vec1, vec2)

	assert.Equal(t, vec1.Size(), multiplied.Size())
	for index, value := range multiplied.Values() {
		assert.Equal(t, vec1.At(index) * vec2.At(index), value)
	}
}

func Test_Dot(t *testing.T) {
	vec1 := NewRandBipolarOfSize(3)
	vec2 := NewRandBipolarOfSize(3)

	dot := Dot(vec1, vec2)

	expected := vec1.At(0) * vec2.At(0) + vec1.At(1) * vec2.At(1) + vec1.At(2) * vec2.At(2)
	assert.Equal(t, expected, dot)
}

func Test_Magnitude(t *testing.T) {
	vec := NewRandBipolarOfSize(3)

    result := vec.Magnitude()

	expected := float32(math.Sqrt(float64(vec.At(0) * vec.At(0) + vec.At(1) * vec.At(1) + vec.At(2) * vec.At(2))))
	assert.Equal(t, expected, result)
}

func Test_Cosine(t *testing.T) {
	vec1 := NewRandBipolarOfSize(3)
	vec2 := NewRandBipolarOfSize(3)

	result := Cosine(vec1, vec2)

	expected := Dot(vec1, vec2) / (vec1.Magnitude() * vec2.Magnitude())
	assert.Equal(t, expected, result)
}

func Test_Add(t *testing.T) {
	vec1 := NewRandBipolarOfSize(3)
	vec2 := NewRandBipolarOfSize(3)
	expectedValue1 := vec1.At(0) + vec2.At(0)
	expectedValue2 := vec1.At(1) + vec2.At(1)
	expectedValue3 := vec1.At(2) + vec2.At(2)

	vec1.Add(vec2)

	assert.Equal(t, vec1.At(0), expectedValue1)
	assert.Equal(t, vec1.At(1), expectedValue2)
	assert.Equal(t, vec1.At(2), expectedValue3)
}

func Test_ToBipolar(t *testing.T) {
	vec := NewEmptyBipolarOfSize(3)
	vec.Set(0, 2)
	vec.Set(1, -3)
	vec.Set(2, 1)

	vec.ToBipolar()

	assert.Equal(t, []float32{1, -1, 1}, vec.Values())
}

func Test_Scale(t *testing.T) {
	vec := NewEmptyBipolarOfSize(2)
	vec.Set(0, 1)
	vec.Set(1, -1)

	vec.ScaleUp(10)

	assert.Equal(t, 10, vec.Size())
    for i, value := range vec.Values() {
        if i < 5 && value != 1 {
            t.Fatalf("Invalid scaled value.")
		}
        if i >= 5 && value != -1 {
			t.Fatalf("Invalid scaled value.")
		}
	}
}
