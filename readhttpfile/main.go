/*package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	response, err := http.Get("https://en.wikipedia.org/wiki/Matrix_(mathematics)")

	if err != nil {
		fmt.Println("error getting information:", err)
		return
	}

	defer response.Body.Close()

	filename := "wikipedia.txt"

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("File created:", filename)
}


*/

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	getinfo("https://en.wikipedia.org/wiki/Iron_Man")

}

func getinfo(hhtplink string) {
	response, err := http.Get(hhtplink)
	if err != nil {
		fmt.Println("error getting information:", err)
		return
	}
	defer response.Body.Close()

	filename := "iron.txt2"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// Parse the HTML content
	z := html.NewTokenizer(response.Body)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			fmt.Println("File created:", filename)
			return
		case tt == html.TextToken:
			// Text token found, extract and write to the file
			text := strings.TrimSpace(string(z.Text()))

			if len(text) > 0 {
				_, err := file.WriteString(text + "\n")
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
