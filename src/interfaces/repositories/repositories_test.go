package repositories_test

import (
	"audioTest/mocks"
	"audioTest/src/interfaces/repositories"
	"errors"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"testing"
)

func TestStoreFile(t *testing.T) {
	t.Run("valid wav file", func(t *testing.T) {
		file, err := ioutil.ReadFile("../../../fixtures/kick.wav")
		if err != nil {
			t.Error(err)
		}

		// Setup mock
		dbHandler := &mocks.DbHandler{}
		dbHandler.On("Execute", mock.Anything).Return()
		fileSystem := &mocks.FileSystem{}
		fileSystem.On("Write", mock.Anything, mock.Anything).Return(nil)

		nats := &mocks.INats{}
		nats.On("Publish", mock.Anything, mock.Anything).Return(nil)

		repository := repositories.AudioRepositoryFactory(dbHandler, fileSystem, nats)
		err = repository.StoreFile(file, "bad.wav")
		if err != nil {
			t.Errorf("Got unexpected error: %s", err.Error())
		}
	})

	t.Run("invalid wav file", func(t *testing.T) {
		file, err := ioutil.ReadFile("../../../fixtures/bad.wav")
		if err != nil {
			t.Error(err)
		}

		// Setup mock
		dbHandler := &mocks.DbHandler{}
		dbHandler.On("Execute", mock.Anything).Return()
		fileSystem := &mocks.FileSystem{}
		fileSystem.On("Write", mock.Anything, mock.Anything).Return(errors.New("fileystem failure"))

		nats := &mocks.INats{}
		nats.On("Publish", mock.Anything, mock.Anything).Return(nil)
		repository := repositories.AudioRepositoryFactory(dbHandler, fileSystem, nats)
		err = repository.StoreFile(file, "bad.wav")
		if err == nil {
			t.Error("Expected error to be thrown")
		}
	})

	t.Run("file system error should fail and return", func(t *testing.T) {
		file, err := ioutil.ReadFile("../../../fixtures/kick.wav")
		if err != nil {
			t.Error(err)
		}

		// Setup mock
		dbHandler := &mocks.DbHandler{}
		dbHandler.On("Execute", mock.Anything).Return()
		fileSystem := &mocks.FileSystem{}
		fileSystem.On("Write", mock.Anything, mock.Anything).Return(errors.New("fileystem failure"))

		nats := &mocks.INats{}
		nats.On("Publish", mock.Anything, mock.Anything).Return(nil)

		repository := repositories.AudioRepositoryFactory(dbHandler, fileSystem, nats)
		err = repository.StoreFile(file, "kick.wav")
		if err == nil {
			t.Error("Expected error to be thrown")
		}
	})
}
