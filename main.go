//Hangman Go by Patrick Chavez

package main

import ( 
	"fmt"
	"os"
	"os/exec"
)

//Print Screen
func printScreen() {
	clear()
	fmt.Println("Hangman Go! version 1.00")
}

//Clear screen
func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	cmd = exec.Command("cmd", "/c", "clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	printScreen()
}