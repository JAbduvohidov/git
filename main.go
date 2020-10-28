package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

var lettersMap = map[rune]string{
	'ф':  "a",
	'и':  "b",
	'с':  "c",
	'в':  "d",
	'у':  "e",
	'а':  "f",
	'п':  "g",
	'р':  "h",
	'ш':  "i",
	'о':  "j",
	'л':  "k",
	'д':  "l",
	'ь':  "m",
	'т':  "n",
	'щ':  "o",
	'з':  "p",
	'й':  "q",
	'к':  "r",
	'ы':  "s",
	'е':  "t",
	'г':  "u",
	'м':  "v",
	'ц':  "w",
	'ч':  "x",
	'н':  "y",
	'я':  "z",
	'Э':  "\"",
	'\'': "'",
	'б':  ",",
	'ю':  ".",
}

func main() {
	args := os.Args

	var arguments []string
	for _, arg := range args[1:] {
		arguments = append(arguments, transliterate([]rune(arg)))
	}

	if arguments == nil {
		arguments = append(arguments, "help")
	}

	cmd := exec.Command("git", arguments...)
	cmd.Dir = "."
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("%s", out)
	}

}

func transliterate(source []rune) (result string) {
	for _, r := range source {
		if !contains(lettersMap, r) {
			result += string(r)
		}
		result += lettersMap[r]
	}
	return
}

func contains(m map[rune]string, r rune) bool {
	for i, _ := range m {
		if i == r {
			return true
		}
		continue
	}
	return false
}
