package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	lowercaseLetters  = [26]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	uppercaseLetters  = [26]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	digits            = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	specialCharacters = [8]string{"!", "@", "#", "$", "%", "^", "&", "*"}
)

func addRandSymbolToBuilder(strBuilder *strings.Builder, characterArray []string, symbolCount int) {
	for count := 0; count < symbolCount; count++ {
		randomIndex := rand.Intn(len(characterArray))

		_, err := strBuilder.WriteString(characterArray[randomIndex])
		if err != nil {
			fmt.Println("Couldn't write character")
			os.Exit(1)
		}
	}
}

func generatePassword(minUpper int, minLower int, minSpecial int, minDigit int) string {
	var finalString strings.Builder

	addRandSymbolToBuilder(&finalString, uppercaseLetters[:], minUpper)
	addRandSymbolToBuilder(&finalString, lowercaseLetters[:], minLower)
	addRandSymbolToBuilder(&finalString, specialCharacters[:], minSpecial)
	addRandSymbolToBuilder(&finalString, digits[:], minDigit)

	var unshuffledStr string = finalString.String()
	runeStr := []rune(unshuffledStr)
	rand.Shuffle(len(runeStr), func(i, j int) {
		runeStr[i], runeStr[j] = runeStr[j], runeStr[i]
	})

	return string(runeStr)

}

func main() {

	minUppercase := flag.Int("u", 5, "Uppercase letters amount")
	minLowercase := flag.Int("l", 5, "Lowercase letters amount")
	minDigits := flag.Int("d", 3, "Digits amount")
	minSpecial := flag.Int("s", 1, "Special symbols amount")
	passwordCount := flag.Int("c", 1, "Password generation amount")

	flag.Parse()
	if *minUppercase < 0 || *minLowercase < 0 || *minDigits < 0 || *minSpecial < 0 || *passwordCount < 0 {
		fmt.Println("Invalid argument, value can't be less than 0")
		os.Exit(1)
	}

	var minLength = *minLowercase + *minUppercase + *minSpecial + *minDigits

	for *passwordCount != 0 {
		var password string = generatePassword(*minUppercase, *minLowercase, *minSpecial, *minDigits)

		if len(password) != minLength {
			fmt.Printf("Password didn't match the minLenght")
			os.Exit(1)
		}

		fmt.Println(password)
		*passwordCount--
	}
}
