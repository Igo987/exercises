package masker

const star = "*"
const URL = "http://"

func GetMasks(link <-chan string, url string) chan string {
	newLink := make(chan string)
	text := <-link
	anyFlag := make(chan bool)
	someFlag := false
	buffer := make([][]string, 0, 100)
	word := make([]string, 0, 100)
	bytes := make([]byte, 0, 100)
	httpURL := make([]string, 0, 100)

	for _, item := range url {
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
		if len(w) <= len(httpURL)+4 {
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
	stringRes := string(bytes)

	go func() {
		newLink <- stringRes
		anyFlag <- true

	}()
	// go func() {
	// 	for {
	// 		select {
	// 		case <-anyFlag:
	// 			close(newLink)
	// 		default:
	// 			log.Println("done")
	// 		}
	// 	}
	// }()
	go func() {
		<-anyFlag
		close(newLink)
	}()

	return newLink
}
