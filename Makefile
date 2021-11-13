migrationCreate:
	migrate create -ext sql -dir persistence/migrations -seq init_schema
migrationUp:
	migrate -path persistence/migrations -database "postgres://olbdkjun:1Wx5AldpwkeozK5lhc4LGWAGwMRm0FsP@kashin.db.elephantsql.com/olbdkjun?sslmode=disable" -verbose up
migrationDown:
	migrate -path persistence/migrations -database "postgres://olbdkjun:1Wx5AldpwkeozK5lhc4LGWAGwMRm0FsP@kashin.db.elephantsql.com/olbdkjun?sslmode=disable" -verbose down
migrationForce:
	migrate -path persistence/migrations -database "postgres://olbdkjun:1Wx5AldpwkeozK5lhc4LGWAGwMRm0FsP@kashin.db.elephantsql.com/olbdkjun?sslmode=disable" -verbose force 2
migrationGoto:
	migrate -path persistence/migrations -database "postgres://olbdkjun:1Wx5AldpwkeozK5lhc4LGWAGwMRm0FsP@kashin.db.elephantsql.com/olbdkjun?sslmode=disable" -verbose goto 2
installSqlc:
	go get github.com/kyleconroy/sqlc/cmd/sqlc
initialiseGoModules:
	go mod init client
dockerBuild:
	docker build -t communication:latest .
dockerRun:
	 docker run --mount source=persian-black-logs,destination=/usr/local/bin/log/ -p 8084:8083 --name communication communication:latest
.PHONY: migrationCreate migrationUp migrationDown migrationForce migrationGoto installSqlc initialiseGoModules dockerRun dockerBuild