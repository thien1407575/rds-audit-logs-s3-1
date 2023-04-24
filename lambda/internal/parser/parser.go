package parser

import (
	"io"
	"rdsauditlogss3/internal/entity"
)

type Parser interface {
	ParseEntries(data io.Reader, logFileTimestamp int64, CheckpointTimestammp int64) ([]*entity.LogEntry, error)
}
