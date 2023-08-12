package app

import "fmt"

type IStorage interface {
	Upload(string, string) (bool, string)
}

type App struct {
	storage IStorage
}

// Image system
func (a *App) UploadImage(imagePath string, destinationPath string) bool {
	boolen, result := a.storage.Upload(imagePath, destinationPath)

	fmt.Println("got this", result)

	return boolen
}

func NewApp(storage IStorage) *App {
	return &App{storage: storage}
}
