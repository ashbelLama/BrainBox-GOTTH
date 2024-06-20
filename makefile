run: build
	@./bin/app

dev: css-watch
	@air

build:
	@go build -o bin/app cmd/main.go

css-watch:
	@npx tailwindcss -i ./view/css/app.css -o public/styles.css --watch