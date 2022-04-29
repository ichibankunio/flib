// Copyright 2020 The Ebiten Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build ignore
// +build ignore

package main

var Time float
var Cursor vec2
var ScreenSize vec2

const pi = 3.1415926535

func Fragment(position vec4, texCoord vec2, color vec4) vec4 {
	st := (position.xy * 2.0 - ScreenSize) / min(ScreenSize.x, ScreenSize.y)
	
    // col := vec4(0)
    // gl_FragColor = vec4(ceil(1.0-length(position)))

	// colored := vec3(step(0,(-4*Time*Time+1)-length(st)))

	// alpha := 1- step(0,(-4*Time*Time+1)-length(st))

	// radius := step(0, cos(Time/2 * 2 * pi)) * cos(Time/2 * 2 * pi)
	// radius := step(0, cos(Time/2 * 2 * pi))4/3*(Time-1)*(Time-1)-1/3
	// alpha := 1 - step(0,radius-length(st))

	alpha := 1 - step(0, cos(Time/1.0 * 2 * pi)-length(st))

	// colored := vec3(circle(st,4))
	// colored := vec3(smoothstep(0.5,0.51, dot(st,st)))
	// inverted := vec3(1, 1, 1) - colored
	// return vec4(ceil(1.0-length(st)))
	// return vec4(colored, 1)
	return vec4(0, 0, 0, alpha)
}

func circle(st vec2, radius float) float {
    dist := st
	return 1-smoothstep(radius-(radius*0.01), radius+(radius*0.01), dot(dist, dist) * 4.0)
}