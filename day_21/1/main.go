package main

import (
	"fmt"
)



var die int
var throws int

func main() {


	p1Position := 6
	p2Position := 4

	p1Score := 0
	p2Score := 0


	result := 0

	for {

		p1Position = (p1Position + roll()) % 10
		if p1Position == 0 {
			p1Score = p1Score + 10
		} else {
			p1Score = p1Score + p1Position
		}

		fmt.Printf("p1 position: %d score: %d\n",p1Position, p1Score)

		if p1Score >= 1000 {
			result = p2Score * throws
			break
		}


		p2Position = (p2Position + roll()) % 10
		if p2Position == 0 {
			p2Score = p2Score + 10
		} else {
			p2Score = p2Score + p2Position
		}
		fmt.Printf("p2 position: %d score: %d\n",p2Position, p2Score)
		if p2Score >= 1000 {
			result = p1Score * throws
			break
		}


	}




	fmt.Printf("result: %d ", result)
}

func roll() (score int){

	for i := 0; i < 3; i++ {
		die = (die + 1) % 100
		if die == 0 {
			score += 100
		} else {
			score += die
		}
	}

	throws += 3

	return score

}