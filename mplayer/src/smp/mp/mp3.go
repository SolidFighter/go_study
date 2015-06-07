package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat int
	progress int
}

func (p *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 music", source)
	p.progress = 0

	for p.progress < 10000 {
		time.Sleep(100 * time.Microsecond)
		fmt.Print(".")
		p.progress += 10
	}

	fmt.Println("\nFinished playing", source)
}

