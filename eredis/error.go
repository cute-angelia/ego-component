package eredis

type Err string

func (e Err) Error() string { return string(e) }

const (
	// ErrInvalidParams  is returned when parameters is invalid.
	ErrInvalidParams = Err("invalid params")

	// ErrNotObtained is returned when a Lock cannot be obtained.
	ErrNotObtained = Err("redislock: not obtained")

	// ErrLockNotHeld is returned when trying to release an inactive Lock.
	ErrLockNotHeld = Err("redislock: lock not held")
)
