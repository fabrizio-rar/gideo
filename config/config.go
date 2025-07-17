package config

import "gideo/src/flags"

type Config struct {
	VideoType  string
	Port       int
	Resolution string
	Bitrate    string
	Framerate  int
	Codec      string
}

func NewConfig(opts *flags.Options) *Config {
	return &Config{
		VideoType:  opts.VideoType,
		Port:       opts.Port,
		Resolution: opts.Resolution,
		Bitrate:    opts.Bitrate,
		Framerate:  opts.Framerate,
		Codec:      opts.Codec,
	}
}