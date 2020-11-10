package basic

import "fmt"

func main()  {
	var a int = 1
	//b := "b"
	fmt.Println(a)
	fmt.Println(&a)
//	var b *string = &a
	fmt.Println(*&a)
}
