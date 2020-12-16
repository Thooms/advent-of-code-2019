package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type Segment struct {
	StartX int64
	StartY int64
	EndX   int64
	EndY   int64
}

func debugDrawSegments(vs, hs []Segment) {
	// find boundaries
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for _, v := range vs {
		if v.StartY < minY {
			minY = v.StartY
		}
		if v.EndY < minY {

		}

	}
}

// returns -1 if they do not cross
func cross(v, h Segment) (int64, int64, int64) {
	vertCross := (v.StartY <= h.StartY && h.StartY <= v.EndY) || (v.EndY <= h.StartY && h.StartY <= v.StartY)
	horCross := (h.StartX <= v.StartX && v.StartX <= h.EndX) || (h.EndX <= v.StartX && v.StartX <= h.StartX)

	if vertCross && horCross {
		crossX, crossY := v.StartX, h.StartY
		return crossX, crossY, crossX + crossY
	}
	return -1, -1, -1
}

// assume everything starts from (0, 0),
// aggregate vertical and horizontal segments
func extractSegments(steps []string) ([]Segment, []Segment) {
	vertSegm, horSegm := []Segment{}, []Segment{}
	currentX, currentY := int64(0), int64(0)
	for _, s := range steps {
		segm := Segment{StartX: currentX, StartY: currentY, EndX: currentX, EndY: currentY}
		n_, _ := strconv.Atoi(s[1:])
		n := int64(n_)
		switch s[0] {
		case 'U':
			nextY := currentY + n
			segm.EndY = nextY
			currentY = nextY
			vertSegm = append(vertSegm, segm)
		case 'D':
			nextY := currentY - n
			segm.EndY = nextY
			currentY = nextY
			vertSegm = append(vertSegm, segm)
		case 'L':
			nextX := currentX - n
			segm.EndX = nextX
			currentX = nextX
			horSegm = append(horSegm, segm)
		case 'R':
			nextX := currentX + n
			segm.EndX = nextX
			currentX = nextX
			horSegm = append(horSegm, segm)
		}
	}
	return vertSegm, horSegm
}

// Find segments, detect crossings, record distance from (0, 0), keep the minimum
func run(steps1, steps2 []string) int64 {
	verSegm1, horSeg1 := extractSegments(steps1)
	verSegm2, horSeg2 := extractSegments(steps2)

	minDist := int64(math.MaxInt64 - 1)

	for _, vs := range verSegm1 {
		for _, hs := range horSeg2 {
			x, y, d := cross(vs, hs)
			if d > 0 {
				fmt.Printf("Found crossing at (%d, %d): distance %d\n", x, y, d)
				if d < minDist {
					minDist = d
				}
			}
		}
	}
	for _, vs := range verSegm2 {
		for _, hs := range horSeg1 {
			x, y, d := cross(vs, hs)
			if d > 0 {
				fmt.Printf("Found crossing at (%d, %d): distance %d\n", x, y, d)
				if d < minDist {
					minDist = d
				}
			}
		}
	}

	return minDist
}

func main() {
	stdin := bufio.NewReader(os.Stdin)
	text, _ := ioutil.ReadAll(stdin)
	lines := strings.Split(strings.TrimSpace(string(text)), "\n")

	fmt.Println(run(
		strings.Split(strings.TrimSpace(string(lines[0])), ","),
		strings.Split(strings.TrimSpace(string(lines[1])), ","),
	))

}
