package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day_04/2/input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	numbers := make([]string, 0)
	game := make([]board, 0)

	b := board{}
	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, ",") {
			numbers = strings.Split(line, ",")
		} else if line == "" {
			if len(b.rows) > 0 {
				game = append(game, b)
				b = board{}
			}
			continue
		} else {
			rowString := strings.Fields(line)
			b.addRow(rowString)
		}
	}

	for _, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		for _, board := range game {
			board.draw(n)
			if lastBingo(game){
				fmt.Printf("Bingo! %d\n", board.score(n))
				fmt.Printf("Board :%v\n", board)
				fmt.Printf("N :%v\n", n)
				os.Exit(1)
			}
		}
	}

}

type board struct {
	rows [][]entry
}

type entry struct {
	value  int
	marked bool
}

func lastBingo(game []board) bool{
	lastBingo := true
	for _, board := range game {
		lastBingo = lastBingo && board.bingo()
	}
	return lastBingo
}

func (e entry) checkAndMark(number int) entry{
	if e.value == number {
		e.marked = true
	}
	return e
}

func (b board) draw(number int) {
	for i, _ := range b.rows {
		for j, e := range b.rows[i] {
			b.rows[i][j] = e.checkAndMark(number)
		}
	}
}

func (b *board) addRow(numbers []string){

	r := make([]entry, 0)
	for _, number := range numbers {
		n, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		r = append(r, entry{n, false})
	}
	b.rows = append(b.rows, r)
}

func (b board) bingo() bool {
	for r:= 0 ; r < 5; r++ {
		isBingo := true
		for c := 0 ; c < 5; c++{
			isBingo = isBingo && b.rows[r][c].marked
		}
		if isBingo {
			return true
		}
	}

	for c:= 0 ; c < 5; c++ {
		isBingo := true
		for r := 0 ; r < 5; r++{
			isBingo = isBingo && b.rows[r][c].marked
		}
		if isBingo {
			return true
		}
	}
	return false
}

func (b board) score(last int) (total int) {
	for i, _ := range b.rows {
		for _, e := range b.rows[i] {
			if !e.marked {
				total += e.value
			}
		}
	}
	return total * last
}