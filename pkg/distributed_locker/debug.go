package distributed_locker

import "lockgate/pkg/util"

func debug(format string, args ...interface{}) {
	util.Debug(format, args...)
}
