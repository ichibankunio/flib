package flib

import (
	"bytes"
	"embed"
	"image"
	"image/jpeg"
	"image/png"
	"log"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"golang.org/x/image/font"
)

//go:embed shaders
var shadersDir embed.FS
var shaders []*ebiten.Shader

const (
	SHADER_CIRCLE = iota
	SHADER_FADEINOUT
)

func init() {
	shadersFS, err := shadersDir.ReadDir("shaders")

	if err != nil {
		log.Fatal(err)
	}

	for _, fs := range shadersFS {
		if !fs.IsDir() {
			file, err := shadersDir.ReadFile("shaders/" + fs.Name())

			if err != nil {
				log.Fatal(err)
			}

			s, err := ebiten.NewShader(file)
			if err != nil {
				panic(err)
			}

			shaders = append(shaders, s)
		}
	}
}

func NewImageFromBytes(byteData []byte) *ebiten.Image {
	r := bytes.NewReader(byteData)

	_, format, err := image.DecodeConfig(r)
	if err != nil {
		panic(err)
	}

	r = bytes.NewReader(byteData)

	var img image.Image
	if format == "png" {
		img, err = png.Decode(r)
		if err != nil {
			panic(err)
		}
	} else if format == "jpeg" {
		img, err = jpeg.Decode(r)
		if err != nil {
			panic(err)
		}
	}
	return ebiten.NewImageFromImage(img)
}

func NewFontFromBytes(byteData []byte, size int) font.Face {
	tt, err := truetype.Parse(byteData)
	if err != nil {
		log.Fatal(err)
	}

	return truetype.NewFace(tt, &truetype.Options{
		Size:    float64(size),
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func NewBGMFromBytes(b []byte, sampleRate int, context *audio.Context) *audio.Player {
	m, _ := mp3.DecodeWithSampleRate(sampleRate, bytes.NewReader(b))
	s := audio.NewInfiniteLoop(m, m.Length())
	p, _ := context.NewPlayer(s)
	// p, _ := audio.NewPlayer(audioContext, s)

	return p
}
