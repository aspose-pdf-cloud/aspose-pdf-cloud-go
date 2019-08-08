# ImageFooter
Represents Pdf image footer.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | Link to the document. | [optional] [default to null]
**Background** | **bool** | Sets or gets a bool value that indicates the content is stamped as background. If the value is true, the stamp content is layed at the bottom. By defalt, the value is false, the stamp content is layed at the top. | [optional] [default to null]
**HorizontalAlignment** | [***HorizontalAlignment**](HorizontalAlignment.md) | Gets or sets Horizontal alignment of stamp on the page.  | [optional] [default to null]
**Opacity** | **float64** | Gets or sets a value to indicate the stamp opacity. The value is from 0.0 to 1.0. By default the value is 1.0. | [optional] [default to null]
**Rotate** | [***Rotation**](Rotation.md) | Sets or gets the rotation of stamp content according Rotation values. Note. This property is for set angles which are multiples of 90 degrees (0, 90, 180, 270 degrees). To set arbitrary angle use RotateAngle property.  If angle set by ArbitraryAngle is not multiple of 90 then Rotate property returns Rotation.None. | [optional] [default to null]
**RotateAngle** | **float64** | Gets or sets rotate angle of stamp in degrees. This property allows to set arbitrary rotate angle.  | [optional] [default to null]
**XIndent** | **float64** | Horizontal stamp coordinate, starting from the left. | [optional] [default to null]
**YIndent** | **float64** | Vertical stamp coordinate, starting from the bottom. | [optional] [default to null]
**Zoom** | **float64** | Zooming factor of the stamp. Allows to scale stamp. | [optional] [default to null]
**FileName** | **string** | Gets or sets the file name. | [optional] [default to null]
**Width** | **float64** | Gets or sets image width. Setting this property allos to scal image horizontally. | [optional] [default to null]
**Height** | **float64** | Gets or sets image height. Setting this image allows to scale image vertically. | [optional] [default to null]
**BottomMargin** | **float64** | Gets or sets bottom margin of stamp. | [optional] [default to null]
**LeftMargin** | **float64** | Gets or sets left margin of stamp. | [optional] [default to null]
**RightMargin** | **float64** | Gets or sets right margin of stamp. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../image_footer.go)


