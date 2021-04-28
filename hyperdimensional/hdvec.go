package hyperdimensional

import (
	"encoding/json"
	"log"
	"math"
	"math/rand"
)

const VectorDefaultSize = 10000

type HdVec struct {
	values                []uint8
	magnitudeCache        float32
	isMagnitudeCacheValid bool
}

func Rand() *HdVec {
	return RandOfSize(VectorDefaultSize)
}

func RandOfSize(size int) *HdVec {
	vec := NewEmptyOfSize(size)

	for index := range vec.values {
		vec.values[index] = uint8(rand.Intn(2))
	}

	return vec
}

func Ones() *HdVec {
	vec := NewEmptyOfSize(VectorDefaultSize)

	for index := range vec.values {
		vec.values[index] = 1
	}

	return vec
}

func NewEmptyOfSize(size int) *HdVec {
	return &HdVec{
		values:                make([]uint8, size),
		magnitudeCache:        0,
		isMagnitudeCacheValid: false,
	}
}

func FromSlice(values []uint8) *HdVec {
	vec := &HdVec{values, 0, false}
	return vec
}

func FromF32Slice(values []float32) *HdVec {
	converted := make([]uint8, len(values))
	for i := 0; i < len(values); i++ {
		converted[i] = uint8(values[i])
	}
	return FromSlice(converted)
}

func FromHDVec(from *HdVec) *HdVec {
	vec := NewEmptyOfSize(from.Size())
	for i := 0; i < from.Size(); i++ {
		vec.values[i] = from.values[i]
	}
	vec.magnitudeCache = from.magnitudeCache
	vec.isMagnitudeCacheValid = from.isMagnitudeCacheValid
	return vec
}

/**
Rotate for:
[ x_0 x_1 x_2 x_3 ... x_n ]
Produces:
[ x_n x_0 x_1 x_2 x_3 ... ]
*/
func Rotate(v *HdVec, count int) *HdVec {
	r := v.Size() - count%v.Size()
	return FromSlice(append(v.values[r:], v.values[:r]...))
}

/*
 * XOR operation
 */
func (v *HdVec) Multiply(v2 *HdVec) *HdVec {
	for i := 0; i < v.Size(); i++ {
		v.values[i] = v.values[i] ^ v2.values[i]
	}
	v.isMagnitudeCacheValid = false
	return v
}

func Add(vectors ...*HdVec) *HdVec {
	if len(vectors) == 0 {
		return nil
	}
	if len(vectors) == 1 {
		return vectors[0]
	}

	threshold := len(vectors)/2 + 1
	if len(vectors)%2 == 0 {
		vectors = append(vectors, Rand())
	}

	result := NewEmptyOfSize(vectors[0].Size())
	for i := 0; i < vectors[0].Size(); i++ {
		sum := 0
		for _, v := range vectors {
			sum += int(v.values[i])
		}

		if sum >= threshold {
			result.values[i] = 1
		} else {
			result.values[i] = 0
		}
	}
	return result
}

/**
 * cos(a, b) = dot(a, b) / magnitude(a) * magnitude(b)
 */
func Cosine(v1, v2 *HdVec) float32 {
	return float32(Dot(v1, v2)) / (v1.Magnitude() * v2.Magnitude())
}

func Dot(v1, v2 *HdVec) int {
	var result int
	for i := 0; i < v1.Size(); i++ {
		result += int(v1.values[i] & v2.values[i])
	}
	return result
}

func (v *HdVec) Magnitude() float32 {
	if v.isMagnitudeCacheValid {
		return v.magnitudeCache
	}
	v.isMagnitudeCacheValid = true

	var result int
	for _, value := range v.values {
		result += int(value)
	}
	v.magnitudeCache = float32(math.Sqrt(float64(result)))
	return v.magnitudeCache
}

func CircularConvolution(v1, v2 *HdVec) *HdVec {
	if v1.Size() != v2.Size() {
		log.Fatalf("vector sizes do not match: %d and %d", v1.Size(), v2.Size())
	}

	result := NewEmptyOfSize(v1.Size())
	for j := 0; j < v1.Size(); j++ {
		for k := 0; k < v1.Size(); k++ {
			v2Index := ((j-k)%v1.Size() + v1.Size()) % v1.Size()
			result.Set(j, result.At(j)+v1.At(k)*v2.At(v2Index))
		}
	}

	return result
}

//func CircularConvolution2(v1, v2 *HdVec) *HdVec {
//	if v1.Size() != v2.Size() {
//		log.Fatalf("vector sizes do not match: %d and %d", v1.Size(), v2.Size())
//	}
//
//	v1fft := fft.FFTReal(v1.ToF64())
//	v2fft := fft.FFTReal(v2.ToF64())
//	mul := make([]complex128, v1.Size())
//	for i := 0; i < v1.Size(); i++ {
//		mul[i] = v1fft[i] * v2fft[i]
//	}
//	c := fft.IFFT(mul)
//	result := make([]float32, v1.Size())
//	for i := 0; i < v1.Size(); i++ {
//		result[i] = float32(real(c[i]))
//	}
//	v := FromF32Slice(result)
//	return v
//}

/**
Involution for:
[ x_0 x_1 x_2 x_3 x_n ... ]
Produces:
[ x_0 x_n ... x_3 x_2 x_1 ]
*/
func (v *HdVec) Involution() *HdVec {
	for i := 1; i < v.Size(); i++ {
		v.Set(i, v.At(v.Size()-i))
	}
	return v
}

func (v *HdVec) Size() int {
	return len(v.values)
}

func (v *HdVec) ToF64() []float64 {
	result := make([]float64, v.Size())
	for i := 1; i < v.Size(); i++ {
		result[i] = float64(v.values[v.Size()-i])
	}
	return result
}

func (v *HdVec) At(index int) uint8 {
	return v.values[index]
}

func (v *HdVec) Set(index int, value uint8) {
	v.values[index] = value
	v.isMagnitudeCacheValid = false
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
