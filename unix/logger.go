package unix

import (
	"github.com/project-flogo/core/support/log"
)

var (
    unixLogger = log.ChildLogger(log.RootLogger(), "unix")
)