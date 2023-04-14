package main

import (
	"context"
	"github.com/BreezeHubs/go-pkg/schedule"
	"log"
)

func main() {
	s := schedule.NewSchedule(context.Background())
	s.Add("", func(ctx context.Context) error {
		for true {

		}
		return nil
	})

	err := s.RunAndGracefullyExit()
	log.Println(err)
}
