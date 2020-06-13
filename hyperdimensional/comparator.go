package hyperdimensional

func FindClosestCosineInMap(target *HdVec, sample map[string]*HdVec) string {
	var match string
	var closestCosine float32 = -2
	for name, vec := range sample {
		cosine := Cosine(target, vec)
		if cosine > closestCosine {
			closestCosine = cosine
			match = name
		}
	}

	return match
}
