package ffmpeg

import (
	"errors"
	"fmt"
	"gideo/config"
	"os/exec"
)

func CheckFFmpeg() error {
	_, err := exec.LookPath("ffmpeg")
	if err != nil {
		return errors.New("ffmpeg not found")
	}
	return nil
}

func CreateContent(conf *config.Config) error { 
	args, err := buildArgs(conf)
	if err != nil {
		return err
	}

	cmd := exec.Command("ffmpeg", args...)
	err = cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func buildArgs(conf *config.Config) ([]string, error) {
	// Here will reside the main logic of processing these flags, need to think it through
	var args []string

	args = append(args, "-f", "lavfi")
	args = append(args, "-i", fmt.Sprintf("testsrc=size=%s:rate=%d", conf.Resolution, conf.Framerate))
	args = append(args, "-c:v", conf.Codec)
	args = append(args, "-b:v", conf.Bitrate)

	outputPath, err := getOutputPath(conf.VideoType)
	if err != nil {
		return nil, err
	}
	// This creates the files on the same directory, maybe a generated folder should be created and the files deleted after use has finished
	// Maybe creating a "gideo_output" directory and then there creating directories for every video/stream
	args = append(args, outputPath)

	return args, nil
}

func getOutputPath(format string) (string, error) {
	switch format{
	case "hls":
		return "output.m3u8", nil
	case "dash":
		return "output.mpd", nil
	case "mp4":
		return "output.mp4", nil
	default:
		return "", errors.New("Unsupported format: " + format)
	}
}