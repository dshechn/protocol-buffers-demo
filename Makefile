.PHONY: pb

default: pb

version:
	@go version

pb: version
	@protoc -I/usr/local/include -I. --go_out=./ addressbook.proto

pbd: version
	@protoc -I/usr/local/include -I. --go_out=./ protocol_type.proto

add: version add_person.go
	go build -o add_person_go add_person.go

list: version list_people.go
	go build -o list_people_go list_people.go

clean:
	rm -f add_person_go list_people_go
