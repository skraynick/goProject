package utils

func getShift(ciphertext string) int {
	// Count the frequency of each letter in the input string
	freq := make(map[rune]int)
	for _, c := range ciphertext {
		if c >= 'A' && c <= 'Z' {
			freq[c]++
		}
	}

	// Compute the expected frequency of letters in English text
	// englishFreq := map[rune]float64{
	//     'E': 12.702, 'T': 9.056, 'A': 8.167, 'O': 7.507, 'I': 6.966,
	//     'N': 6.749, 'S': 6.327, 'H': 6.094, 'R': 5.987, 'D': 4.253,
	//     'L': 4.025, 'C': 2.782, 'U': 2.758, 'M': 2.406, 'W': 2.360,
	//     'F': 2.228, 'G': 2.015, 'Y': 1.974, 'P': 1.929, 'B': 1.492,
	//     'V': 0.978, 'K': 0.772, 'J': 0.153, 'X': 0.150, 'Q': 0.095,
	//     'Z': 0.074,
	// }

	// Find the letter with the highest frequency in the input string
	maxFreq := 0
	var maxLetter rune
	for letter, count := range freq {
		if count > maxFreq {
			maxFreq = count
			maxLetter = letter
		}
	}

	// Calculate the shift necessary to align the two frequency distributions
	shift := int(maxLetter - 'E')
	if shift < 0 {
		shift += 26
	}
	return shift
}
