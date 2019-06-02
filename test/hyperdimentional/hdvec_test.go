package hyperdimentional

import (
	"github.com/gokadin/hyperdimentional/src/hyperdimentional"
	"math"
	"testing"
)

func Test_Size_isCorrect(t *testing.T) {
	// Arrange
    vector := hyperdimentional.NewVecBinomial(10000)

	// Assert
    if vector.Size() != 10000 {
    	t.Fail()
	}
}

func Test_Rotate_isCorrect(t *testing.T) {
	// Arrange
	vector := hyperdimentional.NewVecBinomial(10)

	// Act
	rotated := hyperdimentional.Rotate(vector)

	// Assert
	if rotated.Size() != vector.Size() {
		t.Fatalf("Size does not match.")
	}

	if rotated.At(0) != vector.At(1) || rotated.At(rotated.Size() - 1) != vector.At(0) {
		t.Fatalf("Rotation failed.")
	}
}

func Test_Multiply_isCorrect(t *testing.T) {
	// Arrange
	vec1 := hyperdimentional.NewVecBinomial(10)
	vec2 := hyperdimentional.NewVecBinomial(10)

	// Act
	multiplied := hyperdimentional.Multiply(vec1, vec2)

	// Assert
	if multiplied.Size() != vec1.Size() {
		t.Fatalf("Size does not match.")
	}

	for index, value := range *multiplied.Values() {
		if value != vec1.At(index) * vec2.At(index) {
			t.Fatalf("Multiplication failed.")
		}
	}
}

func Test_Dot_isCorrect(t *testing.T) {
	// Arrange
	vec1 := hyperdimentional.NewVecBinomial(3)
	vec2 := hyperdimentional.NewVecBinomial(3)

	// Act
	dot := hyperdimentional.Dot(vec1, vec2)

	// Assert
	expected := vec1.At(0) * vec2.At(0) + vec1.At(1) * vec2.At(1) + vec1.At(2) * vec2.At(2)
	if dot != expected {
		t.Fatalf("Dot product is incorrect. Should be %f, received %f", expected, dot)
	}
}

func Test_Magnitude_isCorrect(t *testing.T) {
	// Arrange
	vec := hyperdimentional.NewVecBinomial(3)

	// Act
    result := vec.Magnitude()

    // Assert
	expected := math.Sqrt(float64(vec.At(0) * vec.At(0) + vec.At(1) * vec.At(1) + vec.At(2) * vec.At(2)))
	if result != expected {
		t.Fatalf("Magnitude is incorrect. Should be %f, received %f", expected, result)
	}
}

func Test_Cosine_isCorrect(t *testing.T) {
	// Arrange
	vec1 := hyperdimentional.NewVecBinomial(3)
	vec2 := hyperdimentional.NewVecBinomial(3)

	// Act
	result := hyperdimentional.Cosine(vec1, vec2)

	// Assert
	expected := hyperdimentional.Dot(vec1, vec2) / (vec1.Magnitude() * vec2.Magnitude())
	if result != expected {
		t.Fatalf("Cosine is incorrect. Should be %f, received %f", expected, result)
	}
}

func Test_Add_isCorrect(t *testing.T) {
	// Arrange
	vec1 := hyperdimentional.NewVecBinomial(3)
	vec2 := hyperdimentional.NewVecBinomial(3)
	expectedValue1 := vec1.At(0) + vec2.At(0)
	expectedValue2 := vec1.At(1) + vec2.At(1)
	expectedValue3 := vec1.At(2) + vec2.At(2)

	// Act
	vec1.Add(vec2)

	// Assert
	if expectedValue1 != vec1.At(0) || expectedValue2 != vec1.At(1) || expectedValue3 != vec1.At(2) {
		t.Fatalf("Addition is incorrect.")
	}
}
