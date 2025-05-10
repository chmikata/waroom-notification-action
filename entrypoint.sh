#!/bin/bash
set -e

while getopts "c:a:w:t:" opt; do
  case "${opt}" in
    c)
      cmd=${OPTARG}
      args=${cmd}
    ;;
    a)
      args=$(echo "${args} -a ${OPTARG}")
    ;;
    w)
      args=$(echo "${args} -w ${OPTARG}")
    ;;
    t)
      args=$(echo "${args} -t ${OPTARG}")
    ;;
  esac
done

/app/incident-notification ${args}
