#!/bin/bash
# Renames the operation IDs that the spec uses twice, which oapi-codegen turns
# into duplicate Go methods. `GET /screen` keeps `Screen` and
# `GET /addresses/{address}` keeps `GetAddress`; the other two get distinct
# names. Reads api/openapi.json and writes api/openapi.fixed.json. Fails if a
# renamed operation is gone or still clashes, so an upstream change shows up
# here instead of as a compile error.
set -euo pipefail
trap 'on_error "$?" "$LINENO"' ERR

LOG_FILE="${LOG_FILE:-/tmp/$(basename "$0" .sh).log}"
exec > >(tee -a "$LOG_FILE") 2>&1

SPEC_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)/api"
readonly SPEC_DIR
readonly SPEC_IN="${SPEC_DIR}/openapi.json"
readonly SPEC_OUT="${SPEC_DIR}/openapi.fixed.json"

on_error() {
	local exit_code="$1"
	local line="$2"

	printf '{"level":"ERROR","file":"%s","line":%d,"msg":"command failed","exit_code":%d}\n' \
		"${BASH_SOURCE[1]}" "$line" "$exit_code" >&2
}

log() {
	local level="$1"
	shift

	jq -nc \
		--arg time "$(date -u '+%Y-%m-%dT%H:%M:%SZ')" \
		--arg level "$level" \
		--arg file "${BASH_SOURCE[1]##*/}" \
		--arg line "${BASH_LINENO[0]}" \
		--arg func "${FUNCNAME[1]:-main}" \
		--arg msg "$*" \
		'{time:$time,level:$level,file:$file,line:($line|tonumber),func:$func,msg:$msg}' >&2
}

# shellcheck disable=SC2016 # $p and $id are jq variables, not shell ones
readonly JQ_FILTER='
def rename($p; $id):
	if .paths[$p].get == null then error("missing GET " + $p) else . end
	| .paths[$p].get.operationId = $id;

rename("/compliance/screen/{address}"; "GetComplianceScreen")
| rename("/affiliates/address"; "GetAffiliateAddress")
| ([.paths[][] | objects | .operationId // empty] | group_by(.)
	| map(select(length > 1) | .[0])) as $dupes
| if ($dupes | length) > 0
	then error("duplicate operationIds: " + ($dupes | join(", ")))
	else . end
'

main() {
	local tmp
	tmp="$(mktemp "${SPEC_DIR}/.openapi.fixed.XXXXXX")"
	# shellcheck disable=SC2064 # expand tmp now, the local is gone when EXIT fires
	trap "rm -f -- '${tmp}'" EXIT

	jq "$JQ_FILTER" "$SPEC_IN" >"$tmp"

	chmod 0644 "$tmp"
	mv -- "$tmp" "$SPEC_OUT"
	log INFO "renamed duplicate operation IDs out=${SPEC_OUT##*/}"
}

main "$@"
