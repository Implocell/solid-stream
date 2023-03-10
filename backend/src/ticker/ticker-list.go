package ticker

import (
	"math/rand"
	"time"
)

func NewTicker(symbol string, name string, initialValue float64, effect Effect, avgUpdateDelay int64) *Ticker {
	ts := time.Now().UnixMilli()

	return &Ticker{
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

func CreateAllTickers() []*Ticker {
	tickers := []*Ticker{
		NewTicker("CHR", "Chronos", float64(400+rand.Int31n(799)), NORMAL, 150),
		NewTicker("COG", "Cogsworth", float64(200+rand.Int31n(399)), NORMAL, 2023),
		NewTicker("DOC", "Doc Brown", 1985, BACKTOTHEFUTURE, 1000),
		NewTicker("SNI", "Snickers", 613, NORMAL, 330),
		NewTicker("BIC", "B I C", 27, NORMAL, 2110),
		NewTicker("BOI", "Big (the cat) On Ice", 42, NORMAL, 4300),
		NewTicker("UNO", "Uno", 8, NORMAL, 800),
		NewTicker("BRD", "Birdbird exploration", 1284, NORMAL, 240),
		NewTicker("ANO", "Ano", 8, NORMAL, 800),
		NewTicker("BNO", "BNO", 28, NORMAL, 800),
		NewTicker("CNO", "CNO", 83, NORMAL, 800),
		NewTicker("DNO", "DNO", 87, NORMAL, 800),
		NewTicker("ENO", "ENO", 85, NORMAL, 800),
		NewTicker("FNO", "FNO", 81, NORMAL, 800),
		NewTicker("GNO", "GNO", 18, NORMAL, 800),
		NewTicker("HNO", "HNO", 68, NORMAL, 800),
		NewTicker("INO", "INO", 8, NORMAL, 800),
		NewTicker("JNO", "JNO", 8, NORMAL, 800),
		NewTicker("KNO", "KNO", 53, NORMAL, 800),
		NewTicker("LNO", "LNO", 23, NORMAL, 800),
		NewTicker("MNO", "MNO", 8, NORMAL, 800),
		NewTicker("NNO", "NNO", 432, NORMAL, 800),
		NewTicker("ONO", "ONO", 864, NORMAL, 800),
		NewTicker("PNO", "PNO", 8, NORMAL, 800),
		NewTicker("QNO", "QNO", 342, NORMAL, 800),
		NewTicker("RNO", "RNO", 8, NORMAL, 800),
		NewTicker("SNO", "SNO", 122, NORMAL, 800),
		NewTicker("TNO", "TNO", 911, NORMAL, 800),
		NewTicker("VNO", "VNO", 8, NORMAL, 800),
		NewTicker("WNO", "WNO", 75, NORMAL, 800),
		NewTicker("XNO", "XNO", 8, NORMAL, 800),
		NewTicker("YNO", "YNO", 41, NORMAL, 800),
		NewTicker("ZNO", "ZNO", 99, NORMAL, 800),
	}

	return tickers
}
