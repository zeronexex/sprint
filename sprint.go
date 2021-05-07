/* Use it to divide large text file into smaller text files */
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

var aSlice []string

var mainFIle string
var targetPath string
var limit int
var fileEx string

func main() {
	flag.Usage = func() {
		fmt.Println(`sprint <flags> <args>
	flags : 
		-c int
			The size of chunks.
		-fn string
			This name is will be used for created files with numbers added on right incrementally.
	args :
		First argument: is the path where the file you want to split exists
		Second argument: is the path where new chunked files will be created
	ex :
		sprint -c 2 -f raft_small /path/to/file/for/splitting /path/where/new/files/will/be/written
		sprint -c 10 -f "raft small.txt" /path/to/file/for/splitting /path/where/new/files/will/be/written
`+"\n\t\t"+fmt.Sprintf("%v", aurora.BrightCyan("Crafted with")), "🤍", aurora.BrightCyan("by"), aurora.BrightWhite("Rewinter"))
		os.Exit(2)
	}

	var chunk int
	flag.IntVar(&chunk, "c", 0, "The size of chunks.")
	var filename string
	flag.StringVar(&filename, "fn", "", "This name is will be used for created files with numbers added on right incrementally.")

	flag.Parse()
	args := flag.Args()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}
	mainPath := args[0] //	Path of the file containing text

	w := args[1] //	Path where split files will be created

	if mainPath != "" {
		mainFIle = mainPath
	}
	if w != "" {
		targetPath = w
	}
	if chunk != 0 {
		limit = chunk
	}
	if filename != "" {
		fileEx = filename
	}

	file, err := os.Open(mainFIle)
	reError(err)

	defer file.Close()

	x := bufio.NewScanner(file)
	for x.Scan() {
		aSlice = append(aSlice, x.Text())
	}
	xe := 0
	for i := 0; i < len(aSlice); i += limit {
		batch := aSlice[i:min(i+limit, len(aSlice))]
		twoDslice = append(twoDslice, batch)
		xe++
		//fmt.Println(xe, batch)
	}

	fileCount := len(twoDslice)

	fmt.Printf("%v files should be created\n", fileCount)
	fileGenerator(fileCount, targetPath)
	TargetWriter(fileCount, targetPath)

}

func fileGenerator(count int, path string) {

	for i := 0; i < count; i++ {
		file, err := os.Create(filepath.Join(path, fmt.Sprintf("%v%v.txt", fileEx, i)))
		defer file.Close()
		if err != nil {
			log.Fatalln(err)
		}

	}
	fmt.Println("creating...")

}

func TargetWriter(count int, path string) {
	/*for a, x := range twoDslice{
		fmt.Println(a, x)
	}*/

	for i := 0; i < count; i++ {
		xpath := filepath.Join(path, fmt.Sprintf("%v%v.txt", fileEx, i))
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

func reError(e error) {
	if e != nil {
		//panic(e)
		fmt.Println("⚠ All flags must be set. Be sure to set them correctly")
		flag.Usage()


		os.Exit(10)
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

