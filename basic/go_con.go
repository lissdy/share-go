package basic

import (
	"net/http"
	"fmt"
	"time"
)

func main()  {
	start := time.Now()
	var urls []string = []string{
		"https://github.com/appleboy/gorush",
		"https://github.com/gin-gonic/gin",
		"https://github.com/astaxie/beego",
		"https://github.com/helm/helm",
		"https://github.com/argoproj/argo-cd",
	}
	c := make(chan string)
	for _, v := range urls{
		go checkurl(v,c)
	}
	time.Sleep(1000)
	close(c)
	//for msg := range c{
	//	fmt.Println(msg)
	//}

	//msg := <- c
	//fmt.Println(msg)

	for msg := range c{
		fmt.Println(msg)
	}

	fmt.Printf("%.2f s used", time.Since(start).Seconds())

}

func checkurl(url string, c chan string)  {
	_, err := http.Get(url)
	if err!=nil{
		c <- url + " can not be reached\n"
	}else {
		c <- url + " can be reached\n"
	}
}

