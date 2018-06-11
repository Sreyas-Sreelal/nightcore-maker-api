package main

import (
	"os"
	"os/exec"
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"


func convert(name string) string{
	var outDirectory, outputFile,inputDirectory string
	outputFile = RandStringBytesRmndr(7) + ".mp3"
	outDirectory = "downloads/"
	inputDirectory = "uploads/"

	_,error := exec.Command("ffmpeg",
		"-i",
		inputDirectory + name ,
		"-filter:a",
		"atempo=1.0",
		"-filter:a",
		"asetrate=60000", 
		outDirectory + outputFile).CombinedOutput()
	
		if error != nil{
			return ""
		}
	os.Remove(inputDirectory + name )
	return outputFile
}

func init() {
    rand.Seed(time.Now().UnixNano())
}

//RandStringBytesRmndr generates random string
func RandStringBytesRmndr(n int) string {
	
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

