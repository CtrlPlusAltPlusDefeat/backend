module backend

go 1.20

require (
	github.com/aws/aws-lambda-go v1.39.1
	github.com/aws/aws-sdk-go-v2 v1.18.0
	github.com/aws/aws-sdk-go-v2/config v1.18.21
	github.com/aws/aws-sdk-go-v2/credentials v1.13.20
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.10.21
	github.com/aws/aws-sdk-go-v2/service/apigatewaymanagementapi v1.11.8
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.19.4
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.19.6
	github.com/aws/smithy-go v1.13.5
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0

)

require (
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.13.2 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.33 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.27 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.33 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.14.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.12.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.14.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.18.9 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)
