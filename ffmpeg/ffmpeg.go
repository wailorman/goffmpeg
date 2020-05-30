package ffmpeg

import (
	"bytes"
	"strings"

	"github.com/wailorman/goffmpeg/ctxlog"
	"github.com/wailorman/goffmpeg/utils"
)

// Configuration ...
type Configuration struct {
	FfmpegBin  string
	FfprobeBin string
}

// Configure Get and set FFmpeg and FFprobe bin paths
func Configure() (Configuration, error) {
	log := ctxlog.New(ctxlog.DefaultContext)

	var outFFmpeg bytes.Buffer
	var outFFprobe bytes.Buffer

	execFFmpegCommand := utils.GetFFmpegExec()
	execFFprobeCommand := utils.GetFFprobeExec()

	outFFmpeg, err := utils.TestCmd(execFFmpegCommand[0], execFFmpegCommand[1])
	if err != nil {
		return Configuration{}, err
	}

	outFFprobe, err = utils.TestCmd(execFFprobeCommand[0], execFFprobeCommand[1])
	if err != nil {
		return Configuration{}, err
	}

	log.WithField("path", outFFmpeg.String()).
		Debug("Found ffmpeg binary")

	log.WithField("path", outFFprobe.String()).
		Debug("Found ffprobe binary")

	ffmpeg := strings.ReplaceAll(
		outFFmpeg.String(),
		utils.LineSeparator(),
		"",
	)

	ffprobe := strings.ReplaceAll(
		outFFprobe.String(),
		utils.LineSeparator(),
		"",
	)

	cnf := Configuration{ffmpeg, ffprobe}
	return cnf, nil
}
