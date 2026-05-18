package shared_infra

import "github.com/google/uuid"

type InMemUuidGenerator struct {
	expectedUuid uuid.UUID
}

func NewInMemUuidGenerator(expectedUuid uuid.UUID) *InMemUuidGenerator {
	return &InMemUuidGenerator{
		expectedUuid: expectedUuid,
	}
}

func (g *InMemUuidGenerator) Generate() uuid.UUID {
	return g.expectedUuid
}
