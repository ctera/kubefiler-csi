#!/bin/bash

set -euo pipefail


if [ $# -lt 2 ]; then
    echo "Must pass report file and expected coverage"
    exit 1
fi

REPORT_FILE=$1
EXPECTED_COVERAGE=$2

if [ ! -f "${REPORT_FILE}"  ]; then
    echo "File ${REPORT_FILE} does not exist"
    exit 1
fi

if ! [[ "${EXPECTED_COVERAGE}" =~ ^[0-9]+$ ]] ; then
    echo "Expected coverage must be an integer"
    exit 1
fi

readonly coverage=$(go tool cover -func=${REPORT_FILE} | grep total | awk '{print $3}' | awk -F '.' '{print $1}')
if [ ${coverage} -lt ${EXPECTED_COVERAGE} ]; then
    exit 1
fi
