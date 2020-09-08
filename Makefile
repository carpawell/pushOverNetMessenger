buildImage:
	@docker build . -t pushovernetmessenger

startPostgres:
	@docker run --name pushoverdb -e POSTGRES_PASSWORD=admin -e POSTGRES_USER=pushover \
			--network=pushOverNetMessenger \
			-d postgres

startMessenger:
	@docker run -p $(port):$(port) \
			-p $(DBPort):$(DBPort) \
			-e DB_DSN=$(DB_DSN) \
 			--name pushovernetmessenger \
 			pushovernetmessenger \
 			./pushOverNetMessenger \
 			-port=$(port) -app=$(app) -user=$(user)

startMessengerDocker:
	@docker run -p $(port):$(port) \
 			--name pushovernetmessenger \
 			--network=pushOverNetMessenger \
 			pushovernetmessenger \
 			./pushOverNetMessenger \
 			-port=$(port) -app=$(app) -user=$(user) -host=pushoverdb

build:
	@go build -o ./dist/ ./cmd/pushOverNetMessenger

run:
	@docker network create pushOverNetMessenger
	@make buildImage
	@make startPostgres
	@sleep 5
	@make host=$(host) app=$(app) port=$(port) user=$(user) startMessengerDocker

clean:
	@make rmImage
	@make rmPostgres
	@make rmPush
	@docker network rm pushOverNetMessenger

rmImage:
	@docker image rm -f pushovernetmessenger
rmPostgres:
	@docker rm -f pushoverdb
rmPush:
	@docker rm -f pushovernetmessenger
