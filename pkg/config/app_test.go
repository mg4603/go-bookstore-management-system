package config

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockOpen(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	mockDB := &gorm.DB{}
	return mockDB, nil
}

func mockOpenWithError(dialector gorm.Dialector, config *gorm.Config) (*gorm.DB, error) {
	return nil, errors.New("Mocked connection error")
}

func mockLoadEnvSuccess() error {
	return nil
}

func mockLoadEnvFailure() error {
	return errors.New("mocked environment loading error")
}

func TestConnect(t *testing.T) {
	tests := []struct {
		name         string
		envVars      map[string]string
		mockOpenFunc DBOpener
		mockLoader   EnvLoader
		expectError  bool
	}{
		{
			name: "Successful Connection",
			envVars: map[string]string{
				"DB_USER_NAME": "testuser",
				"DB_PASSWORD":  "testpass",
				"DB_NAME":      "testdb",
			},
			mockOpenFunc: mockOpen,
			mockLoader:   mockLoadEnvSuccess,
			expectError:  false,
		},
		{
			name: "Missing environment variables",
			envVars: map[string]string{
				"DB_USER_NAME": "testuser",
				"DB_PASSWORD":  "testpass",
			},
			mockOpenFunc: mockOpenWithError,
			mockLoader:   mockLoadEnvFailure,
			expectError:  true,
		},
		{
			name: "database connection error",
			envVars: map[string]string{
				"DB_USER_NAME": "testuser",
				"DB_PASSWORD":  "testpass",
				"DB_NAME":      "testdb",
			},
			mockOpenFunc: mockOpenWithError,
			mockLoader:   mockLoadEnvFailure,
			expectError:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			for key, value := range tc.envVars {
				err := os.Setenv(key, value)
				assert.NoError(t, err, "failed to set environment variables %s: %v", key, err)

			}
			defer func() {
				for key := range tc.envVars {
					os.Unsetenv(key)
				}
			}()

			err := Connect(tc.mockOpenFunc, tc.mockLoader)

			if tc.expectError {
				assert.Error(t, err, "expected an error but got nil")
			} else {
				assert.NoError(t, err, "expected no error but got: %v", err)
				assert.NotNil(t, db, "expected db to be initialized, but it is nil")
			}
		})
	}
}
