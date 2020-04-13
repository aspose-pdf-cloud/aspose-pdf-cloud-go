# Cell
Represents a cell of the table's row.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsNoBorder** | **bool** | Gets or sets the cell have border. | [optional] [default to null]
**Margin** | [***MarginInfo**](MarginInfo.md) | Gets or sets the padding. | [optional] [default to null]
**Border** | [***BorderInfo**](BorderInfo.md) | Gets or sets the border. | [optional] [default to null]
**BackgroundColor** | [***Color**](Color.md) | Gets or sets the background color. | [optional] [default to null]
**BackgroundImageFile** | **string** | Gets or sets the background image file. | [optional] [default to null]
**BackgroundImageStorageFile** | **string** | Gets or sets path of the background image file from storage. | [optional] [default to null]
**Alignment** | [***HorizontalAlignment**](HorizontalAlignment.md) | Gets or sets the alignment. | [optional] [default to null]
**DefaultCellTextState** | [***TextState**](TextState.md) | Gets or sets the default cell text state. | [optional] [default to null]
**Paragraphs** | [**[]TextRect**](TextRect.md) | Gets or sets the cell&#39;s formatted text. | [optional] [default to null]
**IsWordWrapped** | **bool** | Gets or sets the cell&#39;s text word wrapped. | [optional] [default to null]
**VerticalAlignment** | [***VerticalAlignment**](VerticalAlignment.md) | Gets or sets the vertical alignment. | [optional] [default to null]
**ColSpan** | **int32** | Gets or sets the column span. | [optional] [default to null]
**RowSpan** | **int32** | Gets or sets the row span. | [optional] [default to null]
**Width** | **float64** | Gets or sets the column width. | [optional] [default to null]
**HtmlFragment** | **string** | Gets or sets Html fragment. | [optional] [default to null]
**Images** | [**[]ImageFragment**](ImageFragment.md) | Gets or sets ImageFragment list. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../cell.go)


