#!/bin/bash -e

# set some environment variables
readonly RAY_ROOT=$(cd $(dirname "${BASH_SOURCE}")/.. && pwd -P)
readonly RAY_OUTPUT="${RAY_ROOT}/_output/local"
readonly RAY_OUTPUT_SRCPATH="${RAY_OUTPUT}/src"
readonly RAY_OUTPUT_BINPATH="${RAY_OUTPUT}/bin"

readonly RAY_TARGETS=(
	cmd/sunlight
  )

eval $(go env)

# enable/disable failpoints
toggle_failpoints() {
	FAILPKGS="rayserver/ rayserver/auth/"

	mode="disable"
	if [ ! -z "$FAILPOINTS" ]; then mode="enable"; fi
	if [ ! -z "$1" ]; then mode="$1"; fi

	if which gofail >/dev/null 2>&1; then
		gofail "$mode" $FAILPKGS
	elif [ "$mode" != "disable" ]; then
		echo "FAILPOINTS set but gofail not found"
		exit 1
	fi
}

ray_setup_gopath() {
	# preserve old gopath to support building with unvendored tooling deps (e.g., gofail)
	if [ -n "$GOPATH" ]; then
		GOPATH=":$GOPATH"
	fi
	export GOPATH=${RAY_OUTPUT}

	rm -rf ${RAY_OUTPUT_SRCPATH}
	mkdir -p ${RAY_OUTPUT_SRCPATH}

	ln -s ${RAY_ROOT} ${RAY_OUTPUT_SRCPATH}/ray
}

ray_build_target() {
	toggle_failpoints

	for arg; do
		# echo "target: ${arg}, ${arg##*/}"
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build $GO_BUILD_FLAGS \
		-installsuffix cgo -ldflags "$GO_LDFLAGS" \
		-o ${RAY_OUTPUT_BINPATH}/${arg##*/}.x ray/${arg} || return
	done
}

ray_make_ldflag() {
  local key=${1}
  local val=${2}

  echo "-X ray/cmd/version.${key}=${val}"
}

# Prints the value that needs to be passed to the -ldflags parameter of go build
# in order to set the project on the git tree status.
ray_version_ldflags() {
	local -a ldflags=($(ray_make_ldflag "buildDate" "$(date -u +'%Y-%m-%dT%H:%M:%SZ')"))

	local git_sha=`git rev-parse --short HEAD || echo "GitNotFound"`
	if [ ! -z "$FAILPOINTS" ]; then
		git_sha="$git_sha"-FAILPOINTS
	fi

	ldflags+=($(ray_make_ldflag "gitSHA" "${git_sha}"))

	echo "${ldflags[*]-}"
}

toggle_failpoints

# only build when called directly, not sourced
if echo "$0" | grep "build.sh$" >/dev/null; then
	# force new gopath so builds outside of gopath work
	ray_setup_gopath
	ray_version_ldflags
	#ray_build
	ray_build_target "${RAY_TARGETS[@]}"
fi
