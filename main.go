package main

import (
	"fmt"
	"log"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang/v3/pkg/demoinfocs"
)

func main() {
	f, err := os.Open("./test/cs-demos/003606749594131628213_2016071942.dem")
	if err != nil {
		log.Panic("failed to open demo file: ", err)
	}
	defer f.Close()

	p := dem.NewParser(f)
	defer p.Close()
	
	iterateThroughFrames(&p)

	if err != nil {
		log.Panic("failed to parse demo: ", err)
	}
}

func iterateThroughFrames(p *dem.Parser) {
	var moreframes bool
	var err error

	for moreframes = true; moreframes; moreframes, err = (*p).ParseNextFrame() {
		
		if err != nil {
			log.Panic("ripo")
			break;
		}

		if (*p).CurrentFrame() < 500 {
			continue;
		}

		players := (*p).GameState().Participants().AllByUserID()

		for _, player := range players {
			if player == nil || player.Entity == nil || !player.IsAlive() {
				continue
			}

			// playerVec := { X: player.Position().X, Y: player.Position().Y, Z: player.Position().Z}

			pos := player.Position()

			fmt.Println(player.Name, pos.X, pos.Y, pos.Z)
		}


	}

	fmt.Println("Done")
}
