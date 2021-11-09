package mapreduce

import "sync"

func SumReduce(in []<-chan map[string]int) map[string]int {
	count := map[string]int{}
	var mLock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(in))

	for i := 0; i < len(in); i += 1 {
		go func(n int, c <-chan map[string]int) {
			for m := range c {
				for word := range m {
					mLock.Lock()
					count[word] += m[word]
					mLock.Unlock()
				}
			}
			wg.Done()
		}(i, in[i])
	}

	wg.Wait()
	return count
}
