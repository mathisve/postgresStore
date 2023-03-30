#!/bin/bash

psql \
    "postgres://postgres:password@localhost:5432" \
    -c "DROP TABLE IF EXISTS object CASCADE;"