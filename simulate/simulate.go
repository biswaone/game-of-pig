package simulate

import (
	"fmt"
	"game-of-pig/strategyparse"
	"math/rand"
)

const numberOfGames int = 10

func rollDice() int {
	return rand.Intn(6) + 1
}

func playGame(p int) int {
	score := 0

	for {
		die := rollDice()

		if die == 1 {
			score = 0
			break
		}
		if score >= p {
			return score
		}
		score += die

	}
	return score
}

func simulateBothUsersFixedStrategy(p1, p2 int) {

	player1Wins := 0
	player2Wins := 0

	for i := 0; i < numberOfGames; i++ {
		//set scores to zero after every game
		player1Score := 0
		player2Score := 0
		for {
			if player1Score >= 100 {
				player1Wins += 1
				break
			}
			if player2Score >= 100 {
				player2Wins += 1
				break
			}
			player1Score += playGame(p1)
			player2Score += playGame(p2)
		}
	}
	//TODO: assert that player1Wins + player2Wins == 10
	fmt.Printf("Holding at %v vs Holding at %v: wins: %d/10 (%0.1f%%), losses: %d/10 (%0.1f%%)\n",
		p1, p2, player1Wins, float64(player1Wins)/10*100, player2Wins, float64(player2Wins)/10*100)
}

func simulateFixAndVariableStrategy(p int) {

	for p2 := 1; p2 <= 100; p2++ {
		if p2 != p {
			player1Wins := 0
			player2Wins := 0

			for i := 0; i < numberOfGames; i++ {
				//set scores to zero after every game
				player1Score := 0
				player2Score := 0
				for {

					if player1Score >= 100 {
						player1Wins += 1
						break
					}
					if player2Score >= 100 {
						player2Wins += 1
						break
					}
					player1Score += playGame(p)
					player2Score += playGame(p2)

				}

			}
			fmt.Printf("Holding at %v vs Holding at %v: wins: %d/10 (%0.1f%%), losses: %d/10 (%0.1f%%)\n",
				p, p2, player1Wins, float64(player1Wins)/10*100, player2Wins, float64(player2Wins)/10*100)

		}

	}

}
func simulateBothUsersVariableStrategy() {
	for p1 := 1; p1 <= 100; p1++ { //player 1 holds
		player1Wins := 0 // set the number of wins to 0 after a strategy simulation
		player2Wins := 0
		for p2 := 1; p2 <= 100; p2++ { // for every player 1 holds player 2 plays 99 strategy
			if p1 != p2 { // 10 games are played for every strategy
				for i := 0; i < 10; i++ { //play 10 games for every strategy
					player1Score := 0
					player2Score := 0
					for { // start game
						if player1Score >= 100 {
							player1Wins += 1
							break
						}
						if player2Score >= 100 {
							player2Wins += 1
							break
						}
						player1Score += playGame(p1)
						player2Score += playGame(p2)
					}

				}
			}
		}

		total := player1Wins + player2Wins
		winPercentage := float64(player1Wins) / float64(total)
		lossPercentage := float64(player2Wins) / float64(total)
		fmt.Printf("Result: Wins, losses staying at k = %d: %d/990 (%.2f%%), %d/990 (%.2f%%)\n",
			p1, player1Wins, winPercentage*100, player2Wins, lossPercentage*100)

	}

}
func Simulate(s strategyparse.GameStrategy) {
	//simulate story 1
	if s.Story == 1 {
		simulateBothUsersFixedStrategy(s.Player1, s.Player2)
	}
	//simulate story 2
	if s.Story == 2 {
		simulateFixAndVariableStrategy(s.Player1)

	} //simulate story 3
	if s.Story == 3 {
		simulateBothUsersVariableStrategy()

	}

}
