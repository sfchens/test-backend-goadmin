package utils

import (
	"csf/library/viper"
	"os"
)

type GetDirRes struct {
	Filename string `json:"filename" form:"filename"`
	DirName  string `json:"dir_name" form:"dir_name"`
}

func GetDir(paths string) (dir string, err error) {

	date := NowDate()
	dir = viper.NewViper.GetString("upload.path") + paths + date

	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return
		}
	}
	return
}

func GetPicTureDir(paths string) (dir string, err error) {
	date := NowDate()
	dir = viper.NewViper.GetString("upload.path") + "picture/" + paths + date

	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return
		}
	}
	return
}

func GetVideoDir(paths string) (dir string, err error) {
	date := NowDate()
	dir = viper.NewViper.GetString("upload.path") + "video/" + paths + date

	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return
		}
	}
	return
}
