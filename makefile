run: build
	@./bin/app

build:
	@go build -o bin/app cmd/main.go

air:
	@air

css-watch:
	@npx tailwindcss -i ./view/css/app.css -o public/styles.css --watch

templ-watch:
	@templ generate --watch --proxy=http://localhost:3000
	# @templ generate --watch

run-all:
	tmux new-session -d -s dev-session 'make air' \; \
		split-window -h 'make css-watch' \; \
		split-window -v 'make templ-watch' \; \
		select-layout even-horizontal \; \
		attach