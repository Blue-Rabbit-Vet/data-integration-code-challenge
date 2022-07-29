package routes

import (
	"audioTest/mocks"
	"audioTest/src/infrastructure/DTO"
	"encoding/json"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAudioHandler_GetInfo(t *testing.T) {
	t.Run("empty name query parameter return 400 bad request", func(t *testing.T) {
		// make request
		req, err := http.NewRequest("GET", "/info", nil)
		if err != nil {
			t.Error(err)
		}
		rr := httptest.NewRecorder()

		// Setup Mock
		audioService := &mocks.IAudioService{}
		audioService.On("GetInfo", mock.Anything).Return(nil, nil)

		//Setup router
		router := Routes(
			audioService,
		)

		router.ServeHTTP(rr, req)

		if rr.Code != 400 {
			t.Fatalf("Error status code was expected to be 400 but got %d", rr.Code)
		}
	})

	t.Run("File exists return metadata", func(t *testing.T) {
		g := NewGomegaWithT(t)
		// make request
		req, err := http.NewRequest("GET", "/info?name=mohamed", nil)
		if err != nil {
			t.Error(err)
		}
		rr := httptest.NewRecorder()

		// Response
		mockResponse := &DTO.AudioResponseDTO{
			File:       "myfile.wav",
			SampleRate: 2,
			NumChans:   2,
			BitDepth:   10,
			Duration:   10,
		}
		// Setup Mock
		audioService := &mocks.IAudioService{}
		audioService.On("GetInfo", mock.Anything).Return(mockResponse, nil)

		//Setup router
		router := Routes(
			audioService,
		)

		router.ServeHTTP(rr, req)

		var actualResponse DTO.AudioResponseDTO
		if err := json.NewDecoder(rr.Body).Decode(&actualResponse); err != nil {
			g.Expect(err).ToNot(HaveOccurred())
		}

		g.Expect(&actualResponse).To(Equal(mockResponse))

		if rr.Code != 200 {
			t.Fatalf("Error status code was expected to be 200 but got %d", rr.Code)
		}
	})

	t.Run("File does not exist return 404", func(t *testing.T) {
		// make request
		req, err := http.NewRequest("GET", "/info?name=mohamed", nil)
		if err != nil {
			t.Error(err)
		}
		rr := httptest.NewRecorder()

		// Setup Mock
		audioService := &mocks.IAudioService{}
		audioService.On("GetInfo", mock.Anything).Return(nil, nil)

		//Setup router
		router := Routes(
			audioService,
		)

		router.ServeHTTP(rr, req)

		if rr.Code != 404 {
			t.Fatalf("Error status code was expected to be 404 but got %d", rr.Code)
		}
	})
}
