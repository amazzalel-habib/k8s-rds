// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package clouddirectoryiface provides an interface to enable mocking the Amazon CloudDirectory service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package clouddirectoryiface

import (
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory"
)

// CloudDirectoryAPI provides an interface to enable mocking the
// clouddirectory.CloudDirectory service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudDirectory.
//    func myFunc(svc clouddirectoryiface.CloudDirectoryAPI) bool {
//        // Make svc.AddFacetToObject request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := clouddirectory.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudDirectoryClient struct {
//        clouddirectoryiface.CloudDirectoryAPI
//    }
//    func (m *mockCloudDirectoryClient) AddFacetToObject(input *clouddirectory.AddFacetToObjectInput) (*clouddirectory.AddFacetToObjectOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudDirectoryClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CloudDirectoryAPI interface {
	AddFacetToObjectRequest(*clouddirectory.AddFacetToObjectInput) clouddirectory.AddFacetToObjectRequest

	ApplySchemaRequest(*clouddirectory.ApplySchemaInput) clouddirectory.ApplySchemaRequest

	AttachObjectRequest(*clouddirectory.AttachObjectInput) clouddirectory.AttachObjectRequest

	AttachPolicyRequest(*clouddirectory.AttachPolicyInput) clouddirectory.AttachPolicyRequest

	AttachToIndexRequest(*clouddirectory.AttachToIndexInput) clouddirectory.AttachToIndexRequest

	AttachTypedLinkRequest(*clouddirectory.AttachTypedLinkInput) clouddirectory.AttachTypedLinkRequest

	BatchReadRequest(*clouddirectory.BatchReadInput) clouddirectory.BatchReadRequest

	BatchWriteRequest(*clouddirectory.BatchWriteInput) clouddirectory.BatchWriteRequest

	CreateDirectoryRequest(*clouddirectory.CreateDirectoryInput) clouddirectory.CreateDirectoryRequest

	CreateFacetRequest(*clouddirectory.CreateFacetInput) clouddirectory.CreateFacetRequest

	CreateIndexRequest(*clouddirectory.CreateIndexInput) clouddirectory.CreateIndexRequest

	CreateObjectRequest(*clouddirectory.CreateObjectInput) clouddirectory.CreateObjectRequest

	CreateSchemaRequest(*clouddirectory.CreateSchemaInput) clouddirectory.CreateSchemaRequest

	CreateTypedLinkFacetRequest(*clouddirectory.CreateTypedLinkFacetInput) clouddirectory.CreateTypedLinkFacetRequest

	DeleteDirectoryRequest(*clouddirectory.DeleteDirectoryInput) clouddirectory.DeleteDirectoryRequest

	DeleteFacetRequest(*clouddirectory.DeleteFacetInput) clouddirectory.DeleteFacetRequest

	DeleteObjectRequest(*clouddirectory.DeleteObjectInput) clouddirectory.DeleteObjectRequest

	DeleteSchemaRequest(*clouddirectory.DeleteSchemaInput) clouddirectory.DeleteSchemaRequest

	DeleteTypedLinkFacetRequest(*clouddirectory.DeleteTypedLinkFacetInput) clouddirectory.DeleteTypedLinkFacetRequest

	DetachFromIndexRequest(*clouddirectory.DetachFromIndexInput) clouddirectory.DetachFromIndexRequest

	DetachObjectRequest(*clouddirectory.DetachObjectInput) clouddirectory.DetachObjectRequest

	DetachPolicyRequest(*clouddirectory.DetachPolicyInput) clouddirectory.DetachPolicyRequest

	DetachTypedLinkRequest(*clouddirectory.DetachTypedLinkInput) clouddirectory.DetachTypedLinkRequest

	DisableDirectoryRequest(*clouddirectory.DisableDirectoryInput) clouddirectory.DisableDirectoryRequest

	EnableDirectoryRequest(*clouddirectory.EnableDirectoryInput) clouddirectory.EnableDirectoryRequest

	GetAppliedSchemaVersionRequest(*clouddirectory.GetAppliedSchemaVersionInput) clouddirectory.GetAppliedSchemaVersionRequest

	GetDirectoryRequest(*clouddirectory.GetDirectoryInput) clouddirectory.GetDirectoryRequest

	GetFacetRequest(*clouddirectory.GetFacetInput) clouddirectory.GetFacetRequest

	GetObjectAttributesRequest(*clouddirectory.GetObjectAttributesInput) clouddirectory.GetObjectAttributesRequest

	GetObjectInformationRequest(*clouddirectory.GetObjectInformationInput) clouddirectory.GetObjectInformationRequest

	GetSchemaAsJsonRequest(*clouddirectory.GetSchemaAsJsonInput) clouddirectory.GetSchemaAsJsonRequest

	GetTypedLinkFacetInformationRequest(*clouddirectory.GetTypedLinkFacetInformationInput) clouddirectory.GetTypedLinkFacetInformationRequest

	ListAppliedSchemaArnsRequest(*clouddirectory.ListAppliedSchemaArnsInput) clouddirectory.ListAppliedSchemaArnsRequest

	ListAttachedIndicesRequest(*clouddirectory.ListAttachedIndicesInput) clouddirectory.ListAttachedIndicesRequest

	ListDevelopmentSchemaArnsRequest(*clouddirectory.ListDevelopmentSchemaArnsInput) clouddirectory.ListDevelopmentSchemaArnsRequest

	ListDirectoriesRequest(*clouddirectory.ListDirectoriesInput) clouddirectory.ListDirectoriesRequest

	ListFacetAttributesRequest(*clouddirectory.ListFacetAttributesInput) clouddirectory.ListFacetAttributesRequest

	ListFacetNamesRequest(*clouddirectory.ListFacetNamesInput) clouddirectory.ListFacetNamesRequest

	ListIncomingTypedLinksRequest(*clouddirectory.ListIncomingTypedLinksInput) clouddirectory.ListIncomingTypedLinksRequest

	ListIndexRequest(*clouddirectory.ListIndexInput) clouddirectory.ListIndexRequest

	ListObjectAttributesRequest(*clouddirectory.ListObjectAttributesInput) clouddirectory.ListObjectAttributesRequest

	ListObjectChildrenRequest(*clouddirectory.ListObjectChildrenInput) clouddirectory.ListObjectChildrenRequest

	ListObjectParentPathsRequest(*clouddirectory.ListObjectParentPathsInput) clouddirectory.ListObjectParentPathsRequest

	ListObjectParentsRequest(*clouddirectory.ListObjectParentsInput) clouddirectory.ListObjectParentsRequest

	ListObjectPoliciesRequest(*clouddirectory.ListObjectPoliciesInput) clouddirectory.ListObjectPoliciesRequest

	ListOutgoingTypedLinksRequest(*clouddirectory.ListOutgoingTypedLinksInput) clouddirectory.ListOutgoingTypedLinksRequest

	ListPolicyAttachmentsRequest(*clouddirectory.ListPolicyAttachmentsInput) clouddirectory.ListPolicyAttachmentsRequest

	ListPublishedSchemaArnsRequest(*clouddirectory.ListPublishedSchemaArnsInput) clouddirectory.ListPublishedSchemaArnsRequest

	ListTagsForResourceRequest(*clouddirectory.ListTagsForResourceInput) clouddirectory.ListTagsForResourceRequest

	ListTypedLinkFacetAttributesRequest(*clouddirectory.ListTypedLinkFacetAttributesInput) clouddirectory.ListTypedLinkFacetAttributesRequest

	ListTypedLinkFacetNamesRequest(*clouddirectory.ListTypedLinkFacetNamesInput) clouddirectory.ListTypedLinkFacetNamesRequest

	LookupPolicyRequest(*clouddirectory.LookupPolicyInput) clouddirectory.LookupPolicyRequest

	PublishSchemaRequest(*clouddirectory.PublishSchemaInput) clouddirectory.PublishSchemaRequest

	PutSchemaFromJsonRequest(*clouddirectory.PutSchemaFromJsonInput) clouddirectory.PutSchemaFromJsonRequest

	RemoveFacetFromObjectRequest(*clouddirectory.RemoveFacetFromObjectInput) clouddirectory.RemoveFacetFromObjectRequest

	TagResourceRequest(*clouddirectory.TagResourceInput) clouddirectory.TagResourceRequest

	UntagResourceRequest(*clouddirectory.UntagResourceInput) clouddirectory.UntagResourceRequest

	UpdateFacetRequest(*clouddirectory.UpdateFacetInput) clouddirectory.UpdateFacetRequest

	UpdateObjectAttributesRequest(*clouddirectory.UpdateObjectAttributesInput) clouddirectory.UpdateObjectAttributesRequest

	UpdateSchemaRequest(*clouddirectory.UpdateSchemaInput) clouddirectory.UpdateSchemaRequest

	UpdateTypedLinkFacetRequest(*clouddirectory.UpdateTypedLinkFacetInput) clouddirectory.UpdateTypedLinkFacetRequest

	UpgradeAppliedSchemaRequest(*clouddirectory.UpgradeAppliedSchemaInput) clouddirectory.UpgradeAppliedSchemaRequest

	UpgradePublishedSchemaRequest(*clouddirectory.UpgradePublishedSchemaInput) clouddirectory.UpgradePublishedSchemaRequest
}

var _ CloudDirectoryAPI = (*clouddirectory.CloudDirectory)(nil)