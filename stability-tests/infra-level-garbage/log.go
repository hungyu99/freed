package main

import (
	"github.com/hungyu99/freed/infrastructure/logger"
	"github.com/hungyu99/freed/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("IFLG")
	spawn      = panics.GoroutineWrapperFunc(log)
)
