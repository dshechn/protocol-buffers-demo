.PHONY: pb

default: pb

pb:
	@go version
	@protoc -I/usr/local/include -I. --go_out=./ addressbook.proto

add: add_person.go
	go build -o add_person_go add_person.go

list: list_people.go
	go build -o list_people_go list_people.go

clean:
	rm -f add_person_go list_people_go
