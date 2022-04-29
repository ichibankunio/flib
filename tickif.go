package flib

// import "github.com/hajimehoshi/ebiten/v2"


// type TickF struct {
// 	span int
// 	content func(int, ...interface{}) bool
// 	contentCounter int
// 	repeat int
// }

// func NewTickF(span int, content func(int, ...interface{}) bool) *TickF {
// 	return &TickF{
// 		span: span,
// 		content: content,
// 	}
// }

// func (t *TickF) Update(g *Game, i interface{}) {
// 	if g.counter % t.span == 0 {
// 		if ebiten.MaxTPS() == 30 {
// 			t.content(t.contentCounter, i)
// 		}
// 		if t.content(t.contentCounter, i) {
// 			// t.rewind()
// 		}else {
// 			t.contentCounter++
// 		}
// 	}
	
// }