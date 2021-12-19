#!/bin/bash
docker run --rm -it -d -p 5432:5432 --name postgres -e POSTGRES_PASSWORD=postgres postgres:13.5