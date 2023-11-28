package rpchandlers

import (
	"github.com/hungyu99/freed/infrastructure/logger"
	"github.com/hungyu99/freed/util/panics"
)

var log = logger.RegisterSubSystem("RPCS")
var spawn = panics.GoroutineWrapperFunc(log)
