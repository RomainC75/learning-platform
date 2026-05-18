package shared_domain_uuid

import "github.com/google/uuid"

type UuidGenerator interface {
	Generate() uuid.UUID
}
