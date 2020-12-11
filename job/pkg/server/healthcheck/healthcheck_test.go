package healthcheck

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	v1 "google.golang.org/grpc/health/grpc_health_v1"
)

func TestHealthChecker(t *testing.T) {
	type testSet struct {
		name     string
		expected Status
	}
	testSets := [...]testSet{
		{
			name:     "healthy",
			expected: Healthy,
		},
		{
			name:     "unhealthy",
			expected: Unhealthy,
		},
	}

	for _, tt := range testSets {
		t.Run(tt.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(context.TODO())
			defer cancel()
			checker, statusCh := NewChecker(ctx)
			statusCh <- tt.expected
			actual, err := checker.Check(context.TODO(), nil)
			require.NoError(t, err)
			require.Equal(t, v1.HealthCheckResponse_ServingStatus(tt.expected), actual.Status)
		})
	}
}
