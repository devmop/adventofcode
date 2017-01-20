package main

import "fmt"
import "io"
import "log"

func main() {
	instructions := make([]instruction, 0, 4096)

	for {
		var turn string
		var blocks uint32
		_, err := fmt.Scanf("%1s%d", &turn, &blocks)

		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		instructions = append(instructions, instruction{blocks, turn == "L"})
	}

	res := follow(instructions)

	fmt.Println("Distance is", res.abs())
}

const (
	NORTH = 0
	EAST  = 1
	SOUTH = 2
	WEST  = 3
)

func follow(directions []instruction) position {
	visited := make(map[position]bool)
	var p position
	facing := NORTH

	for _, d := range directions {
		if d.left {
			facing = (facing + 3) % 4
		} else {
			facing = (facing + 1) % 4
		}

		for i := uint32(0); i < d.distance; i++ {
			switch facing {
			case NORTH:
				p.y += 1
			case SOUTH:
				p.y -= 1
			case EAST:
				p.x += 1
			case WEST:
				p.x -= 1
			}

			if visited[p] {
				return p
			}

			visited[p] = true
		}
	}

	return p
}

type position struct {
	x, y int64
}

type instruction struct {
	distance uint32
	left     bool
}

func (p *position) abs() (res int64) {
	if p.x < 0 {
		res += -p.x
	} else {
		res += p.x
	}

	if p.y < 0 {
		res += -p.y
	} else {
		res += p.y
	}

	return
}
