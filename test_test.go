package test

import "fmt"

func Example_strCat() {
	s := "abc"
	ps := &s
	s += "def"
	fmt.Println(s)
	fmt.Println(*ps)
	// Output:
	// abcdef
	// abcdef
}

func Example_numtostr() {
	var num int
	fmt.Sscanf("57", "%d", &num)
	var s string
	s = fmt.Sprint(3.14)
	s = fmt.Sprintf("%x", 13402077)
	fmt.Println(num)
	fmt.Println(s)
	// Output:
	// 57
	// cc7fdd
}


func Example_array() {
	fruits := [...]string{"수박","사과", "바나나", "토마토"}
	for _, fruit := range fruits {
		if HasConsonantSuffix(fruit) {
			fmt.Printf("%s은 맛있다.\n", fruit )	
		} else {

			fmt.Printf("%s는 맛있다.\n", fruit )
		}
	}
	// Output:
	//수박은 맛있다.
	// 사과는 맛있다.
	// 바나나는 맛있다.
	// 토마토는 맛있다.
}
