// Example2.go
package main

import (
	"bufio"
	"container/list"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	//testUserInput2()
	//var name string = "foo , bar , bim"
	//result := strings.Split(name, "\\s*,\\s*")

	// testArray()
	// testArray2()
	//listExample2()
	//mapExample()
	//listMapExample1()
	//listMapExample2()
	/*
		name, age, wage := multiValueReturn()
		fmt.Println(name, age, wage)

		name, age, wage = createStruct()
		fmt.Println(name, age, wage)

		name, age, wage = createStruct()
		fmt.Println(name, age, wage)

		name, age, wage = createStruct()
		fmt.Println(name, age, wage)
	*/

	// deferExample()

	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}
	numbers3 := []int{7, 8, 9}
	sum := sumNumbers(numbers1, numbers2, numbers3)
	fmt.Println("sum = ", sum)

	r := rand.New(rand.NewSource(100))
	l1 := []float64{r.Float64(), r.Float64(), r.Float64()}
	l2 := []float64{r.Float64(), r.Float64(), r.Float64()}
	l3 := []float64{r.Float64(), r.Float64(), r.Float64()}

	fsum := sumFloats(l1, l2, l3)

	fmt.Println("l1 = ", l1)
	fmt.Println("l2 = ", l2)
	fmt.Println("l3 = ", l3)
	fmt.Println("fsum = ", fsum)

	sum = sumList(r.Int(), r.Int(), r.Int(), r.Int())
	fmt.Println("sumList = ", sum)

	//fmt.Println("5! = ", factorial(5))
	// fmt.Println("10! = ", factorial(10))
	fmt.Println(safeDiv(1, 0))
	fmt.Println(safeDiv(2, 1))
	displayPanic()

	var x int = 10
	var y int = 20
	fmt.Println("before swap x = ", x, " y = ", y)
	pointerSwap(&x, &y)
	fmt.Println("after swap x = ", x, " y = ", y)

	person := Person{1, "foo", 10, 100.0, true}
	fmt.Println(person)
	person.setId(10)
	fmt.Println(person)

	person.printProps()
	person.setId(6)
	person.setName("satan")
	person.setAge(66)
	person.setWage(666.6)
	person.setActive(false)
	person.printProps()
}

type Person struct {
	id     int64
	name   string
	age    int32
	wage   float64
	active bool
}

func (person *Person) setId(id int64) {
	person.id = id
}

func (person *Person) getId() int64 {
	return person.id
}

func (p *Person) setName(name string) {
	p.name = name
}

func (p *Person) getName() string {
	return p.name
}

func (p *Person) setAge(age int32) {
	p.age = age
}

func (p *Person) getAge() int32 {
	return p.age
}

func (p *Person) setWage(wage float64) {
	p.wage = wage
}

func (p *Person) getWage() float64 {
	return p.wage
}

func (p *Person) setActive(active bool) {
	p.active = active
}

func (p *Person) isActive() bool {
	return p.active
}

func (p *Person) printProps() {
	fmt.Println("id = ", p.id)
	fmt.Println("name = ", p.name)
	fmt.Println("age = ", p.age)
	fmt.Println("wage = ", p.wage)
	fmt.Println("active =", p.active)
}

func pointerSwap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}

func displayPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("hello world panick")
}
func safeDiv(x int, y int) int {
	defer func() {
		fmt.Println(recover())
	}()

	return x / y
}
func factorial(x int) int {
	if x <= 1 {
		return 1
	}
	return x * factorial(x-1)
}

func sumList(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}

	return sum
}

func sumFloats(l1 []float64, l2 []float64, l3 []float64) float64 {
	var sum float64 = 0
	for _, v := range l1 {
		sum += v
	}

	for _, v := range l2 {
		sum += v
	}

	for _, v := range l3 {
		sum += v
	}

	return sum
}

func sumNumbers(numbers1 []int, numbers2 []int, numbers3 []int) int {
	var sum int = 0

	for _, v := range numbers1 {
		sum += v
	}

	for _, v := range numbers2 {
		sum += v
	}

	for _, v := range numbers3 {
		sum += v
	}

	return sum
}

