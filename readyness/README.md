# Readyness

This package provides a readyness http endpoint (`/ready`) that can be controlled by the use of a `readyness.Service`.

## UseFiberEndpoint

The `UseFiberEndpoint` adds the `/ready` route to your fiber app.

## Service functions

- `IsReady` checks a list of components if they are ready. Returns false as soon as one of those components is not ready.
  If no component is specified, all available components will be checked.
- `SetReady` defines the ready status of a given component as "ready".
- `Register` registers a new component as not ready.
