#!/bin/bash
# wait-for-it.sh: Wait for a service to be ready before starting another service

set -e

host="$1"
shift
cmd="$@"

until pg_isready -h "$host"; do
  >&2 echo "Postgres is unavailable - sleeping"
  sleep 2
done

>&2 echo "Postgres is up - executing command"
exec $cmd
