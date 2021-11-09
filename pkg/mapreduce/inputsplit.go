package mapreduce

import "strings"

// `InputSplit()` split txt into 4 pieces, and then calcutelate it
func InputSplit(txt string, out [4]chan<- string) { // `out` is a write-only channel
	txt = strings.Replace(txt, "\n", ",", -1)
	txt = strings.Replace(txt, "\t", ",", -1)
	txt = strings.Replace(txt, `!`, ",", -1)
	txt = strings.Replace(txt, `.`, ",", -1)
	txt = strings.Replace(txt, `"`, ",", -1)

	arr := strings.Split(txt, ",")
	mod := len(arr) % 4
	for i := 0; i < (4 - mod); i += 1 {
		arr = append(arr, "")
	}

	lenSplit := len(arr) / 4
	for i := 0; i < 4; i += 1 {
		go func(ch chan<- string, lines []string) { // `ch` is a write-only channel
			for _, line := range lines {
				word := strings.Split(line, " ")
				for _, w := range word {
					w = strings.Trim(w, " ")
					ch <- w
				}
			}

			close(ch)
		}(out[i], arr[i*lenSplit:(i+1)*lenSplit])
	}
}
