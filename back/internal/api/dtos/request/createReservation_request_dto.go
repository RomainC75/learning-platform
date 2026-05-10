package dtos

import (
	"time"

	"github.com/google/uuid"
)

type CreateReservationRequest struct {
	ProfessorId uuid.UUID     `json:"professor_id" binding:"required" validate:"required"`
	Date        time.Time     `json:"date" binding:"required" validate:"required"`
	Duration    time.Duration `json:"duration_m" binding:"required" validate:"required"`
}
