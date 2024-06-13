package ipc

import "fmt"

const (
	PingPath      = "/ping"
	PlayPath      = "/transport/play"
	PlayPausePath = "/transport/playpause"
	PausePath     = "/transport/pause"
	StopPath      = "/transport/stop"
	PreviousPath  = "/transport/previous"
	NextPath      = "/transport/next"
	TimePosPath   = "/transport/timepos" // ?s=<seconds>
	VolumePath    = "/volume"            // ?v=<vol>
	ShowPath      = "/window/show"
	QuitPath      = "/window/quit"
)

type Response struct {
	Error string `json:"error"`
}

func SetVolumePath(vol int) string {
	return fmt.Sprintf("%s?v=%d", VolumePath, vol)
}

func SeekToSecondsPath(secs float64) string {
	return fmt.Sprintf("%s?s=%0.2f", TimePosPath, secs)
}
