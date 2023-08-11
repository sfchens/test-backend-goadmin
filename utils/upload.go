package utils

import (
	"csf/library/easy_config"
	"os"
)

type GetDirRes struct {
	Filename string `json:"filename" form:"filename"`
	DirName  string `json:"dir_name" form:"dir_name"`
}

func GetDir(paths string) (dir string, err error) {

	date := NowDate()
	dir = easy_config.Viper.GetString("upload.path") + paths + date

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
	dir = easy_config.Viper.GetString("upload.path") + "picture/" + paths + date

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
	dir = easy_config.Viper.GetString("upload.path") + "video/" + paths + date

	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		if err = os.MkdirAll(dir, 0777); err != nil {
			return
		}
	}
	return
}
