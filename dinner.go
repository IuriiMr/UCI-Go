package main
import (
	"fmt"
	"sync"
	"time"
)

type Stick struct {sync.Mutex}

type Phil struct {
	wg *sync.WaitGroup
	leftCS, rightCS *Stick
}

func (p Phil) eat(allow chan bool, finish chan bool, n int) {

	<- allow

	p.leftCS.Lock()
	p.rightCS.Lock()

	fmt.Println("starting to eat", n)
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("finishing eating", n)

	p.rightCS.Unlock()
	p.leftCS.Unlock()

	finish <- true
	p.wg.Done()
}

func Host(allow, finish chan bool) {

	allow <- true
	allow <- true
	for x := range finish {
		allow <- x
	}
}

func main() {

	var wg sync.WaitGroup
	var allow = make(chan bool, 2)
	var finish = make(chan bool, 2)
	wg.Add(15)

	CSticks := make([]*Stick, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(Stick)
	}

	DPhilos := make([]*Phil, 5)
	for i := 0; i < 5; i++ {
		DPhilos[i] = &Phil{
			&wg,
			CSticks[i],
			CSticks[(i+1)%5]}
	}

	go Host(allow, finish)

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			go DPhilos[i].eat(allow, finish, i+1)
		}
	}
	wg.Wait()
}
