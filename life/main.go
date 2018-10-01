package main

import (
	"gopkg.in/gizak/termui.v1"
	"time"
)

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	type gameStatus struct {
		current bool
		previous  bool
	}

	var gameField = [5][5]gameStatus{}

	// First generation of living cells


	gameField[1][1].current = true
	gameField[1][1].previous = true

	gameField[2][1].current = true
	gameField[2][1].previous = true

	gameField[3][0].current = true
	gameField[3][0].previous = true

	gameField[1][2].current = true
	gameField[1][2].previous = true

	gameField[0][2].current = true
	gameField[0][2].previous = true


	// Cycle for 12 generations
	for i:=0; i < 12; i++ {
		time.Sleep(1000 * time.Millisecond)
		//update status for the current generation
		for i:=0; i < len(gameField); i++ {
			for j:=0; j < len(gameField[i]); j++{
				gameField[i][j].previous = gameField[i][j].current
			}
		}

		for i:=0; i < len(gameField); i++ {
			for j:=0; j < len(gameField[i]); j++{
				var surroundingCells = [][]int{}
				var cords = []int{}
				//cell to the right
				if j+1 < len(gameField[i]) {
					cords = []int{i,j+1}
					surroundingCells = append(surroundingCells, cords)
				} else {
					cords = []int{i, 0}
					surroundingCells = append(surroundingCells, cords)
				}

				//cell to the left
				if j-1 >= 0 && j-1 < len(gameField) {
					cords = []int{i,j-1}
					surroundingCells = append(surroundingCells, cords)
				} else {
					cords = []int{i, len(gameField[i])-1}
					surroundingCells = append(surroundingCells, cords)
				}

				//cells to the bottom
				if i+1 < len(gameField) {
					//bottom cell
					cords = []int{i+1,j}
					surroundingCells = append(surroundingCells, cords)

					//bottom cell to the right
					if j+1 < len(gameField[i]) {
						cords = []int{i+1,j+1}
						surroundingCells = append(surroundingCells, cords)
					} else {
						cords = []int{i+1, 0}
						surroundingCells = append(surroundingCells, cords)
					}

					//bottom cell to the left
					if j-1 >= 0 && j-1 < len(gameField[i]){
						cords = []int{i+1,j-1}
						surroundingCells = append(surroundingCells, cords)
					} else {
						cords = []int{i+1, len(gameField[i])-1}
						surroundingCells = append(surroundingCells, cords)
					}
				} else {
					cords = []int{0, j}
					surroundingCells = append(surroundingCells, cords)
					//bottom cell to the right
					if j+1 < len(gameField[i]) {
						cords = []int{0,j+1}
						surroundingCells = append(surroundingCells, cords)
					} else {
						cords = []int{0, 0}
						surroundingCells = append(surroundingCells, cords)
					}

					//bottom cell to the left
					if j-1 >= 0 && j-1 < len(gameField[i]){
						cords = []int{0,j-1}
						surroundingCells = append(surroundingCells, cords)
					} else {
						cords = []int{0, len(gameField[i])-1}
						surroundingCells = append(surroundingCells, cords)
					}
				}

				//cells to the top
				if i-1 >= 0 && i-1 < len(gameField){
					//top cell
					cords = []int{i-1,j}
					surroundingCells = append(surroundingCells, cords)

					//top cell to the right
					if j+1 <= len(gameField[i])-1 {
						cords = []int{i-1,j+1}
						surroundingCells = append(surroundingCells, cords)
					} else {
						cords = []int{i-1, 0}
						surroundingCells = append(surroundingCells, cords)
					}

					//top cell to the left
					if j-1 >= 0 {
						cords = []int{i-1,j-1}
						surroundingCells = append(surroundingCells, cords)
					} else {
						cords = []int{i-1, len(gameField[i])-1}
						surroundingCells = append(surroundingCells, cords)
					}
				} else {
					cords = []int{len(gameField[i])-1, j}
					surroundingCells = append(surroundingCells, cords)

					//top cell to the right
					if j+1 <= len(gameField[i])-1 {
						cords = []int{len(gameField[i])-1,j+1}
						surroundingCells = append(surroundingCells, cords)
					} else {
						cords = []int{len(gameField[i])-1, 0}
						surroundingCells = append(surroundingCells, cords)
					}

					//top cell to the left
					if j-1 >= 0 {
						cords = []int{len(gameField[i])-1,j-1}
						surroundingCells = append(surroundingCells, cords)
					} else {
						cords = []int{len(gameField[i])-1, len(gameField[i])-1}
						surroundingCells = append(surroundingCells, cords)
					}
				}

				var livingCells = 0
				for k := 0; k < len(surroundingCells); k++ {
					if gameField[surroundingCells[k][0]][surroundingCells[k][1]].previous == true {
						livingCells++
					}
				}

				if gameField[i][j].previous == true {
					if livingCells == 2 || livingCells == 3 {
						gameField[i][j].current = true
					} else {
						gameField[i][j].current = false
					}

				} else {
					if livingCells == 3 {
						gameField[i][j].current = true
					}
				}

				var renderedLife = ""
				for i := 0; i < len(gameField); i++ {
					for j:=0; j < len(gameField[i]); j++{
						if gameField[i][j].current == true {
							renderedLife += "\u25A0 "
						} else {
							renderedLife += "\u25A1 "
						}
					}
					renderedLife += "\n"
				}
				par3 := termui.NewPar(renderedLife)
				par3.Width = 12
				par3.Height = 8
				termui.Render(par3)
			}
		}
	}
	<-termui.EventCh()
}