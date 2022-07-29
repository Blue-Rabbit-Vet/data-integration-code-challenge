package DTO

type AudioResponseDTO struct {
	File       string `json:"file"`
	SampleRate int    `json:"sample_rate"`
	NumChans   int    `json:"num_chans"`
	BitDepth   int    `json:"bit_depth"`
	Duration   int    `json:"duration"`
}
