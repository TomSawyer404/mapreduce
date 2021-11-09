package mapreduce

import "sync"

// `Master()` need to sechedule the tasks.
// In this case, it assigns 4 tasks to 2 ReduceWorker.
func Master(in []<-chan map[string]int, out [2]chan<- map[string]int) {
	var wg sync.WaitGroup
	wg.Add(len(in))
	var mLock sync.Mutex
	i := 0

	for _, ch := range in {
		go func(c <-chan map[string]int) {
			for m := range c {
				mLock.Lock()
				randChoice := i % 2
				mLock.Unlock()

				if randChoice == 0 {
					out[0] <- m
				} else {
					out[1] <- m
				}
			}
			wg.Done()

			mLock.Lock()
			i += 1
			mLock.Unlock()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out[0])
		close(out[1])
	}()
}
