package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// latinLettersTable - мапа, которая содержит в себе определение латинского алфавита
var latinLettersTable = map[string]string{
	"a": "vowel",
	"b": "consonant",
	"c": "consonant",
	"d": "consonant",
	"e": "vowel",
	"f": "consonant",
	"g": "consonant",
	"h": "consonant",
	"i": "vowel",
	"j": "consonant",
	"k": "consonant",
	"l": "consonant",
	"m": "consonant",
	"n": "consonant",
	"o": "vowel",
	"p": "consonant",
	"q": "consonant",
	"r": "consonant",
	"s": "consonant",
	"t": "consonant",
	"u": "vowel",
	"v": "consonant",
	"w": "consonant",
	"x": "consonant",
	"y": "vowel",
	"z": "consonant",
}

// exceptionsTable - мапа, которая содержит в себе исключения для существительных
var exceptionsTable = map[string]string{
	"ventris":   "согласный",
	"ureteris":  "согласный",
	"gastris":   "согласный",
	"matris":    "согласный",
	"tuberis":   "согласный",
	"cadaveris": "согласный",
	"puppis":    "гласная",
	"sitis":     "гласная",
	"fibris":    "гласная",
	"vis":       "гласная",
	"turris":    "гласная",
	"securis":   "гласная",
}

// роды существительного
const (
	manGenus    = "m"
	womanGenus  = "f"
	middleGenus = "n"
)

// CheckVowelType проверяет принадлежность существительного 3 склонения к гласному типу
func CheckVowelType(nomStr, genStr, genus string) bool {
	if (strings.HasSuffix(strings.ToLower(nomStr), "e") || strings.HasSuffix(strings.ToLower(nomStr), "al") || strings.HasSuffix(strings.ToLower(nomStr), "ar")) && genus == "n" {
		return true
	}

	return false
}

// CheckMixedAOrConsonantType проверяет принадлежность существительного 3 склонения к смешанному а, либо согласному типам
func CheckMixedAOrConsonantType(nomStr, genStr string, countNom, countGen int) {
	if strings.HasSuffix(strings.ToLower(genStr), "is") {
		base := genStr[:len(genStr)-len("is")]

		baseLower := strings.ToLower(base)

		consonantCount := 0

		for i := len(baseLower) - 1; i >= 0; i-- {
			letter := string(baseLower[i])
			if letterType, exists := latinLettersTable[letter]; exists {
				if letterType == "consonant" {
					consonantCount++
				} else {
					break
				}
			} else {
				continue
			}

			if consonantCount == 2 {
				break
			}
		}

		if consonantCount == 2 && countGen != countNom {
			fmt.Println("смешанный а")
		} else if consonantCount == 1 && countGen != countNom {
			fmt.Println("согласный")
		} else {
			fmt.Println("Ошибка! Проверьте корректность введённых данных. Возможно, вы не правильно определили род существительного изначально)")
		}
	}
}

// CheckMixedBType проверяет принадлежность существительного 3 склонения к смешанному б типу
func CheckMixedBType(nomStr, genStr, genus string, countNom, countGen int) bool {
	if strings.HasSuffix(strings.ToLower(nomStr), "is") || strings.HasSuffix(strings.ToLower(nomStr), "es") && (countNom == countGen) && (genus == "f" || genus == "m") {
		return true
	}

	return false
}

func main() {
	var nominativeWord, genitiveWord, genus string

	reader := bufio.NewReader(os.Stdin)

	var endLoop string

	for endLoop != "end" {
		fmt.Println("Введите существительное в именительном падеже:")
		nominativeWord, _ = reader.ReadString('\n')
		nominativeWord = strings.TrimSpace(nominativeWord)

		fmt.Println("Введите существительное в родительном падеже:")
		genitiveWord, _ = reader.ReadString('\n')
		genitiveWord = strings.TrimSpace(genitiveWord)

		fmt.Println("Введите род существительного (n/m/f):")
		genus, _ = reader.ReadString('\n')
		genus = strings.TrimSpace(genus)

		var countNominativeVowel int
		var countGenitiveVowel int

		for _, val := range nominativeWord {
			if latinLettersTable[string(val)] == "vowel" {
				countNominativeVowel++
			}
		}

		for _, val := range genitiveWord {
			if latinLettersTable[string(val)] == "vowel" {
				countGenitiveVowel++
			}
		}

		found := false
		for i, _ := range exceptionsTable {
			if genitiveWord == i {
				found = true
				break
			}
		}

		if found {
			fmt.Println(exceptionsTable[string(genitiveWord)])
		} else if CheckVowelType(nominativeWord, genitiveWord, genus) {
			fmt.Println("гласная")
		} else if countNominativeVowel == countGenitiveVowel && CheckMixedBType(nominativeWord, genitiveWord, genus, countNominativeVowel, countGenitiveVowel) {
			fmt.Println("смешанный б")
		} else {
			CheckMixedAOrConsonantType(nominativeWord, genitiveWord, countNominativeVowel, countGenitiveVowel)
		}

		fmt.Println("Если хотите остановиться введите: end, если желаете продолжить введите: go")
		endLoop, _ = reader.ReadString('\n')
		endLoop = strings.TrimSpace(endLoop)
	}

	fmt.Println("Завершение работы программы...")

}
