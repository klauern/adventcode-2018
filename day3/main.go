package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/klauern/adventcode-2018/helpers"
)

type claim struct {
	id   int
	area *area
}

type area struct {
	startX int
	startY int
	width  int
	height int
}

var fabric = [1000][1000]int{}

func mustConv(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return num
}

func parseClaim(line string) *claim {
	strAry := strings.Split(line, " ")
	claimID := mustConv(strings.TrimLeft(strAry[0], "#"))
	xy := strings.Split(strAry[2], ",")
	startX := mustConv(xy[0])
	startY := mustConv(strings.TrimRight(xy[1], ":"))
	sizes := strings.Split(strAry[3], "x")
	width := mustConv(sizes[0])
	height := mustConv(sizes[1])
	return &claim{
		id: claimID,
		area: &area{startX: startX,
			startY: startY,
			width:  width,
			height: height,
		},
	}
}

func (c *claim) overlaps(d *claim) bool {
	return false
}

func (c *claim) findOverlap(d *claim) *area {
	return nil
}

func applyClaimToFabric(c *claim) {
	for i := c.area.startX; i < c.area.startX+c.area.width; i++ {
		for j := c.area.startY; j < c.area.startY+c.area.height; j++ {
			fabric[i][j]++
		}
	}
}

func countFabricOverlap(fabric [1000][1000]int) int {
	total := 0
	for x, a := range fabric {
		for y := range a {
			if fabric[x][y] > 1 {
				total++
			}
		}
	}
	return total
}

func main() {
	inputAry := helpers.ToArray(helpers.MustScanFile("input.txt"))
	for _, v := range inputAry {
		claim := parseClaim(v)
		applyClaimToFabric(claim)
	}
	fmt.Printf("Total Count of overlapped: %v\n", countFabricOverlap(fabric))
}
