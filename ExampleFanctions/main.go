package main
import (
	"fmt"
	)

/*Анонимные функции - это функции, которым не назначено имя. Они отличаются от обычных функций также тем,
	что они могут определяться внутри других функций и также могут иметь доступ к контексту выполнения.
Преимуществом анонимных функций является то, что они имеют доступ к окружению, в котором они вызываются
*/

func main() {
//------------- Сомовызывающейся функция (Рекурсивная функция)
	fmt.Println("Рекурсивная функция = ", AnonimRecursionFunc(4))  

//------------- Вариативные функции
	s := []int{10, 20, 30}
	mult1 := GetMultiples(2, s...)
	mult2 := GetMultiples(3, 1, 2, 3, 4)
	fmt.Println("Вариативные функции = ", mult1, mult2)
	
//------------- Анонимная функция как аргумент функции
	fmt.Println("Анонимная функция как аргумент функции = ", AnonimFunc(10, 25, func(x int, y int) int { return x + y })) 

//------------- Анонимная функция как результат функции
	f := AnonimResultFunction(1) 
	fmt.Println("Анонимная функция как результат функции = ", f(2, 3))
	
//------------- Доступ к окружению
	l := AnonimFunction() 
	fmt.Println("Доступ к окружению = ", l())
	
//------------- Функция как результат другой функции
	g := Select(2) 
	fmt.Println("Функция как результат другой функции = ", g(3, 4))


//------------- Defer
	defer finish()
	fmt.Println("Program has been started")

//------------- Panic
	// fmt.Println(divide(15, 5))
	// fmt.Println(divide(4, 0))
	// fmt.Println("Program has been finished")

}

//------------- Анонимная функция -------------
	// f := func(x, y int) int {
	// 	return x + y
	// }
	// fmt.Println(f(3, 4))  //7


//------------- Сомовызывающейся функция (Рекурсивная функция)
func AnonimRecursionFunc(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * AnonimRecursionFunc(n-1)
}
	
//------------- Вариативные функции
// Обозначают и оператора упаковки, и оператора распаковки.  
// Оператор распаковки всегда стоит после имени переменной.
func GetMultiples(fuctor int, args ...int) []int{
	multiples := make([]int, len(args))
	for index, val := range args {
		multiples[index] = val * fuctor
	}
	return	multiples
}

//------------- Анонимная функция как аргумент функции
func AnonimFunc(x, y int, operation func(int, int) int) int {
	result := operation(x, y)
	return result
}

//------------- Анонимная функция как результат функции
func AnonimResultFunction(n int) func(int, int) int {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}

//------------- Доступ к окружению
func AnonimFunction() func() int {
	var x int = 2
	return func() int {
		x++
		return x * x
	}
}

//------------- Функция как результат другой функции
func Select(n int) func(int, int) int {
	if n == 1 {
		return add
	} else {
		return multiply
	}
}

func add(x, y int) int {
		return x + y
	}
	func multiply(x, y int) int {
		return x * y
	}


//------------- Defer
/*Оператор defer позволяет выполнить определенную функцию в конце программы, 
при этом не важно, где в реальности вызывается эта функция.*/
func finish(){
	fmt.Println("Program has been finished")
}


//------------- Panic
// Оператор Panic позволяет сгенерировать ошибку и выйти из программы:
// func divide(x, y float64) float64{
// 	if y == 0{ 
// 			panic("Division by zero!")
// 	}
// 	return x / y
// }
