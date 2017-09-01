// +build integration

package status

import (
	"testing"

	mbtest "github.com/elastic/beats/metricbeat/mb/testing"
	"github.com/elastic/beats/metricbeat/module/nginxupstream"

	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T) {
	f := mbtest.NewEventFetcher(t, getConfig())
	event, err := f.Fetch()
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	t.Logf("%s/%s event: %+v", f.Module().Name(), f.Name(), event)

	// Check number of fields.
	assert.Equal(t, 10, len(event))
}

func getConfig() map[string]interface{} {
	return map[string]interface{}{
		"module":     "nginxupstream",
		"metricsets": []string{"status"},
		"hosts":      []string{nginxupstream.GetNginxEnvHost()},
	}
}