package mapreduce

// `MapWorker()` counts the number of words, and prepares data for
// `ReduceWorker()`. It stores result into a temp `map[string]int` object,
// aka *key-value pair*. The `key` stands for a `word`, the `value` stands
// for the number of that `word`.
// We will sent the result to `Master` node finally.
func MapWorker(in <-chan string, out chan<- map[string]int) { // `in` is a read-only channel
	count := map[string]int{}
	defer close(out)

	for word := range in {
		if word != "" {
			count[word] += 1
		}
	}

	out <- count
}
