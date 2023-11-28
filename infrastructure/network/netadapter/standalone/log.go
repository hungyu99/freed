package standalone

import (
	"github.com/hungyu99/freed/infrastructure/logger"
	"github.com/hungyu99/freed/util/panics"
)

var log = logger.RegisterSubSystem("NTAR")
var spawn = panics.GoroutineWrapperFunc(log)
