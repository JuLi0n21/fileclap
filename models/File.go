package models

import (
	"time"

	"github.com/google/uuid"
)

type File struct {
	ID     uuid.UUID
	Name   string
	Date   time.Time
	Size   int64
	Folder string
}
