package repositories

import (
	"audioTest/src/domain"
	internal "audioTest/src/infrastructure/nats"
	"bytes"
	"errors"
	"fmt"
	"github.com/go-audio/wav"
	"math"
	"os"
)

type DbHandler interface {
	Execute(statement string)
	Query(statement string) Row
}

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

type FileSystem interface {
	Write(file []byte, fileName string) error
}

type AudioRepository struct {
	dbHandler  DbHandler
	fileSystem FileSystem
	broker     internal.INats
}

func (repo *AudioRepository) GetAudioInfo(filename string) (*domain.Audio, error) {
	row := repo.dbHandler.Query(fmt.Sprintf("select filepath, duration, sampleRate, numChans, bitDepth from files where filename = '%s'", filename))
	var filepath string
	var duration int
	var sampleRate int
	var numChans int
	var bitDepth int

	row.Next()
	row.Scan(&filepath, &duration, sampleRate, numChans, bitDepth)

	audio := &domain.Audio{
		SampleRate: sampleRate,
		NumChans:   numChans,
		BitDepth:   bitDepth,
		Duration:   duration,
	}

	return audio, nil

}

func (repo *AudioRepository) StoreFile(wavFile []byte, fileName string) error {
	reader := bytes.NewReader(wavFile)
	decoder := wav.NewDecoder(reader)
	decoder.ReadInfo()
	duration, err := decoder.Duration()
	if err != nil {
		return err
	}

	timeInSeconds := int(math.Round(duration.Seconds()))

	if !decoder.IsValidFile() {
		return errors.New("file not valid")
	}
	//Write to the database
	repo.dbHandler.Execute(fmt.Sprintf("INSERT INTO files (filename, filepath, duration, sampleRate, numChans, bitDepth)  VALUES ('%s', '%s', %d, %d, %d, %d)",
		fileName,
		GetPath(fileName),
		timeInSeconds,
		int(decoder.SampleRate),
		int(decoder.NumChans),
		int(decoder.BitDepth),
	))

	//Write audio file to the filesystem
	err = repo.fileSystem.Write(wavFile, fileName)
	if err != nil {
		return err
	}

	// Publish successful audio file save
	var message = fmt.Sprintf("File %s successfully uploaded", fileName)
	err = repo.broker.Publish("file.upload", []byte(message))
	if err != nil {
		return err
	}
	return nil
}

func AudioRepositoryFactory(dbHandler DbHandler, fileSystem FileSystem, natsConn internal.INats) *AudioRepository {
	return &AudioRepository{
		dbHandler:  dbHandler,
		fileSystem: fileSystem,
		broker:     natsConn,
	}
}

func GetPath(fileName string) string {
	wd, _ := os.Getwd()

	path := fmt.Sprintf("%s/filestore/%s", wd, fileName)

	return path
}
