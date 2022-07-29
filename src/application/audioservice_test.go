package application

import (
	"audioTest/mocks"
	"audioTest/src/domain"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestAudioService_GetInfo(t *testing.T) {
	t.Run("File exists", func(t *testing.T) {
		// Response
		audio := &domain.Audio{
			SampleRate: 1,
			NumChans:   2,
			BitDepth:   10,
			Duration:   10,
		}

		// Setup Mock
		audioRepo := &mocks.IAudioRepository{}
		audioRepo.On("GetAudioInfo", mock.Anything).Return(audio, nil)

		AudioServiceFactory(audioRepo)
	})


	t.Run("File does not exist", func(t *testing.T) {

	})

}
