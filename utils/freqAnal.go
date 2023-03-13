package utils

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
)

type FrequencyAnalysis struct {
	cypherText string
}

func analyze(cipherText string) map[rune]int {
	frequencies := make(map[rune]int)
	for _, char := range cipherText {
		if unicode.IsLetter(char) {
			frequencies[unicode.ToLower(char)]++
		}
	}
	return frequencies
}

func getFrequencies(cipherText string) map[rune]int {
	freqs := make(map[rune]int)
	for _, c := range cipherText {
		if c >= 'A' && c <= 'Z' {
			freqs[c]++
		}

	}
	return freqs
}

func frequencyMap(cipherString string) map[rune]int {
	freqMap := make(map[rune]int)

	for _, char := range cipherString {
		freqMap[char]++
	}
	return freqMap
}

func frequencyAnalysis(cipherText string) map[rune]int {
	freq := make(map[rune]int)
	for _, ch := range cipherText {
		if ch >= 'A' && ch <= 'Z' {
			freq[ch]++
		} else if ch >= 'a' && ch <= 'z' {
			freq[ch-32]++
		}
	}
	return freq
}

// func thing() string {
// 	re := regexp.MustCompile(`[a-z]+`)
// 	return e.FindAllString(text, -1)
// }

func returnMostFrequentLetter(cipherText string) rune {
	freq := frequencyAnalysis(cipherText)
	var mostFrequent rune
	maxFreq := 0
	for ch, count := range freq {
		if count > maxFreq {
			mostFrequent = ch
			maxFreq = count
		}

	}
	return mostFrequent
}

func returnSkipValue(cipherText string) rune {
	var skipValue rune
	var mostFrequent rune = returnMostFrequentLetter(cipherText)
	mostFrequentList := [13]rune{'E', 'A', 'R', 'I', 'O', 'T', 'N', 'S', 'L', 'C', 'U', 'D', 'P'}

	for k, v := range mostFrequentList {
		println(mostFrequent+1, v, k)

		if mostFrequent+1-v == 3 {
			skipValue = 3
			break
		} else {
			skipValue = mostFrequent + 1 - v

		}
	}

	println(skipValue, 'E', mostFrequent)

	return skipValue
}

func DecryptWithFrequencyAnalysis(cipherText string) string {

	var shift int

	shift = int(returnSkipValue(cipherText))
	//We can't assume we know the shift.
	//use the most frequent.
	//if not e as in 3... then decode using next letter
	//look at output if not close run next shift.
	//run against dictionary of maybe a thousand words
	//If a complete word is found, stop...
	//match encrypted text to the plain text so far and start
	//matching known letters to calculate shift.
	//solve rest of the text.\
	//https://www3.nd.edu/~busiforc/handouts/cryptography/letterfrequencies.html

	//https://ankurraina.medium.com/reading-a-simple-csv-in-go-36d7a269cecd

	//get mapping.. need to sort by highest. after determine if
	//highest is actually e.. math. Difference should be 3, if not it is
	//e. Round two... Use the nest highest value to attemp to decrypt.
	//	freqs := frequencyMap(cipherText)
	// for k, v := range freqs {
	// 	println(string(k), v, k)
	// 		//find shift for none 3 shift
	// 		//Then most frequent is not "e"
	// 		//need to find out what letter it is.

	// 	}
	// }

	return Decrypt(cipherText, int(shift))

}

/*
Not used at this point.
*/
func getSortedFrequencies(cipherText string) []string {
	freqs := analyze(cipherText)
	sorted := make([]string, 0, len(freqs))
	for k, v := range freqs {
		sorted = append(sorted, fmt.Sprintf("%c: %d", k, v))
	}
	sort.Slice(sorted, func(i, j int) bool {
		return strings.Compare(sorted[i], sorted[j]) > 0
	})

	return sorted
}

func getMapping(cipherText string) map[rune]rune {
	freqs := getSortedFrequencies(cipherText)

	mostFrequent := "ETAOINSHR"
	mapping := make(map[rune]rune)
	for i, c := range mostFrequent {
		//need List with sorted by most frequent.. map
		//to the most frequent. Assumptions they should match up.
		mapping[rune(freqs[i][0])] = c
	}
	return mapping
}

func FrequencyAnalysisDecrypt(cipherText string) string {
	mapping := getMapping(cipherText)
	plaintext := make([]rune, len(cipherText))
	for i, c := range cipherText {
		if c >= 'A' && c <= 'Z' {
			plaintext[i] = mapping[c]
		} else {
			plaintext[i] = c
		}
		//println(string(plaintext[i]))

	}
	return string(plaintext)
}
