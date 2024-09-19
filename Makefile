build:
	docker-compose build goauth mysqldb

clean:
	docker container rm goauth goauth_db && docker image rm goauth

fclean: clean
	docker image rm mysql

stop:
	docker stop goauth goauth_db
