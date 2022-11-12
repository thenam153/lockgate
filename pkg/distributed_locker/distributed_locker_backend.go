package distributed_locker

import (
	"errors"
	"time"

	"lockgate"

	"github.com/google/uuid"
)

const (
	DistributedLockLeaseTTLSeconds                 = 10
	DistributedLockPollRetryPeriodSeconds          = 2
	DistributedOptimisticLockingRetryPeriodSeconds = 1
	DistributedLockLeaseRenewPeriodSeconds         = 3
)

var (
	ErrShouldWait               = errors.New("should wait")
	ErrLockAlreadyLeased        = errors.New("lock already leased")
	ErrNoExistingLockLeaseFound = errors.New("no existing lock lease found")
)

func IsErrShouldWait(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == ErrShouldWait.Error()
}

func IsErrLockAlreadyLeased(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == ErrLockAlreadyLeased.Error()
}

func IsErrNoExistingLockLeaseFound(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == ErrNoExistingLockLeaseFound.Error()
}

type DistributedLockerBackend interface {
	Acquire(lockName string, opts AcquireOptions) (lockgate.LockHandle, error)
	RenewLease(handle lockgate.LockHandle) error
	Release(handle lockgate.LockHandle) error
}

type AcquireOptions struct {
	Shared bool `json:"shared"`
}

type LockLeaseRecord struct {
	lockgate.LockHandle
	ExpireAtTimestamp  int64
	SharedHoldersCount int64
	IsShared           bool
}

func NewLockLeaseRecord(lockName string, isShared bool) *LockLeaseRecord {
	return &LockLeaseRecord{
		LockHandle:         lockgate.LockHandle{UUID: uuid.New().String(), LockName: lockName},
		ExpireAtTimestamp:  time.Now().Unix() + DistributedLockLeaseTTLSeconds,
		SharedHoldersCount: 1,
		IsShared:           isShared,
	}
}
