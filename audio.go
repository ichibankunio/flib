package flib

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
)

func PlaySE(player *audio.Player) {
	player.SetVolume(SEVolume)
	player.Play()
}

func PlayBGM(player *audio.Player) {
	player.SetVolume(BGMVolume)
	player.Play()
}