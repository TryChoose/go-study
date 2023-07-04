package main

//获取一个URL
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//获取URL
func main() {
	resp, err := http.Get("https://www.runoob.com/w3cnote/android-tutorial-intro.html")
	if err != nil {
		log.Fatal(err)
		return
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("%s", b)
}
