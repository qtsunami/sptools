package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func ScanRootPath(dir string) {

	var roots = strings.Split(dir, ",")

	fileSize := make(chan int64)

	var wg sync.WaitGroup

	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSize)
	}

	go func() {
		wg.Wait()
		close(fileSize)
	}()

	var nfiles, nbytes int64
	for size := range fileSize {
		nfiles++
		nbytes += size
	}

	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func walkDir(root string, wg *sync.WaitGroup, fileSize chan<- int64) {
	defer wg.Done()
	for _, entry := range dirents(root) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(root, entry.Name())
			go walkDir(subdir, wg, fileSize)
		} else {
			fileSize <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(path string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
