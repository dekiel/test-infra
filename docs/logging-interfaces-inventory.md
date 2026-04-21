# Logging Interfaces Inventory Report

Date: 2026-04-21

## Context

This report documents the inventory of all consumers of `pkg/logging` interfaces from the `kyma-project/test-infra` repository across the entire organization. The goal is to assess the impact of a breaking change in `pkg/logging` to align interfaces with ADR-006 (Logging Standard).

## Target State (ADR-006)

The target interfaces as defined in ADR-006:

```go
type LoggerInterface interface {
    StructuredLoggerInterface
    WithLoggerInterface
    SyncLoggerInterface
}

type StructuredLoggerInterface interface {
    Infow(message string, keysAndValues ...interface{})
    Warnw(message string, keysAndValues ...interface{})
    Errorw(message string, keysAndValues ...interface{})
    Debugw(message string, keysAndValues ...interface{})
}

type WithLoggerInterface interface {
    With(args ...interface{}) LoggerInterface  // was: *zap.SugaredLogger
}

type SyncLoggerInterface interface {
    Sync() error
}
```

## Breaking Changes Summary

| Change | Type | Impact |
|---|---|---|
| Remove `SimpleLoggerInterface` from `LoggerInterface` | Removal | Consumers using `Info()`, `Warn()`, etc. |
| Remove `FormatedLoggerInterface` from `LoggerInterface` | Removal | Consumers using `Infof()`, `Warnf()`, etc. |
| Add `SyncLoggerInterface` to `LoggerInterface` | Addition | Implementations must add `Sync()` |
| `WithLoggerInterface.With()` returns `LoggerInterface` instead of `*zap.SugaredLogger` | Signature change | All consumers of `WithLoggerInterface` |

## Interface Usage — No External Consumers

| Interface | External consumers | Internal consumers (test-infra) |
|---|---|---|
| `SimpleLoggerInterface` | 0 | 0 |
| `FormatedLoggerInterface` | 0 | 0 |

**Removing these interfaces has zero impact on any consumer.**

## Interface Usage — Consumers Requiring Migration

### `StructuredLoggerInterface` — Signature Unchanged

Consumers use this interface via local composition. The method signatures do not change, so these consumers only need an import path update if the package path changes. **No code changes needed if the breaking change is done in-place.**

| Repository | File | Usage pattern |
|---|---|---|
| `kyma/ocm-artifact-builder` | `internal/core/build/service.go` | Local interface composition |
| `kyma/ocm-artifact-builder` | `internal/core/index/reverse_index.go` | Local interface composition |
| `kyma/ocm-artifact-builder` | `pkg/logging/logging.go` | Local interface composition |
| `kyma/ocm-artifact-builder` | `cmd/builder/main.go` | Local interface composition |
| `kyma/dora-integration` | `internal/adapters/github/status.go` | Local interface composition |
| `kyma/dora-integration` | `internal/adapters/github/deployments.go` | Local interface composition |
| `kyma/dora-integration` | `internal/core/dora/dora.go` | Local interface composition |
| `kyma-project/test-infra` | `cmd/image-builder/main.go` | Local interface composition |
| `kyma-project/test-infra` | `cmd/oidc-token-verifier/main.go` | Local interface composition |
| `kyma-project/test-infra` | `pkg/oidc/oidc.go` | Local interface composition |
| `kyma-project/test-infra` | `pkg/tags/tags.go` | Local interface composition |

### `WithLoggerInterface` — Signature Changes (`With()` return type)

These consumers embed `WithLoggerInterface` in local interfaces. The return type change from `*zap.SugaredLogger` to `LoggerInterface` breaks their local interface definition. **Local interfaces must be updated, and the injected logger implementation must provide a `With()` method returning the interface.**

| Repository | File | Usage pattern |
|---|---|---|
| `kyma/sec-scanner-cfg-processor` | `pkg/logging/logger.go` | Local interface composition |
| `kyma/checkmarx` | `pkg/logging/logger.go` | Local interface composition |
| `kyma-project/test-infra` | `cmd/image-builder/main.go` | Local interface composition |
| `kyma-project/test-infra` | `cmd/oidc-token-verifier/main.go` | Local interface composition |
| `kyma-project/test-infra` | `pkg/oidc/oidc.go` | Local interface composition |
| `kyma-project/test-infra` | `pkg/tags/tags.go` | Local interface composition |

### `LoggerInterface` (full) — Direct Usage

These consumers use `logging.LoggerInterface` directly (not via composition). They are affected by removal of `SimpleLoggerInterface` and `FormatedLoggerInterface`, and must switch to the new `LoggerInterface` definition.

| Repository | File | Usage pattern |
|---|---|---|
| `kyma-project/test-infra` | `pkg/gcp/pubsub/types.go` | Direct field type |
| `kyma-project/test-infra` | `pkg/gcp/pubsub/client.go` | Direct parameter type |
| `kyma-project/test-infra` | `pkg/cloudevents/client/client.go` | Direct field type + parameter type |

## Repositories Not Using `pkg/logging`

The following repositories were searched and confirmed to have **zero** imports of `pkg/logging`:

- `kyma/conduit-cli`
- `kyma/kyma-modules`
- `kyma/module-manifests`
- `kyma/neighbors-contracts`
- `kyma/oci-image-builder`
- `kyma/security-dashboard`
- `kyma/security-scanners`
- `kyma/security-scans-modular`
- `kyma/test-infra` (internal mirror)
- `kyma/modg`
- `kyma/image-syncer`
- `kyma/security-scans-issues`
- `kyma/tooling-infra`

## Migration Effort Estimate

| Repository | Effort | Reason |
|---|---|---|
| `kyma-project/test-infra` | Medium | 6 files: update interfaces, switch to new logger factory, update `pubsub` and `cloudevents` |
| `kyma/ocm-artifact-builder` | Low | Bump `go.mod`, no code changes if done in-place |
| `kyma/dora-integration` | Low | Bump `go.mod`, no code changes if done in-place |
| `kyma/sec-scanner-cfg-processor` | Low-Medium | Bump `go.mod`, update local interface + switch logger factory |
| `kyma/checkmarx` | Low-Medium | Bump `go.mod`, update local interface + switch logger factory |

Total: **5 repositories**, **4 external**, all under team control.
