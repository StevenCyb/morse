package translator

import "strings"

var textToMorseMap = map[rune]string{
	// Letters
	'A': ".-", 'B': "-...", 'C': "-.-.", 'D': "-..",
	'E': ".", 'F': "..-.", 'G': "--.", 'H': "....",
	'I': "..", 'J': ".---", 'K': "-.-", 'L': ".-..",
	'M': "--", 'N': "-.", 'O': "---", 'P': ".--.",
	'Q': "--.-", 'R': ".-.", 'S': "...", 'T': "-",
	'U': "..-", 'V': "...-", 'W': ".--", 'X': "-..-",
	'Y': "-.--", 'Z': "--..",

	// Digits
	'0': "-----", '1': ".----", '2': "..---", '3': "...--",
	'4': "....-", '5': ".....", '6': "-....", '7': "--...",
	'8': "---..", '9': "----.",

	// Punctuation
	'.': ".-.-.-", ',': "--..--", '?': "..--..", '\'': ".----.",
	'!': "-.-.--", '/': "-..-.", '(': "-.--.", ')': "-.--.-",
	'&': ".-...", ':': "---...", ';': "-.-.-.", '=': "-...-",
	'+': ".-.-.", '-': "-....-", '_': "..--.-", '"': ".-..-.",
	'$': "...-..-", '@': ".--.-.",
	// Undefined (will be reported as “#” or error)
}

var mapToText = map[string]rune{}

func init() {
	for k, v := range textToMorseMap {
		mapToText[v] = k
	}
}

// TextToMorse converts a text string to its Morse code representation.
func TextToMorse(text string, separator string) string {
	text = strings.ToUpper(text)
	var morse string
	for _, textChar := range text {
		if textChar == ' ' {
			morse += separator
		} else if morseCode, exists := textToMorseMap[textChar]; exists {
			morse += morseCode + separator
		} else {
			morse += "# "
		}
	}

	if len(morse) == 0 {
		return ""
	}

	return morse[:len(morse)-1]
}

// MorseToText converts a Morse code string to its text representation.
func MorseToText(morse string, separator string) string {
	var text string
	if morse == "" {
		return ""
	}

	morseChars := strings.Split(morse, separator)
	for _, morseChar := range morseChars {
		if morseChar == "" {
			text += separator
		} else if textChar, exists := mapToText[morseChar]; exists {
			text += string(textChar)
		} else {
			text += "#"
		}
	}

	return text
}
