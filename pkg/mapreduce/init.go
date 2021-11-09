//! Define some global variables to communicate
//! with goroutines

package mapreduce

var (
	size      = 10
	chtxt1    = make(chan string, size) // channel to MapWorker
	chtxt2    = make(chan string, size)
	chtxt3    = make(chan string, size)
	chtxt4    = make(chan string, size)
	chmap1    = make(chan map[string]int, size) // intermediate results produced by MpaWorker
	chmap2    = make(chan map[string]int, size)
	chmap3    = make(chan map[string]int, size)
	chmap4    = make(chan map[string]int, size)
	chreduce1 = make(chan map[string]int, size) // channel to ReduceWorker
	chreduce2 = make(chan map[string]int, size)
	chtemp1   = make(chan map[string]int, size)
	chtemp2   = make(chan map[string]int, size)
)
