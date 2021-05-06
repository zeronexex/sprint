/* Use it to divide subdomains into text files */
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"log"
	"os"
	"path/filepath"
)

var twoDslice [][]string

var subDomains []string




var alldomains string
var writepath string
var limit int
var fname string

func main() {
	fmt.Println(aurora.Cyan("Crafted with") , "🤍", aurora.Cyan("by") ,aurora.BrightWhite("Rewinter"))
	var cpath string
	flag.StringVar(&cpath, "cpath", "", "Absolute path of the file containing cnames")
	var w string
	flag.StringVar(&w, "w", "","The absolute path of where split files will be created")
	var chunk int
	flag.IntVar(&chunk, "chunk", 0, "Count of cnames each file will contain (except last one)")
	var filename string
	flag.StringVar(&filename, "filename", "", "This name is gonna be used for created files with numbers added on right incrementally")



	flag.Parse()

	if cpath != "" {
		alldomains = cpath
	}
	if w != "" {
		writepath = w
	}
	if chunk != 0 {
		limit = chunk
	}
	if filename != "" {
		fname = filename
	}




	file, err := os.Open(alldomains)
	errcheck(err)

	defer file.Close()

	x := bufio.NewScanner(file)
	for x.Scan() {
		subDomains = append(subDomains, x.Text())
	}
	xe := 0
	for i := 0; i <len(subDomains); i += limit {
		batch := subDomains[i:min(i+limit, len(subDomains))]
		twoDslice = append(twoDslice, batch)
		xe ++
		//fmt.Println(xe, batch)
	}

	fileCount := len(twoDslice)



	fmt.Printf("%v files should be created\n", fileCount)
	files_to_create(fileCount, writepath)
	subDWriter(fileCount, writepath)

}




func files_to_create(count int, path string)  {

	for i := 0; i < count; i ++ {
		file, err := os.Create(filepath.Join(path, fmt.Sprintf("%v%v.txt", fname, i)))
		defer file.Close()
		if err != nil {
			log.Fatalln(err)
		}

	}
	fmt.Println("creating...")

}


func subDWriter(count int, path string)  {
	/*for a, x := range twoDslice{
		fmt.Println(a, x)
	}*/

	for i := 0; i < count; i ++ {
		xpath := filepath.Join(path, fmt.Sprintf("%v%v.txt", fname, i))
		file, err := os.OpenFile(xpath, os.O_RDWR, 0644)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println("split")
		for _, x := range twoDslice[i] {
			_, _ = file.WriteString(x + "\n")
		}
	}
	fmt.Println("Done! GG haxor ;)")
}


func errcheck(e error) {
	if e != nil {
		//panic(e)
		fmt.Println("⚠ All flags must be set. Be sure to set them correctly")
		flag.PrintDefaults()

		os.Exit(10)
	}
}


func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
