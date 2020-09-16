# RedactionAnnotation
Provides RedactionAnnotation.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | Link to the document. | [optional] [default to null]
**Color** | [***Color**](Color.md) | Color of the annotation. | [optional] [default to null]
**Contents** | **string** | Get the annotation content. | [optional] [default to null]
**Modified** | **string** | The date and time when the annotation was last modified. | [optional] [default to null]
**Id** | **string** | Gets ID of the annotation. | [optional] [default to null]
**Flags** | [**[]AnnotationFlags**](AnnotationFlags.md) | Gets Flags of the annotation. | [optional] [default to null]
**Name** | **string** | Gets Name of the annotation. | [optional] [default to null]
**Rect** | [***Rectangle**](Rectangle.md) | Gets Rect of the annotation. | [default to null]
**PageIndex** | **int32** | Gets PageIndex of the annotation. | [optional] [default to null]
**ZIndex** | **int32** | Gets ZIndex of the annotation. | [optional] [default to null]
**HorizontalAlignment** | [***HorizontalAlignment**](HorizontalAlignment.md) | Gets HorizontalAlignment of the annotation. | [optional] [default to null]
**VerticalAlignment** | [***VerticalAlignment**](VerticalAlignment.md) | Gets VerticalAlignment of the annotation. | [optional] [default to null]
**QuadPoint** | [**[]Point**](Point.md) | An array of 8xN numbers specifying the coordinates of content region that is intended to be removed.  | [optional] [default to null]
**FillColor** | [***Color**](Color.md) | Gets or sets color to fill annotation. | [optional] [default to null]
**BorderColor** | [***Color**](Color.md) | Gets or sets color of border which is drawn when redaction is not active. | [optional] [default to null]
**OverlayText** | **string** | Text to print on redact annotation. | [optional] [default to null]
**Repeat** | **bool** | If true overlay text will be repeated on the annotation.  | [optional] [default to null]
**TextAlignment** | [***HorizontalAlignment**](HorizontalAlignment.md) | Gets or sets. Alignment of Overlay Text. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../redaction_annotation.go)


