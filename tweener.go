package flib

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)


type Tweener struct {
	begin float64
	end float64
	duration float64
	easing Easing
}

type Easing int

const (
	EaseLinear Easing = iota
	EaseInQuad
	EaseOutQuad
	EaseInOutQuad

	EaseInCubic
	EaseOutCubic
	EaseInOutCubic

	EaseInQuart
	EaseOutQuart
	EaseInOutQuart

	EaseInQuint
	EaseOutQuint
	EaseInOutQuint

	EaseInSine
	EaseOutSine
	EaseInOutSine

	EaseInExpo
	EaseOutExpo
	EaseInOutExpo

	EaseInCirc
	EaseOutCirc
	EaseInOutCirc

	EaseInBounce
	EaseOutBounce
	EaseOutInBounce
)

func NewTweener(begin float64, end float64, duration float64, easing Easing) *Tweener {
	if ebiten.MaxTPS() == 30 {
		duration /= 2
	}
	return &Tweener{
		begin: begin,
		end: end,
		duration: duration,
		easing: easing,
	}
}

func SecToFrame(sec float64) float64 {
	return sec * float64(ebiten.MaxTPS())
}

func (t *Tweener) Update(cc int) (float64, bool) {
	var value float64
	switch t.easing {
	case EaseLinear:
		value = easeLinear(float64(cc), t.begin, t.end - t.begin, t.duration)

	case EaseInQuad:
		value = easeInQuad(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseOutQuad:
		value = easeOutQuad(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseInOutQuad:
		value = easeInOutQuad(float64(cc), t.begin, t.end - t.begin, t.duration)

	case EaseInCubic:
		value = easeInCubic(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseOutCubic:
		value = easeOutCubic(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseInOutCubic:
		value = easeInOutCubic(float64(cc), t.begin, t.end - t.begin, t.duration)

	case EaseInQuart:
		value = easeInQuart(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseOutQuart:
		value = easeOutQuart(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseInOutQuart:
		value = easeInOutQuart(float64(cc), t.begin, t.end - t.begin, t.duration)

	case EaseInQuint:
		value = easeInQuint(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseOutQuint:
		value = easeOutQuint(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseInOutQuint:
		value = easeInOutQuint(float64(cc), t.begin, t.end - t.begin, t.duration)
	
	case EaseInSine:
		value = easeInSine(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseOutSine:
		value = easeOutSine(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseInOutSine:
		value = easeInOutSine(float64(cc), t.begin, t.end - t.begin, t.duration)
	
	case EaseInExpo:
		value = easeInExpo(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseOutExpo:
		value = easeOutExpo(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseInOutExpo:
		value = easeInOutExpo(float64(cc), t.begin, t.end - t.begin, t.duration)
	
	case EaseInCirc:
		value = easeInCirc(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseOutCirc:
		value = easeOutCirc(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseInOutCirc:
		value = easeInOutCirc(float64(cc), t.begin, t.end - t.begin, t.duration)

	case EaseInBounce:
		value = easeInBounce(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseOutBounce:
		value = easeOutBounce(float64(cc), t.begin, t.end - t.begin, t.duration)
	case EaseOutInBounce:
		value = easeOutInBounce(float64(cc), t.begin, t.end - t.begin, t.duration)
	
	}
	

	return value, value == t.end
}

func easeLinear(t, b, c, d float64) float64 {
	return c * t / d + b
}


func easeInQuad(t, b, c, d float64) float64 {
	t /= d
	return c*t*t + b
}

func easeOutQuad(t, b, c, d float64) float64 {
	t /= d
	return -c*t*(t-2.0) + b
}

func easeInOutQuad(t, b, c, d float64) float64 {
	t /= d/2.0
	if t < 1 {
		return c/2.0*t*t + b

	}
	t = t - 1
	return -c/2.0 * (t*(t-2) - 1) + b
}

func easeInCubic(t, b, c, d float64) float64 {
	t /= d
	return c*t*t*t + b

}

func easeOutCubic(t, b, c, d float64) float64 {
	t /= d
	t = t - 1
	return c*(t*t*t + 1) + b

}

func easeInOutCubic(t, b, c, d float64) float64 {
	t /= d/2.0
	if t < 1 {
		return c/2.0*t*t*t + b

	}
	t = t - 2
	return c/2.0 * (t*t*t + 2) + b

}

func easeInQuart(t, b, c, d float64) float64 {
	t /= d
	return c*t*t*t*t + b

}

func easeOutQuart(t, b, c, d float64) float64 {
	t /= d
	t = t - 1
	return -c*(t*t*t*t - 1) + b

}

func easeInOutQuart(t, b, c, d float64) float64 {
	t /= d/2.0
	if t < 1 {
		return c/2.0*t*t*t*t + b

	}
	t = t - 2
	return -c/2.0 * (t*t*t*t - 2) + b
}

func easeInQuint(t, b, c, d float64) float64 {
	t /= d
	return c*t*t*t*t*t + b

}

func easeOutQuint(t, b, c, d float64) float64 {
	t /= d
	t = t - 1
	return c*(t*t*t*t*t + 1) + b

}

func easeInOutQuint(t, b, c, d float64) float64 {
	t /= d/2.0
	if t < 1 {
		return c/2.0*t*t*t*t*t + b
	}
	t = t - 2
	return c/2.0 * (t*t*t*t*t + 2) + b
}

func easeInSine(t, b, c, d float64) float64 {
	return -c * math.Cos(t/d * (math.Pi/2.0)) + c + b

}

func easeOutSine(t, b, c, d float64) float64 {
	return c * math.Sin(t/d * (math.Pi/2.0)) + b

}
func easeInOutSine(t, b, c, d float64) float64 {
	return -c/2.0 * (math.Cos(math.Pi*t/d) - 1) + b

}

func easeInExpo(t, b, c, d float64) float64 {
	return c * math.Pow(2, 10*(t/d - 1)) + b
}

func easeOutExpo(t, b, c, d float64) float64 {
	return c * (-math.Pow(2.0, -10.0 * t/d) + 1) + b
}

func easeInOutExpo(t, b, c, d float64) float64 {
	t /= d/2.0
	if t < 1 {
		return c/2.0 * math.Pow(2.0, 10.0 * (t-1)) + b

	}
	t = t - 1
	return c/2.0 * (-math.Pow(2,-10*t) + 2) + b
}

func easeInCirc(t, b, c, d float64) float64 {
	t /= d
	return -c * (math.Sqrt(1 - t*t) - 1) + b

}
func easeOutCirc(t, b, c, d float64) float64 {
	t /= d
	t = t - 1
	return c * math.Sqrt(1 - t*t) + b

}

func easeInOutCirc(t, b, c, d float64) float64 {
	t /= d/2.0
	if t < 1 {
		return -c/2.0 * (math.Sqrt(1 - t*t) - 1)

	}
	t = t - 2	
	return c/2.0 * (math.Sqrt(1 - t*t) + 1) + b

}

func easeInBounce(t, b, c, d float64) float64 {
	return c - easeOutBounce(d - t, 0, c, d) + b
}

func easeOutBounce(t, b, c, d float64) float64{
	t = t / d
	if t < 1 / 2.75 {
		return c * (7.5625 * t * t) + b
		
	}else if t < 2 / 2.75 {
		t = t - (1.5 / 2.75)
		return c * (7.5625 * t * t + 0.75) + b
	}else if t < 2.5 / 2.75 {
		t = t - (2.25 / 2.75)
		return c * (7.5625 * t * t + 0.9375) + b
	}else {
		t = t - (2.625 / 2.75)
		return c * (7.5625 * t * t + 0.984375) + b
	}
}

func easeOutInBounce(t, b, c, d float64) float64 {
	if t < d / 2 {
		return easeOutBounce(t * 2, b, c / 2, d)
	}else {
		return easeInBounce((t * 2) - d, b + c / 2, c / 2, d)
	}
}


// func (t Tweener) To(tag string) *Tweener {

// }
