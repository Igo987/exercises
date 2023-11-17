package pkg

import (
	"bufio"
	"log"
	"os"
)

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

type Service struct {
	prod produce
	pres presenty
}

func Run(path_read,path_write string) error { // импортируем
	svc := NewService()
	svc.prod.path = path_read
	svc.pres.path = path_write
	data, err := svc.prod.producer()
	// fmt.Println("data", data)
	if err != nil {
		return err
	}
	/* trunk-ignore(golangci-lint/errcheck) */
	svc.pres.presenter(data)
	return nil
}

func NewService() *Service { // конструктор
	svc := new(Service)
	return svc
}

// поставщик данных (чтение из файла)
type Produce interface {
	producer() (data []string, e error)
}

// обработчик результата (запись в файл)
type Present interface {
	presenter(s []string) (err error)
}

// читать данные из файла (поставщик данных в функцию getMasks)
type produce struct {
	Produce
	path string // откуда читать
}

// видимо, на выходе, после обработки данных функцией getMasks мы данные записываем в файл
type presenty struct {
	Present
	path string // куда записывать
}

func (p produce) producer() ([]string, error) {
	file, err := os.Open(p.path)
	if err != nil {
		log.Fatal(err)
	}
	line_list := []string{}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line_list = append(line_list, GetMasks(line,URL))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	// fmt.Println("producer line_list", line_list)
	return line_list, nil
}

func (p presenty) presenter(s []string) (error) {
  f, err := os.Create(p.path)
	if err != nil {
		return err
	}
	// Не забудьте закрыть файл
	defer f.Close()
	for _, line := range s {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
return nil	
}
