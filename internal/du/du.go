package du

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type Storage struct {
	Index    int
	Path     string
	Name     string
	FileSize int64
}

type Data struct {
	Path   string
	NFile  int
	NBytes int64
}

func Start(dir string) map[int]*Data {

	var roots = strings.Split(dir, ",")

	var storage = make(chan Storage)
	var lret = make(map[int]*Data)

	var wg sync.WaitGroup

	for index, root := range roots {
		lret[index] = &Data{
			Path:   root,
			NFile:  0,
			NBytes: 0,
		}
		wg.Add(1)
		go walkSingleDir(root, index, &wg, storage)
	}

	go func() {
		wg.Wait()
		close(storage)
	}()

	for item := range storage {
		lret[item.Index].NFile = lret[item.Index].NFile + 1
		lret[item.Index].NBytes = lret[item.Index].NBytes + item.FileSize
	}

	return lret
}

func PrintDiskUsage(path string, nfiles int, nsize float64, unit string) {
	fmt.Printf("%s\t %d files  \t%.1f %s\n", path, nfiles, nsize, unit)
}
func walkSingleDir(root string, index int, wg *sync.WaitGroup, storage chan<- Storage) {
	defer wg.Done()
	for _, entry := range dirents(root) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(root, entry.Name())
			go walkSingleDir(subdir, index, wg, storage)
		} else {
			newStorage := Storage{
				index,
				root,
				entry.Name(),
				entry.Size(),
			}
			storage <- newStorage
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
