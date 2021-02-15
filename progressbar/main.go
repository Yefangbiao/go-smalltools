package main

import "math/rand"

var cur int64 = 23414
var total int64 = 1238962178

func main() {
	bar := NewBar(cur, total)
	for i := cur; i < total; {
		next := rand.Int63n(cur)
		i += next
		if i > total {
			i = total
		}
		bar.Display(i)
	}
}
