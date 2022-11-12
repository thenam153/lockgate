package main

import (
	"fmt"
	"os"
	"time"

	"github.com/thenam153/lockgate/pkg/file_locker"

	"github.com/thenam153/lockgate"
)

func do() error {
	// if err := kube.Init(kube.InitOptions{}); err != nil {
	// 	return fmt.Errorf("cannot initialize kube: %s", err)
	// }

	if locker, err := file_locker.NewFileLocker("/tmp/locks"); err != nil {
		return fmt.Errorf("init error: %s", err)
	} else {
		if _, lock, err := locker.Acquire("mylock", lockgate.AcquireOptions{
			OnWaitFunc: func(_ string, doWait func() error) error {
				fmt.Printf("WAITING!\n")
				defer fmt.Printf("DONE!")
				return doWait()
			},
		}); err != nil {
			return fmt.Errorf("acquire mylock error: %s", err)
		} else {
			defer locker.Release(lock)
		}
	}

	fmt.Printf("ACQUIRED!\n")
	time.Sleep(10 * time.Second)

	return nil
}

func main() {
	if err := do(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("GOOD.\n")
}
