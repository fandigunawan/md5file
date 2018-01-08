package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func CalculateChecksum(base, path string) {
	_, err := os.Stat(os.Args[0])
	if err != nil {
		fmt.Println("Invalid path")
		return
	}
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		fileName := path + string(os.PathSeparator) + file.Name()
		if file.IsDir() {
			CalculateChecksum(base, fileName)
		} else {
			hash := md5.New()
			f, _ := os.Open(fileName)
			io.Copy(hash, f)

			fmt.Printf("%x", hash.Sum(nil))
			temp, _ := f.Stat()
			fmt.Printf(" : %d", temp.ModTime().Unix())
			fmt.Printf(" : %d bytes", temp.Size())
			fmt.Println(" : " + strings.Replace(fileName, base, "", -1))

			defer f.Close()
		}
	}

}
func displayHelp(){
	println("Usage : base_path")
	println("Output format : ")
	println(";MD5 Checksum : unix_time : size : file_name")
	println("unix_time is time calculated since UNIX Epoch")
	println("file_name is file path relative to base_path")
}
func main() {
	if len(os.Args) != 2 {
		fmt.Println("Invalid parameter")
		displayHelp()
		return
	}
	println(";MD5 Checksum : unix_time : size : file_name")
	CalculateChecksum(os.Args[1], os.Args[1])
}
