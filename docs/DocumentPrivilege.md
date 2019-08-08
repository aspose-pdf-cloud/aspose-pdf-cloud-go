# DocumentPrivilege
Represents the privileges for accessing Pdf file.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPrint** | **bool** | Sets the permission which allow print or not.  true is allow and false or not set is forbidden. | [optional] [default to null]
**AllowDegradedPrinting** | **bool** | Sets the permission which allow degraded printing or not.  true is allow and false or not set is forbidden. | [optional] [default to null]
**AllowModifyContents** | **bool** | Sets the permission which allow modify contents or not.  true is allow and false or not set is forbidden. | [optional] [default to null]
**AllowCopy** | **bool** | Sets the permission which allow copy or not.  true is allow and false or not set is forbidden. | [optional] [default to null]
**AllowModifyAnnotations** | **bool** | Sets the permission which allow modify annotations or not.  true is allow and false or not set is forbidden. | [optional] [default to null]
**AllowFillIn** | **bool** | Sets the permission which allow fill in forms or not.  true is allow and false or not set is forbidden. | [optional] [default to null]
**AllowScreenReaders** | **bool** | Sets the permission which allow screen readers or not.  true is allow and false or not set is forbidden. | [optional] [default to null]
**AllowAssembly** | **bool** | Sets the permission which allow assembly or not.  true is allow and false or not set is forbidden. | [optional] [default to null]
**PrintAllowLevel** | **int32** | Sets the print level of  document&#39;s privilege. Just as the Adobe Professional&#39;s Printing Allowed settings. 0: None. 1: Low Resolution (150 dpi). 2: High Resolution. | [optional] [default to null]
**ChangeAllowLevel** | **int32** | Sets the change level of  document&#39;s privilege. Just as the Adobe Professional&#39;s Changes Allowed settings. 0: None. 1: Inserting, Deleting and Rotating pages. 2: Filling in form fields and signing existing signature fields. 3: Commenting, filling in form fields, and signing existing signature fields. 4: Any except extracting pages. | [optional] [default to null]
**CopyAllowLevel** | **int32** | Sets the copy level of  document&#39;s privilege. Just as the Adobe Professional&#39;s permission settings. 0: None. 1: Enable text access for screen reader devices for the visually impaired. 2: Enable copying of text, images and other content. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../document_privilege.go)


