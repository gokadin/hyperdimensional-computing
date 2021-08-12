package compression

import (
	"fmt"
	"github.com/gokadin/hyperdimensional-computing/hyperdimensional"
	"math"
	"math/rand"
)

const (
	scaleFactor = 100
	nodeCount   = 3
)

type example struct {
}

func NewExample() *example {
	return &example{}
}

func (e *example) Run() {
	dataSize := 243
	data := make([][]uint8, dataSize)
	positions := make([]*hyperdimensional.HdVec, dataSize)
	for i := 0; i < len(data); i++ {
		data[i] = generateBits(scaleFactor)
		positions[i] = hyperdimensional.Rand()
	}

	y := encodeTree(positions, data, nodeCount)
	decoded := decodeTree(positions, y, nodeCount, len(data))

	fmt.Println("decoded", len(decoded))

	var averageAccuracy float32
	for i := 0; i < len(data); i++ {
		accuracy := compareBits(decoded[i], data[i])
		averageAccuracy += accuracy
		//fmt.Printf("%d -> %.2f%%\n", i, accuracy)
	}
	fmt.Printf("accuracy: %.2f%%", averageAccuracy/float32(len(data)))

	// n < 3 = 1
	// n > 3 = 2 (3^1)
	// n > 9 = 3 (3^2)
	// n > 27 = 4 (3^3)
	// n > 81 = 5 (3^4)
	// n > 243 = 5 (3^5)
}

func encodeTree(positions []*hyperdimensional.HdVec, data [][]uint8, nodeCount int) *hyperdimensional.HdVec {
	numLevels := int(math.Log(float64(len(data))) / math.Log(float64(nodeCount)))

	previousBundle := make([]*hyperdimensional.HdVec, len(data))
	for i := 0; i < len(data); i++ {
		previousBundle[i] = toHdVec(data[i])
	}

	for l := 0; l < numLevels; l++ {
		bundleCount := int(math.Ceil(float64(len(previousBundle)) / float64(nodeCount)))
		currentBundle := make([]*hyperdimensional.HdVec, bundleCount)
		for b := 0; b < bundleCount; b++ {
			lastIndex := b*nodeCount + nodeCount
			if lastIndex >= len(previousBundle) {
				lastIndex = len(previousBundle)
			}
			currentBundle[b] = encodeBlock(positions[l*nodeCount:l*nodeCount+nodeCount], previousBundle[b*nodeCount:lastIndex])
		}
		fmt.Println("encoded level", l, "with", len(currentBundle))
		previousBundle = currentBundle
	}

	if len(previousBundle) != 1 {
		fmt.Println("tree encoding failed")
	}

	return previousBundle[0]
}

func decodeTree(positions []*hyperdimensional.HdVec, tree *hyperdimensional.HdVec, nodeCount, totalCount int) [][]uint8 {
	decoded := []*hyperdimensional.HdVec{tree}
	levelCount := int(math.Log(float64(totalCount))/math.Log(float64(nodeCount))) - 1
	for l := levelCount; l >= 0; l-- {
		levelDecoded := make([]*hyperdimensional.HdVec, len(decoded)*nodeCount)
		for d, encoded := range decoded {
			for v := 0; v < nodeCount; v++ {
				levelDecoded[d*nodeCount+v] = hyperdimensional.Xor(positions[l*nodeCount+v], encoded)
			}
		}
		decoded = levelDecoded
	}

	result := make([][]uint8, len(decoded))
	for i := 0; i < len(decoded); i++ {
		result[i] = toBits(decoded[i])
	}

	return result
}

func encodeBlock(positions, data []*hyperdimensional.HdVec) *hyperdimensional.HdVec {
	bindings := make([]*hyperdimensional.HdVec, len(data))
	for i := 0; i < len(bindings); i++ {
		bindings[i] = hyperdimensional.Xor(positions[i], data[i])
	}
	return hyperdimensional.Add(bindings...)
}

func generateBits(count int) []uint8 {
	result := make([]uint8, count)

	for i := 0; i < count; i++ {
		result[i] = uint8(rand.Intn(2))
	}

	return result
}

func toHdVec(vec []uint8) *hyperdimensional.HdVec {
	factor := hyperdimensional.VectorDefaultSize / len(vec)
	result := hyperdimensional.NewEmptyOfSize(hyperdimensional.VectorDefaultSize)

	for i, value := range vec {
		for j := 0; j < factor; j++ {
			result.Set(i*factor+j, value)
		}
	}

	return result
}

func toBits(vec *hyperdimensional.HdVec) []uint8 {
	factor := hyperdimensional.VectorDefaultSize / scaleFactor
	result := make([]uint8, scaleFactor)

	for i := 0; i < scaleFactor; i++ {
		ones := 0
		for j := i * factor; j < i*factor+factor; j++ {
			if vec.At(j) == 1 {
				ones++
			}
			if ones >= factor/2 {
				result[i] = 1
				break
			}
		}
	}

	return result
}

func compareBits(v1, v2 []uint8) float32 {
	if len(v1) != len(v2) {
		panic("sizes don't match")
	}

	errors := 0
	for i := 0; i < len(v1); i++ {
		if v1[i] != v2[i] {
			errors++
		}
	}

	return 100 - float32(errors)*100/float32(len(v1))
}

func compareVectors(v1, v2 *hyperdimensional.HdVec) float32 {
	errors := 0
	for i := 0; i < v1.Size(); i++ {
		if v1.At(i) != v2.At(i) {
			errors++
		}
	}

	return 100 - float32(errors)*100/float32(v1.Size())
}
