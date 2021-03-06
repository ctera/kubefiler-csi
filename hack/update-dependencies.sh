#!/bin/bash

# Copyright 2021, CTERA Networks.
#
# Portions Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o nounset
set -o errexit
set -o pipefail

# Explicitly opt into go modules, even though we're inside a GOPATH directory
export GO111MODULE=on
# Explicitly clear GOPATH, to ensure nothing this script calls makes use of that path info
export GOPATH=
# Explicitly clear GOFLAGS, since GOFLAGS=-mod=vendor breaks dependency resolution while rebuilding vendor
export GOFLAGS=

cd "$(git rev-parse --show-toplevel)"
trap 'echo "FAILED" >&2' ERR
TMP_DIR="${TMP_DIR:-$(mktemp -d /tmp/update-vendor.XXXX)}"

prune-vendor() {
  find vendor -type f \
    -not -iname "*.c" \
    -not -iname "*.go" \
    -not -iname "*.h" \
    -not -iname "*.proto" \
    -not -iname "*.s" \
    -not -iname "AUTHORS*" \
    -not -iname "CONTRIBUTORS*" \
    -not -iname "COPYING*" \
    -not -iname "LICENSE*" \
    -not -iname "NOTICE*" \
    -delete
}

# ensure_require_replace_directives_for_all_dependencies:
# - ensures all existing 'require' directives have an associated 'replace' directive pinning a version
# - adds explicit 'require' directives for all transitive dependencies
# - adds explicit 'replace' directives for all require directives (existing 'replace' directives take precedence)
function ensure_require_replace_directives_for_all_dependencies() {
  local local_tmp_dir
  local_tmp_dir=$(mktemp -d "${TMP_DIR}/pin_replace.XXXX")

  # collect 'require' directives that actually specify a version
  local require_filter='(.Version != null) and (.Version != "v0.0.0") and (.Version != "v0.0.0-00010101000000-000000000000")'
  # collect 'replace' directives that unconditionally pin versions (old=new@version)
  local replace_filter='(.Old.Version == null) and (.New.Version != null)'

  # Capture local require/replace directives before running any go commands that can modify the go.mod file
  local require_json="${local_tmp_dir}/require.json"
  local replace_json="${local_tmp_dir}/replace.json"
  go mod edit -json | jq -r ".Require // [] | sort | .[] | select(${require_filter})" > "${require_json}"
  go mod edit -json | jq -r ".Replace // [] | sort | .[] | select(${replace_filter})" > "${replace_json}"

  # 1. Ensure require directives have a corresponding replace directive pinning a version
  cat "${require_json}" | jq -r '"-replace \(.Path)=\(.Path)@\(.Version)"'            | xargs -L 100 go mod edit -fmt
  cat "${replace_json}" | jq -r '"-replace \(.Old.Path)=\(.New.Path)@\(.New.Version)"'| xargs -L 100 go mod edit -fmt

  # 2. Add explicit require directives for indirect dependencies
  go list -m -json all | jq -r 'select(.Main != true) | select(.Indirect == true) | "-require \(.Path)@\(.Version)"'          | xargs -L 100 go mod edit -fmt

  # 3. Add explicit replace directives pinning dependencies that aren't pinned yet
  go list -m -json all | jq -r 'select(.Main != true) | select(.Replace == null)  | "-replace \(.Path)=\(.Path)@\(.Version)"' | xargs -L 100 go mod edit -fmt
}

function group_replace_directives() {
  local local_tmp_dir
  local_tmp_dir=$(mktemp -d "${TMP_DIR}/group_replace.XXXX")
  local go_mod_replace="${local_tmp_dir}/go.mod.replace.tmp"
  local go_mod_noreplace="${local_tmp_dir}/go.mod.noreplace.tmp"
  # separate replace and non-replace directives
  cat go.mod | awk "
     # print lines between 'replace (' ... ')' lines
     /^replace [(]/      { inreplace=1; next                   }
     inreplace && /^[)]/ { inreplace=0; next                   }
     inreplace           { print > \"${go_mod_replace}\"; next }

     # print ungrouped replace directives with the replace directive trimmed
     /^replace [^(]/ { sub(/^replace /,\"\"); print > \"${go_mod_replace}\"; next }

     # otherwise print to the noreplace file
     { print > \"${go_mod_noreplace}\" }
  "
  cat "${go_mod_noreplace}" >  go.mod
  echo "replace ("          >> go.mod
  cat "${go_mod_replace}"   >> go.mod
  echo ")"                  >> go.mod

  go mod edit -fmt
}

ensure_require_replace_directives_for_all_dependencies
go mod tidy
ensure_require_replace_directives_for_all_dependencies
group_replace_directives
go mod vendor
#prune-vendor
echo SUCCESS
