# go-pg-domain-hex-actor
You can take a look on orginal videos: https://www.youtube.com/watch?v=VQym87o91f8 and repository: https://github.com/katzien/go-structure-examples
In in repository we just make a simple project which we use domain-hex-actor with postgres database.

# Prepare.
You need to create database `go_pg_domain_hex_actor` and run this script to create `beers` table:
```
CREATE TABLE beers (
    id SERIAL PRIMARY KEY,
    name text,
    brewery text,
    abv integer,
    short_description text,
    created timestamp without time zone
);

-- Indices -------------------------------------------------------

CREATE UNIQUE INDEX beers_pkey ON beers(id int4_ops);
```
Then you need to update `postgresURL` in `main.go` with our connection on your local.

# Run it.
Just run it like a normal golang project and joint with it.
```
go build -o go-pg-domain-hex-actor ./cmd/server/main.go
./go-pg-domain-hex-actor
```

# Enhancement.
In this repository I just make a very simple how we use it first, you can continue with it and let me know if you have any issues on it :).

Some suggestion:
+ Use `https://github.com/golang-migrate/migrate` for create database and migrate.
+ Use `https://github.com/Masterminds/squirrel` for wrap your query.
+ Use `https://github.com/jackc/pgx` to replace `pg` for bester performance and support.
+ Use `https://github.com/spf13/viper` for your environment management.
