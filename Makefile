PHONY: table clean

table:
	aws dynamodb create-table --endpoint-url http://localhost:8000 --cli-input-json file://dynamodb-schema.json

clean:
	aws dynamodb delete-table --endpoint-url http://localhost:8000 --table-name TestTable
