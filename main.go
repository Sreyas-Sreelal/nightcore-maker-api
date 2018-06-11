package main

import (
	"io"
	"os"
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
) 

//SongInfo struct Information about converted song
type SongInfo struct{
	NameOfFile string  `json:"nameoffile"`
	Directory string `json:"directory"`
}
const( 
	uploadsDir = "uploads/"
	downloadsDir="downloads/"
)

//HandleSongUpload handler for recieving uploaded file
func HandleSongUpload(ctx iris.Context){
	file, dataRecieved, err := ctx.FormFile("song")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while uploading: %v", err.Error())
		return
	}

	defer file.Close()
	fname := dataRecieved.Filename

	fileOutput, err := os.OpenFile(uploadsDir+fname,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.Application().Logger().Warnf("Error while preparing the new file: %v", err.Error())
		return
	}

	defer fileOutput.Close()
	io.Copy(fileOutput, file)

	data := SongInfo{
		NameOfFile : convert(fname),
		Directory : downloadsDir,
	}
	
	ctx.JSON(data) 
}

//DownloadFile handler for downloading file
func DownloadFile(ctx iris.Context){
	filename := ctx.Params().Get("filename")
	ctx.SendFile(downloadsDir+filename,filename)
}

func main() {
	Cors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000","http://192.168.137.1:3000"},
		AllowCredentials: true,
	})

	app := iris.New()
	app.Use(Cors)
	
	app.Post("/uploader", iris.LimitRequestBodySize(25<<20), HandleSongUpload)
	app.Get("/downloads/:filename",DownloadFile)
	app.Run(iris.Addr(":5000"))
}
