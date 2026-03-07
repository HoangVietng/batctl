#!/usr/bin/env bash
set -euo pipefail

usage() {
    echo "Usage: $0 [--dry-run]"
    echo ""
    echo "Creates a new CalVer release (YYYY.MM.N)."
    echo "Bumps version in PKGBUILD, commits, tags, and pushes."
    echo ""
    echo "Options:"
    echo "  --dry-run    Show what would happen without making changes"
    exit 1
}

DRY_RUN=false
for arg in "$@"; do
    case "$arg" in
        --dry-run) DRY_RUN=true ;;
        -h|--help) usage ;;
        *) echo "Unknown option: $arg"; usage ;;
    esac
done

REPO_ROOT="$(git rev-parse --show-toplevel)"
cd "$REPO_ROOT"

YEAR=$(date +%Y)
MONTH=$(date +%-m)
PREFIX="v${YEAR}.${MONTH}"

LAST_TAG=$(git tag -l "${PREFIX}.*" --sort=-v:refname | head -n1 || true)

if [[ -z "$LAST_TAG" ]]; then
    PATCH=1
else
    LAST_PATCH="${LAST_TAG##*.}"
    PATCH=$((LAST_PATCH + 1))
fi

NEW_VERSION="${YEAR}.${MONTH}.${PATCH}"
NEW_TAG="v${NEW_VERSION}"

echo "=== Release Plan ==="
echo "  Previous tag: ${LAST_TAG:-"(none this month)"}"
echo "  New version:  ${NEW_VERSION}"
echo "  New tag:      ${NEW_TAG}"
echo ""

if [[ -n "$LAST_TAG" ]]; then
    echo "=== Changelog (since ${LAST_TAG}) ==="
    git log "${LAST_TAG}..HEAD" --oneline --no-decorate
else
    FIRST_TAG=$(git tag -l "v*" --sort=-v:refname | head -n1 || true)
    if [[ -n "$FIRST_TAG" ]]; then
        echo "=== Changelog (since ${FIRST_TAG}) ==="
        git log "${FIRST_TAG}..HEAD" --oneline --no-decorate
    else
        echo "=== Changelog (all commits) ==="
        git log --oneline --no-decorate
    fi
fi
echo ""

if [[ "$DRY_RUN" == "true" ]]; then
    echo "[dry-run] Would update PKGBUILD pkgver to ${NEW_VERSION}"
    echo "[dry-run] Would commit, tag ${NEW_TAG}, and push"
    exit 0
fi

if [[ -n "$(git status --porcelain)" ]]; then
    echo "ERROR: Working tree is dirty. Commit or stash changes first."
    exit 1
fi

sed -i "s/^pkgver=.*/pkgver=${NEW_VERSION}/" PKGBUILD
echo "Updated PKGBUILD pkgver to ${NEW_VERSION}"

git add PKGBUILD
git commit -m "release: ${NEW_TAG}"
git tag -a "${NEW_TAG}" -m "Release ${NEW_VERSION}"

echo ""
echo "Pushing commits and tag..."
git push origin HEAD
git push origin "${NEW_TAG}"

echo ""
echo "=== Done ==="
echo "Tag ${NEW_TAG} pushed. GitHub Actions will create the release."
