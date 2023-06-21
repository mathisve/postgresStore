# postgresStore example

Example to show how you can use postgresStore to upload objects to Postgres / Timescale

```
./postgres.sh
go run .
```

```
psql "postgres://postgres:password@localhost:5432" -c "SELECT * FROM object;"
```
