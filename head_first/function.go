package main

import "fmt"

// Объявление функции «sayHi».
func sayHi() {
	fmt.Println("Hi!")
}

// Вычислить расход краски для стены.
// Объявляем, что paintNeeded возвращает число с плавающей точкой.
func paintNeeded(width float64, height float64) (float64, error) {
	// Если ширина имеет недопустимое значение, вернуть 0 и признак ошибки.
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	// Если высота имеет недопустимое значение, вернуть 0 и признак ошибки.
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	// Функция возвращает расход краски вместо того, чтобы выводить его.
	return area / 10, nil
}

func main() {
	// Вызываем «sayHi».
	sayHi()

	// Вычислить расход краски для первой стены.
	//paintNeeded(4.2, 3.0)
	// Вычислить расход краски для второй стены.
	//paintNeeded(5.2, 3.5)

	// Вычислить расход краски для стены.
	// Объявляем переменные для хранения расхода краски для текущей стены, а также для общего расхода по всем стенам.
	var amount, total float64
	// Вызываем paintNeeded и сохраняем возвращаемое значение
	amount, err := paintNeeded(4.2, 3.0) // err - Добавляем вторую переменную для второго возвращаемого значения
	// Выводим значение ошибки (или «nil», если ошибки не было).
	fmt.Println(err)

	// Выводим расход для первой стены.
	fmt.Printf("%0.2f liters needed\n", amount)
	// Прибавляем расход для текущий стены к total.
	total += amount

	// Выводим общий расход по всем стенам.
	fmt.Printf("Total: %0.2f liters\n", total)

	//fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	//fmt.Println("-----------------------------")
	//fmt.Printf("%12s | %2d\n", "Stamps", 50)
	//fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	//fmt.Printf("%12s | %2d\n", "Tape", 99)

}
