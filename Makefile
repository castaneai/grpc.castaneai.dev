generate:
	protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/echo.proto

deploy:
	ko resolve -f cloudrun.yaml | gcloud beta run services replace - --platform=managed --region=asia-northeast1
