# ComboBoxField
Provides ComboBoxField.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**[]Link**](Link.md) | Link to the document. | [optional] [default to null]
**PartialName** | **string** | Field name. | [optional] [default to null]
**Rect** | [***Rectangle**](Rectangle.md) | Field rectangle. | [optional] [default to null]
**Value** | **string** | Field value. | [optional] [default to null]
**PageIndex** | **int32** | Page index. | [default to null]
**Height** | **float64** | Gets or sets height of the field. | [optional] [default to null]
**Width** | **float64** | Gets or sets width of the field. | [optional] [default to null]
**ZIndex** | **int32** | Z index. | [optional] [default to null]
**IsGroup** | **bool** | Is group. | [optional] [default to null]
**Parent** | [***FormField**](FormField.md) | Gets field parent. | [optional] [default to null]
**IsSharedField** | **bool** | Property for Generator support. Used when field is added to header or footer. If true, this field will created once and it&#39;s appearance will be visible on all pages of the document. If false, separated field will be created for every document page. | [optional] [default to null]
**Flags** | [**[]AnnotationFlags**](AnnotationFlags.md) | Gets Flags of the field. | [optional] [default to null]
**Color** | [***Color**](Color.md) | Color of the annotation. | [optional] [default to null]
**Contents** | **string** | Get the field content. | [optional] [default to null]
**Margin** | [***MarginInfo**](MarginInfo.md) | Gets or sets a outer margin for paragraph (for pdf generation) | [optional] [default to null]
**Highlighting** | [***LinkHighlightingMode**](LinkHighlightingMode.md) | Field highlighting mode. | [optional] [default to null]
**HorizontalAlignment** | [***HorizontalAlignment**](HorizontalAlignment.md) | Gets HorizontalAlignment of the field. | [optional] [default to null]
**VerticalAlignment** | [***VerticalAlignment**](VerticalAlignment.md) | Gets VerticalAlignment of the field. | [optional] [default to null]
**Border** | [***Border**](Border.md) | Gets or sets annotation border characteristics. | [optional] [default to null]
**MultiSelect** | **bool** | Gets or sets multiselection flag. | [optional] [default to null]
**Selected** | **int32** | Gets or sets index of selected item. Numbering of items is started from 1. | [optional] [default to null]
**Options** | [**[]Option**](Option.md) | Gets collection of options of the combobox. | [optional] [default to null]
**ActiveState** | **string** | Gets or sets current annotation appearance state. | [optional] [default to null]
**Editable** | **bool** | Gets or sets editable status of the field. | [optional] [default to null]
**SpellCheck** | **bool** | Gets or sets spellchaeck activiity status. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../combo_box_field.go)


