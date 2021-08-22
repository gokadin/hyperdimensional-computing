package compression

import (
	"fmt"
	hd "github.com/gokadin/hyperdimensional-computing/hyperdimensional/fhrr"
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
	dataSize := 9
	data := make([][]uint8, dataSize)
	positions := make([]*hd.HdVec, dataSize)
	for i := 0; i < len(data); i++ {
		data[i] = generateBits(scaleFactor)
		positions[i] = hd.Rand()
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

func encodeTree(positions []*hd.HdVec, data [][]uint8, nodeCount int) *hd.HdVec {
	numLevels := int(math.Log(float64(len(data))) / math.Log(float64(nodeCount)))

	previousBundle := make([]*hd.HdVec, len(data))
	for i := 0; i < len(data); i++ {
		previousBundle[i] = toHdVec(data[i])
	}

	for l := 0; l < numLevels; l++ {
		bundleCount := int(math.Ceil(float64(len(previousBundle)) / float64(nodeCount)))
		currentBundle := make([]*hd.HdVec, bundleCount)
		for b := 0; b < bundleCount; b++ {
			lastIndex := b*nodeCount + nodeCount
			if lastIndex >= len(previousBundle) {
				lastIndex = len(previousBundle)
			}
			currentBundle[b] = encodeBlock(positions[l*nodeCount:l*nodeCount+nodeCount], previousBundle[b*nodeCount:lastIndex])
		}
		previousBundle = currentBundle
	}

	if len(previousBundle) != 1 {
		fmt.Println("tree encoding failed")
	}

	return previousBundle[0]
}

func decodeTree(positions []*hd.HdVec, tree *hd.HdVec, nodeCount, totalCount int) [][]uint8 {
	decoded := []*hd.HdVec{tree}
	levelCount := int(math.Log(float64(totalCount))/math.Log(float64(nodeCount))) - 1
	for l := levelCount; l >= 0; l-- {
		levelDecoded := make([]*hd.HdVec, len(decoded)*nodeCount)
		for d, encoded := range decoded {
			for v := 0; v < nodeCount; v++ {
				levelDecoded[d*nodeCount+v] = hd.Unbind(positions[l*nodeCount+v], encoded)
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

func encodeBlock(positions, data []*hd.HdVec) *hd.HdVec {
	bindings := make([]*hd.HdVec, len(data))
	for i := 0; i < len(bindings); i++ {
		bindings[i] = hd.Bind(positions[i], data[i])
	}
	return hd.Bundle(bindings...)
}

func generateBits(count int) []uint8 {
	result := make([]uint8, count)

	for i := 0; i < count; i++ {
		result[i] = uint8(rand.Intn(2))
	}

	return result
}

func toHdVec(vec []uint8) *hd.HdVec {
	factor := hd.VectorDefaultSize / len(vec)
	result := hd.NewEmptyOfSize(hd.VectorDefaultSize)

	for i, value := range vec {
		for j := 0; j < factor; j++ {
			if value == 1 {
				result.Set(i*factor+j, math.Pi-0.0000001)
			} else {
				result.Set(i*factor+j, -math.Pi+0.000001)
			}
		}
	}

	return result
}

func toBits(vec *hd.HdVec) []uint8 {
	factor := hd.VectorDefaultSize / scaleFactor
	result := make([]uint8, scaleFactor)

	for i := 0; i < scaleFactor; i++ {
		ones := 0
		for j := i * factor; j < i*factor+factor; j++ {
			if vec.At(j) >= 0 {
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

func compareVectors(v1, v2 *hd.HdVec) float32 {
	errors := 0
	for i := 0; i < v1.Size(); i++ {
		if v1.At(i) != v2.At(i) {
			errors++
		}
	}

	return 100 - float32(errors)*100/float32(v1.Size())
}
