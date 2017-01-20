package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	var sum uint

	for {
		var room, check string
		var id uint
		_, err := fmt.Scanf("%s%d%s", &room, &id, &check)

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if real(room, check) {
			fmt.Println(decrypt(room, id), id)
			sum += id
		}
	}

	fmt.Println(sum)
}

func decrypt(room string, id uint) string {
	result := make([]rune, len(room))

	for i, v := range room {
		if v == '-' {
			result[i] = ' '
			continue
		}

		r := v - 'a'
		r += rune(id)
		r = r % 26
		result[i] = r + 'a'
	}

	return string(result)
}

func real(room, check string) bool {
	var counts [26]uint32

	for _, r := range room {
		if r != '-' {
			counts[r-'a']++
		}
	}

	best := make([]uint32, 0, 5)

	sort := func() {
		for i := len(best) - 1; i > 0 && counts[best[i-1]] < counts[best[i]]; i-- {
			best[i-1], best[i] = best[i], best[i-1]
		}
	}

	for i := uint32(0); i < 5; i++ {
		best = append(best, i)
		sort()
	}

	for i := uint32(5); i < 26; i++ {
		if counts[best[4]] < counts[i] {
			best[4] = i
			sort()
		}
	}

	var result [5]rune

	for i, v := range best {
		result[i] = 'a' + rune(v)
	}

	return string(result[:]) == check
}
