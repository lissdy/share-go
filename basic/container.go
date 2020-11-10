package basic

import (
	"fmt"
	"reflect"
)

func main()  {
	var arr [3]int = [3]int{1,2,3}
	fmt.Println(reflect.TypeOf(arr))
	var sli []int=[]int{4,5,6}
//	var ss *[]int=&sli

	fmt.Println(&arr)
	//var sli1 []int = make([]int,3,10)
	sli = append(sli,10)
	fmt.Println(cap(sli))
	fmt.Println(sli)
	//arr = append(arr,5)
	fmt.Println(reflect.TypeOf(sli))

	dic := map[int]int{}
	dic[1]=1
	delete(dic,1)


}
