package log

import (
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLog(t *testing.T) {
	log.WithFields(log.Fields{
		"id":17722017,
		"name":"彭奏章",
	}).Warn("Failed to retrieve super manager")
}
