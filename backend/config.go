package backend

import (
	"os"

	"github.com/google/uuid"
	"github.com/pelletier/go-toml/v2"
)

type ServerConnection struct {
	Hostname    string
	AltHostname string
	Username    string
	LegacyAuth  bool
}

type ServerConfig struct {
	ServerConnection
	ID       uuid.UUID
	Nickname string
	Default  bool
}

type AppConfig struct {
	WindowWidth        int
	WindowHeight       int
	LastCheckedVersion string
	EnableSystemTray   bool
	CloseToSystemTray  bool
	StartupPage        string

	// Experimental - may be removed in future
	FontNormalTTF string
	FontBoldTTF   string
}

type AlbumPageConfig struct {
	TracklistColumns []string
}

type AlbumsPageConfig struct {
	SortOrder string
}

type ArtistPageConfig struct {
	InitialView      string
	TracklistColumns []string
}

type FavoritesPageConfig struct {
	InitialView      string
	TracklistColumns []string
}

type NowPlayingPageConfig struct {
	TracklistColumns []string
}

type PlaylistPageConfig struct {
	TracklistColumns []string
}

type PlaylistsPageConfig struct {
	InitialView string
}

type TracksPageConfig struct {
	TracklistColumns []string
}

type LocalPlaybackConfig struct {
	AudioDeviceName     string
	AudioExclusive      bool
	InMemoryCacheSizeMB int
	Volume              int
}

type ScrobbleConfig struct {
	Enabled              bool
	ThresholdTimeSeconds int
	ThresholdPercent     int
}

type ReplayGainConfig struct {
	Mode            string
	PreampGainDB    float64
	PreventClipping bool
}

type ThemeConfig struct {
	Appearance string
}

type Config struct {
	Application    AppConfig
	Servers        []*ServerConfig
	AlbumPage      AlbumPageConfig
	AlbumsPage     AlbumsPageConfig
	ArtistPage     ArtistPageConfig
	FavoritesPage  FavoritesPageConfig
	NowPlayingPage NowPlayingPageConfig
	PlaylistPage   PlaylistPageConfig
	PlaylistsPage  PlaylistsPageConfig
	TracksPage     TracksPageConfig
	LocalPlayback  LocalPlaybackConfig
	Scrobbling     ScrobbleConfig
	ReplayGain     ReplayGainConfig
	Theme          ThemeConfig
}

var SupportedStartupPages = []string{"Albums", "Favorites", "Playlists"}

func DefaultConfig(appVersionTag string) *Config {
	return &Config{
		Application: AppConfig{
			WindowWidth:        1000,
			WindowHeight:       800,
			LastCheckedVersion: appVersionTag,
			EnableSystemTray:   true,
			CloseToSystemTray:  false,
			StartupPage:        "Albums",
		},
		AlbumPage: AlbumPageConfig{
			TracklistColumns: []string{"Artist", "Time", "Plays", "Favorite", "Rating"},
		},
		AlbumsPage: AlbumsPageConfig{
			SortOrder: string("Recently Added"),
		},
		ArtistPage: ArtistPageConfig{
			InitialView:      "Discography",
			TracklistColumns: []string{"Album", "Time", "Plays", "Favorite", "Rating"},
		},
		FavoritesPage: FavoritesPageConfig{
			TracklistColumns: []string{"Artist", "Album", "Time", "Plays"},
			InitialView:      "Albums",
		},
		NowPlayingPage: NowPlayingPageConfig{
			TracklistColumns: []string{"Artist", "Album", "Time", "Plays"},
		},
		PlaylistPage: PlaylistPageConfig{
			TracklistColumns: []string{"Artist", "Album", "Time", "Plays"},
		},
		PlaylistsPage: PlaylistsPageConfig{
			InitialView: "List",
		},
		TracksPage: TracksPageConfig{
			TracklistColumns: []string{"Artist", "Album", "Time", "Plays"},
		},
		LocalPlayback: LocalPlaybackConfig{
			// "auto" is the name to pass to MPV for autoselecting the output device
			AudioDeviceName:     "auto",
			AudioExclusive:      false,
			InMemoryCacheSizeMB: 30,
			Volume:              100,
		},
		Scrobbling: ScrobbleConfig{
			Enabled:              true,
			ThresholdTimeSeconds: 240,
			ThresholdPercent:     50,
		},
		ReplayGain: ReplayGainConfig{
			Mode:            ReplayGainNone,
			PreampGainDB:    0.0,
			PreventClipping: true,
		},
		Theme: ThemeConfig{
			Appearance: "Dark",
		},
	}
}

func ReadConfigFile(filepath, appVersionTag string) (*Config, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c := DefaultConfig(appVersionTag)
	if err := toml.NewDecoder(f).Decode(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Config) WriteConfigFile(filepath string) error {
	b, err := toml.Marshal(c)
	if err != nil {
		return err
	}
	os.WriteFile(filepath, b, 0644)

	return nil
}
