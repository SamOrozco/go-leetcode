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
			for _, char := range word  {
				if val, ok := morseCodeMap[char]; ok {
					bldr.WriteString(val)
				}
			}
			uniqueResponses[bldr.String()] = true
		}
		return len(uniqueResponses)
}
p a c k a g e   m a i n 
 
 
 
 i m p o r t   " s t r i n g s " 
 
 
 
 f u n c   m a i n ( )   { 
 
 	 i f   u n i q u e M o r s e R e p r e s e n t a t i o n s ( [ ] s t r i n g { " g i n " ,   " z e n " ,   " g i g " ,   " m s g " } )   ! =   2   { 
 
 	 	 p r i n t l n ( " f a i l e d " ) 
 
 	 } 
 
 } 
 
 
 
 v a r   m o r s e C o d e M a p   =   m a p [ r u n e ] s t r i n g { 
 
 	 ' a ' :   " . - " , 
 
 	 ' b ' :   " - . . . " , 
 
 	 ' c ' :   " - . - . " , 
 
 	 ' d ' :   " - . . " , 
 
 	 ' e ' :   " . " , 
 
 	 ' f ' :   " . . - . " , 
 
 	 ' g ' :   " - - . " , 
 
 	 ' h ' :   " . . . . " , 
 
 	 ' i ' :   " . . " , 
 
 	 ' j ' :   " . - - - " , 
 
 	 ' k ' :   " - . - " , 
 
 	 ' l ' :   " . - . . " , 
 
 	 ' m ' :   " - - " , 
 
 	 ' n ' :   " - . " , 
 
 	 ' o ' :   " - - - " , 
 
 	 ' p ' :   " . - - . " , 
 
 	 ' q ' :   " - - . - " , 
 
 	 ' r ' :   " . - . " , 
 
 	 ' s ' :   " . . . " , 
 
 	 ' t ' :   " - " , 
 
 	 ' u ' :   " . . - " , 
 
 	 ' v ' :   " . . . - " , 
 
 	 ' w ' :   " . - - " , 
 
 	 ' x ' :   " - . . - " , 
 
 	 ' y ' :   " - . - - " , 
 
 	 ' z ' :   " - - . . " , 
 
 } 
 
 
 
 f u n c   u n i q u e M o r s e R e p r e s e n t a t i o n s ( w o r d s   [ ] s t r i n g )   i n t   { 
 
 	 i f   l e n ( w o r d s )   = =   0   { 
 
 	 	 r e t u r n   0 
 
 	 } 
 
 
 
 	 u n i q u e R e s p o n s e s   : =   m a k e ( m a p [ s t r i n g ] b o o l ) 
 
 	 / /   i t e r a t e   e a c h   w o r d   i n   o u r   w o r d s   s l i c e 
 
 	 f o r   _ ,   w o r d   : =   r a n g e   w o r d s   { 
 
 	 	 / /   i t e r a t e   e a c h   c h a r   f o r   e a c h   w o r d   a n d   b u i l d   m o r s e   c o d e   s t r i n g 
 
 	 	 b l d r   : =   s t r i n g s . B u i l d e r { } 
 
 	 	 f o r   _ ,   c h a r   : =   r a n g e   w o r d   { 
 
 	 	 	 i f   v a l ,   o k   : =   m o r s e C o d e M a p [ c h a r ] ;   o k   { 
 
 	 	 	 	 b l d r . W r i t e S t r i n g ( v a l ) 
 
 	 	 	 } 
 
 	 	 } 
 
 	 	 u n i q u e R e s p o n s e s [ b l d r . S t r i n g ( ) ]   =   t r u e 
 
 	 } 
 
 	 r e t u r n   l e n ( u n i q u e R e s p o n s e s ) 
 
 } 
 
 