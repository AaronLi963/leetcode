package main

func compress(chars []byte) int {
	l := len(chars)
	if l <= 1 {
		return l
	}
	pos := 1
	count := 1
	for i := 1; i < l; i++ {
		if chars[i] == chars[i-1] {
			count += 1
		} else {
			if count != 1 {
				b := IntToBytes(count)
				for j := range b {
					chars[pos] = b[j]
					pos++
				}
				count = 1
			}
			chars[pos] = chars[i]
			pos++

		}
		if i == (l - 1) {
			if count != 1 {
				b := IntToBytes(count)
				for j := range b {
					chars[pos] = b[j]
					pos++
				}
			}
		}
	}
	return pos
}

func IntToBytes(i int) []byte {
	b := []byte{}
	for i > 0 {
		mod := i % 10
		b = append([]byte{byte(mod) + byte('0')}, b...)
		i /= 10
	}
	return b
}
