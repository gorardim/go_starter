gen_api:
	go run cmd/xapi/xapi/main.go --dir=api/api --api-title="open api"

gen_admin_api:
	go run cmd/xapi/xapi/main.go --dir=api/admin --api-title="admin api"

gen_job:
	go run cmd/xapi/xnsq/main.go --dir=api/job

gen_model:
	go run cmd/genmodel/main.go

gen_repo:
	go run cmd/genrepo/main.go

mysql_schema_exporter:
	go run cmd/mysql_schema_exporter/main.go

install:
	go install ./cmd/travel_schema


deploy_dev:
	ssh travel "cd /root/shell && sh ./deploy_backend_go.sh master"
