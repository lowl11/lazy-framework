package models

type Http2Config struct {
	MaxConcurrentStreams int
	MaxReadFrameSize     int
}
