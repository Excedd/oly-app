package interfaces

import (
	"context"
	d "oly-backend/domain"
)

type Service interface {
	SetPRs(ctx context.Context, prs d.PersonalRecord) error
	GetPRs(ctx context.Context) (d.PersonalRecord, error)
}
