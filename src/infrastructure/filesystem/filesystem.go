package filesystem

import (
	"audioTest/src/interfaces/repositories"
	"fmt"
	"io/ioutil"
	"os"
)

type FileSystem struct {
}

func (f FileSystem) Write(file []byte, fileName string) error {
	mode := int(0777)
	fileMode := os.FileMode(mode)
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	path := fmt.Sprintf("%s/filestore/%s", wd, fileName)
	err = ioutil.WriteFile(path, file, fileMode)
	if err != nil {
		return err
	}

	return nil
}

func FileSystemFactory() repositories.FileSystem{
	return &FileSystem{}
}
