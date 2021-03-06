// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUrlEncode(t *testing.T) {

	toEncode := "testing 1 2 3"
	encoded := UrlEncode(toEncode)

	require.Equal(t, encoded, "testing%201%202%203")

	toEncode = "testing123"
	encoded = UrlEncode(toEncode)

	require.Equal(t, encoded, "testing123")

	toEncode = "testing$#~123"
	encoded = UrlEncode(toEncode)

	require.Equal(t, encoded, "testing%24%23~123")
}
