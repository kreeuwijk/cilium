// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package loader

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/cilium/cilium/pkg/datapath/linux/config"
	"github.com/cilium/cilium/pkg/testutils"
)

func TestWrap(t *testing.T) {
	var (
		realEPBuffer   bytes.Buffer
		templateBuffer bytes.Buffer
	)

	realEP := testutils.NewTestEndpoint()
	template := wrap(&realEP, nil)
	cfg := &config.HeaderfileWriter{}

	// Write the configuration that should be the same, and verify it is.
	err := cfg.WriteTemplateConfig(&realEPBuffer, &realEP)
	require.Nil(t, err)
	err = cfg.WriteTemplateConfig(&templateBuffer, template)
	require.Nil(t, err)
	require.Equal(t, realEPBuffer.String(), templateBuffer.String())

	// Write with the static data, and verify that the buffers differ.
	// Note this isn't an overly strong test because it only takes one
	// character to change for this test to pass, but we would ideally
	// define every bit of static data differently in the templates.
	realEPBuffer.Reset()
	templateBuffer.Reset()
	err = cfg.WriteEndpointConfig(&realEPBuffer, &realEP)
	require.Nil(t, err)
	err = cfg.WriteEndpointConfig(&templateBuffer, template)
	require.Nil(t, err)

	require.NotEqual(t, realEPBuffer.String(), templateBuffer.String())
}
