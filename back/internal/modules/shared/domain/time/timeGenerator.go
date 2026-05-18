package shared_domain_time

import "time"

type TimeGenerator interface {
	Now() time.Time
}
