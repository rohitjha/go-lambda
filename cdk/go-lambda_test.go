package main

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestGoLambdaStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := NewPythonLambdaStack(app, "LambdaStack", nil)

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	assert.NotNil(t, template)
}
