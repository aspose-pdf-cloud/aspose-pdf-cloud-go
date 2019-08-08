# Bookmark
Provides link to bookmark.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | Link to the document. | [optional] [default to null]
**Title** | **string** | Get the Title; | [optional] [default to null]
**Italic** | **bool** | Is bookmark italic. | [optional] [default to null]
**Bold** | **bool** | Is bookmark bold. | [optional] [default to null]
**Color** | [***Color**](Color.md) | Get the color | [optional] [default to null]
**Action** | **string** | Gets or sets the action bound with the bookmark. If PageNumber is presented the action can not be specified. The action type includes: &quot;GoTo&quot;, &quot;GoToR&quot;, &quot;Launch&quot;, &quot;Named&quot;. | [optional] [default to null]
**Level** | **int32** | Gets or sets bookmark&#39;s hierarchy level. | [optional] [default to null]
**Destination** | **string** | Gets or sets bookmark&#39;s destination page. Required if action is set as string.Empty. | [optional] [default to null]
**PageDisplay** | **string** | Gets or sets the type of display bookmark&#39;s destination page. | [optional] [default to null]
**PageDisplayBottom** | **int32** | Gets or sets the bottom coordinate of page display. | [optional] [default to null]
**PageDisplayLeft** | **int32** | Gets or sets the left coordinate of page display. | [optional] [default to null]
**PageDisplayRight** | **int32** | Gets or sets the right coordinate of page display. | [optional] [default to null]
**PageDisplayTop** | **int32** | Gets or sets the top coordinate of page display. | [optional] [default to null]
**PageDisplayZoom** | **int32** | Gets or sets the zoom factor of page display. | [optional] [default to null]
**PageNumber** | **int32** | Gets or sets the number of bookmark&#39;s destination page.  | [optional] [default to null]
**RemoteFile** | **string** | Gets or sets the file (path) which is required for &quot;GoToR&quot; action of bookmark. | [optional] [default to null]
**Bookmarks** | [***Bookmarks**](Bookmarks.md) | The children bookmarks. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../bookmark.go)


