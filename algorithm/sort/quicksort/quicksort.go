package main

func quickSort(s []int, left int, right int) {
	if left >= right {
		return
	}

	v := s[left]
	l := left
	r := right
	for l < r {
		for l < r && s[r] > v {
			r--
		}
		if l < r {
			s[l] = s[r]
		}

		for l < r && s[l] < v {
			l++
		}
		if l < r {
			s[r] = s[l]
		}
	}

	s[l] = v

	quickSort(s, left, l-1)
	quickSort(s, l+1, right)
}

func quickSort3(s []int, left int, right int) {
	if left >= right {
		return
	}

	// 三数取中，减少最坏情况概率
	mid := left + (right-left)/2
	if s[left] > s[mid] {
		s[left], s[mid] = s[mid], s[left]
	}
	if s[mid] > s[right] {
		s[mid], s[right] = s[right], s[mid]
	}
	if s[left] < s[mid] {
		s[left], s[mid] = s[mid], s[left]
	}
	// pivot 放在 left
	v := s[left]

	l := left
	r := right

	for l < r {
		for l < r && s[r] >= v {
			r--
		}
		if l < r {
			s[l] = s[r]
		}

		for l < r && s[l] <= v {
			l++
		}
		if l < r {
			s[r] = s[l]
		}
	}

	s[l] = v

	quickSort(s, left, l-1)
	quickSort(s, l+1, right)
}
