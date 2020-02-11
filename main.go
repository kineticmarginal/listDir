package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"sort"
)

var (
	d = flag.String("d", ".", "Directory to process")
	a = flag.Bool("a", false, "Print all info")
	h = flag.Bool("h", false, "Print size in human-readable format ")
	sorted = flag.String("sort", "", "Sort files in directory")
)

func hrSize(fsize int64) string {
	if (fsize < 1024) {
		return strconv.Itoa(int(fsize)) + "B"
	} else if (fsize >= 1024 && fsize < 1024*1024){
		fsize /= 1024
		return strconv.Itoa(int(fsize)) + "KB"
	} else if (fsize >= 1024*1024 && fsize < 1024*1024*1024){
		fsize /= 1024*1024
		return strconv.Itoa(int(fsize)) + "MB"
	}
	return "0B"
}

func printAll(file os.FileInfo) {
	time := file.ModTime().Format("Jan 06 15:4")
	fSize := strconv.Itoa(int(file.Size()))
	if *h {
		fmt.Printf("%s %s %s \n", hrSize(int64(file.Size())), time, file.Name())
	} else {
		fmt.Printf("%s %s %s \n", fSize, time, file.Name())
	}
}

type SortByDate []os.FileInfo

func (ss SortByDate) Len() int{
	return len(ss)
}

func (ss SortByDate) Swap(i int, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

func (ss SortByDate) Less(i int, j int) bool {
	return ss[i].ModTime().UnixNano() < ss[j].ModTime().UnixNano()
}

func main() {
	flag.Parse()
	files, _ := ioutil.ReadDir(*d)
	if *sorted == "date" {
		sort.Sort(SortByDate(files))
	}
	for _, file := range files {
		if *a {
			printAll(file)
		} else {
			fmt.Println(file.Name())
		}
	}
}
