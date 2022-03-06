all: errors wrappingerrors customerrors castingerrors defer panic recover

# ======================
# error handling

errors:
	go run ./cmd/errors/main.go

wrappingerrors:
	go run ./cmd/wrappingerrors/main.go

customerrors:
	go run ./cmd/customerrors/main.go

castingerrors:
	go run ./cmd/castingerrors/main.go

defer:
	go run ./cmd/defer/main.go

panic:
	go run ./cmd/panic/main.go 0

recover:
	go run ./cmd/recover/main.go
