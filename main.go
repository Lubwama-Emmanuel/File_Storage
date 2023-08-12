package main

import (
	"github.com/Lubwama-Emmanuel/File_Storage/app"
	filesystem "github.com/Lubwama-Emmanuel/File_Storage/file_system"
)

func main() {
	storage := filesystem.NewFileSystem("test")

	imageApp := app.NewApp(storage)

	imageApp.UploadImage("./hero_2.jpg", "destination")
}
