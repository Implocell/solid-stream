package ticker

import (
	"math/rand"
	"time"
)

func NewTicker(symbol string, name string, initialValue float64, effect Effect, avgUpdateDelay int64) Ticker {
	ts := time.Now().UnixMilli()

	return Ticker{
		Symbol:         symbol,
		Name:           name,
		Value:          initialValue,
		InitialValue:   initialValue,
		Effect:         effect,
		Updated:        ts,
		NextUpdate:     ts + avgUpdateDelay,
		AvgUpdateDelay: avgUpdateDelay,
	}
}

func CreateAllTickers() []Ticker {
	tickers := []Ticker{
		NewTicker("CHR", "Chronos", float64(400+rand.Int31n(799)), NORMAL, 100),
		NewTicker("COG", "Cogsworth", float64(200+rand.Int31n(399)), NORMAL, 2023),
		NewTicker("DOC", "Doc Brown", 1985, BACKTOTHEFUTURE, 1000),
	}

	return tickers
}
