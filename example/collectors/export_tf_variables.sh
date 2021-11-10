#!/usr/bin/env bash

set -eo pipefail
echo "Exporting variables as secrets"

export TF_VAR_secret="${APPD_SECRET}"
export TF_VAR_controller_url="${APPD_CONTROLLER_BASE_URL}"
echo "Done"