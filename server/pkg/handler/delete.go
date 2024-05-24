package handler

import (
	"os"
	"time"
)

func deleteFile(filename string) {
	<-time.After(1 * time.Second)
	Locallog.Info("Deleting file: ", filename)

	err := os.Remove("res/" + filename)
	if err != nil {
		Locallog.Warn("Try to delete file in ten seconds ", filename, err)
		<-time.After(10 * time.Second)
		err := os.Remove("res/" + filename)
		if err != nil {
			Locallog.Error(err)
			return
		}
	}
}

func DeleteLoop() {
	Locallog.Debug("entering delete loop")
	for {
		Locallog.Debug("entering delete loop: for")

		timebefore := time.Now()
		sleepTime := time.Hour
		time.Sleep(sleepTime)
		dir, err := os.ReadDir("res/")
		if err != nil {
			Locallog.Error(err)
		}
		for i := 0; i < len(dir); i++ {
			info, err := dir[i].Info()
			if err != nil {
				Locallog.Error(err)
			}
			if timebefore.Sub(info.ModTime()) > sleepTime {
				Locallog.Info("Deleting ", dir[i].Name())
				err := os.RemoveAll("res/" + dir[i].Name())
				if err != nil {
					Locallog.Error(err)
				}
			}
		}
	}
}
