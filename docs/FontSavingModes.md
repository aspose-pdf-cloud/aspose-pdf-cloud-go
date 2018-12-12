# FontSavingModes
Enumerates modes that can be used for saving of fonts
referenced in saved PDF 
            

## Enum
Name | Type | Value | Description
------------ | ------------- | ------------- | -------------
**FontSavingModesAlwaysSaveAsWOFF** | **string** | "AlwaysSaveAsWOFF" | All referenced fonts will be saved and referenced as WOFF-fonts
**FontSavingModesAlwaysSaveAsTTF** | **string** | "AlwaysSaveAsTTF" | All referenced fonts will be saved and referenced as TTF-fonts
**FontSavingModesAlwaysSaveAsEOT** | **string** | "AlwaysSaveAsEOT" | All referenced fonts will be saved and referenced as EOT-fonts
**FontSavingModesSaveInAllFormats** | **string** | "SaveInAllFormats" | All referenced fonts will be saved (and referenced in CSS) as 3 independent files : EOT, TTH,WOFF. It increases size of output data but makes output sutable for overhelming majority of web brawsers

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../font_saving_modes.go)


