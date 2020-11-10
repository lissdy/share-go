package basic

import (
	"net/http"
	"fmt"
	"time"
	"sync"
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
	checkURLs(urls)
	fmt.Printf("%.2f s used", time.Since(start).Seconds())

}

func checkURLs(urls []string)  {
	c := make(chan string)
	var wg sync.WaitGroup
	for _, url := range urls{
		wg.Add(1)
		go checkURL(url,c,&wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()



	for msg := range c {
		fmt.Println(msg)
	}
	fmt.Println(111)
}

func checkURL(url string, c chan string, wg *sync.WaitGroup)  {
	defer (*wg).Done()
	_, err := http.Get(url)
	if err!=nil{
		c <- url + " can not be reached\n"
	}else {
		c <- url + " can be reached\n"
	}
}

