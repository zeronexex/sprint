
package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/logrusorgru/aurora"
	"os"
)

var aRay []string

func main()  {

	flag.Usage = func() {
		fmt.Println("Example:\n\ttwim base.txt file2.txt file3.txt ...\n\tUse 'twim -s' for silent mode")
	}

	var s bool
	flag.BoolVar(&s, "s", false, "Silent mode")
	flag.Parse()

	if flag.NArg() <= 0 {
		flag.Usage()
		os.Exit(2)
	}

	xp := flag.Args()
	w := xp[0]
	x := xp[0:]

	if len(xp) <= 0 {
		flag.Usage()
		os.Exit(2)
	} else {

		file, err := os.Open(w)
		if err != nil {
			fmt.Println("Something went wrong while opening the first file")
		}
		defer file.Close()

		scan := bufio.NewScanner(file)
		for scan.Scan() {
			aRay = append(aRay, scan.Text())
		}


		ff, err := os.OpenFile(w, os.O_APPEND|os.O_WRONLY, 0600)

		if err != nil {
			fmt.Println(aurora.BrightRed("Can't open file to write"))
		}
		defer file.Close()


		for _, gFile := range x {
			fl, err := os.Open(gFile)
			if err != nil {
				fmt.Println(aurora.BrightRed("Can't open file provided to read from"))
			}

			scan := bufio.NewScanner(fl)
			for scan.Scan() {
				st := scan.Text()
				if stringInSlice(st, aRay) == false {

					aRay = append(aRay, st)
					_, err = ff.WriteString("\n" + st)
					if !s {
						fmt.Println(st)
					}
				}

			}
		}
	}

}



func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
