# Health

This package provides a health http endpoint (`/health`) that can be controlled by the use of a `health.Service`.

## UseFiberEndpoint

The `UseFiberEndpoint` adds the `/health` route to your fiber app.

## Service functions

- `IsHealthy` checks a list of components if they are healthy. Returns false as soon as one of those components is unhealthy.
  If no component is specified, all available components will be checked.
- `SetHealthy` defines the health status of a given component as "healthy".
- `SetUnhealthy` defines the health status of a given component as "not healthy" and requires a reason.
