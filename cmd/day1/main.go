package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Move struct {
	Direction string
	Distance  int
}

type Dial struct {
	Position int
	Size     int
	Moves    []*Move
	Zeros    int
}

func (d *Dial) String() string {
	return fmt.Sprintf("pos: %d, moves: %d, zeros: %d", d.Position, len(d.Moves), d.Zeros)
}

func (d *Dial) IsZero() bool {
	return d.Position == 0
}

func (d *Dial) Move(m *Move) {
	d.Moves = append(d.Moves, m)

	switch m.Direction {
	case "R":
		d.Zeros += (d.Position + m.Distance) / d.Size
		d.Position = (d.Position + m.Distance) % d.Size
	case "L":
		if m.Distance >= d.Position {
			if d.Position > 0 {
				d.Zeros += 1 + (m.Distance-d.Position)/d.Size
			} else {
				d.Zeros += m.Distance / d.Size
			}
		}
		d.Position = ((d.Position-m.Distance)%d.Size + d.Size) % d.Size
	}
}

func ParseMove(input string) (*Move, error) {
	var m = &Move{}
	m.Direction = input[0:1]
	dist, err := strconv.Atoi(input[1:])
	if err != nil {
		return m, err
	}
	m.Distance = dist
	return m, nil
}

func (m Move) String() string {
	return fmt.Sprintf("%s:%d", m.Direction, m.Distance)
}

func main() {
	var dial = &Dial{Position: 50, Size: 100}
	// var counter = 0

	fp, err := os.Open("cmd/day1/data/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(fp)

	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		if err != nil {
			break
		}
		m, err := ParseMove(line)
		if err != nil {
			log.Fatal(err)
		}
		dial.Move(m)
		fmt.Println(m, dial)

	}

	fmt.Println("number of zeros", dial.Zeros)
}
