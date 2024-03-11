// pass_fail - сообщает, сдал ли пользователь экзамен.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter a grade: ")        // Запрашиваем у поль-зователя значение.
	reader := bufio.NewReader(os.Stdin)   // Создать «буферизованное средство чтения» для получе-ния текста с клавиатуры.
	input, err := reader.ReadString('\n') // Возвращает текст, введенный пользователем до нажатия клавиши Enter.
	// Сообщить о любых ошибках при преобразовании.
	if err != nil {
		log.Fatal(err)
	}

	input = strings.TrimSpace(input)            // Удаляем символ новой строки из входной строки.
	grade, err := strconv.ParseFloat(input, 64) // Преобразовать строку в значение float64.
	// Как и в случае с ReadString, сообщить о любых ошибках при преобразовании.
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failng"
	}

	fmt.Println("A grade of ", grade, " is ", status)

	// Задача с "холодильника"
	//fileInfo, err := os.Stat("my.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(fileInfo.Size())
}
