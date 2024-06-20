run: build
	@./bin/app

dev:
	@air

build:
	@go build -o bin/app cmd/main.go

css-watch:
	@npx tailwindcss -i ./view/css/app.css -o public/styles.css --watch

templ-watch:
	@templ generate --watch --proxy=http://localhost:3000
	# @templ generate --watch