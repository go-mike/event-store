package abstractions

import "time"

type Clock interface {
	Now() time.Time
}
