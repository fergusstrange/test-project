package testproject

import (
	"log"
	"testing"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/assert"
)

func Test_DatabaseStarts(t *testing.T) {
	db := embeddedpostgres.NewDatabase(embeddedpostgres.DefaultConfig().
		RuntimePath("/home/runner/work/test-project/test-project"))
	err := db.Start()
	assert.NoError(t, err)

	log.Println("I started, now it's time to stop...")

	err = db.Stop()
	assert.NoError(t, err)
}
