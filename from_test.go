package n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var deploymentFile = "test/helm/deployment.yaml"

func TestFromHelmFile(t *testing.T) {
	_, err := FromHelmFile(deploymentFile)
	assert.Nil(t, err)
}
