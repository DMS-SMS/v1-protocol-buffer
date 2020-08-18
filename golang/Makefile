.PHONY: auth-admin
auth-admin:
	protoc --proto_path=../proto/auth --go_out=./auth --micro_out=./auth auth-admin.proto

.PHONY: auth-student
auth-student:
	protoc --proto_path=../proto/auth --go_out=./auth --micro_out=./auth auth-student.proto


.PHONY: outing
outing:
	protoc --proto_path=../proto --go_out=./outing --micro_out=./outing outing.proto

.PHONY: club
club:
	protoc --proto_path=../proto --go_out=./club --micro_out=./club club.proto

.PHONY: announcement
announcement:
	protoc --proto_path=../proto --go_out=./announcement --micro_out=./announcement announcement.proto

.PHONY: schedule
schedule:
	protoc --proto_path=../proto --go_out=./schedule --micro_out=./schedule schedule.proto
