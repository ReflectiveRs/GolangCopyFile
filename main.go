package main

import (
	"log"
	"io/ioutil"
	"io"
	"os"
	"time"
)
var Dir string
var Dircp string

func main() {
	Dir = " " //// Catalog copy
	Dircp = " " //// Catalog save

	logf, err := os.OpenFile("logfile", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0660) //// log
	if err != nil {
	}
	defer logf.Close()
	log.SetOutput(logf)

	for {
		log.Println("Syncing files")
		files_copy, err := ioutil.ReadDir(Dir)
		if err != nil {
			log.Println(err)
		}
		for _, files := range files_copy {
			if files.IsDir() == true {
				os.MkdirAll(Dircp+"/"+files.Name(), 0777)
				copyfolder(Dir+"/"+files.Name(), Dircp+"/"+files.Name())
				continue
			}
			copy(Dir+"/"+files.Name(), Dircp+"/"+files.Name())
		}
		time.Sleep(15 * time.Minute) //// time between copy files
	}
}

func copy(src, dst string) {
	if _, err := os.Stat(dst); os.IsNotExist(err) {
		log.Println("Copy the new file: " + src)
		file1, err := os.Open(src)
		if err != nil {
			log.Println(err)
		}
		defer file1.Close()

		file2, err := os.Create(dst)
		if err != nil {
			log.Println(err)
		}
		defer file2.Close()

		_, err = io.Copy(file2, file1)
		if err != nil {
			log.Println(err)
		}
		file2.Close()
	}
}
func copyfolder(dir1, dir2 string) {
	dir_file, err := ioutil.ReadDir(dir1)
	if err != nil {
		log.Println(err)
	}
	for _, file := range dir_file {
		if file.IsDir() == true {
			os.MkdirAll(dir2+"/"+file.Name(), 0777)
			copyfolder(dir1+"/"+file.Name(), dir2+"/"+file.Name())
		}
		copy(dir1+"/"+file.Name(), dir2+"/"+file.Name())
	}
}
