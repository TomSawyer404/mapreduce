package mapreduce

// `ReduceWorker()` summaries total number of words
func ReduceWorker(in <-chan map[string]int, out chan<- map[string]int) {
	count := map[string]int{}
	defer close(out)

	for n := range in {
		for word := range n {
			count[word] += n[word]
		}
	}

	out <- count
}