func deferExample() {
	defer printTwo()
	defer printThree()
	printOne()

}

func printOne() {
	fmt.Println("function One")
}

func printTwo() {
	fmt.Println("function Two")
}

func printThree() {
	fmt.Println("function Three")
}

func createStruct() (string, int, float64) {
	r := rand.New(rand.NewSource(100))
	name := "foo " + strconv.FormatInt(r.Int63(), 10)
	age := r.Int()
	wage := r.Float64()
	return name, age, wage
}

func multiValueReturn() (string, int, float64) {
	name := "foo"
	age := 10
	wage := 100.0
	return name, age, wage
}

func listMapExample2() {
	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}

	fmt.Println(matrix)

	people := []struct {
		name string
		age  int32
		wage float64
	}{
		{"foo", 10, 100.0},
		{"bar", 20, 200.0},
	}

	fmt.Println(people)

	for _, person := range people {
		fmt.Println(person.name, person.age, person.wage)
	}
}
func listMapExample1() {
	list1 := []int{1, 2, 3, 4}
	fmt.Println(list1)

	for _, el := range list1 {
		fmt.Print(el, " ")
	}
	fmt.Println()

	map1 := map[string]string{
		"foo":  "bar",
		"edu":  "tilos",
		"pako": "deko",
	}

	fmt.Println(map1)

	for k, v := range map1 {
		fmt.Println(k, v)
	}

	var list2 [3]string
	list2[0] = "foo"
	list2[1] = "bar"
	list2[2] = "bim"

	fmt.Println(list2)

	var map2 = make(map[string]string)
	map2["foo"] = "bar"
	map2["bim"] = "pako"

	fmt.Println(map2)

	list3 := list.New()
	list3.PushBack("foo")
	list3.PushBack(true)
	list3.PushBack(10)

	for el := list3.Front(); el != nil; el = el.Next() {
		fmt.Print(el.Value, " ")
	}
	fmt.Println()

	var id int64
	var name string
	var age int64
	var wage float64
	var active bool

	var temp string
	fmt.Println("insert id: ")
	fmt.Scanln(&temp)
	id, _ = strconv.ParseInt(temp, 64, 10)
	fmt.Println("insert name: ")
	fmt.Scanln(&temp)
	name = temp
	fmt.Println("insert age: ")
	fmt.Scanln(&temp)
	age, _ = strconv.ParseInt(temp, 64, 10)
	fmt.Println("insert wage: ")
	fmt.Scanln(&temp)
	wage, _ = strconv.ParseFloat(temp, 64)
	fmt.Println("insert active: ")
	fmt.Scanln(&temp)
	active, _ = strconv.ParseBool(temp)

	fmt.Println("<<all props>>")
	fmt.Println("id = ", id)
	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
	fmt.Println("wage = ", wage)
	fmt.Println("active = ", active)

	mapInMap := map[string]map[string]string{
		"foo": map[string]string{"foo1": "1", "foo2": "2", "foo3": "3"},
		"bar": map[string]string{"bar1": "1", "bar2": "2", "bar3": "3"},
	}

	fmt.Println(mapInMap["foo"]["foo1"])
	fmt.Println(mapInMap["bar"]["bar2"])
}

func mapExample() {
	props := make(map[string]string)
	props["id"] = strconv.FormatInt(1, 10)
	props["name"] = "foo"
	props["age"] = strconv.FormatInt(18, 10)
	props["wage"] = strconv.FormatFloat(100.0, 'g', -1, 64)
	props["active"] = strconv.FormatBool(true)

	fmt.Println("id = ", props["id"])
	fmt.Println("name = ", props["name"])
	fmt.Println("age = ", props["age"])
	fmt.Println("wage = ", props["wage"])
	fmt.Println("active = ", props["active"])

	mapInMap := map[string]map[string]string{
		"foo": map[string]string{"foo1": "1", "foo2": "2"},
		"bar": map[string]string{"bar1": "1", "bar2": "2"},
	}

	fmt.Println(mapInMap["foo"]["foo1"])
	fmt.Println(mapInMap["bar"]["bar2"])

	//map another example
	map2 := map[string]string{
		"foo":  "bar",
		"edu":  "tilos",
		"pako": "bim",
	}

	fmt.Println(map2)
	fmt.Println(map2["foo"], map2["edu"], map2["pako"])
}

