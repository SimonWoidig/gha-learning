package main

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Str("data", uuid.New().String()).Msg("")
}
