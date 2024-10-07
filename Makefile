#name app
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}
mysqldumb:
	sudo docker exec -it mysql_container mysqldump -uroot -proot1234 --databases shopdevdo --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > migrations/shopdevgo.sql