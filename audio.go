package flib

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
)

var volumeBGM float64 = 1.0
var volumeSE float64 = 1.0

func GetBGMVolume() float64 {
	return volumeBGM
}

func GetSEVolume() float64 {
	return volumeSE
}

func SetBGMVolume(v float64) {
	volumeBGM = v
}

func SetSEVolume(v float64) {
	volumeSE = v
}

func PlaySE(player *audio.Player) {
	player.SetVolume(volumeSE)
	player.Play()
}

func PlayBGM(player *audio.Player) {
	player.SetVolume(volumeBGM)
	player.Play()
}