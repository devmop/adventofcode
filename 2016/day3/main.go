package main

import "fmt"
import "io"

func main() {
	var t [3]triangle

	var count uint

	for {
		fmt.Scan(&t[0].a, &t[1].a, &t[2].a)
		fmt.Scan(&t[0].b, &t[1].b, &t[2].b)
		_, err := fmt.Scan(&t[0].c, &t[1].c, &t[2].c)

		if err == io.EOF {
			break
		}

		for i := 0; i < 3; i++ {
			if t[i].poss() {
				count++
			}
		}
	}

	fmt.Println(count)
}

type triangle struct {
	a, b, c uint
}

func (t *triangle) poss() bool {
	a := t.a
	b := t.b
	c := t.c
	return a+b > c && a+c > b && b+c > a
}
