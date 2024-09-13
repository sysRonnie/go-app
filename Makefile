run: 
	@templ generate
	@go run cmd/main.go

css:
	npx tailwindcss -i view/css/app.css -o public/styles.css --watch
