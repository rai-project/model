package model

import (
	"github.com/rai-project/config"
	"github.com/rai-project/logger"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

var (
	log      *logrus.Entry
	validate = validator.New()
)

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "model")
	})
}
