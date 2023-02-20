package cmd

type CaesarCipher struct {
	shift int
}

func (c *CaesarCipher) Encrypt(plainText string) string {
	var cipherText string
	for _, ch := range plainText {
		if ch >= 'A' && ch <= 'Z' {
			cipherText += string((int(ch-'A'+rune(c.shift)) % 26) + 'A')
		} else if ch >= 'a' && ch <= 'z' {
			cipherText += string((int(ch-'a'+rune(c.shift)) % 26) + 'a')
		} else {
			cipherText += string(ch)
		}
	}
	return cipherText
}

func (c *CaesarCipher) Decrypt(cipherText string) string {
	var plainText string
	for _, ch := range cipherText {
		if ch >= 'A' && ch <= 'Z' {
			plainText += string((int(ch-'A'-rune(c.shift)+26) % 26) + 'A')
		} else if ch >= 'a' && ch <= 'z' {
			plainText += string((int(ch-'a'-rune(c.shift)+26) % 26) + 'a')
		} else {
			plainText += string(ch)
		}
	}
	return plainText
}
