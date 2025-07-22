# SignatureCustomAppearance
An abstract class which represents signature custom appearance object.

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FontFamilyName** | **string** | Gets/sets font family name. It should be existed in the document. Default value: Arial. | [optional] [default to null]
**FontSize** | **float64** | Gets/sets font size. Default value: 10. | [optional] [default to null]
**Rotation** | [***Rotation**](Rotation.md) | Gets or sets signature rotation. | [default to null]
**ShowContactInfo** | **bool** | Gets/sets contact info visibility. Default value: true. | [default to null]
**ShowReason** | **bool** | Gets/sets reason visibility. Default value: true. | [default to null]
**ShowLocation** | **bool** | Gets/sets location visibility. Default value: true. | [default to null]
**ContactInfoLabel** | **string** | Gets/sets contact info label. Default value: &quot;Contact&quot;. | [optional] [default to null]
**ReasonLabel** | **string** | Gets/sets reason label. Default value: &quot;Reason&quot;. | [optional] [default to null]
**LocationLabel** | **string** | Gets/sets location label. Default value: &quot;Location&quot;. | [optional] [default to null]
**DigitalSignedLabel** | **string** | Gets/sets digital signed label. Default value: &quot;Digitally signed by&quot;. | [optional] [default to null]
**DateSignedAtLabel** | **string** | Gets/sets date signed label. Default value: &quot;Date&quot;. | [optional] [default to null]
**DateTimeLocalFormat** | **string** | Gets/sets datetime local format. Default value: &quot;yyyy.MM.dd HH:mm:ss zzz&quot;. | [optional] [default to null]
**DateTimeFormat** | **string** | Gets/sets datetime format. Default value: &quot;yyyy.MM.dd HH:mm:ss&quot;. | [optional] [default to null]
**BackgroundColor** | [***Color**](Color.md) | Gets/sets background color. | [optional] [default to null]
**ForegroundColor** | [***Color**](Color.md) | Gets/sets foreground color. | [optional] [default to null]
**UseDigitalSubjectFormat** | **bool** | Gets/sets subject format usage. | [default to null]
**DigitalSubjectFormat** | [**[]SignatureSubjectNameElements**](SignatureSubjectNameElements.md) | Gets/sets subject format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)[[View Source]](../signature_custom_appearance.go)


