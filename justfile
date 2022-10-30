set export
set shell := ["fish", "-c"]
browse := if os() == "linux" { "xdg-open" } else { "open" }
copy        := if os() == "linux" { "xsel -ib"} else { "pbcopy" }

argocd_port  := "30950"

default:
    @just --list --unsorted

# Build nx
build-all:
  pnpm nx affected --target=build --parallel --max-parallel=10 --prod
test-all:
  pnpm nx affected --target=test --parallel --max-parallel=10
e2e-all:
  pnpm nx affected --target=e2e --parallel --max-parallel=10
lint-all:
  pnpm nx affected --target=lint --parallel --max-parallel=10

ci: build-all test-all e2e-all lint-all

create-cluster:
  -ctlptl create cluster kind --registry=ctlptl-registry

tilt-up:
  tilt up
