// Copyright 2023 StreamNative, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"oxia/common"
	"testing"
)

func TestCall_LogLevel_Default(t *testing.T) {
	var captured zerolog.Level
	rootCmd.SetArgs([]string{})
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		captured = common.LogLevel
		return nil
	}
	err := rootCmd.Execute()
	assert.Equal(t, common.DefaultLogLevel, captured)
	assert.NoError(t, err)
}

func TestCall_LogLevel(t *testing.T) {
	tests := []struct {
		name          string
		level         string
		expectedErr   error
		expectedLevel zerolog.Level
	}{
		{"disabled", "disabled", nil, zerolog.Disabled},
		{"trace", "trace", nil, zerolog.TraceLevel},
		{"debug", "debug", nil, zerolog.DebugLevel},
		{"info", "info", nil, zerolog.InfoLevel},
		{"warn", "warn", nil, zerolog.WarnLevel},
		{"error", "error", nil, zerolog.ErrorLevel},
		{"fatal", "fatal", nil, zerolog.FatalLevel},
		{"panic", "panic", nil, zerolog.PanicLevel},
		{"junk", "junk", LogLevelError("junk"), zerolog.InfoLevel},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			invoked := false
			rootCmd.SetArgs(append([]string{"-l"}, test.level))
			rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
				invoked = true
				assert.Equal(t, test.expectedLevel, common.LogLevel)
				return nil
			}
			err := rootCmd.Execute()
			if err == nil {
				assert.True(t, invoked)
			} else {
				assert.False(t, invoked)
			}
			assert.ErrorIs(t, err, test.expectedErr)
		})
	}
}
