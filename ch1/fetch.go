/*
 * @program: The-Go-Programming-Language
 * @author: Leon
 * @create: 2020-09-12 16:01
 */
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			fmt.Println(!strings.HasPrefix(url, "http://"))
			url = "http://" + url

		}
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch:%v\v", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
		fmt.Println(resp.StatusCode)
	}
}
