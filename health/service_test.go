package health_test

import (
	"testing"

	"github.com/Becklyn/go-wire-core/app"
	"github.com/Becklyn/go-wire-core/health"
	"github.com/fraym/golog"
	"github.com/stretchr/testify/assert"
)

func Test_IsHealthy_ReturnsHealthyByDefault(t *testing.T) {
	logger := golog.NewZerologLogger()
	lifecycle := app.NewLifecycle()
	service := health.New(logger, lifecycle)

	healthy, _ := service.IsHealthy()
	assert.True(t, healthy) // must be healthy by default
}

func Test_IsHealthy_ReturnsNoReasonByDefault(t *testing.T) {
	logger := golog.NewZerologLogger()
	lifecycle := app.NewLifecycle()
	service := health.New(logger, lifecycle)

	_, reason := service.IsHealthy()
	assert.Equal(t, "", reason) // must have no reason by default
}

func Test_IsHealthy_ForAUnchangedComponent_ReturnsHealthy(t *testing.T) {
	logger := golog.NewZerologLogger()
	lifecycle := app.NewLifecycle()
	service := health.New(logger, lifecycle)

	healthy, _ := service.IsHealthy("foo")
	assert.True(t, healthy) // must be healthy by default
}

func Test_IsHealthy_ForAUnchangedComponent_ReturnsNoReason(t *testing.T) {
	logger := golog.NewZerologLogger()
	lifecycle := app.NewLifecycle()
	service := health.New(logger, lifecycle)

	_, reason := service.IsHealthy("foo")
	assert.Equal(t, "", reason) // must have no reason by default
}

func Test_IsHealthy_ReturnsUnhealthy_IfAnyComponentIsUnhealthy(t *testing.T) {
	logger := golog.NewZerologLogger()
	lifecycle := app.NewLifecycle()
	service := health.New(logger, lifecycle)

	service.SetUnhealthy("foo", "can't foo without bar")
	service.SetHealthy("bar")

	healthy, _ := service.IsHealthy()
	assert.True(t, !healthy) // must be unhealthy
}

func Test_IsHealthy_ForAHealthyComponent_ReturnsHealthy(t *testing.T) {
	logger := golog.NewZerologLogger()
	lifecycle := app.NewLifecycle()
	service := health.New(logger, lifecycle)

	service.SetHealthy("foo")

	healthy, _ := service.IsHealthy("foo")
	assert.True(t, healthy) // must be healthy
}

func Test_IsHealthy_ForAnUnhelathyComopnent_ReturnsUnhealthy(t *testing.T) {
	logger := golog.NewZerologLogger()
	lifecycle := app.NewLifecycle()
	service := health.New(logger, lifecycle)

	service.SetUnhealthy("foo", "can't foo without bar")

	healthy, _ := service.IsHealthy("foo")
	assert.True(t, !healthy) // must be unhealthy
}

func Test_IsHealthy_ForAnUnhelathyComopnent_ReturnsReason(t *testing.T) {
	logger := golog.NewZerologLogger()
	lifecycle := app.NewLifecycle()
	service := health.New(logger, lifecycle)

	service.SetUnhealthy("foo", "can't foo without bar")

	_, reason := service.IsHealthy("foo")
	assert.Equal(t, "can't foo without bar", reason) // must be the reason of foo
}
