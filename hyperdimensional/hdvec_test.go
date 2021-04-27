package hyperdimensional

import (
	"github.com/stretchr/testify/suite"
	"math"
	"math/rand"
	"testing"
	"time"
)

type HDVecTestSuite struct {
	suite.Suite
}

func TestHDVecTestSuite(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	suite.Run(t, new(HDVecTestSuite))
}

func (suite *HDVecTestSuite) Test_default_size() {
	vector := Rand()

	suite.Equal(VectorDefaultSize, vector.Size())
}

func (suite *HDVecTestSuite) Test_Rand_InitializesWithFalseMagnitudeCacheValidity() {
	vector := Rand()

	suite.False(vector.isMagnitudeCacheValid)
}

func (suite *HDVecTestSuite) Test_RandOfSize_InitializesWithFalseMagnitudeCacheValidity() {
	vector := RandOfSize(100)

	suite.False(vector.isMagnitudeCacheValid)
}

func (suite *HDVecTestSuite) Test_Ones_InitializesWithFalseMagnitudeCacheValidity() {
	vector := Ones()

	suite.False(vector.isMagnitudeCacheValid)
}

func (suite *HDVecTestSuite) Test_NewEmptyOfSize_InitializesWithFalseMagnitudeCacheValidity() {
	vector := NewEmptyOfSize(100)

	suite.False(vector.isMagnitudeCacheValid)
}

func (suite *HDVecTestSuite) Test_FromSlice_InitializesWithFalseMagnitudeCacheValidity() {
	vector := FromSlice([]uint8{0, 1, 0, 1, 1})

	suite.False(vector.isMagnitudeCacheValid)
}

func (suite *HDVecTestSuite) Test_Rotate() {
	vector := FromSlice([]uint8{1, 2, 3, 4, 5})
	expected := FromSlice([]uint8{5, 1, 2, 3, 4})

	rotated := Rotate(vector, 1)

	suite.Equal(expected.values, rotated.values)
}

func (suite *HDVecTestSuite) Test_Rotate_ofZero() {
	vector := RandOfSize(3)

	rotated := Rotate(vector, 0)

	suite.True(Equal(vector, rotated))
}

func (suite *HDVecTestSuite) Test_Rotate_isCorrectWhenCountIsMoreThanOne() {
	vector := FromSlice([]uint8{1, 2, 3, 4, 5})
	expected := FromSlice([]uint8{3, 4, 5, 1, 2})

	rotated := Rotate(vector, 3)

	suite.Equal(expected.values, rotated.values)
}

func (suite *HDVecTestSuite) Test_Multiply() {
	vec1 := FromSlice([]uint8{1, 1, 0, 0})
	vec2 := FromSlice([]uint8{1, 0, 1, 0})

	vec1.Multiply(vec2)

	suite.Equal(len(vec2.values), vec1.Size())
	suite.Equal(uint8(0), vec1.At(0))
	suite.Equal(uint8(1), vec1.At(1))
	suite.Equal(uint8(1), vec1.At(2))
	suite.Equal(uint8(0), vec1.At(3))
}

func (suite *HDVecTestSuite) Test_Multiply_invalidatesMagnitudeCache() {
	vec1 := FromSlice([]uint8{1, 1, 0, 0})
	vec1.isMagnitudeCacheValid = true
	vec2 := FromSlice([]uint8{1, 0, 1, 0})

	vec1.Multiply(vec2)

	suite.False(vec1.isMagnitudeCacheValid)
}

func (suite *HDVecTestSuite) Test_Dot() {
	vec1 := RandOfSize(3)
	vec2 := RandOfSize(3)

	dot := Dot(vec1, vec2)

	expected := int(vec1.At(0)*vec2.At(0) + vec1.At(1)*vec2.At(1) + vec1.At(2)*vec2.At(2))
	suite.Equal(expected, dot)
}

func (suite *HDVecTestSuite) Test_Magnitude() {
	vec := RandOfSize(3)

	result := vec.Magnitude()

	expected := float32(math.Sqrt(float64(vec.At(0)*vec.At(0) + vec.At(1)*vec.At(1) + vec.At(2)*vec.At(2))))
	suite.Equal(expected, result)
}

func (suite *HDVecTestSuite) Test_Cosine() {
	vec1 := RandOfSize(100)
	vec2 := RandOfSize(100)

	result := Cosine(vec1, vec2)

	dot := Dot(vec1, vec2)
	var expected float32
	if dot != 0 {
		expected = float32(dot) / (vec1.Magnitude() * vec2.Magnitude())
	}
	suite.Equal(expected, result)
}

func (suite *HDVecTestSuite) Test_Add_oddVectors() {
	vec1 := FromSlice([]uint8{1, 1, 1, 0})
	vec2 := FromSlice([]uint8{1, 1, 0, 0})
	vec3 := FromSlice([]uint8{1, 0, 0, 0})

	result := Add(vec1, vec2, vec3)

	suite.Equal(4, result.Size())
	suite.Equal(uint8(1), result.At(0))
	suite.Equal(uint8(1), result.At(1))
	suite.Equal(uint8(0), result.At(2))
	suite.Equal(uint8(0), result.At(3))
}

func (suite *HDVecTestSuite) Test_Add_oddMultipleVectors() {
	vec1 := FromSlice([]uint8{1, 1, 1, 1, 1, 0})
	vec2 := FromSlice([]uint8{1, 1, 1, 1, 0, 0})
	vec3 := FromSlice([]uint8{1, 1, 1, 0, 0, 0})
	vec4 := FromSlice([]uint8{1, 1, 0, 0, 0, 0})
	vec5 := FromSlice([]uint8{1, 0, 0, 0, 0, 0})

	result := Add(vec1, vec2, vec3, vec4, vec5)

	suite.Equal(6, result.Size())
	suite.Equal(uint8(1), result.At(0))
	suite.Equal(uint8(1), result.At(1))
	suite.Equal(uint8(1), result.At(2))
	suite.Equal(uint8(0), result.At(3))
	suite.Equal(uint8(0), result.At(4))
	suite.Equal(uint8(0), result.At(5))
}

func (suite *HDVecTestSuite) Test_Add_evenVectors() {
	vec1 := FromSlice([]uint8{1, 1, 0, 0})
	vec2 := FromSlice([]uint8{1, 0, 1, 0})

	result := Add(vec1, vec2)

	suite.Equal(4, result.Size())
	suite.Equal(uint8(1), result.At(0))
	suite.Equal(uint8(0), result.At(3))
	// cannot verify indices 1 and 2 since randomness is involved
}

func (suite *HDVecTestSuite) TestEqual_whenEqual() {
	a := NewEmptyOfSize(3)
	a.Set(0, 1)
	a.Set(1, 2)
	a.Set(2, 3)

	b := NewEmptyOfSize(3)
	b.Set(0, 1)
	b.Set(1, 2)
	b.Set(2, 3)

	suite.True(Equal(a, b))
}

func (suite *HDVecTestSuite) TestEqual_whenNotEqual() {
	a := NewEmptyOfSize(3)
	a.Set(0, 1)
	a.Set(1, 2)
	a.Set(2, 3)

	b := NewEmptyOfSize(3)
	b.Set(0, 1)
	b.Set(1, 2)
	b.Set(2, 4)

	suite.False(Equal(a, b))
}

func (suite *HDVecTestSuite) Test_Set_invalidatesMagnitudeCache() {
	a := RandOfSize(50)
	a.isMagnitudeCacheValid = true

	a.Set(0, 1)

	suite.False(a.isMagnitudeCacheValid)
}
