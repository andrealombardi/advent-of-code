package main

import (
	"fmt"
)


var x,y,xVel, yVel,xDrag,yDrag, minX, maxX, minY, maxY int
func main() {

	//target area: x=14..50, y=-267..-225

	result := 0
	steps := 1000
	minX = 14
	maxX = 50
	minY = -267
	maxY = -225

	for i := 0; i < steps; i++ {
		for j := -1000; j < steps; j++ {
			xVel = i
			yVel = j
			xDrag = 0
			yDrag = 0
			x = 0
			y = 0

			for {
				step()
				if hit() {
					fmt.Printf("x:%d, y:%d, xVel:%d, yVel:%d\n", x,y,xVel,yVel)
					result++
					break
				} else if overshoot() {
					break
				}
			}

		}

	}

	fmt.Printf("result: %d ", result)

}


func step() {
	x = x + xVel -xDrag
	if xVel > xDrag {
		xDrag++
	}

	y = y + yVel -yDrag
	yDrag++

	//fmt.Printf("%d,%d\n", y,x)
}

func hit() bool{
	return x >= minX && x <= maxX && y >= minY && y <= maxY
}
func overshoot() bool {
	return x > maxX || y < minY
}