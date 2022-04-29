package flib

import "github.com/hajimehoshi/ebiten/v2"

type TickF struct {
	span int
	content func(int, []interface{}) bool
	contentCounter int
	repeat int
	Pause bool
}

type Tick struct {
	span float64
	content func(int, []interface{}) bool
	contentCounter int
	repeat int
	Pause bool
	startCall bool
	deltaTimeSum float64
}

func (t *TickF) Rewind() {
	t.contentCounter = 0
	t.Pause = true
}

func (t *TickF) Update(g *Game, i ...interface{}) {
	if !t.Pause {
		if g.Counter % t.span == 0 {
			if ebiten.MaxTPS() == 30 {
				t.content(t.contentCounter, i)
			}
			if t.content(t.contentCounter, i) || (t.contentCounter == t.repeat - 1 && t.repeat != -1) {
				t.Rewind()
			}else {
				t.contentCounter++
			}
		}
	}
}

func NewTickF(spanFrame int, pause bool, repeat int, content func(int, []interface{}) bool) *TickF {
	return &TickF{
		span: spanFrame,
		content: content,
		Pause: pause,
		repeat: repeat,
	}
}

func (t *Tick) Rewind() {
	var d float64
	if t.startCall {
		d = t.span
	}
	t.deltaTimeSum = d
	t.contentCounter = 0
	t.Pause = true
}

func (t *Tick) Update(g *Game, i ...interface{}) {
	if !t.Pause {
		t.deltaTimeSum += g.deltaTime
		if t.deltaTimeSum >= t.span {
			if t.content(t.contentCounter, i) || (t.contentCounter == t.repeat - 1 && t.repeat != -1) {
				t.Rewind()
			}else {
				t.deltaTimeSum = 0
				t.contentCounter ++
			}
			
		}
	}
}

func NewTick(spanSec float64, pause bool, repeat int, startCall bool, content func(int, []interface{}) bool) *Tick {
	var d float64
	if startCall {
		d = spanSec
	}

	return &Tick{
		span: spanSec,
		content: content,
		startCall: startCall,
		deltaTimeSum: d,
		Pause: pause, 
	}
}


