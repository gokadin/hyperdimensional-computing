package hyperdimentional

import (
	"github.com/gokadin/hyperdimentional/src/hyperdimentional"
	"math"
	"testing"
)

func Test_Size_isCorrect(t *testing.T) {
	// Arrange
    vector := hyperdimentional.New(10000)

	// Assert
    if len(vector) != 10000 {
    	t.Fail()
	}
}

func Test_Rotate_isCorrect(t *testing.T) {
	// Arrange
	vector := hyperdimentional.New(10)

	// Act
	rotated := hyperdimentional.Rotate(vector)

	// Assert
	if len(rotated) != len(vector) {
		t.Fatalf("Size does not match.")
	}

	if rotated[0] != vector[1] || rotated[len(rotated) - 1] != vector[0] {
		t.Fatalf("Rotation failed.")
	}
}

func Test_Multiply_isCorrect(t *testing.T) {
	// Arrange
	vec1 := hyperdimentional.New(10)
	vec2 := hyperdimentional.New(10)

	// Act
	multiplied := hyperdimentional.Multiply(vec1, vec2)

	// Assert
	if len(multiplied) != len(vec1) {
		t.Fatalf("Size does not match.")
	}

	for index, value := range multiplied {
		if value != vec1[index] * vec2[index] {
			t.Fatalf("Multiplication failed.")
		}
	}
}

func Test_Dot_isCorrect(t *testing.T) {
	// Arrange
	vec1 := hyperdimentional.New(3)
	vec2 := hyperdimentional.New(3)

	// Act
	dot := hyperdimentional.Dot(vec1, vec2)

	// Assert
	expected := vec1[0] * vec2[0] + vec1[1] * vec2[1] + vec1[2] * vec2[2]
	if dot != expected {
		t.Fatalf("Dot product is incorrect. Should be %f, received %f", expected, dot)
	}
}

func Test_Magnitude_isCorrect(t *testing.T) {
	// Arrange
	vec := hyperdimentional.New(3)

	// Act
    result := vec.Magnitude()

    // Assert
	expected := math.Sqrt(float64(vec[0] * vec[0] + vec[1] * vec[1] + vec[2] * vec[2]))
	if result != expected {
		t.Fatalf("Magnitude is incorrect. Should be %f, received %f", expected, result)
	}
}

func Test_Cosine_isCorrect(t *testing.T) {
	// Arrange
	vec1 := hyperdimentional.New(3)
	vec2 := hyperdimentional.New(3)

	// Act
	result := hyperdimentional.Cosine(vec1, vec2)

	// Assert
	expected := hyperdimentional.Dot(vec1, vec2) / (vec1.Magnitude() * vec2.Magnitude())
	if result != expected {
		t.Fatalf("Cosine is incorrect. Should be %f, received %f", expected, result)
	}
}

