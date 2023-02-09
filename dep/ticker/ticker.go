package ticker

import (
	"fmt"
	"sync"
	"time"
)

var myTicker *MyTicker
var once sync.Once

type MyTicker struct {
	tickers map[string]*Ticker
	lock    sync.Mutex
}

/**
 * @description: 获取单例
 * @param {type}
 * @return:
 */
func GetInstance() *MyTicker {
	once.Do(func() {
		myTicker = &MyTicker{
			tickers: make(map[string]*Ticker),
		}
	})
	return myTicker
}

func (this *MyTicker) AddTicker(name string, t time.Duration, fn func()) *Ticker {
	if _, ok := this.tickers[name]; ok {
		return nil
	}

	ticker := newTicker(name, t, fn)

	this.lock.Lock()
	defer this.lock.Unlock()
	this.tickers[name] = ticker
	ticker.Start()

	return ticker
}

func (this *MyTicker) GetTicker(name string) *Ticker {
	if ticker, ok := this.tickers[name]; ok {
		return ticker
	}

	return nil
}

type Ticker struct {
	Name     string
	Ticker   *time.Ticker
	Duration time.Duration
	Fn       func()
	Running  bool
	Close    chan bool
}

func newTicker(name string, t time.Duration, fn func()) *Ticker {
	return &Ticker{
		Name:     name,
		Fn:       fn,
		Duration: t,
		Close:    make(chan bool, 1),
	}
}

func (this *Ticker) Start() {
	if this.Running {
		return
	}

	if this.Fn == nil {
		fmt.Println("Fn is nil ", this.Name)
		return
	}

	if this.Ticker == nil {
		this.Ticker = time.NewTicker(this.Duration)
	} else {
		this.Ticker.Reset(this.Duration)
	}

	this.Running = true
	this.Fn()
	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Printf("tick err :%s ", err)
				return
			}
		}()

		for {
			select {
			case <-this.Ticker.C:
				this.Fn()
			case <-this.Close:
				return
			}
		}
	}()
}

func (this *Ticker) Stop() {
	if this.Running {
		this.Running = false
		this.Ticker.Stop()
		this.Close <- true
	}
}
