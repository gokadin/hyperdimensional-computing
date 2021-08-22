//package hyperdimensional
//
//func FindClosestCosineInMap(target *HdVec, sample map[string]*HdVec) string {
//	var match string
//	var closestCosine float32 = -2
//	for name, vec := range sample {
//		cosine := Cosine(target, vec)
//		if cosine > closestCosine {
//			closestCosine = cosine
//			match = name
//		}
//	}
//
//	return match
//}
//
//func FindClosestCosineInMapInt(target *HdVec, sample map[int]*HdVec) int {
//	var match int
//	var closestCosine float32 = -2
//	for key, vec := range sample {
//		cosine := Cosine(target, vec)
//		if cosine > closestCosine {
//			closestCosine = cosine
//			match = key
//		}
//	}
//
//	return match
//}
