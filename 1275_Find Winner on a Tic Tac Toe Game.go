/*
Tic-tac-toe is played by two players A and B on a 3 x 3 grid. The rules of Tic-Tac-Toe are:

Players take turns placing characters into empty squares ' '.
The first player A always places 'X' characters, while the second player B always places 'O' characters.
'X' and 'O' characters are always placed into empty squares, never on filled ones.
The game ends when there are three of the same (non-empty) character filling any row, column, or diagonal.
The game also ends if all squares are non-empty.
No more moves can be played if the game is over.
Given a 2D integer array moves where moves[i] = [rowi, coli] indicates that the ith move will be played on grid[rowi][coli]. return the winner of the game if it exists (A or B). In case the game ends in a draw return "Draw". If there are still movements to play return "Pending".

You can assume that moves is valid (i.e., it follows the rules of Tic-Tac-Toe), the grid is initially empty, and A will play first.
*/

package main

type game struct {
	grid        [4][4]int
	emptyFields int
}

func NewGame() *game {
	return &game{
		emptyFields: 9,
	}
}

func (g *game) move(line int, colomn int, player int) {
	g.grid[line][colomn] = player
	g.emptyFields--
	g.grid[line][3]--
	g.grid[3][colomn]--
}

func (g *game) checkLine(line int) int {
	if g.grid[line][3] != -3 {
		return 0
	}
	if g.grid[line][0] == g.grid[line][1] && g.grid[line][1] == g.grid[line][2] {
		return g.grid[line][0]
	}
	return 0
}

func (g *game) checkColomn(colomn int) int {
	if g.grid[3][colomn] != -3 {
		return 0
	}
	if g.grid[0][colomn] == g.grid[1][colomn] && g.grid[1][colomn] == g.grid[2][colomn] {
		return g.grid[0][colomn]
	}
	return 0
}
func (g *game) checkDiagonal() int {
	if g.grid[0][0] == g.grid[1][1] && g.grid[1][1] == g.grid[2][2] {
		return g.grid[0][0]
	}
	if g.grid[0][2] == g.grid[1][1] && g.grid[1][1] == g.grid[2][0] {
		return g.grid[1][1]
	}
	return 0
}

func (g *game) checkWinner() int {
	for i := 0; i < 3; i++ {
		winner := g.checkLine(i)
		if winner != 0 {
			return winner
		}
	}
	for i := 0; i < 3; i++ {
		winner := g.checkColomn(i)
		if winner != 0 {
			return winner
		}
	}
	winner := g.checkDiagonal()
	return winner
}

func tictactoe(moves [][]int) string {
	currGame := NewGame()
	for i, v := range moves {
		if i%2 == 0 {
			currGame.move(v[0], v[1], 1)
		} else {
			currGame.move(v[0], v[1], 2)
		}
	}
	winner := currGame.checkWinner()
	if winner == 0 {
		if currGame.emptyFields != 0 {
			return "Pending"
		} else {
			return "Draw"
		}
	} else if winner == 1 {
		return "A"
	} else {
		return "B"
	}

}
