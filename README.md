# sprint (multi-module monorepo)

The Sprint framework and its capability modules, consolidated into **one repository
with one Go module per component** — the same model as `go.arpabet.com/store`.
A consuming app still pulls only the modules (and therefore the heavy deps) it
actually imports; this repo just removes the 17-separate-repos coordination tax.

This was produced by **Track B** of the composition plan, on top of the completed
**Track A** (everything unified on `glue v1.5.0` / `store v1.1.0` / `go 1.25`).

## Layout

All modules live under the `go.arpabet.com/sprint/` path prefix (a subdir monorepo
requires this — flat paths like `go.arpabet.com/cert` cannot live in subdirectories
of one repo). Directory = module subpath = release-tag prefix.

```
sprint/              go.arpabet.com/sprint/sprint            framework API (was go.arpabet.com/sprint)
sprintpb/            go.arpabet.com/sprint/sprintpb          framework protos
sprintframework/     go.arpabet.com/sprint/sprintframework   implementation + "everything" bundle
cert/  certmod/  certpb/      cert API / ACME+issuer impl / protos
dns/   dnsmod/               dns API / whois+resolver impl
fs/    fsmod/                filesystem API / impl
nat/   natmod/               NAT API / upnp+pmp impl
seal/  sealmod/              crypto-seal API / impl
raftapi/  raftmod/  raftpb/  raft API / impl / protos
```

**Stays external** (consumed as published deps, not folded in):
`go.arpabet.com/raft-badger` (released standalone as v1.0.1; no internal deps),
`go.arpabet.com/store` (+ `store/providers/*`), `glue`, `uuid`, `base62`,
`properties`.

## Local development

`go.work` ties the 17 modules together. Until the **first coordinated release**
publishes these new paths, each module also carries bootstrap `replace
go.arpabet.com/sprint/X => ../X` directives so it builds before anything is tagged.
`release.sh` strips those replaces and pins real versions at release time; after the
first release `go.work` alone is enough (as in the `store` repo).

```bash
go work sync
go build ./...        # from any module dir, e.g. cd cert && go build ./...
```

## Releasing

```bash
./release.sh --dry-run v1.0.0     # show the plan
./release.sh v1.0.0               # bump internal requires, strip bootstrap replaces, tag, push
```

Every module is tagged `‹subdir›/vX.Y.Z` (Go multi-module convention), e.g.
`cert/v1.0.0`, `sprintframework/v1.0.0`. One shared version moves everything; a
module carrying an extra change takes a per-module patch override
(`./release.sh v1.0.0 certmod=v1.0.1`).

## Adoption / migration

- **Vanity:** point `go.arpabet.com/sprint/*` (`?go-get=1`) at this repo.
- **Existing apps don't break:** the old `go.arpabet.com/cert@v1.x` (etc.) tags stay
  immutable on the module proxy, so anything already pinned keeps building. New
  development moves imports to `go.arpabet.com/sprint/cert`.
- The 17 original single-module repos can be archived once consumers have migrated.
