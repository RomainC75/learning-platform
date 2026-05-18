package schedule_unit

import (
	schedule_domain "language-learning/internal/modules/schedule/domain"
	utils_time "language-learning/utils/time"
	"time"

	"github.com/google/uuid"
)

var (
	professorUuid = uuid.MustParse("22ce4a72-2978-4f19-9181-e59963add95b")
	studentUuid   = uuid.MustParse("3d30a9c8-5768-4d73-ab5d-8d36cab7f03f")
	studentUuid2  = uuid.MustParse("d8faff5b-8692-4e76-a240-ae77f66db979")
	bookingUuid   = uuid.MustParse("76d0f1d9-a7fe-4783-bd81-b51fc409aab9")
	professor     = schedule_domain.NewProfessor(professorUuid, "John", "Doe", schedule_domain.Schedule{})
	student       = schedule_domain.NewStudent(studentUuid, "Jane", "Smith")
)

var (
	studentDate1  = time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC)
	studentDate2  = time.Date(2026, time.May, 9, 10, 30, 0, 0, time.UTC)
	professorDate = time.Date(2026, time.May, 9, 11, 0, 0, 0, time.UTC)
	july25        = time.Date(2026, time.July, 25, 12, 0, 0, 0, time.UTC)
	nowjuly1      = time.Date(2026, time.July, 1, 12, 0, 0, 0, time.UTC)
	nowjuly2      = time.Date(2026, time.July, 2, 12, 0, 0, 0, time.UTC)
	duration1h    = utils_time.MustParseDuration("1h")
	duration4h    = utils_time.MustParseDuration("4h")
	duration1Week = utils_time.MustParseDuration("168h")
)
