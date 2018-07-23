//********************************************************************
//golang练习例子
//********************************************************************
//使用select的多路复用
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		//case <-tick:
		case <-time.After(10 * time.Second)
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("launch ...")
}

//-----------------------------------------------------------------------------------
import fmt
func main(){
ch := make(chan int, 2)
for i := 0; i < 10; i++ {
	select{
	case x := <-ch:
		fmt.Println(x)
	case ch <- i:	
	}
}}
//-------------------------------------------------------------------------------------
import (
	"flag"
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
)
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir(){
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		}else{
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	//确定初始目录
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0{
		roots = []string{"."}
	}
	
	//遍历文件树
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots{
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	
	//输出结果
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
//----------------------------------------------------------------------






