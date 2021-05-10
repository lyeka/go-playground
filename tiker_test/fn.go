package tiker

import (
	"fmt"
	"time"
)

func TickDo() {
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			_ = fmt.Sprintf("hello")
		default:
			continue
		}
	}
}

func TickDoWithoutDefault() {
	ticker := time.NewTicker(200 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			_ = fmt.Sprintf("hello")
		}
	}
}

func TickDoWithSleep() {
	ticker := time.NewTicker(200 * time.Millisecond)
	flat := false
	for {
		select {
		case <-ticker.C:
			_ = fmt.Sprintf("hello")
		default:
			if flat {

			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func LongTimeTickDo() {
	go TickDo()
	time.Sleep(3 * time.Second)
}

func LongTimeTickDoWithoutDefault() {
	go TickDoWithoutDefault()
	time.Sleep(3 * time.Second)
}

func LongTimeTickDoWithSleep() {
	go TickDoWithSleep()
	time.Sleep(3 * time.Second)
}
