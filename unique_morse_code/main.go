package main

import "strings"

func main() {
        if uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}) != 2 {
                println("failed")
        }
}

var morseCodeMap = map[rune]string{
        'a': ".-",
        'b': "-...",
        'c': "-.-.",
        'd': "-..",
        'e': ".",
        'f': "..-.",
        'g': "--.",
        'h': "....",
        'i': "..",
        'j': ".---",
        'k': "-.-",
        'l': ".-..",
        'm': "--",
        'n': "-.",
        'o': "---",
        'p': ".--.",
        'q': "--.-",
        'r': ".-.",
        's': "...",
        't': "-",
        'u': "..-",
        'v': "...-",
        'w': ".--",
        'x': "-..-",
        'y': "-.--",
        'z': "--..",
}

func uniqueMorseRepresentations(words []string) int {
        if len(words) == 0 {
                return 0
        }

        uniqueResponses := make(map[string]bool)
        // iterate each word in our words slice
        for _, word := range words {
                // iterate each char for each word and build morse code string
                bldr := strings.Builder{}
                for _, char := range word {
                        if val, ok := morseCodeMap[char]; ok {
                                bldr.WriteString(val)
                        }
                }
                uniqueResponses[bldr.String()] = true
        }
        return len(uniqueResponses)
}