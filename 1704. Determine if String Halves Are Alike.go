package main

func halvesAreAlike(s string) bool {
	dic := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	offset := len(s) / 2
	counter := 0
	for i := 0; i < offset; i++ {
		if dic[s[i]] {
			counter++
		}
		if dic[s[i+offset]] {
			counter--
		}
	}
	return counter == 0
}
