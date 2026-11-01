# go-slurmstack.github.io

Landing page of [slurmstack](https://go-slurmstack.github.io/), a Slurm cluster
on OpenStack from one Go binary. Served at <https://go-slurmstack.github.io>
and built with [Hugo](https://gohugo.io): a single page (`layouts/index.html`;
the cluster's roles are driven by `[[params.nodes]]` in `hugo.toml`).

Documentation lives in [go-slurmstack/docs](https://github.com/go-slurmstack/docs),
served at <https://go-slurmstack.github.io/docs/>. The slurmstack source
repository is private.

`.github/workflows/deploy.yml` builds the page on every pull request and
deploys it to GitHub Pages (Actions deployment) on every push to `main`.

## Local preview

```sh
hugo server      # http://localhost:1313
```

## The design is not here

It comes from the shared Hugo module
[`go-fleettools/landingkit`](https://github.com/go-fleettools/landingkit),
pinned in `go.mod`; `kit.go` keeps `go mod tidy` from deleting that
requirement. `[params.brand]` in `hugo.toml` is read from the logo. A file under
`layouts/partials/` named like one of the kit's would silently override it.
