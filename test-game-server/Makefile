
.PHONY: run
run: setup-proto

.PHONY: setup-proto
setup-proto:
	protoc -I ./proto \
	--go_out=./pb \
	--go_opt=module=github.com/CA22-game-creators/cookingbomb-gameserver/test-game-server/proto \
	--go-grpc_out=./pb \
	--go-grpc_opt=require_unimplemented_servers=false \
	--go-grpc_opt=module=github.com/CA22-game-creators/cookingbomb-gameserver/test-game-server/proto \
	./proto/*.proto
	go mod tidy
