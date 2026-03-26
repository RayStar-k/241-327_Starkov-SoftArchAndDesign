#!/bin/sh

echo "Starting Guitar Shop Backend..."

echo "Waiting for PostgreSQL to be ready..."
until nc -z postgres-service 5432; do
  echo "PostgreSQL is unavailable - sleeping"
  sleep 2
done

echo "PostgreSQL is up - continuing..."

if [ "$GENERATE_TEST_DATA" = "true" ]; then
  echo "Generating test data..."
  ./gentestdata
  echo "Test data generation completed!"
fi

echo "Starting API server..."
exec ./main
