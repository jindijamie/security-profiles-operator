// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a repository creation template. This template is used to define the
// settings for repositories created by Amazon ECR on your behalf. For example,
// repositories created through pull through cache actions. For more information,
// see [Private repository creation templates]in the Amazon Elastic Container Registry User Guide.
//
// [Private repository creation templates]: https://docs.aws.amazon.com/AmazonECR/latest/userguide/repository-creation-templates.html
func (c *Client) CreateRepositoryCreationTemplate(ctx context.Context, params *CreateRepositoryCreationTemplateInput, optFns ...func(*Options)) (*CreateRepositoryCreationTemplateOutput, error) {
	if params == nil {
		params = &CreateRepositoryCreationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateRepositoryCreationTemplate", params, optFns, c.addOperationCreateRepositoryCreationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateRepositoryCreationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRepositoryCreationTemplateInput struct {

	// A list of enumerable strings representing the Amazon ECR repository creation
	// scenarios that this template will apply towards. The two supported scenarios are
	// PULL_THROUGH_CACHE and REPLICATION
	//
	// This member is required.
	AppliedFor []types.RCTAppliedFor

	// The repository namespace prefix to associate with the template. All
	// repositories created using this namespace prefix will have the settings defined
	// in this template applied. For example, a prefix of prod would apply to all
	// repositories beginning with prod/ . Similarly, a prefix of prod/team would
	// apply to all repositories beginning with prod/team/ .
	//
	// To apply a template to all repositories in your registry that don't have an
	// associated creation template, you can use ROOT as the prefix.
	//
	// There is always an assumed / applied to the end of the prefix. If you specify
	// ecr-public as the prefix, Amazon ECR treats that as ecr-public/ . When using a
	// pull through cache rule, the repository prefix you specify during rule creation
	// is what you should specify as your repository creation template prefix as well.
	//
	// This member is required.
	Prefix *string

	// The ARN of the role to be assumed by Amazon ECR. This role must be in the same
	// account as the registry that you are configuring. Amazon ECR will assume your
	// supplied role when the customRoleArn is specified. When this field isn't
	// specified, Amazon ECR will use the service-linked role for the repository
	// creation template.
	CustomRoleArn *string

	// A description for the repository creation template.
	Description *string

	// The encryption configuration to use for repositories created using the template.
	EncryptionConfiguration *types.EncryptionConfigurationForRepositoryCreationTemplate

	// The tag mutability setting for the repository. If this parameter is omitted,
	// the default setting of MUTABLE will be used which will allow image tags to be
	// overwritten. If IMMUTABLE is specified, all image tags within the repository
	// will be immutable which will prevent them from being overwritten.
	ImageTagMutability types.ImageTagMutability

	// The lifecycle policy to use for repositories created using the template.
	LifecyclePolicy *string

	// The repository policy to apply to repositories created using the template. A
	// repository policy is a permissions policy associated with a repository to
	// control access permissions.
	RepositoryPolicy *string

	// The metadata to apply to the repository to help you categorize and organize.
	// Each tag consists of a key and an optional value, both of which you define. Tag
	// keys can have a maximum character length of 128 characters, and tag values can
	// have a maximum length of 256 characters.
	ResourceTags []types.Tag

	noSmithyDocumentSerde
}

type CreateRepositoryCreationTemplateOutput struct {

	// The registry ID associated with the request.
	RegistryId *string

	// The details of the repository creation template associated with the request.
	RepositoryCreationTemplate *types.RepositoryCreationTemplate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateRepositoryCreationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateRepositoryCreationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateRepositoryCreationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "CreateRepositoryCreationTemplate"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpCreateRepositoryCreationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRepositoryCreationTemplate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateRepositoryCreationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "CreateRepositoryCreationTemplate",
	}
}
