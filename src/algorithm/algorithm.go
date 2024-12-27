package algorithm

import (
	"context"
	"sync"

	m "github.com/LeonDavidZipp/Pathfinder/src/models"
)

// handles solving & everything related to it
func SolveWrapper(ctx context.Context, mp *m.Map) (*m.Solution, error) {
	bot := m.NewBot(mp.Start)
	wg := &sync.WaitGroup{}
	sol := make(chan *m.Solution, 42)

	// start solving
	wg.Add(1)
	go Solve(*bot, wg, sol)

	// wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(sol)
	}()

	// while waiting, find shortest of incoming solutions
	var res *m.Solution = nil
	for {
		select {
		case s, ok := <-sol:
			if !ok {
				return res, nil
			} else if res == nil || s.Steps < res.Steps {
				res = s
			}
		case <-ctx.Done():
			return res, nil
		}
	}
}

// recursive solving function
func Solve(bot m.Bot, wg *sync.WaitGroup, sol chan<- *m.Solution) {
	defer wg.Done()
	for {
		dirs := bot.CountPaths()
		switch len(dirs) {
		case 0:
			return
		case 1:
			bot.Move(dirs[0])
			if bot.Pos.Type == m.End {
				sol <- &m.Solution{
					Steps: bot.Steps,
					Route: bot.Route,
				}
				return
			}
		default:
			for _, d := range dirs[1:] {
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
				go Solve(b, wg, sol)
			}
			bot.Move(dirs[0])
			if bot.Pos.Type == m.End {
				sol <- &m.Solution{
					Steps: bot.Steps,
					Route: bot.Route,
				}
				return
			}
		}
	}
}
