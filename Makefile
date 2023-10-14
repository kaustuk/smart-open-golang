populate-local-file:
	aws --endpoint-url=http://localhost:4566 s3 cp test-file/testfile s3://test-kaustuk-bucket/
