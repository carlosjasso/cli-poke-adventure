package prompt

import "fmt"

type foregroundColor string

const (
	green  foregroundColor = "\033[32m"
	yellow foregroundColor = "\033[33m"
	white  foregroundColor = "\033[0m"
)

func Display() {
	fmt.Printf("%sPokeAdventure %s> %s", green, yellow, white)
}
