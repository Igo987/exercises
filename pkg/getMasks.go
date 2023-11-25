package masker

const star = "*"
const URL = "http://"

func GetMasks(text, separator string) (new string) { // можно было вторым возвращаемым значением задать error, но не придумал что может пойти не так
	someFlag := false
	buffer := make([][]string, 0)
	word := make([]string, 0)
	bytes := make([]byte, 0)
	httpURL := make([]string, 0)

	for _, item := range separator {
		httpURL = append(httpURL, string(item))
	}
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
		for j := 0; j < len(httpURL); j++ {
			if (w[j]) == httpURL[j] {
				someFlag = true
			} else {
				someFlag = false
				if !(someFlag) {
					break
				}
			}
		}
		if len(w) <= len(httpURL)+4 { // да, число 4 магическое: (если после httpURL(http://) меньше 5 символов,то, ну наверное, это не URL адрес и можно это слово не проверять)
			for _, oneWord := range w {
				word := []byte(oneWord)
				bytes = append(bytes, word...)
			}

		} else if someFlag {

			for indexLetter := range w {

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
	return new
}
