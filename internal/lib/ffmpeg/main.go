package ffmpeg

import (
	"fmt"
	"video-tool/internal/lib/processes"
)

const op = "/internal/lib/ffmpeg: "

type Input struct {
	FilePath string
	Params   []string
}

type Output struct {
	FilePath string
	Params   []string
}

type Ffmpeg struct {
	Inputs     []Input
	Output     Output
	AudioCodec string
	VideoCodec string
	Process    *processes.Process
}

type Ffmpeger interface {
	Start() error
	Kill() error
	getCmd() ([]string, error)
}

func New(inputs []Input, output Output, audioCodec string, videoCodec string) (*Ffmpeg, error) {
	ffmpeg := Ffmpeg{
		Inputs:     inputs,
		Output:     output,
		AudioCodec: audioCodec,
		VideoCodec: videoCodec,
	}
	cmd, err := ffmpeg.getCmd()
	if err != nil {
		return nil, err
	}
	process := processes.New(cmd...)
	ffmpeg.Process = process
	return &ffmpeg, nil
}

func (f *Ffmpeg) getCmd() ([]string, error) {
	if f == nil || f.Inputs == nil || &f.Output == nil {
		return nil, fmt.Errorf("%s ffmpeg.getCmd(): Ffmpeg object hasn't been initialized properly", op)
	}

	var cmd []string
	cmd = append(cmd, "ffmpeg")
	for _, input := range f.Inputs {
		inputParams := input.Params
		inputParams = append(inputParams, "-i", input.FilePath)
		cmd = append(cmd, inputParams...)
	}
	if f.AudioCodec != "" {
		aCodec := f.AudioCodec
		cmd = append(cmd, "-c:a", aCodec)
	}

	if f.VideoCodec != "" {
		vCodec := f.VideoCodec
		cmd = append(cmd, "-c:v", vCodec)
	}

	outputParams := f.Output.Params
	outputParams = append(outputParams, f.Output.FilePath)
	cmd = append(cmd, outputParams...)
	fmt.Println(cmd)
	return cmd, nil
}

func (f *Ffmpeg) Start() error {
	if f.Process.Active {
		return nil
	}
	err := f.Process.Spawn()
	return err
}

func (f *Ffmpeg) Kill() error {
	if !f.Process.Active {
		return nil
	}
	err := f.Process.Kill()
	return err
}
