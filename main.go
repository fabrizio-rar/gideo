package main

import (
	"fmt"
	"gideo/config"
	"gideo/src/ffmpeg"
	"gideo/src/flags"
	"gideo/src/server"
	"os"
)

func main() {
	flagOpts, err := flags.ParseFlags()
	if err != nil {
		fmt.Println("Error parsing flags:", err)
		os.Exit(1)
	}

	// Initialize configuration from flags
	conf := config.NewConfig(flagOpts)

	// Check ffmpeg
	err = ffmpeg.CheckFFmpeg()
	if err != nil {
		fmt.Println("FFmpeg not found in PATH", err)
		os.Exit(1)
	}

	// Create content
	err = ffmpeg.CreateContent(conf)
	if err != nil {
		fmt.Println("Error creating content:", err)
		os.Exit(1)
	}

	// Serve the generated content
	// TODO: This is wrong, it should have the output directory as the generated one
	err = server.ServeContent(conf.Port, "output")
	if err != nil {
		fmt.Println("Error serving content:", err)
		os.Exit(1)
	}
}