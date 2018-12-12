# FontEncodingRules
This enumeration defines rules which tune encoding logic

## Enum
Name | Type | Value | Description
------------ | ------------- | ------------- | -------------
**FontEncodingRulesDefault** | **string** | "Default" | Leave encoding logic "as is" - in accordance with PDF specification
**FontEncodingRulesDecreaseToUnicodePriorityLevel** | **string** | "DecreaseToUnicodePriorityLevel" | ToUnicode is a special mechanism which helps to decode input codes to unicode symbols. According to specification it must be used first of all mechanisms to get unicode symbols for specific input code. But some documents has non-standard fonts and to convert these documents correctly it may be necessary to decrease ToUnicode priority and use another mechanisms to decode input codes.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../font_encoding_rules.go)


