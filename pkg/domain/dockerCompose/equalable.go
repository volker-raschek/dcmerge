package dockerCompose

type Equalable interface {
	Equal(equalable Equalable) bool
}

// Contains returns true when sliceA is in sliceB.
func Contains[R Equalable](sliceA, sliceB []R) bool {
	switch {
	case sliceA == nil && sliceB == nil:
		return true
	case sliceA != nil && sliceB == nil:
		return false
	case sliceA == nil && sliceB != nil:
		return false
	default:
	LOOP:
		for i := range sliceA {
			for j := range sliceB {
				if sliceA[i].Equal(sliceB[j]) {
					continue LOOP
				}
			}
			return false
		}
		return true
	}
}

// Equal returns true when sliceA and sliceB are equal.
func Equal[R Equalable](sliceA, sliceB []R) bool {
	return Contains(sliceA, sliceB) &&
		Contains(sliceB, sliceA) &&
		len(sliceA) == len(sliceB)
}

// Equal returns true when booth string maps of Equalable are equal.
func EqualStringMap[R Equalable](mapA, mapB map[string]R) bool {
	equalFunc := func(mapA, mapB map[string]R) bool {
	LOOP:
		for keyA, valueA := range mapA {
			for keyB, valueB := range mapB {
				if keyA == keyB &&
					valueA.Equal(valueB) {
					continue LOOP
				}
			}
			return false
		}
		return true
	}

	return equalFunc(mapA, mapB) && equalFunc(mapB, mapA)
}

// ExistsInMap returns true if object of type any exists under the passed name.
func ExistsInMap[T any](m map[string]T, name string) bool {
	switch m {
	case nil:
		return false
	default:
		_, present := m[name]
		return present
	}
}
