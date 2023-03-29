package readyness_test

import (
	"testing"

	"github.com/Becklyn/go-wire-core/readyness"
	"github.com/fraym/golog"
	"github.com/stretchr/testify/assert"
)

func Test_IsReady_ReturnsHealthyByDefault(t *testing.T) {
	logger := golog.NewZerologLogger()
	service := readyness.New(logger)

	ready, _ := service.IsReady()
	assert.True(t, ready) // must be ready by default
}

func Test_IsReady_ReturnsNoComponentByDefault(t *testing.T) {
	logger := golog.NewZerologLogger()
	service := readyness.New(logger)

	_, component := service.IsReady()
	assert.Equal(t, "", component) // must have no component by default
}

func Test_IsReady_ForANewComponent_ReturnsNotReady(t *testing.T) {
	logger := golog.NewZerologLogger()
	service := readyness.New(logger)

	ready, _ := service.IsReady("foo")
	assert.False(t, ready) // must be not ready by default
}

func Test_IsReady_ForANewComponent_ReturnsComponent(t *testing.T) {
	logger := golog.NewZerologLogger()
	service := readyness.New(logger)

	_, component := service.IsReady("foo")
	assert.Equal(t, "foo", component) // must have component by default
}

func Test_IsReady_ReturnsNotReady_IfAnyComponentIsNotReady(t *testing.T) {
	logger := golog.NewZerologLogger()
	service := readyness.New(logger)

	service.Register("foo")
	service.SetReady("bar")

	ready, _ := service.IsReady()
	assert.True(t, !ready) // must be not ready
}

func Test_IsReady_ForAReadyComponent_ReturnsReady(t *testing.T) {
	logger := golog.NewZerologLogger()
	service := readyness.New(logger)

	service.SetReady("foo")

	ready, _ := service.IsReady("foo")
	assert.True(t, ready) // must be ready
}

func Test_IsReady_ForANotReadyComopnent_ReturnsNotReady(t *testing.T) {
	logger := golog.NewZerologLogger()
	service := readyness.New(logger)

	service.Register("foo")

	ready, _ := service.IsReady("foo")
	assert.True(t, !ready) // must be not ready
}

func Test_IsReady_ForANotReadyComopnent_ReturnsComponent(t *testing.T) {
	logger := golog.NewZerologLogger()
	service := readyness.New(logger)

	service.SetReady("foo")

	_, component := service.IsReady("foo", "bar")
	assert.Equal(t, "bar", component) // must be the bar component
}
