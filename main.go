package main

import (
	"fmt"
)

var star = "*"

func getMasks(text, separator string) (new string) { // можно было вторым возвращаемым значением задать error, но не придумал что может пойти не так
	someFlag := false
	buffer := make([][]string, 0)
	word := make([]string, 0)
	bytes := make([]byte, 0)
	httpURL := make([]string, 0)

	for _, item := range separator {
		httpURL = append(httpURL, string(item))
	}
	// fmt.Println(httpURL, len(httpURL))
	for index, item := range text {
		if string(item) != " " {
			word = append(word, (string(text[index])))

		} else {
			buffer = append(buffer, word)
			word = []string{}
		}
	}
	buffer = append(buffer, word)

	for ind, w := range buffer {
		if len(w) <= len(httpURL)+4 { // да, число 4 магическое
			for _, oneWord := range w {
				word := []byte(oneWord)
				bytes = append(bytes, word...)
			}

		} else if (w[0] == httpURL[0]) && (w[1] == httpURL[1]) && (w[2] == httpURL[2]) && (w[3] == httpURL[3]) && (w[4] == httpURL[4]) { // здесь сделать нормально
		cycle:
			for indexLetter := range w {

				for j := 0; j < len(httpURL); j++ {
					if (w[j]) == httpURL[j] {
						someFlag = true
					} else {
						someFlag = false
						if !(someFlag) {
							break cycle
						}
					}
				}
				if someFlag {
					if indexLetter < len(httpURL) {
						bytes = append(bytes, []byte(httpURL[indexLetter])...)
					} else {
						bytes = append(bytes, []byte(star)...)
					}
				}
			}
		} else {
			for _, oneWord := range w {
				word := []byte(oneWord)
				bytes = append(bytes, word...)
			}
		}
		if ind < (len(buffer) - 1) {
			space := []byte(" ")
			bytes = append(bytes, space...)
		}
	}
	new = string(bytes)
	fmt.Println(new)
	return string(new)
}

func main() {
	getMasks("http://hehee see you", "http://")
}
