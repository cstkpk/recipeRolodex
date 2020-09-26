initDB: # Clear DB and re-initialize with seed data
	@if brew services list | grep -q "mariadb           stopped"; then \
		echo "MariaDB is not running. Exiting program."; exit 1; \
	fi;
	@echo "Initializing database..."
	cd ./db; bash initDB.sh
	@echo "Done!"
