package resource

type ResourceType string

const (
	Unmapped                                 ResourceType = "unmapped"
	Unnecessary                              ResourceType = "unnecessary"
	AccessAnalyzerAnalyzer                   ResourceType = "accessanalyzer/analyzer"
	AcmCertificate                           ResourceType = "acm/certificate"
	AmazonMQBroker                           ResourceType = "mq/broker"
	AmazonMQBrokerConfiguration              ResourceType = "mq/brokerconfiguration"
	AmplifyApp                               ResourceType = "amplify/app"
	AmplifyBranch                            ResourceType = "amplify/branch"
	AmplifyDomain                            ResourceType = "amplify/domain"
	ApiGatewayApiKey                         ResourceType = "apig/apikey"
	ApiGatewayAuthorizer                     ResourceType = "apig/authorizer"
	ApiGatewayBasePathMapping                ResourceType = "apig/basepathmapping"
	ApiGatewayDeployment                     ResourceType = "apig/deployment"
	ApiGatewayDocumentationPart              ResourceType = "apig/documentationpart"
	ApiGatewayDomainName                     ResourceType = "apig/domainname"
	ApiGatewayMethod                         ResourceType = "apig/method"
	ApiGatewayModel                          ResourceType = "apig/model"
	ApiGatewayRequestValidator               ResourceType = "apig/requestvalidator"
	ApiGatewayResource                       ResourceType = "apig/resource"
	ApiGatewayRestApi                        ResourceType = "apig/restapi"
	ApiGatewayVpcLink                        ResourceType = "apig/vpclink"
	ApiGatewayStage                          ResourceType = "apig/stage"
	ApiGatewayUsagePlan                      ResourceType = "apig/usageplan"
	ApiGatewayUsagePlanKey                   ResourceType = "apig/usageplankey"
	ApiGatewayV2Api                          ResourceType = "apigv2/api"
	ApiGatewayV2ApiMapping                   ResourceType = "apigv2/apimapping"
	ApiGatewayV2Authorizer                   ResourceType = "apigv2/authorizer"
	ApiGatewayV2Deployment                   ResourceType = "apigv2/deployment"
	ApiGatewayV2DomainName                   ResourceType = "apigv2/domainname"
	ApiGatewayV2Integration                  ResourceType = "apigv2/integration"
	ApiGatewayV2Model                        ResourceType = "apigv2/model"
	ApiGatewayV2Route                        ResourceType = "apigv2/route"
	ApiGatewayV2Stage                        ResourceType = "apigv2/stage"
	AppMeshMesh                              ResourceType = "appmesh/mesh"
	AppMeshRoute                             ResourceType = "appmesh/route"
	AppMeshVirtualGateway                    ResourceType = "appmesh/virtualgateway"
	AppMeshVirtualNode                       ResourceType = "appmesh/virtualnode"
	AppMeshVirtualRouter                     ResourceType = "appmesh/virtualrouter"
	AppMeshVirtualService                    ResourceType = "appmesh/virtualservice"
	AppSyncApiCache                          ResourceType = "appsync/apicache"
	AppSyncApiKey                            ResourceType = "appsync/apikey"
	AppSyncDataSource                        ResourceType = "appsync/datasource"
	AppSyncFunction                          ResourceType = "appsync/function"
	AppSyncGraphQLApi                        ResourceType = "appsync/graphqlapi"
	AppSyncGraphQLSchema                     ResourceType = "appsync/graphqlschema" //TODO: this one
	AppSyncResolver                          ResourceType = "appsync/resolver"
	ApplicationAutoScalingScalablePolicy     ResourceType = "applicationautoscaling/scalablepolicy"
	ApplicationAutoScalingScalableTarget     ResourceType = "applicationautoscaling/scalabletarget"
	AthenaDataCatalog                        ResourceType = "athena/datacatalog"
	AthenaNamedQuery                         ResourceType = "athena/namedquery"
	AthenaWorkGroup                          ResourceType = "athena/workgroup"
	AutoscalingGroup                         ResourceType = "autoscaling/group"
	AutoscalingLaunchConfig                  ResourceType = "autoscaling/launchconfig"
	AutoscalingLifecycleHook                 ResourceType = "autoscaling/lifecyclehook"
	AutoscalingPolicy                        ResourceType = "autoscaling/policy"
	AutoscalingScheduledAction               ResourceType = "autoscaling/scheduledaction"
	BackupPlan                               ResourceType = "backup/plan"
	BackupSelection                          ResourceType = "backup/selection"
	BackupVault                              ResourceType = "backup/vault"
	BatchComputeEnvironment                  ResourceType = "batch/computeenvironment"
	BatchJobDefinition                       ResourceType = "batch/jobdefinition"
	BatchJobQueue                            ResourceType = "batch/jobqueue"
	BudgetsBudget                            ResourceType = "budgets/budget"
	Cloud9Environment                        ResourceType = "cloud9/environment"
	CloudFormationStack                      ResourceType = "cloudformation/stack"
	CloudFormationStackSet                   ResourceType = "cloudformation/stackset"
	CloudFrontCachePolicy                    ResourceType = "cloudfront/cachepolicy"
	CloudFrontDistribution                   ResourceType = "cloudfront/distribution"
	CloudFrontKeyGroup                       ResourceType = "cloudfront/keygroup"
	CloudFrontOriginAccessIdentity           ResourceType = "cloudfront/originaccessidentity"
	CloudFrontOriginRequestPolicy            ResourceType = "cloudfront/originrequestpolicy"
	CloudFrontPublicKey                      ResourceType = "cloudfront/publickey"
	CloudFrontStreamingDistribution          ResourceType = "cloudfront/streamingdistribution"
	CloudtrailTrail                          ResourceType = "cloudtrail/trail"
	CloudwatchAlarm                          ResourceType = "cloudwatch/alarm"
	CloudwatchDashboard                      ResourceType = "cloudwatch/dashboard"
	CodeBuildProject                         ResourceType = "codebuild/project"
	CodeBuildSourceCredential                ResourceType = "codebuild/sourcecredential"
	CodeCommitRepository                     ResourceType = "codecommit/repository"
	CodeDeployApplication                    ResourceType = "codedeploy/application"
	CodeDeployDeploymentConfig               ResourceType = "codedeploy/deploymentconfig"
	CodeDeployDeploymentGroup                ResourceType = "codedeploy/deploymentgroup"
	CodeStarProject                          ResourceType = "codestar/project"
	CodePipelinePipeline                     ResourceType = "codepipeline/pipeline"
	CodePipelineWebhook                      ResourceType = "codepipeline/webhook"
	CognitoIdentityPool                      ResourceType = "cognito/identitypool"
	CognitoUserPool                          ResourceType = "cognito/userpool"
	CognitoUserPoolClient                    ResourceType = "cognito/userpoolclient"
	CognitoUserPoolGroup                     ResourceType = "cognito/userpoolgroup"
	CognitoUserPoolIdentityProvider          ResourceType = "cognito/userpoolidentityprovider"
	CognitoUserPoolResourceServer            ResourceType = "cognito/userpoolresourceserver"
	ConfigAggregationAuthorization           ResourceType = "config/aggregationauthorizer"
	ConfigConfigurationAggregator            ResourceType = "config/configurationaggregator"
	ConfigConfigurationRecorder              ResourceType = "config/configurationrecorder"
	ConfigConformancePack                    ResourceType = "config/conformancepack"
	ConfigDeliveryChannel                    ResourceType = "config/deliverychannel"
	ConfigOrganizationConfigRule             ResourceType = "config/organizationconfigrule"
	ConfigOrganizationConformancePack        ResourceType = "config/organizationconformancepack"
	ConfigRemediationConfiguration           ResourceType = "config/remediationconfiguration"
	ConfigRule                               ResourceType = "config/rule"
	DAXCluster                               ResourceType = "dax/cluster"
	DAXParameterGroup                        ResourceType = "dax/parametergroup"
	DAXSubnetGroup                           ResourceType = "dax/subnetgroup"
	DMSEndpoint                              ResourceType = "dms/endpoint"
	DMSReplicationInstance                   ResourceType = "dms/replicationinstance"
	DMSReplicationSubnetGroup                ResourceType = "dms/replicationsubnetgroup"
	DMSReplicationTask                       ResourceType = "dms/replicationtask"
	DocDBCluster                             ResourceType = "docdb/cluster"
	DocDBInstance                            ResourceType = "docdb/instance"
	DocDBParameterGroup                      ResourceType = "docdb/parametergroup"
	DocDBSubnetGroup                         ResourceType = "docdb/subnetgroup"
	DynamoDbBackup                           ResourceType = "ddb/backup"
	DynamoDbTable                            ResourceType = "ddb/table"
	DynamoDbStreamStream                     ResourceType = "ddbstream/stream"
	Ec2CustomerGateway                       ResourceType = "ec2/customergateway"
	Ec2DHCPOption                            ResourceType = "ec2/dhcpoption"
	Ec2Eip                                   ResourceType = "ec2/eip"
	Ec2FlowLog                               ResourceType = "ec2/flowlog"
	Ec2Image                                 ResourceType = "ec2/image"
	Ec2Instance                              ResourceType = "ec2/instance"
	Ec2InternetGateway                       ResourceType = "ec2/internetgateway"
	Ec2KeyPair                               ResourceType = "ec2/keypair"
	Ec2LaunchTemplate                        ResourceType = "ec2/launchtemplate"
	Ec2NetworkACL                            ResourceType = "ec2/networkacl"
	Ec2NetworkInterface                      ResourceType = "ec2/networkinterface"
	Ec2NatGateway                            ResourceType = "ec2/natgateway"
	Ec2Route                                 ResourceType = "ec2/route"
	Ec2RouteTable                            ResourceType = "ec2/routetable"
	Ec2SecurityGroup                         ResourceType = "ec2/securitygroup"
	Ec2Snapshot                              ResourceType = "ec2/snapshot"
	Ec2SpotFleet                             ResourceType = "ec2/spotfleet"
	Ec2Subnet                                ResourceType = "ec2/subnet"
	Ec2TransitGateway                        ResourceType = "ec2/transitgateway"
	Ec2TransitGatewayAttachment              ResourceType = "ec2/transitgatewayattachment"
	Ec2TransitGatewayRouteTable              ResourceType = "ec2/transitgatewayroutetable"
	Ec2Volume                                ResourceType = "ec2/volume"
	Ec2Vpc                                   ResourceType = "ec2/vpc"
	Ec2VpcEndpoint                           ResourceType = "ec2/vpcendpoint"
	Ec2VpcEndpointService                    ResourceType = "ec2/vpcendpointservice"
	Ec2VpcEndpointConnectionNotification     ResourceType = "ec2/vpcendpointconnectionnotification"
	Ec2VpcPeering                            ResourceType = "ec2/vpcpeering"
	Ec2VpnConnection                         ResourceType = "ec2/vpnconnection"
	Ec2VpnGateway                            ResourceType = "ec2/vpngateway"
	EcrRepository                            ResourceType = "ecr/repository"
	EcsCapacityProvider                      ResourceType = "ecs/capacityprovider"
	EcsCluster                               ResourceType = "ecs/cluster"
	EcsService                               ResourceType = "ecs/service"
	EcsTask                                  ResourceType = "ecs/task"
	EcsTaskDefinition                        ResourceType = "ecs/taskdefinition"
	EfsAccessPoint                           ResourceType = "efs/accesspoint"
	EfsFileSystem                            ResourceType = "efs/filesystem"
	EfsMountTarget                           ResourceType = "efs/mounttarget"
	EksCluster                               ResourceType = "eks/cluster"
	EksFargateProfile                        ResourceType = "eks/fargateprofile"
	EksNodeGroup                             ResourceType = "eks/nodegroup"
	ElasticacheCluster                       ResourceType = "elasticache/cluster"
	ElasticacheParameterGroup                ResourceType = "elasticache/parametergroup"
	ElasticacheSecurityGroup                 ResourceType = "elasticache/securitygroup"
	ElasticacheSnapshot                      ResourceType = "elasticache/snapshot"
	ElasticacheSubnetGroup                   ResourceType = "elasticache/subnetgroup"
	ElasticacheReplicationGroup              ResourceType = "elasticache/replicationgroup"
	ElasticsearchDomain                      ResourceType = "elasticsearch/domain"
	ElasticBeanstalkApplication              ResourceType = "elasticbeanstalk/application"
	ElasticBeanstalkEnvironment              ResourceType = "elasticbeanstalk/environment"
	ElbLoadBalancer                          ResourceType = "elb/loadbalancer"
	ElbV2LoadBalancer                        ResourceType = "elbv2/loadbalancer"
	ElbV2TargetGroup                         ResourceType = "elbv2/targetgroup"
	ElbV2Listener                            ResourceType = "elbv2/listener"
	ElbV2ListenerRule                        ResourceType = "elbv2/listenerrule"
	EmrCluster                               ResourceType = "emr/cluster"
	EmrInstanceFleetConfig                   ResourceType = "emr/instancefleetconfig"
	EmrInstanceGroupConfig                   ResourceType = "emr/instancegroupconfig"
	EmrSecurityConfiguration                 ResourceType = "emr/securityconfiguration"
	EmrStep                                  ResourceType = "emr/step"
	EventsBus                                ResourceType = "events/eventbus"
	EventsRule                               ResourceType = "events/rule"
	FirehoseDeliveryStream                   ResourceType = "firehose/stream"
	FSxBackup                                ResourceType = "fsx/backup"
	FSxFileSystem                            ResourceType = "fsx/filesystem"
	GlueClassifier                           ResourceType = "glue/classifier"
	GlueConnection                           ResourceType = "glue/connection"
	GlueCrawler                              ResourceType = "glue/crawler"
	GlueDatabase                             ResourceType = "glue/database"
	GlueJob                                  ResourceType = "glue/job"
	GlueTable                                ResourceType = "glue/table"
	GlueTrigger                              ResourceType = "glue/trigger"
	GlueWorkflow                             ResourceType = "glue/workflow"
	GreengrassConnectorDefinition            ResourceType = "greengrass/connectordefinition"
	GreengrassConnectorDefinitionVersion     ResourceType = "greengrass/connectordefinitionversion"
	GreengrassCoreDefinition                 ResourceType = "greengrass/coredefinition"
	GreengrassCoreDefinitionVersion          ResourceType = "greengrass/coredefinitionversion"
	GreengrassDeviceDefinition               ResourceType = "greengrass/devicedefinition"
	GreengrassDeviceDefinitionVersion        ResourceType = "greengrass/devicedefinitionversion"
	GreengrassFunctionDefinition             ResourceType = "greengrass/functiondefinition"
	GreengrassFunctionDefinitionVersion      ResourceType = "greengrass/functiondefinitionversion"
	GreengrassLoggerDefinition               ResourceType = "greengrass/loggerdefinition"
	GreengrassLoggerDefinitionVersion        ResourceType = "greengrass/loggerdefinitionversion"
	GreengrassResourceDefinition             ResourceType = "greengrass/resourcedefinition"
	GreengrassResourceDefinitionVersion      ResourceType = "greengrass/resourcedefinitionversion"
	GreengrassSubscriptionDefinition         ResourceType = "greengrass/subscriptiondefinition"
	GreengrassSubscriptionDefinitionVersion  ResourceType = "greengrass/subscriptiondefinitionversion"
	GreengrassGroup                          ResourceType = "greengrass/group"
	GreengrassGroupVersion                   ResourceType = "greengrass/groupversion"
	GuardDutyDetector                        ResourceType = "guardduty/detector"
	GuardDutyMember                          ResourceType = "guardduty/member"
	IamAccessKey                             ResourceType = "iam/accesskey"
	IamInstanceProfile                       ResourceType = "iam/instanceprofile"
	IamGroup                                 ResourceType = "iam/group"
	IamPolicy                                ResourceType = "iam/policy"
	IamRole                                  ResourceType = "iam/role"
	IamUser                                  ResourceType = "iam/user"
	ImageBuilderComponent                    ResourceType = "imagebuilder/component"
	ImageBuilderDistributionConfiguration    ResourceType = "imagebuilder/distributionconfiguration"
	ImageBuilderImage                        ResourceType = "imagebuilder/image"
	ImageBuilderImagePipeline                ResourceType = "imagebuilder/imagepipeline"
	ImageBuilderImageRecipe                  ResourceType = "imagebuilder/imagerecipe"
	ImageBuilderInfrastructureConfiguration  ResourceType = "imagebuilder/infrastructureconfiguration"
	IoTCACertificate                         ResourceType = "iot/cacertificate"
	IoTCertificate                           ResourceType = "iot/certificate"
	IoTPolicy                                ResourceType = "iot/policy"
	IoTThing                                 ResourceType = "iot/thing"
	IoTThingGroup                            ResourceType = "iot/thinggroup"
	IoTThingType                             ResourceType = "iot/thingtype"
	IotTopicRule                             ResourceType = "iot/topicrule"
	IotTopicRuleDestination                  ResourceType = "iot/topicruledestination"
	IoTSiteWiseGateway                       ResourceType = "iotsitewise/gateway"
	IoTSiteWiseAsset                         ResourceType = "iotsitewise/asset"
	IoTSiteWiseAssetModel                    ResourceType = "iotsitewise/assetmodel"
	KafkaCluster                             ResourceType = "kafka/cluster"
	KinesisStream                            ResourceType = "kinesis/stream"
	KmsAlias                                 ResourceType = "kms/alias"
	KmsKey                                   ResourceType = "kms/key"
	LambdaAlias                              ResourceType = "lambda/alias"
	LambdaEventSourceMapping                 ResourceType = "lambda/eventsourcemapping"
	LambdaEventInvokeConfig                  ResourceType = "lambda/eventinvokeconfig" // TODO: remove?
	LambdaFunction                           ResourceType = "lambda/function"
	LambdaLayer                              ResourceType = "lambda/layer"
	LambdaLayerVersion                       ResourceType = "lambda/layerversion"
	LambdaVersion                            ResourceType = "lambda/version"
	LogGroup                                 ResourceType = "logs/loggroup"
	LogStream                                ResourceType = "logs/logstream"
	LogMetricFilter                          ResourceType = "logs/metricfilter"
	LogQueryDefinition                       ResourceType = "logs/querydefinition"
	LogSubscriptionFilter                    ResourceType = "logs/subsciptionfilter"
	NeptuneDbCluster                         ResourceType = "neptune/dbcluster"
	NeptuneDbClusterParameterGroup           ResourceType = "neptune/dbclusterparametergroup"
	NeptuneDbClusterSnapshot                 ResourceType = "neptune/dbclustersnapshot"
	NeptuneDbInstance                        ResourceType = "neptune/dbinstance"
	NeptuneDbParameterGroup                  ResourceType = "neptune/dbparametergroup"
	NeptuneDbSubnetGroup                     ResourceType = "neptune/dbsubnetgroup"
	QLDBLedger                               ResourceType = "qldb/ledger"
	RdsDbCluster                             ResourceType = "rds/dbcluster"
	RdsDbClusterParameterGroup               ResourceType = "rds/dbclusterparametergroup"
	RdsDbClusterSnapshot                     ResourceType = "rds/dbclustersnapshot"
	RdsDbInstance                            ResourceType = "rds/dbinstance"
	RdsDbParameterGroup                      ResourceType = "rds/dbparametergroup"
	RdsDbProxy                               ResourceType = "rds/dbproxy"
	RdsDbProxyTargetGroup                    ResourceType = "rds/dbproxytargetgroup"
	RDSDbSnapshot                            ResourceType = "rds/dbsnapshot"
	RdsDbSubnetGroup                         ResourceType = "rds/dbsubnetgroup"
	RedshiftCluster                          ResourceType = "redshift/cluster"
	RedshiftParameterGroup                   ResourceType = "redshift/parametergroup"
	RedshiftSecurityGroup                    ResourceType = "redshift/securitygroup"
	RedshiftSnapshot                         ResourceType = "redshift/snapshot"
	RedshiftSubnetGroup                      ResourceType = "redshift/subnetgroup"
	Route53HealthCheck                       ResourceType = "route53/healthcheck"
	Route53HostedZone                        ResourceType = "route53/hostedzone"
	Route53RecordSet                         ResourceType = "route53/recordset"
	S3Bucket                                 ResourceType = "s3/bucket"
	SagemakerEndpoint                        ResourceType = "sagemaker/endpoint"
	SagemakerEndpointConfig                  ResourceType = "sagemaker/endpointconfig"
	SagemakerModel                           ResourceType = "sagemaker/model"
	SagemakerNotebookInstance                ResourceType = "sagemaker/notebookinstance"
	SagemakerNotebookInstanceLifecycleConfig ResourceType = "sagemaker/notebookinstancelifecycleconfig"
	SecretManagerSecret                      ResourceType = "secretmanager/secret"
	ServiceCatalogAcceptedPortfolioShare     ResourceType = "servicecatalog/acceptedportfolioshare"
	ServiceCatalogPortfolio                  ResourceType = "servicecatalog/portfolio"
	ServiceDiscoveryInstance                 ResourceType = "servicediscovery/instance"
	ServiceDiscoveryNamespace                ResourceType = "servicediscovery/namespace"
	ServiceDiscoveryService                  ResourceType = "servicediscovery/service"
	SesConfigurationSet                      ResourceType = "ses/configurationset"
	SesConfigurationSetEventDestination      ResourceType = "ses/configurationseteventdestination"
	SesReceiptFilter                         ResourceType = "ses/receiptfilter"
	SesReceiptRule                           ResourceType = "ses/receiptrule"
	SesReceiptRuleSet                        ResourceType = "ses/receiptruleset"
	SesTemplate                              ResourceType = "ses/template"
	SignerSigningProfile                     ResourceType = "signer/signingprofile"
	SnsSubscription                          ResourceType = "sns/subscription"
	SnsTopic                                 ResourceType = "sns/topic"
	SqsQueue                                 ResourceType = "sqs/queue"
	SsmAssociation                           ResourceType = "ssm/association"
	SsmDocument                              ResourceType = "ssm/document"
	SsmMaintenanceWindow                     ResourceType = "ssm/maintenancewindow"
	SsmMaintenanceWindowTask                 ResourceType = "ssm/maintenancewindowtask"
	SsmParameter                             ResourceType = "ssm/parameter"
	SsmPatchBaseline                         ResourceType = "ssm/patchbaseline"
	StepFunctionStateMachine                 ResourceType = "stepfunction/statemachine"
	TransferServer                           ResourceType = "transfer/server"
	TransferUser                             ResourceType = "transfer/user"
	WafByteMatchSet                          ResourceType = "waf/bytematchset"
	WafIpSet                                 ResourceType = "waf/ipset"
	WafRule                                  ResourceType = "waf/rule"
	WafSizeConstraintSet                     ResourceType = "waf/sizeconstraintset"
	WafSqlInjectionMatchSet                  ResourceType = "waf/sqlinjectionmatchset"
	WafWebACL                                ResourceType = "waf/webacl"
	WafXssMatchSet                           ResourceType = "waf/xssmatchset"
	Wafv2IpSet                               ResourceType = "wafv2/ipset"
	Wafv2RegexPatternSet                     ResourceType = "wafv2/regexpatternset"
	Wafv2RuleGroup                           ResourceType = "wafv2/rulegroup"
	Wafv2WebACL                              ResourceType = "wafv2/webacl"
	WafRegionalByteMatchSet                  ResourceType = "wafregional/bytematchset"
	WafRegionalGeoMatchSet                   ResourceType = "wafregional/geomatchset"
	WafRegionalIpSet                         ResourceType = "wafregional/ipset"
	WafRegionalRateBasedRule                 ResourceType = "wafregional/ratebasedrule"
	WafRegionalRegexPatternSet               ResourceType = "wafregional/regexpatternset"
	WafRegionalRule                          ResourceType = "wafregional/rule"
	WafRegionalSizeConstraint                ResourceType = "wafregional/sizeconstraintset"
	WafRegionalSqlInjectionMatchSet          ResourceType = "wafregional/sqlinjectionmatchset"
	WafRegionalWebACL                        ResourceType = "wafregional/webacl"
	WafRegionalXssMatchSet                   ResourceType = "wafregional/xssmatchset"
	WorkspacesWorkspace                      ResourceType = "workspaces/workspace"
)

func (rt ResourceType) String() string {
	return string(rt)
}
