package hyperdimensional

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

const (
	a = 97
	b = 98
	c = 99
)

type EncoderTestSuite struct {
	suite.Suite
	letters map[int]*HdVec
	encoder *Encoder
}

func (suite *EncoderTestSuite) SetupTest() {
	suite.letters = make(map[int]*HdVec, 3)
	suite.letters[a] = Rand()
	suite.letters[b] = Rand()
	suite.letters[c] = Rand()
	suite.encoder = NewEncoder(suite.letters)
}

func TestEncoderTestSuite(t *testing.T) {
	suite.Run(t, new(EncoderTestSuite))
}

func (suite *EncoderTestSuite) TestEncodeVec_initializesARandomVec() {
	suite.NotNil(suite.encoder.randomVec)
}

func (suite *EncoderTestSuite) TestEncodeVec_gramFactor1() {
	text := "abc"
	expected := Add(suite.letters[a], suite.letters[b], suite.letters[c])

	result := suite.encoder.EncodeVec(text, 1)

	suite.True(Equal(expected, result))
}

func (suite *EncoderTestSuite) TestEncodeVec_gramFactor2() {
	text := "abca"
	expected := Add(
		Rotate(suite.letters[a], 1).Multiply(suite.letters[b]),
		Rotate(suite.letters[b], 1).Multiply(suite.letters[c]),
		Rotate(suite.letters[c], 1).Multiply(suite.letters[a]),
	)

	result := suite.encoder.EncodeVec(text, 2)

	suite.True(Equal(expected, result))
}

func (suite *EncoderTestSuite) TestEncodeVec_evenNumberOfGramsUsesTheRandomVecToBalanceTheAddOperation() {
	text := "abc"
	expected := Add(
		Rotate(suite.letters[a], 1).Multiply(suite.letters[b]),
		Rotate(suite.letters[b], 1).Multiply(suite.letters[c]),
		suite.encoder.randomVec,
	)

	result := suite.encoder.EncodeVec(text, 2)

	suite.True(Equal(expected, result))
}

func (suite *EncoderTestSuite) TestEncodeVec_gramFactor3() {
	text := "abcac"
	expected := Add(
		Rotate(suite.letters[a], 2).Multiply(Rotate(suite.letters[b], 1)).Multiply(suite.letters[c]),
		Rotate(suite.letters[b], 2).Multiply(Rotate(suite.letters[c], 1)).Multiply(suite.letters[a]),
		Rotate(suite.letters[c], 2).Multiply(Rotate(suite.letters[a], 1)).Multiply(suite.letters[c]),
	)

	result := suite.encoder.EncodeVec(text, 3)

	suite.True(Equal(expected, result))
}

func (suite *EncoderTestSuite) TestEncodeVec_gramFactor3OnTextOfLength1() {
	text := "a"
	expected := FromHDVec(suite.letters[a])

	result := suite.encoder.EncodeVec(text, 3)

	suite.True(Equal(expected, result))
}

func (suite *EncoderTestSuite) TestEncodeVec_gramFactor3OnTextOfLength2() {
	text := "ab"
	expected := FromHDVec(Rotate(suite.letters[a], 1).Multiply(suite.letters[b]))

	result := suite.encoder.EncodeVec(text, 3)

	suite.True(Equal(expected, result))
}

func (suite *EncoderTestSuite) TestEncodeVec_textOfLength0ReturnsOnes() {
	result := suite.encoder.EncodeVec("", 3)

	suite.Equal(Ones(), result)
}

func (suite *EncoderTestSuite) TestEncodeVec_panicsOnGramFactorZero() {
	suite.Panics(func() {
		suite.encoder.EncodeVec("test", 0)
	})
}

func (suite *EncoderTestSuite) TestEncodeVec_accuracy() {
	result := suite.encoder.EncodeVec("", 3)

	suite.Equal(Ones(), result)
}
