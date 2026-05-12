package schedule_unit

import (
	utils_time "language-learning/utils/time"
	"time"

	"github.com/google/uuid"
)

var (
	studentUuid      = uuid.MustParse("3d30a9c8-5768-4d73-ab5d-8d36cab7f03f")
	professorUuid    = uuid.MustParse("22ce4a72-2978-4f19-9181-e59963add95b")
	otherStudentUuid = uuid.MustParse("d8faff5b-8692-4e76-a240-ae77f66db979")
)

var (
	studentDate1  = time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC)
	studentDate2  = time.Date(2026, time.May, 9, 10, 30, 0, 0, time.UTC)
	professorDate = time.Date(2026, time.May, 9, 11, 0, 0, 0, time.UTC)
	duration1h    = utils_time.MustParseDuration("1h")
)
