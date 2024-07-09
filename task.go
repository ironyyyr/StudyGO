package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const lookFor = ".txt"

func main() {
	makeOut(uniq(readInp()))
}

// Чтение, вввод
func readInp() []string {
	var (
		counterTxt int
		lines      []string
	)

	//Подсчет количества файлов input/output
	checkInputOutput := os.Args[1:]
	for _, elem := range checkInputOutput {
		if strings.HasSuffix(elem, lookFor) {
			counterTxt++
		}
	}

	//Чтение

	//Если файла на вход нет
	if counterTxt == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		//Если только файл на вход
	} else if counterTxt == 1 {
		file, err := os.Open(os.Args[len(os.Args)-1])
		if err != nil {
			fmt.Print("Ошибка чтения файла")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		//Если файл и на вход, и на выход
	} else if counterTxt == 2 {
		file, err := os.Open(os.Args[len(os.Args)-2])
		if err != nil {
			fmt.Print("Ошибка чтения файла")
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
	}
	return lines
}

// Вывод или запись в файл
func makeOut(exitLines []string) {
	var counterTxt int
	//Подсчет количества файлов input/output
	checkInputOutput := os.Args[1:]
	for _, elem := range checkInputOutput {
		if strings.HasSuffix(elem, lookFor) {
			counterTxt++
		}
	}

	//Проверка output - файла на существование
	if counterTxt == 2 {
		_, err := os.Stat(os.Args[len(os.Args)-1])
		if err != nil {
			if os.IsNotExist(err) {
			}
		} else {
			os.Remove(os.Args[len(os.Args)-1])
		}
	}

	//Если файлы и на вход, и на выход
	if counterTxt == 2 {
		outFile, err := os.Create(os.Args[len(os.Args)-1])
		if err != nil {
			fmt.Println("Ошибка создания файла")
			os.Exit(1)
		}
		defer outFile.Close()
		for _, elem := range exitLines {
			fmt.Fprintf(outFile, elem+"\n")
		}
		//Если нет команды записывать результат в файл
	} else if counterTxt < 2 {
		for _, elem := range exitLines {
			fmt.Println(elem)
		}
	}
}

// Форматирование данных
func uniq(lines []string) []string {
	var (
		numFields, numChars, counter int
		editLines                    []string
		exitLines                    []string
		mElem                        string
		jElem                        string
	)

	useC := flag.Bool("c", false, "Подсчитать количество поторениий строки во входных данных")
	useD := flag.Bool("d", false, "Вывести только те строки, которые повторились во входных данных")
	useU := flag.Bool("u", false, "Вывести только те строки, которые не повторились во входных данных")
	useI := flag.Bool("i", false, "Не учитывать регистр букв")
	flag.IntVar(&numFields, "f", 0, "Количество неучитываемых первых слов в строке")
	flag.IntVar(&numChars, "s", 0, "Количество неучитываемых первых символов в строке")
	flag.Parse()

	if *useI {
		for i := 0; i <= len(lines)-2; i++ {
			counter = 1
			if len(strings.Join(strings.Split(lines[i], " ")[numFields:], " ")) > numChars {
				mElem = strings.ToLower(strings.Join(strings.Split(lines[i], " ")[numFields:], " ")[numChars:])
			} else {
				mElem = lines[i]
			}
			for j := i + 1; j <= len(lines)-1; j++ {
				if len(strings.Join(strings.Split(lines[j], " ")[numFields:], " ")) > numChars {
					jElem = strings.ToLower(strings.Join(strings.Split(lines[j], " ")[numFields:], " ")[numChars:])
				} else {
					jElem = lines[j]
				}
				if mElem == jElem {
					counter++
				} else {
					editLines = append(editLines, strconv.Itoa(counter)+" "+lines[i])
					i = j
					counter = 1
					break
				}
			}
			editLines = append(editLines, strconv.Itoa(counter)+" "+lines[i])
		}
	} else {
		for i := 0; i <= len(lines)-1; i++ {
			counter = 1
			if len(strings.Join(strings.Split(lines[i], " ")[numFields:], " ")) > numChars {
				mElem = strings.Join(strings.Split(lines[i], " ")[numFields:], " ")[numChars:]
			} else {
				mElem = lines[i]
			}
			for j := i + 1; j <= len(lines)-1; j++ {
				if len(strings.Join(strings.Split(lines[j], " ")[numFields:], " ")) > numChars {
					jElem = strings.Join(strings.Split(lines[j], " ")[numFields:], " ")[numChars:]
				} else {
					jElem = lines[j]
				}
				if mElem == jElem {
					counter++
				} else {
					editLines = append(editLines, strconv.Itoa(counter)+" "+lines[i])
					i = j
					counter = 1
					break
				}
			}
			editLines = append(editLines, strconv.Itoa(counter)+" "+lines[i])
		}
	}

	if *useC && *useD && *useU || *useU && *useD || *useD && *useC || *useC && *useU {
		fmt.Print("Некорректное чтнение условия исполнения программы")
	} else {
		switch {
		case *useC:
			for _, elem := range editLines {
				exitLines = append(exitLines, elem)
			}
		case *useD:
			for _, elem := range editLines {
				num, err := strconv.Atoi(strings.Split(elem, " ")[0])
				if err != nil {
					fmt.Println("Ошибка перевода строки")
				}
				if num >= 2 {
					exitLines = append(exitLines, strings.Join(strings.Split(elem, " ")[1:], " "))
				}
			}
		case *useU:
			for _, elem := range editLines {
				num, err := strconv.Atoi(strings.Split(elem, " ")[0])
				if err != nil {
					fmt.Println("Ошибка перевода строки")
				}
				if num < 2 {
					exitLines = append(exitLines, strings.Join(strings.Split(elem, " ")[1:], " "))
				}
			}
		default:
			for _, elem := range editLines {
				exitLines = append(exitLines, strings.Join(strings.Split(elem, " ")[1:], " "))
			}
		}
	}
	return exitLines
}
