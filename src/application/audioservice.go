package application

import (
	"audioTest/src/domain"
	"audioTest/src/infrastructure/DTO"
)

type IAudioService interface {
	StoreFile([]byte, string) error
	GetInfo(string) (*DTO.AudioResponseDTO, error)
	GetInfoList(int) (*[]DTO.AudioResponseDTO, error)
}

type AudioService struct {
	audioRepository domain.IAudioRepository
}



func (service *AudioService) GetInfo(fileName string) (*DTO.AudioResponseDTO, error) {
	metadata, err := service.audioRepository.GetAudioInfo(fileName)
	if err != nil {
		return nil, err
	}

	return convertToDTO(metadata, fileName), nil
}

func (service *AudioService) StoreFile(bytes []byte, fileName string) error {
	err := service.audioRepository.StoreFile(bytes, fileName)

	return err
}

func AudioServiceFactory(audioRepository domain.IAudioRepository) *AudioService {
	return &AudioService{
		audioRepository: audioRepository,
	}
}

func (service *AudioService) GetInfoList(duration int) (*[]DTO.AudioResponseDTO, error) {
	panic("implement me")
}

func convertToDTO(audio *domain.Audio, filename string) *DTO.AudioResponseDTO {
	return &DTO.AudioResponseDTO{
		File:       filename,
		SampleRate: audio.SampleRate,
		NumChans:   audio.NumChans,
		BitDepth:   audio.BitDepth,
		Duration:   audio.Duration,
	}
}
