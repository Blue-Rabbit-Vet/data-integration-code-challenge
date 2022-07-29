package domain

type IAudioRepository interface {
	StoreFile([]byte, string) error
	GetAudioInfo(string) (*Audio, error)
}

type Audio struct {
	SampleRate int
	NumChans int
	BitDepth int
	Duration int
}



