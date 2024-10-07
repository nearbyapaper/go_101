package song

import "fmt"

type Song struct {
	name   string
	artist string
}

// Constructor for Song
func New(songName string, artistName string) (*Song, error) {
	// Validate inputs
	if songName == "" || artistName == "" {
		return nil, fmt.Errorf("Song name and artist name cannot be empty")
	}

	// Create a new Song
	newSong := Song{
		name:   songName,
		artist: artistName,
	}

	return &newSong, nil
}

// Method to play the song
func (s Song) Play() {
	fmt.Printf("Playing %s by %s\n", s.name, s.artist)
}

// Method to pause the song
func (s Song) Pause() {
	fmt.Printf("Pausing %s by %s\n", s.name, s.artist)
}

// Method to stop the song (fixed typo)
func (s Song) Stop() {
	fmt.Printf("Stopping %s by %s\n", s.name, s.artist)
}
