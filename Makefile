gen:
		protoc --go_out=. *.proto

clean:
		rm pb/*.go

run:
		go run main.go
