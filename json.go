package flib

import (
	// "encoding/json"
	// "log"
	// "image/color"
)

type textJson struct {
	name string `json:"name"`
	txt string `json:"txt`
	x float64 `json:"x"`
	y float64 `json:"y"`
	padX float64 `json:"padX"`
	padY float64 `json:"padY"`
	font string `json:"font"`
	hidden bool `json:"hidden"`
	centerX float64 `json:"centerX"`
	color colorJson `json:"color"`
}

type colorJson struct {
	r uint32 `json:"r"`
	g uint32 `json:"g"`
	b uint32 `json:"b"`
	a uint32 `json:"a"`
}

// func newTextArrayWithJson(b []byte) []*TxtSpr {
// 	// JSONデコード
// 	var textsJson []textJson
// 	if err := json.Unmarshal(b, &textsJson); err != nil {
// 		log.Fatal(err)
// 	}

// 	var texts []*TxtSpr
// 	for i, v := range textsJson {
// 		texts = append(texts, NewTxtSpr(v.txt, v.x, v.y, color.RGBA{255, 0, 0, 255}, v.font, v.padX, v.padY, v.hidden))
// 	}
// }