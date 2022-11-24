package stringsx

func SliceContain(sli []string, str string) bool {
	for _, s := range sli {
		if s == str {
			return true
		}
	}
	return false
}

func MergeStringSlice(sliA []string, sliB []string) (sli []string) {
	if sliA == nil || len(sliA) == 0 {
		return sliB
	}
	if sliB == nil || len(sliB) == 0 {
		return sliA
	}
	for _, ele := range sliA {
		if !SliceContain(sliB, ele) {
			sliB = append(sliB, ele)
		}
	}
	return sliB
}
