package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
)

func main() {
	flag.Parse()
	if len(flag.Args()) < 2 {
		return
	}

	word := flag.Args()[0]
	path := flag.Args()[1]
	scan(word, path)
}

func scan(word, path string) {
	files, _ := os.ReadDir(path)
	var dQueue, fQueue []string
	for _, file := range files {
		fn := path + `/` + file.Name()
		if file.IsDir() {
			dQueue = append(dQueue, fn)
		} else {
			extension := strings.Split(file.Name(), ".")[1]
			if strings.ToUpper(extension) == "TXT" {
				fQueue = append(fQueue, fn)
			}
		}
	}

	for i := range dQueue {
		scan(word, dQueue[i])
	}

	wg := sync.WaitGroup{}
	wg.Add(len(fQueue))
	for i := range fQueue {
		go func(fn string) {
			defer wg.Done()

			f, _ := os.Open(fn)
			scanner := bufio.NewScanner(f)
			line := 1
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), word) {
					fmt.Println(fn, line)
				}
				line++
			}
		}(fQueue[i])
	}
	wg.Wait()
}
