package hyperdimentional

import (
	"github.com/gokadin/hyperdimentional/src/hyperdimentional"
	"testing"
)

func Test_Size_isCorrect(t *testing.T) {
    hdvec := hyperdimentional.NewHdVecBinomial(10000)

    if hdvec.Size() != 10000 {
    	t.Fail()
	}
}

func Test_Rotate_isCorrect(t *testing.T) {
	hdvec := hyperdimentional.NewHdVecBinomial(10)

	rotated := hyperdimentional.Rotate(hdvec)

	if rotated.Size() != hdvec.Size() {
		t.Fatalf("Size does not match.")
	}

	if rotated.First() != hdvec.Get(1) || rotated.Last() != hdvec.First() {
		t.Fatalf("Rotation failed.")
	}
}

func Test_Multiply_isCorrect(t *testing.T) {
	hdvecA := hyperdimentional.NewHdVecBinomial(10)
	hdvecB := hyperdimentional.NewHdVecBinomial(10)

	multiplied := hyperdimentional.Multiply(hdvecA, hdvecB)

	if multiplied.Size() != hdvecA.Size() {
		t.Fatalf("Size does not match.")
	}

	for index, value := range *multiplied.Values() {
		if value != hdvecA.Get(index) * hdvecB.Get(index) {
			t.Fatalf("Multiplication failed.")
		}
	}
}