func listExample2() {
	listNumbers := list.New()
	listNumbers.PushBack(1)
	listNumbers.PushBack(2)
	listNumbers.PushBack(3)
	fmt.Println("listNumbers: ")
	for el := listNumbers.Front(); el != nil; el = el.Next() {
		fmt.Print(el.Value, " ")
	}

	fmt.Println()

	listNames := list.New()
	listNames.PushBack("foo")
	listNames.PushBack("bar")
	listNames.PushBack("bim")

	fmt.Println("listNames len = ", listNames.Len())
	for el := listNames.Front(); el != nil; el = el.Next() {
		fmt.Print(el.Value, " ")
	}
	fmt.Println()

	listNames.Remove(listNames.Front())
	listNames.Remove(listNames.Front())
	listNames.Remove(listNames.Front())

	fmt.Println("listNames len after 3times Remove = ", listNames.Len())
}

func listExample() {
	listNames := list.New()
	listNames.PushBack("foo")
	listNames.PushBack("bar")
	listNames.PushBack("bim")
	//fmt.Println(listNames)
	fmt.Println("listNames: ")

	for el := listNames.Front(); el != nil; el = el.Next() {
		fmt.Print(el.Value.(string), " ")
	}

	fmt.Println()

	fmt.Println("List len = ", listNames.Len())

	listNames.Remove(listNames.Front())
	listNames.Remove(listNames.Front())
	listNames.Remove(listNames.Front())

	fmt.Println("size of listNames after remove = ", listNames.Len())
}

func structArrayExample() {
	names := getNames()
	fmt.Println(strings.Join(names, " , "))

	numbers := getNumbers()
	fmt.Println("numbers : ")
	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
	fmt.Println()

	booleans := getBools()
	fmt.Println("booleans : ")
	for i := 0; i < len(booleans); i++ {
		fmt.Print(booleans[i], " ")
	}
	fmt.Println()

	people := []struct {
		name string
		age  int32
		wage float64
	}{
		{"foo", 10, 100.0},
		{"bar", 20, 200.0},
		{"bim", 30, 300.0},
	}

	fmt.Println(people)

	for i := 0; i < len(people); i++ {
		person := people[i]
		fmt.Println(person.name, person.age, person.wage)
	}
}
func getBools() []bool {
	booleans := []bool{
		true, false, false,
	}

	return booleans
}

func getNumbers() []int64 {
	numbers := []int64{
		1, 2, 3, 4, 5,
	}

	return numbers
}

func getNames() []string {
	names := []string{
		"foo", "bar", "bim",
	}

	return names
}

func testArray2() {
	var names [3]string
	names[0] = "foo"
	names[1] = "bar"
	names[2] = "bim"
	fmt.Println(names)

	names2 := []string{"foo", "bar", "bim"}
	fmt.Println(names2)

	people := []struct {
		name string
		age  int32
		wage float64
	}{
		{"foo", 10, 100.0},
		{"bar", 20, 200.0},
		{"bim", 30, 300.0},
	}

	for _, person := range people {
		fmt.Println(person)
	}

	matrix := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}

	fmt.Println("matrix = ", matrix)
	fmt.Println("matrix (method2 ) : ")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}

		fmt.Println()
	}

	namesMatrix := [][]string{
		[]string{"foo", "bar"},
		[]string{"edu", "tilos"},
	}

	fmt.Println("namesMatrix = ", namesMatrix)
	for _, name := range namesMatrix {
		for _, n := range name {
			fmt.Print(n, " ")
		}
		fmt.Println()
	}

	for i := 0; i < len(namesMatrix); i++ {
		for j := 0; j < len(namesMatrix[i]); j++ {
			fmt.Print(namesMatrix[i][j], " ")
		}
		fmt.Println()
	}

	for i := 0; i < len(namesMatrix); i++ {
		fmt.Println(strings.Join(namesMatrix[i], " , "))
	}

}

