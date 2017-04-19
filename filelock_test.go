package filelock

import (
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestLockAndUnlock(t *testing.T) {
	f, err := ioutil.TempFile("", "lock")
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
	defer func() {
		err = os.Remove(f.Name())
		if err != nil {
			t.Fatal(err)
		}
	}()

	// lock the file
	l, err := New(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer l.Destroy()
	err = l.Lock()
	if err != nil {
		t.Fatal(err)
	}

	// try lock a locked file
	dupl, err := New(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	got, err := dupl.TryLock()
	if got {
		t.Errorf("unexpected acquired lock")
	}
	if err != ErrLocked {
		t.Errorf("err = %v, want %v", err, ErrLocked)
	}

	// unlock the file
	err = l.Unlock()
	if err != nil {
		t.Fatal(err)
	}

	// try lock the unlocked file
	got, err = dupl.TryLock()
	if !got {
		t.Errorf("lock should have been acquired")
	}
	if err != nil {
		t.Errorf("err = %v, want %v", err, nil)
	}
	defer dupl.Destroy()

	// blocking on locked file
	locked := make(chan struct{}, 1)
	go func() {
		l.Lock()
		locked <- struct{}{}
	}()

	select {
	case <-locked:
		t.Error("unexpected unblocking")
	case <-time.After(100 * time.Millisecond):
	}

	// unlock
	err = dupl.Unlock()
	if err != nil {
		t.Fatal(err)
	}

	// the previously blocked routine should be unblocked
	select {
	case <-locked:
	case <-time.After(1 * time.Second):
		t.Error("unexpected blocking")
	}
}
