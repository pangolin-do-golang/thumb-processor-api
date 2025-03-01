package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name       string
		envVars    map[string]string
		want       *Config
		wantErr    bool
		wantErrMsg string // Add a field for specific error message check
	}{
		{
			name: "Success",
			envVars: map[string]string{
				"S3_BUCKET":     "test-bucket",
				"SQS_QUEUE_URL": "https://sqs.us-east-1.amazonaws.com/123456789012/test-queue",
				"DB_USERNAME":   "test_db_username",
				"DB_PASSWORD":   "test_db_password",
				"DB_HOST":       "test_db_host",
				"DB_PORT":       "test_db_port",
				"DB_NAME":       "test_db_name",
			},
			want: &Config{
				S3:  S3{Bucket: "test-bucket"},
				SQS: SQS{QueueURL: "https://sqs.us-east-1.amazonaws.com/123456789012/test-queue"},
				DB: Database{
					User:     "test_db_username",
					Password: "test_db_password",
					Host:     "test_db_host",
					Port:     "test_db_port",
					Name:     "test_db_name",
				},
			},
			wantErr: false,
		},
		{
			name: "MissingS3Bucket",
			envVars: map[string]string{
				"SQS_QUEUE_URL": "https://sqs.us-east-1.amazonaws.com/123456789012/test-queue",
			},
			want: &Config{
				SQS: SQS{QueueURL: "https://sqs.us-east-1.amazonaws.com/123456789012/test-queue"}, // S3 should be empty
			},
			wantErr: false, // env package doesn't return error if some vars are missing
		},
		{
			name: "MissingSQSQueueURL",
			envVars: map[string]string{
				"S3_BUCKET": "test-bucket",
			},
			want: &Config{
				S3: S3{Bucket: "test-bucket"}, // SQS should be empty
			},
			wantErr: false, // env package doesn't return error if some vars are missing
		},
		{
			name:    "EmptyEnvVars",
			envVars: map[string]string{},
			want:    &Config{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variables for the test
			for k, v := range tt.envVars {
				os.Setenv(k, v)
				defer os.Unsetenv(k) // Ensure cleanup after the test
			}

			got, err := Load()

			if tt.wantErr {
				assert.Error(t, err)
				if tt.wantErrMsg != "" {
					assert.Contains(t, err.Error(), tt.wantErrMsg) // check specific error message
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
