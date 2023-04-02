package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ______________________________ Array & Slice example ______________________________
	var list []int64
	l := len(list)
	fmt.Println("Example Len = ", l) //0
	c := cap(list)
	fmt.Println("Example Cap = ", c) //0

	var list2 []int64
	list2 = []int64{1, 2, 3, 4, 5}
	l2 := len(list2)
	fmt.Println("Example Len2 = ", l2) //5
	c2 := cap(list2)
	fmt.Println("Example Cap2 = ", c2) //5

	//С помошью Make можно создать срез
	//1) параметр тип среза,        		[int64]
	//2) параметр размером(длина) среза,	[len]
	//3) параметр емкость среза.			[cap]

	var list3 []int64
	list3 = make([]int64, 0, 5) // []
	l3 := len(list3)
	fmt.Println("Example Len3 = ", l3) //0
	c3 := cap(list3)
	fmt.Println("Example Cap3 = ", c3) //5

	var list4 []int64
	list4 = make([]int64, 5, 5) //[0,0,0,0,0]
	l4 := len(list4)
	fmt.Println("Example Len4 = ", l4) //5
	c4 := cap(list4)
	fmt.Println("Example Cap4 = ", c4) //5

	var list5 []int64
	list5 = []int64{1, 2, 3, 4} //[1,2,3,4]
	//Функция Append добавляет элемент к срезу
	list5 = append(list5, 5)
	fmt.Println("Use Append = ", list5) //[1,2,3,4,5]
	list5 = make([]int64, 0, 3)
	fmt.Println("Use Make = ", list5) //len: 0, cap: 3
	list5 = append(list5, 1)          // [1] len: 1 cap: 3   //Увеличиваем Len на 1
	list5 = append(list5, 2)          // [1,2] len: 2 cap: 3
	list5 = append(list5, 3)          // [1,2,3] len: 3 cap: 3
	// Если Len > cap, создается новый массив, значения переносятся туда и размер массива увеличивается в 2 раза больше предыдущего
	list5 = append(list5, 4) // [1,2,3,4] len: 4 cap: 6
	l5 := len(list5)
	fmt.Println("Example Len5 = ", l5) //5
	c5 := cap(list5)
	fmt.Println("Example Cap5 = ", c5) //8

	//Пример String
	var list6 []int64
	list6 = make([]int64, 0, 3)
	fmt.Println("Len =", len(list6), "Cap =", cap(list6))

	list6 = append(list6, 1)
	fmt.Println("Len =", len(list6), "Cap =", cap(list6))

	list6 = append(list6, 2)
	fmt.Println("Len =", len(list6), "Cap =", cap(list6))

	list6 = append(list6, 3)
	fmt.Println("Len =", len(list6), "Cap =", cap(list6))

	list6 = append(list6, 4)
	fmt.Println("Len =", len(list6), "Cap =", cap(list6))
	fmt.Println(list6)

	//work with string type
	word := "\tsome text\n"
	title := `\tsome text\n`
	fmt.Println("work with string = ", word, title)

	bute := 'c'
	rune := '漢'
	fmt.Println("Bute = ", bute, "Rune = ", rune)

	str := "13245"
	fmt.Println("[0] = ", str[0])
	fmt.Println("[1] = ", str[1])
	fmt.Println("Len = ", len(str))

	str2 := "стройка"
	fmt.Println("Len = ", len(str2))
	fmt.Println("[2] =", str2[2])
	newStr := str2[2:]
	fmt.Println(str2, "[2:] = ", newStr)

	// ------------------- Пример копирование -------------------
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println("Copy slice1, slice2", slice1, slice2)

	// ------------------- Пример map -------------------
	x := make(map[string]int)
	x["key"] = 10
	x["id"] = 15
	fmt.Println(x)
	delete(x, "id")
	fmt.Println("Delete x = ", x)

	elements := make(map[string]string)
	elements["O"] = "Oxygen"
	elements["Ne"] = "Neon"
	elements["Li"] = "Lithium"
	elements["N"] = "Nitrogen"
	fmt.Println(elements)

	// ------------------- Пример с использовании append() -------------------
	number := []int{}
	for i := 0; i < 4; i++ {
		number = append(number, i)
	}
	fmt.Println("Use Append = ", number)

	// ------------------- Пример с использовании cap() -------------------
	numbers := make([]int, 4)
	for i := 0; i < cap(numbers); i++ {
		numbers[i] = i
	}
	fmt.Println(numbers)
	fmt.Println("Use Cap = ", cap(numbers), " = ", reflect.TypeOf(cap(numbers)))

	// -------------------Пример с многомерным массивом -------------------
	sea := [][]string{{"shark", "octopus", "squid", "mantis shrimp"}, {"Sammy", "Jesse", "Drew", "Jamie"}}
	fmt.Println("Sea[1][2]", sea[1][2])
	fmt.Println("Sea[0][2]", sea[0][2])

	// ------------------- Удаление елемента в масиве -------------------
	user := [8]string{"Sara", "Alice", "J", "Liza", "Kate", "Lara", "Kis", "freya"}
	pre := int(len(user))
	// ------------------- Пример с Циклом -------------------
	for i := 0; i < pre; i++ {
		var item string
		item += user[i]
		fmt.Print(i, " = ", item, "\n")
	}
	// ------------------- Пример с Циклом -------------------
	var total string
	for _, value := range user {
		total += value
		fmt.Printf("%q\n", total)
	}
	// ------------------- Пример с Циклом -------------------
	for i := 0; i < pre; i++ {
		var link string
		link += user[i]
		println(i, "=", link)
		fmt.Println("TypeOf = ", reflect.TypeOf(pre))
	}
	// ------------------- Remove elements slice&array -------------------
	fmt.Println("Array = ", user)
	Users := user[:]
	n := 3
	Users = append(user[:n], user[n+1:]...)
	fmt.Println("Use Append = ", Users)

	// ------------------- Выборочнный вывод масива -------------------
	select1 := user[2:6]
	select2 := user[:4]
	select3 := user[3:]
	select4 := user[:]
	fmt.Println("Array [2:6] = ", select1)
	fmt.Println("Array [:4] = ", select2)
	fmt.Println("Array [3:] = ", select3)
	fmt.Println("Array [:] = ", select4)

	// ------------------- Изменения значения в масиве -------------------
	var number2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("number = ", number2)
	fmt.Println("number[2] = ", number2[2])
	number2[2] = 1995
	fmt.Println("Number[2] = ", number2[2])

	// ------------------- Задаенм индексы в рукчную -------------------
	number3 := []int{1: 1, 3: 2, 2: 17}
	fmt.Println("Array = ", number3)
	fmt.Println("Array [2] = ", number3[2])
	colors := [3]string{2: "blue", 0: "red", 1: "yellow"}
	fmt.Println("Array Colors = ", colors)
	fmt.Println("Colors [2] = ", colors[2])

	// ------------------- Добавление элементов с помощью append() -------------------
	var users2 []string = make([]string, 3)
	users2 = append(users2, "bob")
	users2[0] = "adam"
	users2[1] = "rex"
	users2[2] = "roxy"
	fmt.Println("Добавление элементов = ", users2[2])
}
