module github.com/trek10inc/awsets/cmd/awsets

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.26.0
	github.com/aws/aws-sdk-go-v2/config v0.1.1
	github.com/aws/aws-sdk-go-v2/service/accessanalyzer v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/acm v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/amplify v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigateway v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/apigatewayv2 v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/applicationautoscaling v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/appmesh v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/appsync v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/athena v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/autoscaling v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/backup v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/batch v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/budgets v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloud9 v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudformation v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudfront v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/codebuild v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/codecommit v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/codedeploy v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/codepipeline v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentity v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/configservice v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/databasemigrationservice v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dax v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/docdb v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ec2 v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecr v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ecs v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/efs v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/eks v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticache v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/elasticsearchservice v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/emr v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/firehose v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/fsx v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/glue v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/greengrass v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/iam v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/iot v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/kafka v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/kinesis v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/kms v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/lambda v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/mq v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/neptune v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/qldb v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/rds v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/redshift v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/route53 v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/s3 v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sagemaker v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/servicecatalog v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sfn v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sns v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/sqs v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssm v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/waf v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafregional v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/wafv2 v0.26.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/workspaces v0.26.0 // indirect
	github.com/emicklei/dot v0.11.0
	github.com/jmespath/go-jmespath v0.4.0
	github.com/trek10inc/awsets v0.4.1
	github.com/urfave/cli/v2 v2.2.0
	go.etcd.io/bbolt v1.3.5
)

replace github.com/trek10inc/awsets => ../..
