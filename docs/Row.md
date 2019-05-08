# Row
Represents a row of the table.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundColor** | [***Color**](Color.md) | Gets or sets the background color. | [optional] [default to null]
**Border** | [***BorderInfo**](BorderInfo.md) | Gets or sets the border. | [optional] [default to null]
**Cells** | [**[]Cell**](Cell.md) | Sets the cells of the row. | [default to null]
**DefaultCellBorder** | [***BorderInfo**](BorderInfo.md) | Gets default cell border; | [optional] [default to null]
**MinRowHeight** | **float64** | Gets height for row; | [optional] [default to null]
**FixedRowHeight** | **float64** | Gets fixed row height - row may have fixed height; | [optional] [default to null]
**IsInNewPage** | **bool** | Gets fixed row is in new page - page with this property should be printed to next page Default false; | [optional] [default to null]
**IsRowBroken** | **bool** | Gets is row can be broken between two pages | [optional] [default to null]
**DefaultCellTextState** | [***TextState**](TextState.md) | Gets or sets default text state for row cells | [optional] [default to null]
**DefaultCellPadding** | [***MarginInfo**](MarginInfo.md) | Gets or sets default margin for row cells | [optional] [default to null]
**VerticalAlignment** | [***VerticalAlignment**](VerticalAlignment.md) | Gets or sets the vertical alignment. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../row.go)


