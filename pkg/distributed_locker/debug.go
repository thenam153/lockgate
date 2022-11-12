package distributed_locker

import "github.com/thenam153/lockgate/pkg/util"

func debug(format string, args ...interface{}) {
	util.Debug(format, args...)
}
