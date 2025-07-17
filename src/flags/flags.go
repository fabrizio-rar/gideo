package flags

import "flag"

type Options struct {
	VideoType  string
	Port       int
	Resolution string
	Bitrate    string
	Framerate  int
	Codec      string
}

func ParseFlags() (*Options, error) {
	videoType := flag.String("type", "hls", "Video/stream type")
	port := flag.Int("port", 8080, "Port to serve on")
	resolution := flag.String("resolution", "1920x1080", "Video resolution")
	bitrate := flag.String("bitrate", "1000k", "Video bitrate")
	framerate := flag.Int("framerate", 30, "Video framerate")
	codec := flag.String("codec", "libx264", "Video codec")
	flag.Parse()

	return &Options{
		VideoType:  *videoType,
		Port:       *port,
		Resolution: *resolution,
		Bitrate:    *bitrate,
		Framerate:  *framerate,
		Codec:      *codec,
	}, nil
}