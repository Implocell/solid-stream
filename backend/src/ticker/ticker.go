package ticker

import (
	"math"
	"math/rand"
	"time"
)

type Effect string

const (
	NORMAL          Effect = "NORMAL"
	REVERSE         Effect = "REVERSE"
	SLOW            Effect = "SLOW"
	BACKTOTHEFUTURE Effect = "BACKTOTHEFUTURE"
	// https://imageio.forbes.com/blogs-images/jvchamary/files/2019/04/back_future_timelines-1200x900.jpg?format=jpg&width=960
	// or https://static.wikia.nocookie.net/bttf/images/e/ea/Bttf.png/revision/latest?cb=20200108043202
)

type Ticker struct {
	Symbol         string  `json:"symbol"`
	Name           string  `json:"name"`
	Value          float64 `json:"value"`
	InitialValue   float64 `json:"-"`
	Updated        int64   `json:"updated"`
	Effect         Effect  `json:"-"`
	NextUpdate     int64   `json:"-"`
	AvgUpdateDelay int64   `json:"-"`
}

const ONE_WEEK_IN_MILLI = 7 * 24 * 60 * 60 * 1000

var GAME_START int64 = time.Now().UnixMilli()
var GAME_END int64 = time.Now().UnixMilli() + ONE_WEEK_IN_MILLI
var GAME_DURATION int64 = GAME_END - GAME_START

/*
@TODO figure out how to generate a determistic result over time,
despite not being able to reliably run at the same timestamp
across multiple games with the same seed.

Just using sine is a bit boring... SAD!
*/
func (t *Ticker) GenerateUpdate(timestamp int64) {
	timeSinceStart := float64(timestamp - GAME_START)

	// gameProgressRatio := float64(timestamp-GAME_START) / float64(GAME_DURATION)

	t.Value = (1 + math.Sin(timeSinceStart)/100) * t.InitialValue
	t.Updated = timestamp
	t.NextUpdate = timestamp + int64((rand.Float64()+0.5)*float64(t.AvgUpdateDelay))
}
