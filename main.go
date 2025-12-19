package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Player struct {
	Name string
	Score int
	Word string
	Guesser bool
}

func clearTerminal() {
	var cmd *exec.Cmd

	if os.Getenv("OS") == "Windows_NT" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}


func printCentered(text string) {
	width := 80
	padding := (width - len(text)) / 2
	fmt.Println(strings.Repeat(" ", padding) + text)
}

func showRules() {
	fmt.Println("===================================")
	fmt.Println("🎮 HANGMAN - TWO PLAYER GAME RULES")
	fmt.Println("===================================")

	fmt.Println("\n👥 Players:")
	fmt.Println("- The game is played by two players: P1 and P2.")

	fmt.Println("\n🔁 Rounds:")
	fmt.Println("- The number of rounds is selected before the game starts.")
	fmt.Println("- Each round has two turns:")
	fmt.Println("  1. P1 chooses a word, P2 guesses.")
	fmt.Println("  2. P2 chooses a word, P1 guesses.")

	fmt.Println("\n🔤 Guessing Rules:")
	fmt.Println("- The guessing player guesses one letter at a time.")
	fmt.Println("- If the letter exists in the word, it is revealed in all positions.")
	fmt.Println("- If the letter does not exist, the kill count increases.")

	fmt.Println("\n☠️ Kill Count (HANGMAN):")
	fmt.Println("- A maximum of 7 wrong guesses are allowed.")
	fmt.Println("- Each wrong guess completes one letter of 'HANGMAN'.")
	fmt.Println("- If 'HANGMAN' is fully completed, the guessing player loses the turn.")

	fmt.Println("\n🏆 Scoring:")
	fmt.Println("- If the guessing player completes the word, they get 1 point.")
	fmt.Println("- If 'HANGMAN' is completed first, the word chooser gets 1 point.")

	fmt.Println("\n📊 End of Game:")
	fmt.Println("- After all rounds, the player with the highest score wins.")
	fmt.Println("- If scores are equal, the game ends in a draw.")

	fmt.Println("\n⚙️ Additional Rules:")
	fmt.Println("- Only alphabetic letters are allowed.")
	fmt.Println("- Repeated guesses do not increase the kill count.")
	fmt.Println("- Letter case is ignored (A = a).")

	fmt.Println("===================================")
}

func getPlayerNames(p1 *Player, p2 *Player) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\nPlayers! Please enter your names")

	fmt.Print("Player 1: ")
	p1.Name, _ = reader.ReadString('\n')
	p1.Name = strings.TrimSpace(p1.Name)
	p1.Score = 0
	p1.Guesser = false

	fmt.Print("Player 2: ")
	p2.Name, _ = reader.ReadString('\n')
	p2.Name = strings.TrimSpace(p2.Name)
	p2.Score = 0  
	p2.Guesser = true           
}



func main() {
	var p1 Player
	var p2 Player

	clearTerminal()
	printCentered("💀 WELCOME TO HANGMAN 💀")
	showRules()

	getPlayerNames(&p1, &p2)

	fmt.Println("\nWelcome " + p1.Name + " and " + p2.Name)

	fmt.Println("\nPress ENTER to continue...")
	fmt.Scanln()
}	