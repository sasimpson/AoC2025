package main

import (
	"bufio"
	"fmt"
	"io"
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

func (g *Group) Invalid() []int {
	var invalids []int
	if g.Start > g.End {
		return []int{}
	}

	for i := g.Start; i <= g.End; i++ {
		x := strconv.Itoa(i)
		if len(x)%2 == 0 && (x[:len(x)/2] == x[len(x)/2:]) {
			invalids = append(invalids, i)
		}
	}
	return invalids
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
	fp, err := os.Open("cmd/day2/data/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)
	var groups []*Group

	for {
		groupStr, err := reader.ReadString(',')
		if err != nil {
			if err == io.EOF {
				groups = append(groups, NewGroup(groupStr))
			}
			break
		}
		groups = append(groups, NewGroup(strings.TrimSuffix(groupStr, ",")))
	}

	var groupSum int
	for _, group := range groups {
		fmt.Println(group, group.Invalid())
		for _, i := range group.Invalid() {
			groupSum += i
		}
	}
	fmt.Println("sum of invalid ids", groupSum)
}
