package schedule_unit

import (
	schedule_domain "language-learning/internal/modules/schedule/domain"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	utils_time "language-learning/utils/time"
	"time"

	"github.com/google/uuid"
)

var (
	professorUuid = uuid.MustParse("22ce4a72-2978-4f19-9181-e59963add95b")
	studentUuid   = uuid.MustParse("3d30a9c8-5768-4d73-ab5d-8d36cab7f03f")
	studentUuid2  = uuid.MustParse("d8faff5b-8692-4e76-a240-ae77f66db979")
	bookingUuid   = uuid.MustParse("76d0f1d9-a7fe-4783-bd81-b51fc409aab9")
	booked1Uuid   = uuid.MustParse("c0a6543f-41f9-4e8d-8502-24571b1b477c")
	booked2Uuid   = uuid.MustParse("0aef3256-8475-4b91-b675-ce072516c701")
	// wa
	// : monday 10-14
	// : thursday 14-18
	// ae
	// : 2026/07/01 - 1w
	// booked
	// : 2026/05/11 12-13
	// : 2026/05/12 12-13
	weeklyAvalabilities = []schedule_domain.WeeklyAvailability{
		schedule_domain.NewWeeklyAvailability(
			time.Monday,
			shared_domain_time.NewTimeRange(
				shared_domain_time.MustLocalTime24(10, 0),
				utils_time.MustParseDuration("4h"),
			),
		),
		schedule_domain.NewWeeklyAvailability(
			time.Tuesday,
			shared_domain_time.NewTimeRange(
				shared_domain_time.MustLocalTime24(14, 0),
				utils_time.MustParseDuration("4h"),
			),
		),
	}
	availabilityExceptions = []schedule_domain.AvailabilityException{
		schedule_domain.NewAvailabilityException(
			shared_domain_time.NewDateTimeRange(
				july1,
				duration1Week,
			),
		),
	}
	bookedList = []schedule_domain.Booking{
		schedule_domain.NewBooking(
			booked1Uuid,
			shared_domain_time.NewDateTimeRange(date_2026_may_11_12h00, duration1h),
			professor,
			student,
			nowMarch1,
		),
		schedule_domain.NewBooking(
			booked2Uuid,
			shared_domain_time.NewDateTimeRange(date_2026_may_12_14h00, duration1h),
			professor,
			student,
			nowMarch1,
		),
	}
	professor = schedule_domain.NewProfessor(professorUuid, "John", "Doe", schedule_domain.NewSchedule(weeklyAvalabilities, availabilityExceptions))
	student   = schedule_domain.NewStudent(studentUuid, "Jane", "Smith")
)

var (
	date_2026_may_20_12h00 = time.Date(2026, time.May, 20, 12, 0, 0, 0, time.UTC)
	date_2026_may_19_14h00 = time.Date(2026, time.May, 19, 14, 0, 0, 0, time.UTC)
	date_2026_may_11_12h00 = time.Date(2026, time.May, 11, 12, 0, 0, 0, time.UTC)
	date_2026_may_12_14h00 = time.Date(2026, time.May, 12, 14, 0, 0, 0, time.UTC)
	date_2026_july_6_12h00 = time.Date(2026, time.July, 6, 12, 0, 0, 0, time.UTC)
	studentDate1           = time.Date(2026, time.May, 9, 12, 0, 0, 0, time.UTC)
	studentDate2           = time.Date(2026, time.May, 9, 10, 30, 0, 0, time.UTC)
	professorDate          = time.Date(2026, time.May, 9, 11, 0, 0, 0, time.UTC)
	july25                 = time.Date(2026, time.July, 25, 12, 0, 0, 0, time.UTC)
	july1                  = time.Date(2026, time.July, 1, 0, 0, 0, 0, time.UTC)
	nowMarch1              = time.Date(2026, time.March, 1, 12, 0, 0, 0, time.UTC)
	nowjuly2               = time.Date(2026, time.July, 2, 12, 0, 0, 0, time.UTC)
	duration1h             = utils_time.MustParseDuration("1h")
	duration4h             = utils_time.MustParseDuration("4h")
	duration1Week          = utils_time.MustParseDuration("168h")
)
