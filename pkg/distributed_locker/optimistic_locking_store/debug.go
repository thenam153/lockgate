package optimistic_locking_store

import "lockgate/pkg/util"

func debug(format string, args ...interface{}) {
	util.Debug(format, args...)
}
