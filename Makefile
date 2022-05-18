dev:
	npm run dev

run:
	npm run build
	go run *.go 

install:
	npm install
	go mod vendor