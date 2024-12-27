package algorithm

import (
	"fmt"
	"sync"

	m "github.com/LeonDavidZipp/Pathfinder/src/models"
)

// handles solving & everything related to it
func SolveWrapper(mp *m.Map) (*m.Solution, error) {
	bot := m.NewBot(mp.Start)
	wg := &sync.WaitGroup{}
	sol := make(chan *m.Solution, 42)
	err := make(chan error, 42)

	wg.Add(1)
	go Solve(*bot, wg, sol, err)
	wg.Wait()

	close(sol)
	close(err)

	switch len(sol) {
	case 0:
		return nil, fmt.Errorf("no solution found")
	default:
		shortest := <-sol
		for s := range sol {
			if s.Steps < shortest.Steps {
				shortest = s
			}
		}
		return shortest, nil
	}
}

// recursive solving function
func Solve(bot m.Bot, wg *sync.WaitGroup, sol chan<- *m.Solution, err chan<- error) {
	defer wg.Done()
	for {
		cnt, dirs := bot.CountPaths()
		switch cnt {
		case 0:
			err <- fmt.Errorf("no path found")
			return
		case 1:
			bot.Move(dirs[0])
			if bot.Pos.Type == m.End {
				sol <- &m.Solution{
					Route: bot.Route,
					Steps: bot.Steps,
				}
				return
			}
		default:
			for _, d := range dirs {
				b := m.CopyBot(&bot)
				b.Move(d)
				if b.Pos.Type == m.End {
					sol <- &m.Solution{
						Route: bot.Route,
						Steps: bot.Steps,
					}
					return
				}
				wg.Add(1)
				go Solve(b, wg, sol, err)
			}
		}
	}
}
