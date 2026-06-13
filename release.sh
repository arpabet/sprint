#!/usr/bin/env bash
#
# Copyright (c) 2026 Karagatan LLC.
# SPDX-License-Identifier: BUSL-1.1
#
# Coordinated release for the go.arpabet.com/sprint multi-module monorepo.
#
# One shared version moves every module; an interface change ripples into all of
# them. A module carrying an extra change takes a higher patch via a per-module
# override:
#
#     ./release.sh v1.3.0 certmod=v1.3.1
#
# Modules are discovered automatically (every dir with a go.mod) and tagged with
# the multi-module convention "<subdir>/vX.Y.Z" (e.g. cert/v1.3.0). Before tagging,
# internal `require go.arpabet.com/sprint/X` lines are pinned to the release
# version and the local-dev `replace go.arpabet.com/sprint/X => ../X` bootstrap
# directives are stripped (consumers ignore replaces anyway; this keeps published
# go.mods clean). Re-add nothing by hand — `go.work` covers local dev post-release.
#
# Usage: ./release.sh [--dry-run] [--no-push] <version> [module=version ...]
#
set -euo pipefail

PREFIX="go.arpabet.com/sprint"
REMOTE="origin"
DRY_RUN=0; NO_PUSH=0

die() { echo "error: $*" >&2; exit 1; }
semver_ok() { [[ "$1" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[0-9A-Za-z.]+)?$ ]]; }

# ---- args ----------------------------------------------------------------------
declare -A OVERRIDE
VERSION=""
for a in "$@"; do
	case "$a" in
		--dry-run) DRY_RUN=1 ;;
		--no-push) NO_PUSH=1 ;;
		*=v*)      OVERRIDE["${a%%=*}"]="${a#*=}" ;;
		v*)        VERSION="$a" ;;
		*)         die "unrecognized arg: $a" ;;
	esac
done
[[ -n "$VERSION" ]] || die "usage: ./release.sh [--dry-run] [--no-push] <version> [module=version ...]"
semver_ok "$VERSION" || die "'$VERSION' is not vMAJOR.MINOR.PATCH"
for m in "${!OVERRIDE[@]}"; do semver_ok "${OVERRIDE[$m]}" || die "override '$m=${OVERRIDE[$m]}' not valid semver"; done

[[ -z "$(git status --porcelain)" ]] || die "working tree is dirty; commit or stash first"

# ---- discover modules ----------------------------------------------------------
mapfile -t MODULES < <(find . -name go.mod -not -path './.*' | sed 's#/go.mod$##; s#^\./##' | sort)
[[ "${#MODULES[@]}" -gt 0 ]] || die "no modules found"

ver_for() { local m="$1"; echo "${OVERRIDE[$m]:-$VERSION}"; }

echo "Release plan (shared $VERSION):"
for m in "${MODULES[@]}"; do printf "  %-22s -> %s/%s\n" "$m" "$m" "$(ver_for "$m")"; done

# ---- rewrite go.mod: pin internal requires, strip bootstrap replaces ----------
for m in "${MODULES[@]}"; do
	gm="$m/go.mod"
	# strip local-dev replace directives for internal modules
	perl -i -ne "print unless m{^replace \Q$PREFIX\E/}" "$gm"
	# pin every internal require to its release version
	for dep in "${MODULES[@]}"; do
		dv="$(ver_for "$dep")"
		perl -i -pe "s{(\Q$PREFIX/$dep\E)\s+v\S+}{\$1 $dv}g" "$gm"
	done
done

if [[ "$DRY_RUN" -eq 1 ]]; then
	echo "--- dry run: go.mod changes below, nothing committed ---"
	git --no-pager diff -- '**/go.mod'
	git checkout -- . 2>/dev/null || true
	exit 0
fi

# ---- commit + tag --------------------------------------------------------------
git add -A
git commit -m "release $VERSION"
TAGS=()
for m in "${MODULES[@]}"; do t="$m/$(ver_for "$m")"; git tag "$t"; TAGS+=("$t"); done
echo "tagged: ${TAGS[*]}"

if [[ "$NO_PUSH" -eq 1 ]]; then
	echo "--no-push: created commit + tags locally; not pushed"
	exit 0
fi
git push "$REMOTE" HEAD
git push "$REMOTE" "${TAGS[@]}"
echo "released $VERSION"