func testArray() {
	var names [3]string
	names[0] = "foo"
	names[1] = "bar"
	names[2] = "bim"

	fmt.Println(names[0], names[1], names[2])

	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	var s []int = numbers[0:3]
	fmt.Println("s = ", s)

	people := []struct {
		id   int64
		name string
		age  int32
	}{
		{1, "foo", 10},
		{2, "bar", 20},
	}

	fmt.Println("people = ", people)

	people2 := []struct {
		name   string
		age    int32
		wage   float64
		active bool
	}{
		{"foo", 10, 100.0, true},
		{"bar", 20, 200.0, false},
		{"bim", 30, 300.0, true},
	}

	fmt.Println(people2)

	numbers2 := []int{1, 2, 3, 4, 5, 6}
	numbers3 := numbers2[:2]
	numbers4 := numbers2[2:]

	fmt.Println("numbers2 = ", numbers2)
	fmt.Println("numbers3 = ", numbers3)
	fmt.Println("numbers4 = ", numbers4)

	fmt.Println("numbers2 , len = ", len(numbers2), "cap = ", cap(numbers2), "self = ", numbers2)

	numbers5 := make([]int, 5)
	fmt.Println(numbers5)

	matrix := [][]int64{
		[]int64{1, 2, 3},
		[]int64{4, 5, 6},
		[]int64{7, 8, 9},
	}

	fmt.Println(matrix)

	fmt.Println("another way of printing matrix: ")
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}

	namesMatrix := [][]string{
		[]string{"foo", "bar"},
		[]string{"edu", "tilos"},
		[]string{"pako", "deko"},
	}

	fmt.Println("printing namesMatrix(Method1):")
	for i := 0; i < len(namesMatrix); i++ {
		for j := 0; j < len(namesMatrix[i]); j++ {
			fmt.Print(namesMatrix[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println("printing namesMatrix(Method2): ")
	for i := 0; i < len(namesMatrix); i++ {
		fmt.Println(strings.Join(namesMatrix[i], " , "))
	}

	var names2 []string
	names2 = append(names2, "foo")
	names2 = append(names2, "bim")
	names2 = append(names2, "bar")
	fmt.Println(strings.Join(names2, ", "))

	var names3 []string
	names3 = append(names3, "foo", "bar", "bim")
	fmt.Println(strings.Join(names3, " , "))

	names3 = append(names3, "edu", "tilos", "newPako")
	for _, oneName := range names3 {
		fmt.Print(oneName, " ")
	}

	fmt.Println()

	names4 := make([]string, 5)

	for i, name := range names4 {
		names4[i] = name + strconv.Itoa(i)
	}

	fmt.Println("names4 = ", names4)

}

func testUserInput2() {
	var temp string
	var name string
	var age int64
	var wage float64
	var active bool

	fmt.Println("insert name: ")
	fmt.Scanln(&temp)
	name = temp

	fmt.Println("insert age: ")
	fmt.Scanln(&temp)
	age, _ = strconv.ParseInt(temp, 10, 64)

	fmt.Println("insert wage: ")
	fmt.Scanln(&temp)
	wage, _ = strconv.ParseFloat(temp, 64)

	fmt.Println("insert active: ")
	fmt.Scanln(&temp)
	active, _ = strconv.ParseBool(temp)

	fmt.Println("<<Result: >>")
	fmt.Println("name = ", name)
	fmt.Println("age = ", age)
	fmt.Println("wage = ", wage)
	fmt.Println("active = ", active)
}

func testUserInput() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("insert name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("insert age: ")
	_age, _ := reader.ReadString('\n')
	//age := int(_age)

	fmt.Println("name = ", name, "age = ", _age)
}
