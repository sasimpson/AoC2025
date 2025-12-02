package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Group struct {
	raw   string
	Start int
	End   int
}

func (g *Group) String() string {
	return fmt.Sprintf("%d-%d", g.Start, g.End)
}

func NewGroup(raw string) *Group {
	ids := strings.Split(raw, "-")
	if len(ids) != 2 {
		return nil
	}
	start, err := strconv.Atoi(ids[0])
	if err != nil {
		return nil
	}
	end, err := strconv.Atoi(ids[1])
	if err != nil {
		return nil
	}
	return &Group{raw: raw, Start: start, End: end}
}

func main() {
	fp, err := os.Open("cmd/day2/data/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	for {
		groupStr, err := reader.ReadString(',')
		if err != nil {
			break
		}
		group := NewGroup(strings.TrimSuffix(groupStr, ","))
		fmt.Println(group)
	}
}
