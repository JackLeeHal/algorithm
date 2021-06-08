package binarytree

func Quicksort(in []int) []int {
	sort(in, 0, len(in)-1)

	return in
}

func sort(in []int, start, end int) {
	if start < end {
		// partition
		p := in[end]
		i := start
		for j := start; j < end; j++ {
			if in[j] < p {
				in[i], in[j] = in[j], in[i]
				i++
			}
		}

		in[i], in[end] = in[end], in[i]
		sort(in, 0, i-1)
		sort(in, i+1, end)
	}
}
