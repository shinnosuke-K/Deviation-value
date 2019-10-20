
dbstart:
	postgres -D /usr/local/var/postgres &

dbstop:
	pkill postgres

run:

	go run server.go


