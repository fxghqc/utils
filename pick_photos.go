package main

import (
	"io/ioutil"
	"os"
)

func main() {
	root := "D:/Share/girls"
	files, _ := ioutil.ReadDir(root)
	for _, v := range files {
		if v.IsDir() {
			name := v.Name()
			images, _ := ioutil.ReadDir(root + "/" + name)
			largest := LargestFile(images)
			os.Rename(root+"/"+name+"/"+images[largest].Name(),
				  root+"/ok/"+name)
		}
	}
}

func LargestFile(files []os.FileInfo) int {
	index := 0
	size := files[0].Size()
	for i, file := range files {
		if file.Size() > size {
			size = file.Size()
			index = i
		}
	}
	return index
}