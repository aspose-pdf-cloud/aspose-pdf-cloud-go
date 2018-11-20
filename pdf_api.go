// --------------------------------------------------------------------------------------------------------------------
// <copyright company="Aspose" file="PdfApi.cs">
//   Copyright (c) 2018 Aspose.PDF Cloud
// </copyright>
// <summary>
//   Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
// 
//  The above copyright notice and this permission notice shall be included in all
//  copies or substantial portions of the Software.
// 
//  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
//  SOFTWARE.
// </summary>
// --------------------------------------------------------------------------------------------------------------------

package asposepdfcloud

import (
	"io/ioutil"
	_url "net/url"
	"net/http"
	"strings"
	"os"
	"encoding/json"
	"fmt"
)

type PdfApiService service

/* Create Instance of PdfApiService
 @param appSid string Application SID
 @param appKey string Application Key
 @param basePath string Base service path. Set "" for default
 @return *PdfApiService */
 func NewPdfApiService(appSid string, appKey string, basePath string) *PdfApiService {
	config := NewConfiguration(appSid, appKey, basePath)
	client := NewAPIClient(config)
	return client.PdfApi
}


/* PdfApiService Delete document annotation by ID
 @param name The document name.
 @param annotationId The annotation ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) DeleteAnnotation(name string, annotationId string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete all annotations from the document
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) DeleteDocumentAnnotations(name string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete all link annotations from the document
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) DeleteDocumentLinkAnnotations(name string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete document field by name.
 @param name The document name.
 @param fieldName The field name/
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) DeleteField(name string, fieldName string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/fields/{fieldName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fieldName"+"}", fmt.Sprintf("%v", fieldName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete image from document page.
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) DeleteImage(name string, imageId string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete document page link annotation by ID
 @param name The document name.
 @param linkId The link ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) DeleteLinkAnnotation(name string, linkId string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/links/{linkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkId"+"}", fmt.Sprintf("%v", linkId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete document page by its number.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) DeletePage(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete all annotations from the page
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) DeletePageAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete all link annotations from the page
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) DeletePageLinkAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete custom document properties.
 @param name 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) 
     @param "folder" (string) 
 @return AsposeResponse*/
func (a *PdfApiService) DeleteProperties(name string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/documentproperties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Delete document property.
 @param name 
 @param propertyName 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) 
     @param "folder" (string) 
 @return AsposeResponse*/
func (a *PdfApiService) DeleteProperty(name string, propertyName string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page circle annotation by ID.
 @param name The document name.
 @param annotationId The annotation ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return CircleAnnotationResponse*/
func (a *PdfApiService) GetCircleAnnotation(name string, annotationId string, localVarOptionals map[string]interface{}) (CircleAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  CircleAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/circle/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read common document info.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return DocumentResponse*/
func (a *PdfApiService) GetDocument(name string, localVarOptionals map[string]interface{}) (DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read documant page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AnnotationsInfoResponse*/
func (a *PdfApiService) GetDocumentAnnotations(name string, localVarOptionals map[string]interface{}) (AnnotationsInfoResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AnnotationsInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document attachment info by its index.
 @param name The document name.
 @param attachmentIndex The attachment index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AttachmentResponse*/
func (a *PdfApiService) GetDocumentAttachmentByIndex(name string, attachmentIndex int32, localVarOptionals map[string]interface{}) (AttachmentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AttachmentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/attachments/{attachmentIndex}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attachmentIndex"+"}", fmt.Sprintf("%v", attachmentIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document attachments info.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AttachmentsResponse*/
func (a *PdfApiService) GetDocumentAttachments(name string, localVarOptionals map[string]interface{}) (AttachmentsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AttachmentsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/attachments"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document bookmark/bookmarks (including children).
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "bookmarkPath" (string) The bookmark path. Leave it empty if you want to get all the bookmarks in the document.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return *os.File*/
func (a *PdfApiService) GetDocumentBookmarks(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/bookmarks"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["bookmarkPath"], "string", "bookmarkPath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["bookmarkPath"].(string); localVarOk {
		localVarQueryParams.Add("bookmarkPath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document circle annotations.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return CircleAnnotationsResponse*/
func (a *PdfApiService) GetDocumentCircleAnnotations(name string, localVarOptionals map[string]interface{}) (CircleAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  CircleAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/circle"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document free text annotations.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return FreeTextAnnotationsResponse*/
func (a *PdfApiService) GetDocumentFreeTextAnnotations(name string, localVarOptionals map[string]interface{}) (FreeTextAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FreeTextAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/freetext"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document line annotations.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return LineAnnotationsResponse*/
func (a *PdfApiService) GetDocumentLineAnnotations(name string, localVarOptionals map[string]interface{}) (LineAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LineAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/line"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document polyline annotations.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return PolyLineAnnotationsResponse*/
func (a *PdfApiService) GetDocumentPolyLineAnnotations(name string, localVarOptionals map[string]interface{}) (PolyLineAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolyLineAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/polyline"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document polygon annotations.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return PolygonAnnotationsResponse*/
func (a *PdfApiService) GetDocumentPolygonAnnotations(name string, localVarOptionals map[string]interface{}) (PolygonAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolygonAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/polygon"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document properties.
 @param name 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) 
     @param "folder" (string) 
 @return DocumentPropertiesResponse*/
func (a *PdfApiService) GetDocumentProperties(name string, localVarOptionals map[string]interface{}) (DocumentPropertiesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentPropertiesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/documentproperties"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document property by name.
 @param name 
 @param propertyName 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) 
     @param "folder" (string) 
 @return DocumentPropertyResponse*/
func (a *PdfApiService) GetDocumentProperty(name string, propertyName string, localVarOptionals map[string]interface{}) (DocumentPropertyResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document square annotations.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return SquareAnnotationsResponse*/
func (a *PdfApiService) GetDocumentSquareAnnotations(name string, localVarOptionals map[string]interface{}) (SquareAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SquareAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/square"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document text annotations.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return TextAnnotationsResponse*/
func (a *PdfApiService) GetDocumentTextAnnotations(name string, localVarOptionals map[string]interface{}) (TextAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/text"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Download a specific file 
 @param path Path of the file including the file name and extension e.g. /file.ext
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "versionId" (string) File&#39;s version
     @param "storage" (string) User&#39;s storage name
 @return *os.File*/
func (a *PdfApiService) GetDownload(path string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/storage/file"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["versionId"], "string", "versionId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("path", parameterToString(path, ""))
	if localVarTempParam, localVarOk := localVarOptionals["versionId"].(string); localVarOk {
		localVarQueryParams.Add("versionId", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Download document attachment content by its index.
 @param name The document name.
 @param attachmentIndex The attachment index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return *os.File*/
func (a *PdfApiService) GetDownloadDocumentAttachmentByIndex(name string, attachmentIndex int32, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/attachments/{attachmentIndex}/download"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"attachmentIndex"+"}", fmt.Sprintf("%v", attachmentIndex), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert EPUB file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.epub)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetEpubInStorageToPdf(srcPath string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/epub"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Get document field by name.
 @param name The document name.
 @param fieldName The field name/
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return FieldResponse*/
func (a *PdfApiService) GetField(name string, fieldName string, localVarOptionals map[string]interface{}) (FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/fields/{fieldName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fieldName"+"}", fmt.Sprintf("%v", fieldName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Get document fields.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return FieldsResponse*/
func (a *PdfApiService) GetFields(name string, localVarOptionals map[string]interface{}) (FieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FieldsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page free text annotation by ID.
 @param name The document name.
 @param annotationId The annotation ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return FreeTextAnnotationResponse*/
func (a *PdfApiService) GetFreeTextAnnotation(name string, annotationId string, localVarOptionals map[string]interface{}) (FreeTextAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FreeTextAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/freetext/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert HTML file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.zip)
 @param htmlFileName Name of HTML file in ZIP.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "height" (float64) Page height
     @param "width" (float64) Page width
     @param "isLandscape" (bool) Is page landscaped
     @param "marginLeft" (float64) Page margin left
     @param "marginBottom" (float64) Page margin bottom
     @param "marginRight" (float64) Page margin right
     @param "marginTop" (float64) Page margin top
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetHtmlInStorageToPdf(srcPath string, htmlFileName string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/html"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["height"], "float64", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["width"], "float64", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["isLandscape"], "bool", "isLandscape"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginLeft"], "float64", "marginLeft"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginBottom"], "float64", "marginBottom"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginRight"], "float64", "marginRight"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginTop"], "float64", "marginTop"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	localVarQueryParams.Add("htmlFileName", parameterToString(htmlFileName, ""))
	if localVarTempParam, localVarOk := localVarOptionals["height"].(float64); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["width"].(float64); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["isLandscape"].(bool); localVarOk {
		localVarQueryParams.Add("isLandscape", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginLeft"].(float64); localVarOk {
		localVarQueryParams.Add("marginLeft", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginBottom"].(float64); localVarOk {
		localVarQueryParams.Add("marginBottom", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginRight"].(float64); localVarOk {
		localVarQueryParams.Add("marginRight", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginTop"].(float64); localVarOk {
		localVarQueryParams.Add("marginTop", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document image by ID.
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return ImageResponse*/
func (a *PdfApiService) GetImage(name string, imageId string, localVarOptionals map[string]interface{}) (ImageResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ImageResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document image in GIF format
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return *os.File*/
func (a *PdfApiService) GetImageExtractAsGif(name string, imageId string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}/extract/gif"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document image in JPEG format
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return *os.File*/
func (a *PdfApiService) GetImageExtractAsJpeg(name string, imageId string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}/extract/jpeg"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document image in PNG format
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return *os.File*/
func (a *PdfApiService) GetImageExtractAsPng(name string, imageId string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}/extract/png"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document image in TIFF format
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return *os.File*/
func (a *PdfApiService) GetImageExtractAsTiff(name string, imageId string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}/extract/tiff"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document images.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return ImagesResponse*/
func (a *PdfApiService) GetImages(name string, pageNumber int32, localVarOptionals map[string]interface{}) (ImagesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ImagesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/images"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert LaTeX file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.tex)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetLaTeXInStorageToPdf(srcPath string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/latex"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page line annotation by ID.
 @param name The document name.
 @param annotationId The annotation ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return LineAnnotationResponse*/
func (a *PdfApiService) GetLineAnnotation(name string, annotationId string, localVarOptionals map[string]interface{}) (LineAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LineAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/line/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document link annotation by ID.
 @param name The document name.
 @param linkId The link ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return LinkAnnotationResponse*/
func (a *PdfApiService) GetLinkAnnotation(name string, linkId string, localVarOptionals map[string]interface{}) (LinkAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LinkAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/links/{linkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkId"+"}", fmt.Sprintf("%v", linkId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Get the file listing of a specific folder 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "path" (string) Start with name of storage e.g. root folder &#39;/&#39;or some folder &#39;/folder1/..&#39;
     @param "storage" (string) User&#39;s storage name
 @return FilesResponse*/
func (a *PdfApiService) GetListFiles(localVarOptionals map[string]interface{}) (FilesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FilesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/storage/folder"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["path"], "string", "path"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["path"].(string); localVarOk {
		localVarQueryParams.Add("path", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert MHT file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.mht)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetMhtInStorageToPdf(srcPath string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/mht"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page info.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return DocumentPageResponse*/
func (a *PdfApiService) GetPage(name string, pageNumber int32, localVarOptionals map[string]interface{}) (DocumentPageResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentPageResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read documant page annotations. Returns only FreeTextAnnotations, TextAnnotations, other annotations will implemented next releases.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AnnotationsInfoResponse*/
func (a *PdfApiService) GetPageAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AnnotationsInfoResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AnnotationsInfoResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page circle annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return CircleAnnotationsResponse*/
func (a *PdfApiService) GetPageCircleAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (CircleAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  CircleAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/circle"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to Bmp image and return resulting file in response.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPageConvertToBmp(name string, pageNumber int32, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/bmp"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to Emf image and return resulting file in response.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPageConvertToEmf(name string, pageNumber int32, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/emf"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to Gif image and return resulting file in response.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPageConvertToGif(name string, pageNumber int32, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/gif"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to Jpeg image and return resulting file in response.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPageConvertToJpeg(name string, pageNumber int32, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/jpeg"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to Png image and return resulting file in response.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPageConvertToPng(name string, pageNumber int32, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/png"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to Tiff image  and return resulting file in response.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPageConvertToTiff(name string, pageNumber int32, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/tiff"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page free text annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return FreeTextAnnotationsResponse*/
func (a *PdfApiService) GetPageFreeTextAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (FreeTextAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FreeTextAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/freetext"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page line annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return LineAnnotationsResponse*/
func (a *PdfApiService) GetPageLineAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (LineAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LineAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/line"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page link annotation by ID.
 @param name The document name.
 @param pageNumber The page number.
 @param linkId The link ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return LinkAnnotationResponse*/
func (a *PdfApiService) GetPageLinkAnnotation(name string, pageNumber int32, linkId string, localVarOptionals map[string]interface{}) (LinkAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LinkAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/links/{linkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkId"+"}", fmt.Sprintf("%v", linkId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page link annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return LinkAnnotationsResponse*/
func (a *PdfApiService) GetPageLinkAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (LinkAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LinkAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page polyline annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return PolyLineAnnotationsResponse*/
func (a *PdfApiService) GetPagePolyLineAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (PolyLineAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolyLineAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/polyline"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page polygon annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return PolygonAnnotationsResponse*/
func (a *PdfApiService) GetPagePolygonAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (PolygonAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolygonAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/polygon"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page square annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return SquareAnnotationsResponse*/
func (a *PdfApiService) GetPageSquareAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (SquareAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SquareAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/square"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read page text items.
 @param name The document name.
 @param pageNumber Number of page (starting from 1).
 @param lLX 
 @param lLY 
 @param uRX 
 @param uRY 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "format" ([]string) List of formats for search.
     @param "regex" (string) Formats are specified as a regular expression.
     @param "splitRects" (bool) Split result fragments (default is true).
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return TextRectsResponse*/
func (a *PdfApiService) GetPageText(name string, pageNumber int32, lLX float64, lLY float64, uRX float64, uRY float64, localVarOptionals map[string]interface{}) (TextRectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextRectsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/text"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["regex"], "string", "regex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["splitRects"], "bool", "splitRects"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["format"].([]string); localVarOk {
		localVarQueryParams.Add("format", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["regex"].(string); localVarOk {
		localVarQueryParams.Add("regex", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["splitRects"].(bool); localVarOk {
		localVarQueryParams.Add("splitRects", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("LLX", parameterToString(lLX, ""))
	localVarQueryParams.Add("LLY", parameterToString(lLY, ""))
	localVarQueryParams.Add("URX", parameterToString(uRX, ""))
	localVarQueryParams.Add("URY", parameterToString(uRY, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page text annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return TextAnnotationsResponse*/
func (a *PdfApiService) GetPageTextAnnotations(name string, pageNumber int32, localVarOptionals map[string]interface{}) (TextAnnotationsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextAnnotationsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/text"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document pages info.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return DocumentPagesResponse*/
func (a *PdfApiService) GetPages(name string, localVarOptionals map[string]interface{}) (DocumentPagesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentPagesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert PCL file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.pcl)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPclInStorageToPdf(srcPath string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/pcl"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to DOC format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "addReturnToLineEnd" (bool) Add return to line end.
     @param "format" (string) Allows to specify .doc or .docx file format.
     @param "imageResolutionX" (int32) Image resolution X.
     @param "imageResolutionY" (int32) Image resolution Y.
     @param "maxDistanceBetweenTextLines" (float64) Max distance between text lines.
     @param "mode" (string) Allows to control how a PDF document is converted into a word processing document.
     @param "recognizeBullets" (bool) Recognize bullets.
     @param "relativeHorizontalProximity" (float64) Relative horizontal proximity.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToDoc(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/doc"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["addReturnToLineEnd"], "bool", "addReturnToLineEnd"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["format"], "string", "format"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["imageResolutionX"], "int32", "imageResolutionX"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["imageResolutionY"], "int32", "imageResolutionY"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxDistanceBetweenTextLines"], "float64", "maxDistanceBetweenTextLines"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["mode"], "string", "mode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["recognizeBullets"], "bool", "recognizeBullets"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["relativeHorizontalProximity"], "float64", "relativeHorizontalProximity"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["addReturnToLineEnd"].(bool); localVarOk {
		localVarQueryParams.Add("addReturnToLineEnd", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["format"].(string); localVarOk {
		localVarQueryParams.Add("format", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["imageResolutionX"].(int32); localVarOk {
		localVarQueryParams.Add("imageResolutionX", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["imageResolutionY"].(int32); localVarOk {
		localVarQueryParams.Add("imageResolutionY", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxDistanceBetweenTextLines"].(float64); localVarOk {
		localVarQueryParams.Add("maxDistanceBetweenTextLines", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["mode"].(string); localVarOk {
		localVarQueryParams.Add("mode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["recognizeBullets"].(bool); localVarOk {
		localVarQueryParams.Add("recognizeBullets", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["relativeHorizontalProximity"].(float64); localVarOk {
		localVarQueryParams.Add("relativeHorizontalProximity", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to EPUB format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "contentRecognitionMode" (string) ?roperty tunes conversion for this or that desirable method of recognition of content.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToEpub(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/epub"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["contentRecognitionMode"], "string", "contentRecognitionMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["contentRecognitionMode"].(string); localVarOk {
		localVarQueryParams.Add("contentRecognitionMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to Html format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "additionalMarginWidthInPoints" (int32) Defines width of margin that will be forcibly left around that output HTML-areas.
     @param "compressSvgGraphicsIfAny" (bool) The flag that indicates whether found SVG graphics(if any) will be compressed(zipped) into SVGZ format during saving.
     @param "convertMarkedContentToLayers" (bool) If attribute ConvertMarkedContentToLayers set to true then an all elements inside a PDF marked content (layer) will be put into an HTML div with \&quot;data-pdflayer\&quot; attribute specifying a layer name. This layer name will be extracted from optional properties of PDF marked content. If this attribute is false (by default) then no any layers will be created from PDF marked content.
     @param "defaultFontName" (string) Specifies the name of an installed font which is used to substitute any document font that is not embedded and not installed in the system. If null then default substitution font is used.
     @param "documentType" (string) Result document type.
     @param "fixedLayout" (bool) The value indicating whether that HTML is created as fixed layout.
     @param "imageResolution" (int32) Resolution for image rendering.
     @param "minimalLineWidth" (int32) This attribute sets minimal width of graphic path line. If thickness of line is less than 1px Adobe Acrobat rounds it to this value. So this attribute can be used to emulate this behavior for HTML browsers.
     @param "preventGlyphsGrouping" (bool) This attribute switch on the mode when text glyphs will not be grouped into words and strings This mode allows to keep maximum precision during positioning of glyphs on the page and it can be used for conversion documents with music notes or glyphs that should be placed separately each other. This parameter will be applied to document only when the value of FixedLayout attribute is true.
     @param "splitCssIntoPages" (bool) When multipage-mode selected(i.e &#39;SplitIntoPages&#39; is &#39;true&#39;), then this attribute defines whether should be created separate CSS-file for each result HTML page.
     @param "splitIntoPages" (bool) The flag that indicates whether each page of source document will be converted into it&#39;s own target HTML document, i.e whether result HTML will be splitted into several HTML-pages.
     @param "useZOrder" (bool) If attribute UseZORder set to true, graphics and text are added to resultant HTML document accordingly Z-order in original PDF document. If this attribute is false all graphics is put as single layer which may cause some unnecessary effects for overlapped objects.
     @param "antialiasingProcessing" (string) The parameter defines required antialiasing measures during conversion of compound background images from PDF to HTML.
     @param "cssClassNamesPrefix" (string) When PDFtoHTML converter generates result CSSs, CSS class names (something like \&quot;.stl_01 {}\&quot; ... \&quot;.stl_NN {}) are generated and used in result CSS. This property allows forcibly set class name prefix.
     @param "explicitListOfSavedPages" ([]int32) With this property You can explicitely define what pages of document should be converted. Pages in this list must have 1-based numbers. I.e. valid numbers of pages must be taken from range (1...[NumberOfPagesInConvertedDocument]) Order of appearing of pages in this list does not affect their order in result HTML page(s) - in result pages allways will go in order in which they are present in source PDF.
     @param "fontEncodingStrategy" (string) Defines encoding special rule to tune PDF decoding for current document.
     @param "fontSavingMode" (string) Defines font saving mode that will be used during saving of PDF to desirable format.
     @param "htmlMarkupGenerationMode" (string) Sometimes specific reqirments to generation of HTML markup are present. This parameter defines HTML preparing modes that can be used during conversion of PDF to HTML to match such specific requirments.
     @param "lettersPositioningMethod" (string) The mode of positioning of letters in words in result HTML.
     @param "pagesFlowTypeDependsOnViewersScreenSize" (bool) If attribute &#39;SplitOnPages&#x3D;false&#39;, than whole HTML representing all input PDF pages will be put into one big result HTML file. This flag defines whether result HTML will be generated in such way that flow of areas that represent PDF pages in result HTML will depend on screen resolution of viewer.
     @param "partsEmbeddingMode" (string) It defines whether referenced files (HTML, Fonts,Images, CSSes) will be embedded into main HTML file or will be generated as apart binary entities.
     @param "rasterImagesSavingMode" (string) Converted PDF can contain raster images This parameter defines how they should be handled during conversion of PDF to HTML.
     @param "removeEmptyAreasOnTopAndBottom" (bool) Defines whether in created HTML will be removed top and bottom empty area without any content (if any).
     @param "saveShadowedTextsAsTransparentTexts" (bool) Pdf can contain texts that are shadowed by another elements (f.e. by images) but can be selected to clipboard in Acrobat Reader (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML to mimic behaviour of Acrobat Reader (othervise such texts are usually saved as hidden, not available for copying to clipboard).
     @param "saveTransparentTexts" (bool) Pdf can contain transparent texts that can be selected to clipboard (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML.
     @param "specialFolderForAllImages" (string) The path to directory to which must be saved any images if they are encountered during saving of document as HTML. If parameter is empty or null then image files(if any) wil be saved together with other files linked to HTML It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file.
     @param "specialFolderForSvgImages" (string) The path to directory to which must be saved only SVG-images if they are encountered during saving of document as HTML. If parameter is empty or null then SVG files(if any) wil be saved together with other image-files (near to output file) or in special folder for images (if it specified in SpecialImagesFolderIfAny option). It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file.
     @param "trySaveTextUnderliningAndStrikeoutingInCss" (bool) PDF itself does not contain underlining markers for texts. It emulated with line situated under text. This option allows converter try guess that this or that line is a text&#39;s underlining and put this info into CSS instead of drawing of underlining graphically.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToHtml(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/html"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["additionalMarginWidthInPoints"], "int32", "additionalMarginWidthInPoints"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["compressSvgGraphicsIfAny"], "bool", "compressSvgGraphicsIfAny"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["convertMarkedContentToLayers"], "bool", "convertMarkedContentToLayers"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["defaultFontName"], "string", "defaultFontName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["documentType"], "string", "documentType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fixedLayout"], "bool", "fixedLayout"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["imageResolution"], "int32", "imageResolution"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["minimalLineWidth"], "int32", "minimalLineWidth"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["preventGlyphsGrouping"], "bool", "preventGlyphsGrouping"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["splitCssIntoPages"], "bool", "splitCssIntoPages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["splitIntoPages"], "bool", "splitIntoPages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["useZOrder"], "bool", "useZOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["antialiasingProcessing"], "string", "antialiasingProcessing"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["cssClassNamesPrefix"], "string", "cssClassNamesPrefix"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fontEncodingStrategy"], "string", "fontEncodingStrategy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fontSavingMode"], "string", "fontSavingMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["htmlMarkupGenerationMode"], "string", "htmlMarkupGenerationMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["lettersPositioningMethod"], "string", "lettersPositioningMethod"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pagesFlowTypeDependsOnViewersScreenSize"], "bool", "pagesFlowTypeDependsOnViewersScreenSize"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["partsEmbeddingMode"], "string", "partsEmbeddingMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["rasterImagesSavingMode"], "string", "rasterImagesSavingMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["removeEmptyAreasOnTopAndBottom"], "bool", "removeEmptyAreasOnTopAndBottom"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["saveShadowedTextsAsTransparentTexts"], "bool", "saveShadowedTextsAsTransparentTexts"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["saveTransparentTexts"], "bool", "saveTransparentTexts"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["specialFolderForAllImages"], "string", "specialFolderForAllImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["specialFolderForSvgImages"], "string", "specialFolderForSvgImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["trySaveTextUnderliningAndStrikeoutingInCss"], "bool", "trySaveTextUnderliningAndStrikeoutingInCss"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["additionalMarginWidthInPoints"].(int32); localVarOk {
		localVarQueryParams.Add("additionalMarginWidthInPoints", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["compressSvgGraphicsIfAny"].(bool); localVarOk {
		localVarQueryParams.Add("compressSvgGraphicsIfAny", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["convertMarkedContentToLayers"].(bool); localVarOk {
		localVarQueryParams.Add("convertMarkedContentToLayers", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["defaultFontName"].(string); localVarOk {
		localVarQueryParams.Add("defaultFontName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["documentType"].(string); localVarOk {
		localVarQueryParams.Add("documentType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fixedLayout"].(bool); localVarOk {
		localVarQueryParams.Add("fixedLayout", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["imageResolution"].(int32); localVarOk {
		localVarQueryParams.Add("imageResolution", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["minimalLineWidth"].(int32); localVarOk {
		localVarQueryParams.Add("minimalLineWidth", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["preventGlyphsGrouping"].(bool); localVarOk {
		localVarQueryParams.Add("preventGlyphsGrouping", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["splitCssIntoPages"].(bool); localVarOk {
		localVarQueryParams.Add("splitCssIntoPages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["splitIntoPages"].(bool); localVarOk {
		localVarQueryParams.Add("splitIntoPages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["useZOrder"].(bool); localVarOk {
		localVarQueryParams.Add("useZOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["antialiasingProcessing"].(string); localVarOk {
		localVarQueryParams.Add("antialiasingProcessing", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["cssClassNamesPrefix"].(string); localVarOk {
		localVarQueryParams.Add("cssClassNamesPrefix", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["explicitListOfSavedPages"].([]int32); localVarOk {
		localVarQueryParams.Add("explicitListOfSavedPages", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fontEncodingStrategy"].(string); localVarOk {
		localVarQueryParams.Add("fontEncodingStrategy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fontSavingMode"].(string); localVarOk {
		localVarQueryParams.Add("fontSavingMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["htmlMarkupGenerationMode"].(string); localVarOk {
		localVarQueryParams.Add("htmlMarkupGenerationMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["lettersPositioningMethod"].(string); localVarOk {
		localVarQueryParams.Add("lettersPositioningMethod", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pagesFlowTypeDependsOnViewersScreenSize"].(bool); localVarOk {
		localVarQueryParams.Add("pagesFlowTypeDependsOnViewersScreenSize", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["partsEmbeddingMode"].(string); localVarOk {
		localVarQueryParams.Add("partsEmbeddingMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["rasterImagesSavingMode"].(string); localVarOk {
		localVarQueryParams.Add("rasterImagesSavingMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["removeEmptyAreasOnTopAndBottom"].(bool); localVarOk {
		localVarQueryParams.Add("removeEmptyAreasOnTopAndBottom", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["saveShadowedTextsAsTransparentTexts"].(bool); localVarOk {
		localVarQueryParams.Add("saveShadowedTextsAsTransparentTexts", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["saveTransparentTexts"].(bool); localVarOk {
		localVarQueryParams.Add("saveTransparentTexts", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["specialFolderForAllImages"].(string); localVarOk {
		localVarQueryParams.Add("specialFolderForAllImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["specialFolderForSvgImages"].(string); localVarOk {
		localVarQueryParams.Add("specialFolderForSvgImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["trySaveTextUnderliningAndStrikeoutingInCss"].(bool); localVarOk {
		localVarQueryParams.Add("trySaveTextUnderliningAndStrikeoutingInCss", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to LaTeX format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "pagesCount" (int32) Pages count.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToLaTeX(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/latex"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["pagesCount"], "int32", "pagesCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["pagesCount"].(int32); localVarOk {
		localVarQueryParams.Add("pagesCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to MOBIXML format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToMobiXml(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/mobixml"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to PdfA format and returns resulting file in response content
 @param name The document name.
 @param type_ Type of PdfA format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToPdfA(name string, type_ string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/pdfa"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("type", parameterToString(type_, ""))
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to PPTX format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "separateImages" (bool) Separate images.
     @param "slidesAsImages" (bool) Slides as images.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToPptx(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/pptx"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["separateImages"], "bool", "separateImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["slidesAsImages"], "bool", "slidesAsImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["separateImages"].(bool); localVarOk {
		localVarQueryParams.Add("separateImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["slidesAsImages"].(bool); localVarOk {
		localVarQueryParams.Add("slidesAsImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to SVG format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "compressOutputToZipArchive" (bool) Specifies whether output will be created as one zip-archive.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToSvg(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/svg"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["compressOutputToZipArchive"], "bool", "compressOutputToZipArchive"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["compressOutputToZipArchive"].(bool); localVarOk {
		localVarQueryParams.Add("compressOutputToZipArchive", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to TIFF format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "brightness" (float64) Image brightness.
     @param "compression" (string) Tiff compression. Possible values are: LZW, CCITT4, CCITT3, RLE, None.
     @param "colorDepth" (string) Image color depth. Possible valuse are: Default, Format8bpp, Format4bpp, Format1bpp.
     @param "leftMargin" (int32) Left image margin.
     @param "rightMargin" (int32) Right image margin.
     @param "topMargin" (int32) Top image margin.
     @param "bottomMargin" (int32) Bottom image margin.
     @param "orientation" (string) Image orientation. Possible values are: None, Landscape, Portait.
     @param "skipBlankPages" (bool) Skip blank pages flag.
     @param "width" (int32) Image width.
     @param "height" (int32) Image height.
     @param "xResolution" (int32) Horizontal resolution.
     @param "yResolution" (int32) Vertical resolution.
     @param "pageIndex" (int32) Start page to export.
     @param "pageCount" (int32) Number of pages to export.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToTiff(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/tiff"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["brightness"], "float64", "brightness"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["compression"], "string", "compression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["colorDepth"], "string", "colorDepth"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["leftMargin"], "int32", "leftMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["rightMargin"], "int32", "rightMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["topMargin"], "int32", "topMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["bottomMargin"], "int32", "bottomMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["orientation"], "string", "orientation"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skipBlankPages"], "bool", "skipBlankPages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xResolution"], "int32", "xResolution"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["yResolution"], "int32", "yResolution"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageIndex"], "int32", "pageIndex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageCount"], "int32", "pageCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["brightness"].(float64); localVarOk {
		localVarQueryParams.Add("brightness", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["compression"].(string); localVarOk {
		localVarQueryParams.Add("compression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["colorDepth"].(string); localVarOk {
		localVarQueryParams.Add("colorDepth", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["leftMargin"].(int32); localVarOk {
		localVarQueryParams.Add("leftMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["rightMargin"].(int32); localVarOk {
		localVarQueryParams.Add("rightMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["topMargin"].(int32); localVarOk {
		localVarQueryParams.Add("topMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["bottomMargin"].(int32); localVarOk {
		localVarQueryParams.Add("bottomMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["orientation"].(string); localVarOk {
		localVarQueryParams.Add("orientation", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skipBlankPages"].(bool); localVarOk {
		localVarQueryParams.Add("skipBlankPages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["xResolution"].(int32); localVarOk {
		localVarQueryParams.Add("xResolution", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["yResolution"].(int32); localVarOk {
		localVarQueryParams.Add("yResolution", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageIndex"].(int32); localVarOk {
		localVarQueryParams.Add("pageIndex", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageCount"].(int32); localVarOk {
		localVarQueryParams.Add("pageCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to XLS format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "insertBlankColumnAtFirst" (bool) Insert blank column at first
     @param "minimizeTheNumberOfWorksheets" (bool) Minimize the number of worksheets
     @param "scaleFactor" (float64) Scale factor
     @param "uniformWorksheets" (bool) Uniform worksheets
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToXls(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/xls"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["insertBlankColumnAtFirst"], "bool", "insertBlankColumnAtFirst"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["minimizeTheNumberOfWorksheets"], "bool", "minimizeTheNumberOfWorksheets"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["scaleFactor"], "float64", "scaleFactor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["uniformWorksheets"], "bool", "uniformWorksheets"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["insertBlankColumnAtFirst"].(bool); localVarOk {
		localVarQueryParams.Add("insertBlankColumnAtFirst", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["minimizeTheNumberOfWorksheets"].(bool); localVarOk {
		localVarQueryParams.Add("minimizeTheNumberOfWorksheets", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["scaleFactor"].(float64); localVarOk {
		localVarQueryParams.Add("scaleFactor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["uniformWorksheets"].(bool); localVarOk {
		localVarQueryParams.Add("uniformWorksheets", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to XML format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToXml(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/xml"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to XPS format and returns resulting file in response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPdfInStorageToXps(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/xps"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page polyline annotation by ID.
 @param name The document name.
 @param annotationId The annotation ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return PolyLineAnnotationResponse*/
func (a *PdfApiService) GetPolyLineAnnotation(name string, annotationId string, localVarOptionals map[string]interface{}) (PolyLineAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolyLineAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/polyline/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page polygon annotation by ID.
 @param name The document name.
 @param annotationId The annotation ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return PolygonAnnotationResponse*/
func (a *PdfApiService) GetPolygonAnnotation(name string, annotationId string, localVarOptionals map[string]interface{}) (PolygonAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolygonAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/polygon/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert PS file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.ps)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetPsInStorageToPdf(srcPath string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/ps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page square annotation by ID.
 @param name The document name.
 @param annotationId The annotation ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return SquareAnnotationResponse*/
func (a *PdfApiService) GetSquareAnnotation(name string, annotationId string, localVarOptionals map[string]interface{}) (SquareAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SquareAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/square/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert SVG file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.svg)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "adjustPageSize" (bool) Adjust page size
     @param "height" (float64) Page height
     @param "width" (float64) Page width
     @param "isLandscape" (bool) Is page landscaped
     @param "marginLeft" (float64) Page margin left
     @param "marginBottom" (float64) Page margin bottom
     @param "marginRight" (float64) Page margin right
     @param "marginTop" (float64) Page margin top
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetSvgInStorageToPdf(srcPath string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/svg"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["adjustPageSize"], "bool", "adjustPageSize"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "float64", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["width"], "float64", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["isLandscape"], "bool", "isLandscape"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginLeft"], "float64", "marginLeft"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginBottom"], "float64", "marginBottom"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginRight"], "float64", "marginRight"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginTop"], "float64", "marginTop"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["adjustPageSize"].(bool); localVarOk {
		localVarQueryParams.Add("adjustPageSize", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(float64); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["width"].(float64); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["isLandscape"].(bool); localVarOk {
		localVarQueryParams.Add("isLandscape", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginLeft"].(float64); localVarOk {
		localVarQueryParams.Add("marginLeft", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginBottom"].(float64); localVarOk {
		localVarQueryParams.Add("marginBottom", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginRight"].(float64); localVarOk {
		localVarQueryParams.Add("marginRight", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginTop"].(float64); localVarOk {
		localVarQueryParams.Add("marginTop", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document text.
 @param name The document name.
 @param lLX 
 @param lLY 
 @param uRX 
 @param uRY 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "format" ([]string) List of formats for search.
     @param "regex" (string) Formats are specified as a regular expression.
     @param "splitRects" (bool) Split result fragments (default is true).
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return TextRectsResponse*/
func (a *PdfApiService) GetText(name string, lLX float64, lLY float64, uRX float64, uRY float64, localVarOptionals map[string]interface{}) (TextRectsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextRectsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/text"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["regex"], "string", "regex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["splitRects"], "bool", "splitRects"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["format"].([]string); localVarOk {
		localVarQueryParams.Add("format", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["regex"].(string); localVarOk {
		localVarQueryParams.Add("regex", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["splitRects"].(bool); localVarOk {
		localVarQueryParams.Add("splitRects", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	localVarQueryParams.Add("LLX", parameterToString(lLX, ""))
	localVarQueryParams.Add("LLY", parameterToString(lLY, ""))
	localVarQueryParams.Add("URX", parameterToString(uRX, ""))
	localVarQueryParams.Add("URY", parameterToString(uRY, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Read document page text annotation by ID.
 @param name The document name.
 @param annotationId The annotation ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return TextAnnotationResponse*/
func (a *PdfApiService) GetTextAnnotation(name string, annotationId string, localVarOptionals map[string]interface{}) (TextAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/text/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Verify signature document.
 @param name The document name.
 @param signName Sign name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return SignatureVerifyResponse*/
func (a *PdfApiService) GetVerifySignature(name string, signName string, localVarOptionals map[string]interface{}) (SignatureVerifyResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SignatureVerifyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/verifySignature"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("signName", parameterToString(signName, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert web page to PDF format and return resulting file in response. 
 @param url Source url
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "height" (float64) Page height
     @param "width" (float64) Page width
     @param "isLandscape" (bool) Is page landscaped
     @param "marginLeft" (float64) Page margin left
     @param "marginBottom" (float64) Page margin bottom
     @param "marginRight" (float64) Page margin right
     @param "marginTop" (float64) Page margin top
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetWebInStorageToPdf(url string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/web"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["height"], "float64", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["width"], "float64", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["isLandscape"], "bool", "isLandscape"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginLeft"], "float64", "marginLeft"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginBottom"], "float64", "marginBottom"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginRight"], "float64", "marginRight"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginTop"], "float64", "marginTop"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("url", parameterToString(url, ""))
	if localVarTempParam, localVarOk := localVarOptionals["height"].(float64); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["width"].(float64); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["isLandscape"].(bool); localVarOk {
		localVarQueryParams.Add("isLandscape", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginLeft"].(float64); localVarOk {
		localVarQueryParams.Add("marginLeft", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginBottom"].(float64); localVarOk {
		localVarQueryParams.Add("marginBottom", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginRight"].(float64); localVarOk {
		localVarQueryParams.Add("marginRight", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginTop"].(float64); localVarOk {
		localVarQueryParams.Add("marginTop", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Get number of words per document page.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return WordCountResponse*/
func (a *PdfApiService) GetWordsPerPage(name string, localVarOptionals map[string]interface{}) (WordCountResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  WordCountResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/wordCount"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and returns resulting file response content
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetXfaPdfInStorageToAcroForm(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/xfatoacroform"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert XML file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.xml)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xslFilePath" (string) Full XSL source filename (ex. /folder1/folder2/template.xsl)
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetXmlInStorageToPdf(srcPath string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/xml"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["xslFilePath"], "string", "xslFilePath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["xslFilePath"].(string); localVarOk {
		localVarQueryParams.Add("xslFilePath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert XPS file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.xps)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetXpsInStorageToPdf(srcPath string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/xps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert XslFo file (located on storage) to PDF format and return resulting file in response. 
 @param srcPath Full source filename (ex. /folder1/folder2/template.xslfo)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
 @return *os.File*/
func (a *PdfApiService) GetXslFoInStorageToPdf(srcPath string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/create/xslfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"multipart/form-data",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Append document to existing one.
 @param name The original document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "appendDocument" (AppendDocument) with the append document data.
     @param "appendFile" (string) Append file server path.
     @param "startPage" (int32) Appending start page.
     @param "endPage" (int32) Appending end page.
     @param "storage" (string) The documents storage.
     @param "folder" (string) The original document folder.
 @return DocumentResponse*/
func (a *PdfApiService) PostAppendDocument(name string, localVarOptionals map[string]interface{}) (DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/appendDocument"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["appendFile"], "string", "appendFile"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startPage"], "int32", "startPage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endPage"], "int32", "endPage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["appendFile"].(string); localVarOk {
		localVarQueryParams.Add("appendFile", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startPage"].(int32); localVarOk {
		localVarQueryParams.Add("startPage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endPage"].(int32); localVarOk {
		localVarQueryParams.Add("endPage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["appendDocument"].(AppendDocument); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Create field.
 @param name The document name.
 @param page Document page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "field" (Field) with the field data.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostCreateField(name string, page int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("page", parameterToString(page, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["field"].(Field); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Document&#39;s replace text method.
 @param name 
 @param textReplace 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) 
     @param "folder" (string) 
 @return TextReplaceResponse*/
func (a *PdfApiService) PostDocumentTextReplace(name string, textReplace TextReplaceListRequest, localVarOptionals map[string]interface{}) (TextReplaceResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextReplaceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/text/replace"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &textReplace
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Insert image to document page.
 @param name The document name.
 @param pageNumber The page number.
 @param llx Coordinate lower left X.
 @param lly Coordinate lower left Y.
 @param urx Coordinate upper right X.
 @param ury Coordinate upper right Y.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "imageFilePath" (string) Path to image file if specified. Request content is used otherwise.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "image" (*os.File) Image file.
 @return AsposeResponse*/
func (a *PdfApiService) PostInsertImage(name string, pageNumber int32, llx float64, lly float64, urx float64, ury float64, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/images"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["imageFilePath"], "string", "imageFilePath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("llx", parameterToString(llx, ""))
	localVarQueryParams.Add("lly", parameterToString(lly, ""))
	localVarQueryParams.Add("urx", parameterToString(urx, ""))
	localVarQueryParams.Add("ury", parameterToString(ury, ""))
	if localVarTempParam, localVarOk := localVarOptionals["imageFilePath"].(string); localVarOk {
		localVarQueryParams.Add("imageFilePath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["image"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Move page to new position.
 @param name The document name.
 @param pageNumber The page number.
 @param newIndex The new page position/index.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostMovePage(name string, pageNumber int32, newIndex int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/movePage"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("newIndex", parameterToString(newIndex, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Optimize document.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "options" (OptimizeOptions) The optimization options.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostOptimizeDocument(name string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/optimize"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["options"].(OptimizeOptions); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add document page circle annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param annotations The array of annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostPageCircleAnnotations(name string, pageNumber int32, annotations []CircleAnnotation, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/circle"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotations
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add document page free text annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param annotations The array of annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostPageFreeTextAnnotations(name string, pageNumber int32, annotations []FreeTextAnnotation, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/freetext"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotations
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add document page line annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param annotations The array of annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostPageLineAnnotations(name string, pageNumber int32, annotations []LineAnnotation, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/line"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotations
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add document page link annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param links Array of link anotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostPageLinkAnnotations(name string, pageNumber int32, links []LinkAnnotation, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/links"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &links
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add document page polyline annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param annotations The array of annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostPagePolyLineAnnotations(name string, pageNumber int32, annotations []PolyLineAnnotation, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/polyline"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotations
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add document page polygon annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param annotations The array of annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostPagePolygonAnnotations(name string, pageNumber int32, annotations []PolygonAnnotation, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/polygon"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotations
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add document page square annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param annotations The array of annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostPageSquareAnnotations(name string, pageNumber int32, annotations []SquareAnnotation, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/square"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotations
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add document page text annotations.
 @param name The document name.
 @param pageNumber The page number.
 @param annotations The array of annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostPageTextAnnotations(name string, pageNumber int32, annotations []TextAnnotation, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/annotations/text"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotations
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Page&#39;s replace text method.
 @param name 
 @param pageNumber 
 @param textReplaceListRequest 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) 
     @param "folder" (string) 
 @return TextReplaceResponse*/
func (a *PdfApiService) PostPageTextReplace(name string, pageNumber int32, textReplaceListRequest TextReplaceListRequest, localVarOptionals map[string]interface{}) (TextReplaceResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextReplaceResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/text/replace"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &textReplaceListRequest
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Sign document.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "signature" (Signature) Signature object containing signature data.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostSignDocument(name string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/sign"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["signature"].(Signature); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Sign page.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "signature" (Signature) Signature object containing signature data.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PostSignPage(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/sign"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["signature"].(Signature); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Split document to parts.
 @param name Document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "format" (string) Resulting documents format.
     @param "from" (int32) Start page if defined.
     @param "to" (int32) End page if defined.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return SplitResultResponse*/
func (a *PdfApiService) PostSplitDocument(name string, localVarOptionals map[string]interface{}) (SplitResultResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SplitResultResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/split"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["format"], "string", "format"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["from"], "int32", "from"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["to"], "int32", "to"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["format"].(string); localVarOk {
		localVarQueryParams.Add("format", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["from"].(int32); localVarOk {
		localVarQueryParams.Add("from", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["to"].(int32); localVarOk {
		localVarQueryParams.Add("to", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add new page to end of the document.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return DocumentPagesResponse*/
func (a *PdfApiService) PutAddNewPage(name string, localVarOptionals map[string]interface{}) (DocumentPagesResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentPagesResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add text to PDF document page.
 @param name The document name.
 @param pageNumber Number of page (starting from 1).
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "paragraph" (Paragraph) Paragraph data.
     @param "folder" (string) Document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutAddText(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/text"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["paragraph"].(Paragraph); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Replace document circle annotation
 @param name The document name.
 @param annotationId The annotation ID.
 @param annotation Annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return CircleAnnotationResponse*/
func (a *PdfApiService) PutCircleAnnotation(name string, annotationId string, annotation CircleAnnotation, localVarOptionals map[string]interface{}) (CircleAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  CircleAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/circle/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotation
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Upload a specific file 
 @param path Path where to upload including filename and extension e.g. /file.ext or /Folder 1/file.ext
 @param file File to upload
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "versionId" (string) Source file&#39;s version
     @param "storage" (string) User&#39;s storage name
 @return AsposeResponse*/
func (a *PdfApiService) PutCreate(path string, file *os.File, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/storage/file"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["versionId"], "string", "versionId"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("path", parameterToString(path, ""))
	if localVarTempParam, localVarOk := localVarOptionals["versionId"].(string); localVarOk {
		localVarQueryParams.Add("versionId", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File) = file
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Create empty document.
 @param name The new document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The new document folder.
 @return DocumentResponse*/
func (a *PdfApiService) PutCreateDocument(name string, localVarOptionals map[string]interface{}) (DocumentResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert EPUB file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.epub)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "dstFolder" (string) The destination document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutEpubInStorageToPdf(name string, srcPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/epub"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Flatten form fields in document.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutFieldsFlatten(name string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/fields/flatten"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Replace document free text annotation
 @param name The document name.
 @param annotationId The annotation ID.
 @param annotation Annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return FreeTextAnnotationResponse*/
func (a *PdfApiService) PutFreeTextAnnotation(name string, annotationId string, annotation FreeTextAnnotation, localVarOptionals map[string]interface{}) (FreeTextAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FreeTextAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/freetext/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotation
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert HTML file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.zip)
 @param htmlFileName Name of HTML file in ZIP.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "height" (float64) Page height
     @param "width" (float64) Page width
     @param "isLandscape" (bool) Is page landscaped
     @param "marginLeft" (float64) Page margin left
     @param "marginBottom" (float64) Page margin bottom
     @param "marginRight" (float64) Page margin right
     @param "marginTop" (float64) Page margin top
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutHtmlInStorageToPdf(name string, srcPath string, htmlFileName string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/html"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["height"], "float64", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["width"], "float64", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["isLandscape"], "bool", "isLandscape"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginLeft"], "float64", "marginLeft"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginBottom"], "float64", "marginBottom"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginRight"], "float64", "marginRight"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginTop"], "float64", "marginTop"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	localVarQueryParams.Add("htmlFileName", parameterToString(htmlFileName, ""))
	if localVarTempParam, localVarOk := localVarOptionals["height"].(float64); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["width"].(float64); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["isLandscape"].(bool); localVarOk {
		localVarQueryParams.Add("isLandscape", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginLeft"].(float64); localVarOk {
		localVarQueryParams.Add("marginLeft", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginBottom"].(float64); localVarOk {
		localVarQueryParams.Add("marginBottom", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginRight"].(float64); localVarOk {
		localVarQueryParams.Add("marginRight", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginTop"].(float64); localVarOk {
		localVarQueryParams.Add("marginTop", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document image in GIF format to folder
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "destFolder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutImageExtractAsGif(name string, imageId string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}/extract/gif"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destFolder"], "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destFolder"].(string); localVarOk {
		localVarQueryParams.Add("destFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document image in JPEG format to folder
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "destFolder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutImageExtractAsJpeg(name string, imageId string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}/extract/jpeg"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destFolder"], "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destFolder"].(string); localVarOk {
		localVarQueryParams.Add("destFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document image in PNG format to folder
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "destFolder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutImageExtractAsPng(name string, imageId string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}/extract/png"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destFolder"], "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destFolder"].(string); localVarOk {
		localVarQueryParams.Add("destFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document image in TIFF format to folder
 @param name The document name.
 @param imageId Image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "destFolder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutImageExtractAsTiff(name string, imageId string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}/extract/tiff"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destFolder"], "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destFolder"].(string); localVarOk {
		localVarQueryParams.Add("destFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert image file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param imageTemplates Image templates
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutImageInStorageToPdf(name string, imageTemplates ImageTemplatesRequest, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/images"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &imageTemplates
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document images in GIF format to folder.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "destFolder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutImagesExtractAsGif(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/images/extract/gif"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destFolder"], "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destFolder"].(string); localVarOk {
		localVarQueryParams.Add("destFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document images in JPEG format to folder.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) 
     @param "folder" (string) The document folder.
     @param "destFolder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutImagesExtractAsJpeg(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/images/extract/jpeg"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destFolder"], "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destFolder"].(string); localVarOk {
		localVarQueryParams.Add("destFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document images in PNG format to folder.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "destFolder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutImagesExtractAsPng(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/images/extract/png"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destFolder"], "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destFolder"].(string); localVarOk {
		localVarQueryParams.Add("destFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Extract document images in TIFF format to folder.
 @param name The document name.
 @param pageNumber The page number.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "destFolder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutImagesExtractAsTiff(name string, pageNumber int32, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/images/extract/tiff"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["destFolder"], "string", "destFolder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["destFolder"].(string); localVarOk {
		localVarQueryParams.Add("destFolder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert LaTeX file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.tex)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutLaTeXInStorageToPdf(name string, srcPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/latex"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Replace document line annotation
 @param name The document name.
 @param annotationId The annotation ID.
 @param annotation Annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return LineAnnotationResponse*/
func (a *PdfApiService) PutLineAnnotation(name string, annotationId string, annotation LineAnnotation, localVarOptionals map[string]interface{}) (LineAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LineAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/line/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotation
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Replace document page link annotations
 @param name The document name.
 @param linkId The link ID.
 @param link Link anotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return LinkAnnotationResponse*/
func (a *PdfApiService) PutLinkAnnotation(name string, linkId string, link LinkAnnotation, localVarOptionals map[string]interface{}) (LinkAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  LinkAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/links/{linkId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"linkId"+"}", fmt.Sprintf("%v", linkId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &link
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Merge a list of documents.
 @param name Resulting documen name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "mergeDocuments" (MergeDocuments) with a list of documents.
     @param "storage" (string) Resulting document storage.
     @param "folder" (string) Resulting document folder.
 @return *os.File*/
func (a *PdfApiService) PutMergeDocuments(name string, localVarOptionals map[string]interface{}) (*os.File,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  *os.File
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/merge"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["mergeDocuments"].(MergeDocuments); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert MHT file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.mht)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutMhtInStorageToPdf(name string, srcPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/mht"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add page stamp.
 @param name The document name.
 @param pageNumber The page number.
 @param stamp with data.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutPageAddStamp(name string, pageNumber int32, stamp Stamp, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/stamp"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &stamp
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to bmp image and upload resulting file to storage.
 @param name The document name.
 @param pageNumber The page number.
 @param outPath The out path of result image.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPageConvertToBmp(name string, pageNumber int32, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/bmp"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to emf image and upload resulting file to storage.
 @param name The document name.
 @param pageNumber The page number.
 @param outPath The out path of result image.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPageConvertToEmf(name string, pageNumber int32, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/emf"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to gif image and upload resulting file to storage.
 @param name The document name.
 @param pageNumber The page number.
 @param outPath The out path of result image.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPageConvertToGif(name string, pageNumber int32, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/gif"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to Jpeg image and upload resulting file to storage.
 @param name The document name.
 @param pageNumber The page number.
 @param outPath The out path of result image.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPageConvertToJpeg(name string, pageNumber int32, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/jpeg"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to png image and upload resulting file to storage.
 @param name The document name.
 @param pageNumber The page number.
 @param outPath The out path of result image.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPageConvertToPng(name string, pageNumber int32, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/png"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert document page to Tiff image and upload resulting file to storage.
 @param name The document name.
 @param pageNumber The page number.
 @param outPath The out path of result image.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "width" (int32) The converted image width.
     @param "height" (int32) The converted image height.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPageConvertToTiff(name string, pageNumber int32, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/pages/{pageNumber}/convert/tiff"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"pageNumber"+"}", fmt.Sprintf("%v", pageNumber), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert PCL file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.pcl)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPclInStorageToPdf(name string, srcPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/pcl"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to DOC format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.doc)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "addReturnToLineEnd" (bool) Add return to line end.
     @param "format" (string) Allows to specify .doc or .docx file format.
     @param "imageResolutionX" (int32) Image resolution X.
     @param "imageResolutionY" (int32) Image resolution Y.
     @param "maxDistanceBetweenTextLines" (float64) Max distance between text lines.
     @param "mode" (string) Allows to control how a PDF document is converted into a word processing document.
     @param "recognizeBullets" (bool) Recognize bullets.
     @param "relativeHorizontalProximity" (float64) Relative horizontal proximity.
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToDoc(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/doc"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["addReturnToLineEnd"], "bool", "addReturnToLineEnd"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["format"], "string", "format"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["imageResolutionX"], "int32", "imageResolutionX"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["imageResolutionY"], "int32", "imageResolutionY"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxDistanceBetweenTextLines"], "float64", "maxDistanceBetweenTextLines"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["mode"], "string", "mode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["recognizeBullets"], "bool", "recognizeBullets"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["relativeHorizontalProximity"], "float64", "relativeHorizontalProximity"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["addReturnToLineEnd"].(bool); localVarOk {
		localVarQueryParams.Add("addReturnToLineEnd", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["format"].(string); localVarOk {
		localVarQueryParams.Add("format", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["imageResolutionX"].(int32); localVarOk {
		localVarQueryParams.Add("imageResolutionX", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["imageResolutionY"].(int32); localVarOk {
		localVarQueryParams.Add("imageResolutionY", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxDistanceBetweenTextLines"].(float64); localVarOk {
		localVarQueryParams.Add("maxDistanceBetweenTextLines", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["mode"].(string); localVarOk {
		localVarQueryParams.Add("mode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["recognizeBullets"].(bool); localVarOk {
		localVarQueryParams.Add("recognizeBullets", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["relativeHorizontalProximity"].(float64); localVarOk {
		localVarQueryParams.Add("relativeHorizontalProximity", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to EPUB format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.epub)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "contentRecognitionMode" (string) ?roperty tunes conversion for this or that desirable method of recognition of content.
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToEpub(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/epub"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["contentRecognitionMode"], "string", "contentRecognitionMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["contentRecognitionMode"].(string); localVarOk {
		localVarQueryParams.Add("contentRecognitionMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to Html format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.html)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "additionalMarginWidthInPoints" (int32) Defines width of margin that will be forcibly left around that output HTML-areas.
     @param "compressSvgGraphicsIfAny" (bool) The flag that indicates whether found SVG graphics(if any) will be compressed(zipped) into SVGZ format during saving.
     @param "convertMarkedContentToLayers" (bool) If attribute ConvertMarkedContentToLayers set to true then an all elements inside a PDF marked content (layer) will be put into an HTML div with \&quot;data-pdflayer\&quot; attribute specifying a layer name. This layer name will be extracted from optional properties of PDF marked content. If this attribute is false (by default) then no any layers will be created from PDF marked content.
     @param "defaultFontName" (string) Specifies the name of an installed font which is used to substitute any document font that is not embedded and not installed in the system. If null then default substitution font is used.
     @param "documentType" (string) Result document type.
     @param "fixedLayout" (bool) The value indicating whether that HTML is created as fixed layout.
     @param "imageResolution" (int32) Resolution for image rendering.
     @param "minimalLineWidth" (int32) This attribute sets minimal width of graphic path line. If thickness of line is less than 1px Adobe Acrobat rounds it to this value. So this attribute can be used to emulate this behavior for HTML browsers.
     @param "preventGlyphsGrouping" (bool) This attribute switch on the mode when text glyphs will not be grouped into words and strings This mode allows to keep maximum precision during positioning of glyphs on the page and it can be used for conversion documents with music notes or glyphs that should be placed separately each other. This parameter will be applied to document only when the value of FixedLayout attribute is true.
     @param "splitCssIntoPages" (bool) When multipage-mode selected(i.e &#39;SplitIntoPages&#39; is &#39;true&#39;), then this attribute defines whether should be created separate CSS-file for each result HTML page.
     @param "splitIntoPages" (bool) The flag that indicates whether each page of source document will be converted into it&#39;s own target HTML document, i.e whether result HTML will be splitted into several HTML-pages.
     @param "useZOrder" (bool) If attribute UseZORder set to true, graphics and text are added to resultant HTML document accordingly Z-order in original PDF document. If this attribute is false all graphics is put as single layer which may cause some unnecessary effects for overlapped objects.
     @param "antialiasingProcessing" (string) The parameter defines required antialiasing measures during conversion of compound background images from PDF to HTML.
     @param "cssClassNamesPrefix" (string) When PDFtoHTML converter generates result CSSs, CSS class names (something like \&quot;.stl_01 {}\&quot; ... \&quot;.stl_NN {}) are generated and used in result CSS. This property allows forcibly set class name prefix.
     @param "explicitListOfSavedPages" ([]int32) With this property You can explicitely define what pages of document should be converted. Pages in this list must have 1-based numbers. I.e. valid numbers of pages must be taken from range (1...[NumberOfPagesInConvertedDocument]) Order of appearing of pages in this list does not affect their order in result HTML page(s) - in result pages allways will go in order in which they are present in source PDF.
     @param "fontEncodingStrategy" (string) Defines encoding special rule to tune PDF decoding for current document.
     @param "fontSavingMode" (string) Defines font saving mode that will be used during saving of PDF to desirable format.
     @param "htmlMarkupGenerationMode" (string) Sometimes specific reqirments to generation of HTML markup are present. This parameter defines HTML preparing modes that can be used during conversion of PDF to HTML to match such specific requirments.
     @param "lettersPositioningMethod" (string) The mode of positioning of letters in words in result HTML.
     @param "pagesFlowTypeDependsOnViewersScreenSize" (bool) If attribute &#39;SplitOnPages&#x3D;false&#39;, than whole HTML representing all input PDF pages will be put into one big result HTML file. This flag defines whether result HTML will be generated in such way that flow of areas that represent PDF pages in result HTML will depend on screen resolution of viewer.
     @param "partsEmbeddingMode" (string) It defines whether referenced files (HTML, Fonts,Images, CSSes) will be embedded into main HTML file or will be generated as apart binary entities.
     @param "rasterImagesSavingMode" (string) Converted PDF can contain raster images This parameter defines how they should be handled during conversion of PDF to HTML.
     @param "removeEmptyAreasOnTopAndBottom" (bool) Defines whether in created HTML will be removed top and bottom empty area without any content (if any).
     @param "saveShadowedTextsAsTransparentTexts" (bool) Pdf can contain texts that are shadowed by another elements (f.e. by images) but can be selected to clipboard in Acrobat Reader (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML to mimic behaviour of Acrobat Reader (othervise such texts are usually saved as hidden, not available for copying to clipboard).
     @param "saveTransparentTexts" (bool) Pdf can contain transparent texts that can be selected to clipboard (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML.
     @param "specialFolderForAllImages" (string) The path to directory to which must be saved any images if they are encountered during saving of document as HTML. If parameter is empty or null then image files(if any) wil be saved together with other files linked to HTML It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file.
     @param "specialFolderForSvgImages" (string) The path to directory to which must be saved only SVG-images if they are encountered during saving of document as HTML. If parameter is empty or null then SVG files(if any) wil be saved together with other image-files (near to output file) or in special folder for images (if it specified in SpecialImagesFolderIfAny option). It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file.
     @param "trySaveTextUnderliningAndStrikeoutingInCss" (bool) PDF itself does not contain underlining markers for texts. It emulated with line situated under text. This option allows converter try guess that this or that line is a text&#39;s underlining and put this info into CSS instead of drawing of underlining graphically.
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToHtml(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/html"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["additionalMarginWidthInPoints"], "int32", "additionalMarginWidthInPoints"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["compressSvgGraphicsIfAny"], "bool", "compressSvgGraphicsIfAny"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["convertMarkedContentToLayers"], "bool", "convertMarkedContentToLayers"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["defaultFontName"], "string", "defaultFontName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["documentType"], "string", "documentType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fixedLayout"], "bool", "fixedLayout"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["imageResolution"], "int32", "imageResolution"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["minimalLineWidth"], "int32", "minimalLineWidth"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["preventGlyphsGrouping"], "bool", "preventGlyphsGrouping"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["splitCssIntoPages"], "bool", "splitCssIntoPages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["splitIntoPages"], "bool", "splitIntoPages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["useZOrder"], "bool", "useZOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["antialiasingProcessing"], "string", "antialiasingProcessing"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["cssClassNamesPrefix"], "string", "cssClassNamesPrefix"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fontEncodingStrategy"], "string", "fontEncodingStrategy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fontSavingMode"], "string", "fontSavingMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["htmlMarkupGenerationMode"], "string", "htmlMarkupGenerationMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["lettersPositioningMethod"], "string", "lettersPositioningMethod"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pagesFlowTypeDependsOnViewersScreenSize"], "bool", "pagesFlowTypeDependsOnViewersScreenSize"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["partsEmbeddingMode"], "string", "partsEmbeddingMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["rasterImagesSavingMode"], "string", "rasterImagesSavingMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["removeEmptyAreasOnTopAndBottom"], "bool", "removeEmptyAreasOnTopAndBottom"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["saveShadowedTextsAsTransparentTexts"], "bool", "saveShadowedTextsAsTransparentTexts"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["saveTransparentTexts"], "bool", "saveTransparentTexts"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["specialFolderForAllImages"], "string", "specialFolderForAllImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["specialFolderForSvgImages"], "string", "specialFolderForSvgImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["trySaveTextUnderliningAndStrikeoutingInCss"], "bool", "trySaveTextUnderliningAndStrikeoutingInCss"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["additionalMarginWidthInPoints"].(int32); localVarOk {
		localVarQueryParams.Add("additionalMarginWidthInPoints", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["compressSvgGraphicsIfAny"].(bool); localVarOk {
		localVarQueryParams.Add("compressSvgGraphicsIfAny", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["convertMarkedContentToLayers"].(bool); localVarOk {
		localVarQueryParams.Add("convertMarkedContentToLayers", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["defaultFontName"].(string); localVarOk {
		localVarQueryParams.Add("defaultFontName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["documentType"].(string); localVarOk {
		localVarQueryParams.Add("documentType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fixedLayout"].(bool); localVarOk {
		localVarQueryParams.Add("fixedLayout", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["imageResolution"].(int32); localVarOk {
		localVarQueryParams.Add("imageResolution", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["minimalLineWidth"].(int32); localVarOk {
		localVarQueryParams.Add("minimalLineWidth", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["preventGlyphsGrouping"].(bool); localVarOk {
		localVarQueryParams.Add("preventGlyphsGrouping", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["splitCssIntoPages"].(bool); localVarOk {
		localVarQueryParams.Add("splitCssIntoPages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["splitIntoPages"].(bool); localVarOk {
		localVarQueryParams.Add("splitIntoPages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["useZOrder"].(bool); localVarOk {
		localVarQueryParams.Add("useZOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["antialiasingProcessing"].(string); localVarOk {
		localVarQueryParams.Add("antialiasingProcessing", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["cssClassNamesPrefix"].(string); localVarOk {
		localVarQueryParams.Add("cssClassNamesPrefix", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["explicitListOfSavedPages"].([]int32); localVarOk {
		localVarQueryParams.Add("explicitListOfSavedPages", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fontEncodingStrategy"].(string); localVarOk {
		localVarQueryParams.Add("fontEncodingStrategy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fontSavingMode"].(string); localVarOk {
		localVarQueryParams.Add("fontSavingMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["htmlMarkupGenerationMode"].(string); localVarOk {
		localVarQueryParams.Add("htmlMarkupGenerationMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["lettersPositioningMethod"].(string); localVarOk {
		localVarQueryParams.Add("lettersPositioningMethod", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pagesFlowTypeDependsOnViewersScreenSize"].(bool); localVarOk {
		localVarQueryParams.Add("pagesFlowTypeDependsOnViewersScreenSize", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["partsEmbeddingMode"].(string); localVarOk {
		localVarQueryParams.Add("partsEmbeddingMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["rasterImagesSavingMode"].(string); localVarOk {
		localVarQueryParams.Add("rasterImagesSavingMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["removeEmptyAreasOnTopAndBottom"].(bool); localVarOk {
		localVarQueryParams.Add("removeEmptyAreasOnTopAndBottom", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["saveShadowedTextsAsTransparentTexts"].(bool); localVarOk {
		localVarQueryParams.Add("saveShadowedTextsAsTransparentTexts", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["saveTransparentTexts"].(bool); localVarOk {
		localVarQueryParams.Add("saveTransparentTexts", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["specialFolderForAllImages"].(string); localVarOk {
		localVarQueryParams.Add("specialFolderForAllImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["specialFolderForSvgImages"].(string); localVarOk {
		localVarQueryParams.Add("specialFolderForSvgImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["trySaveTextUnderliningAndStrikeoutingInCss"].(bool); localVarOk {
		localVarQueryParams.Add("trySaveTextUnderliningAndStrikeoutingInCss", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to LaTeX format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.tex)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "pagesCount" (int32) Pages count.
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToLaTeX(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/latex"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["pagesCount"], "int32", "pagesCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["pagesCount"].(int32); localVarOk {
		localVarQueryParams.Add("pagesCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to MOBIXML format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.mobixml)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToMobiXml(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/mobixml"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to PdfA format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.pdf)
 @param type_ Type of PdfA format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToPdfA(outPath string, type_ string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/pdfa"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	localVarQueryParams.Add("type", parameterToString(type_, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to PPTX format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.pptx)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "separateImages" (bool) Separate images.
     @param "slidesAsImages" (bool) Slides as images.
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToPptx(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/pptx"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["separateImages"], "bool", "separateImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["slidesAsImages"], "bool", "slidesAsImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["separateImages"].(bool); localVarOk {
		localVarQueryParams.Add("separateImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["slidesAsImages"].(bool); localVarOk {
		localVarQueryParams.Add("slidesAsImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to SVG format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.svg)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToSvg(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/svg"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to TIFF format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.tiff)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "brightness" (float64) Image brightness.
     @param "compression" (string) Tiff compression. Possible values are: LZW, CCITT4, CCITT3, RLE, None.
     @param "colorDepth" (string) Image color depth. Possible valuse are: Default, Format8bpp, Format4bpp, Format1bpp.
     @param "leftMargin" (int32) Left image margin.
     @param "rightMargin" (int32) Right image margin.
     @param "topMargin" (int32) Top image margin.
     @param "bottomMargin" (int32) Bottom image margin.
     @param "orientation" (string) Image orientation. Possible values are: None, Landscape, Portait.
     @param "skipBlankPages" (bool) Skip blank pages flag.
     @param "width" (int32) Image width.
     @param "height" (int32) Image height.
     @param "xResolution" (int32) Horizontal resolution.
     @param "yResolution" (int32) Vertical resolution.
     @param "pageIndex" (int32) Start page to export.
     @param "pageCount" (int32) Number of pages to export.
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToTiff(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/tiff"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["brightness"], "float64", "brightness"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["compression"], "string", "compression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["colorDepth"], "string", "colorDepth"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["leftMargin"], "int32", "leftMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["rightMargin"], "int32", "rightMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["topMargin"], "int32", "topMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["bottomMargin"], "int32", "bottomMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["orientation"], "string", "orientation"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skipBlankPages"], "bool", "skipBlankPages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xResolution"], "int32", "xResolution"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["yResolution"], "int32", "yResolution"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageIndex"], "int32", "pageIndex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageCount"], "int32", "pageCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["brightness"].(float64); localVarOk {
		localVarQueryParams.Add("brightness", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["compression"].(string); localVarOk {
		localVarQueryParams.Add("compression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["colorDepth"].(string); localVarOk {
		localVarQueryParams.Add("colorDepth", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["leftMargin"].(int32); localVarOk {
		localVarQueryParams.Add("leftMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["rightMargin"].(int32); localVarOk {
		localVarQueryParams.Add("rightMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["topMargin"].(int32); localVarOk {
		localVarQueryParams.Add("topMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["bottomMargin"].(int32); localVarOk {
		localVarQueryParams.Add("bottomMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["orientation"].(string); localVarOk {
		localVarQueryParams.Add("orientation", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skipBlankPages"].(bool); localVarOk {
		localVarQueryParams.Add("skipBlankPages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["xResolution"].(int32); localVarOk {
		localVarQueryParams.Add("xResolution", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["yResolution"].(int32); localVarOk {
		localVarQueryParams.Add("yResolution", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageIndex"].(int32); localVarOk {
		localVarQueryParams.Add("pageIndex", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageCount"].(int32); localVarOk {
		localVarQueryParams.Add("pageCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to XLS format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.xls)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "insertBlankColumnAtFirst" (bool) Insert blank column at first
     @param "minimizeTheNumberOfWorksheets" (bool) Minimize the number of worksheets
     @param "scaleFactor" (float64) Scale factor
     @param "uniformWorksheets" (bool) Uniform worksheets
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToXls(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/xls"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["insertBlankColumnAtFirst"], "bool", "insertBlankColumnAtFirst"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["minimizeTheNumberOfWorksheets"], "bool", "minimizeTheNumberOfWorksheets"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["scaleFactor"], "float64", "scaleFactor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["uniformWorksheets"], "bool", "uniformWorksheets"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["insertBlankColumnAtFirst"].(bool); localVarOk {
		localVarQueryParams.Add("insertBlankColumnAtFirst", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["minimizeTheNumberOfWorksheets"].(bool); localVarOk {
		localVarQueryParams.Add("minimizeTheNumberOfWorksheets", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["scaleFactor"].(float64); localVarOk {
		localVarQueryParams.Add("scaleFactor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["uniformWorksheets"].(bool); localVarOk {
		localVarQueryParams.Add("uniformWorksheets", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to XML format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.xml)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToXml(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/xml"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (in request content) to XPS format and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.xps)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInRequestToXps(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/xps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to DOC format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.doc)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "addReturnToLineEnd" (bool) Add return to line end.
     @param "format" (string) Allows to specify .doc or .docx file format.
     @param "imageResolutionX" (int32) Image resolution X.
     @param "imageResolutionY" (int32) Image resolution Y.
     @param "maxDistanceBetweenTextLines" (float64) Max distance between text lines.
     @param "mode" (string) Allows to control how a PDF document is converted into a word processing document.
     @param "recognizeBullets" (bool) Recognize bullets.
     @param "relativeHorizontalProximity" (float64) Relative horizontal proximity.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToDoc(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/doc"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["addReturnToLineEnd"], "bool", "addReturnToLineEnd"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["format"], "string", "format"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["imageResolutionX"], "int32", "imageResolutionX"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["imageResolutionY"], "int32", "imageResolutionY"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["maxDistanceBetweenTextLines"], "float64", "maxDistanceBetweenTextLines"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["mode"], "string", "mode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["recognizeBullets"], "bool", "recognizeBullets"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["relativeHorizontalProximity"], "float64", "relativeHorizontalProximity"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["addReturnToLineEnd"].(bool); localVarOk {
		localVarQueryParams.Add("addReturnToLineEnd", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["format"].(string); localVarOk {
		localVarQueryParams.Add("format", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["imageResolutionX"].(int32); localVarOk {
		localVarQueryParams.Add("imageResolutionX", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["imageResolutionY"].(int32); localVarOk {
		localVarQueryParams.Add("imageResolutionY", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["maxDistanceBetweenTextLines"].(float64); localVarOk {
		localVarQueryParams.Add("maxDistanceBetweenTextLines", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["mode"].(string); localVarOk {
		localVarQueryParams.Add("mode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["recognizeBullets"].(bool); localVarOk {
		localVarQueryParams.Add("recognizeBullets", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["relativeHorizontalProximity"].(float64); localVarOk {
		localVarQueryParams.Add("relativeHorizontalProximity", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to EPUB format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.epub)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "contentRecognitionMode" (string) ?roperty tunes conversion for this or that desirable method of recognition of content.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToEpub(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/epub"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["contentRecognitionMode"], "string", "contentRecognitionMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["contentRecognitionMode"].(string); localVarOk {
		localVarQueryParams.Add("contentRecognitionMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to Html format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.html)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "additionalMarginWidthInPoints" (int32) Defines width of margin that will be forcibly left around that output HTML-areas.
     @param "compressSvgGraphicsIfAny" (bool) The flag that indicates whether found SVG graphics(if any) will be compressed(zipped) into SVGZ format during saving.
     @param "convertMarkedContentToLayers" (bool) If attribute ConvertMarkedContentToLayers set to true then an all elements inside a PDF marked content (layer) will be put into an HTML div with \&quot;data-pdflayer\&quot; attribute specifying a layer name. This layer name will be extracted from optional properties of PDF marked content. If this attribute is false (by default) then no any layers will be created from PDF marked content.
     @param "defaultFontName" (string) Specifies the name of an installed font which is used to substitute any document font that is not embedded and not installed in the system. If null then default substitution font is used.
     @param "documentType" (string) Result document type.
     @param "fixedLayout" (bool) The value indicating whether that HTML is created as fixed layout.
     @param "imageResolution" (int32) Resolution for image rendering.
     @param "minimalLineWidth" (int32) This attribute sets minimal width of graphic path line. If thickness of line is less than 1px Adobe Acrobat rounds it to this value. So this attribute can be used to emulate this behavior for HTML browsers.
     @param "preventGlyphsGrouping" (bool) This attribute switch on the mode when text glyphs will not be grouped into words and strings This mode allows to keep maximum precision during positioning of glyphs on the page and it can be used for conversion documents with music notes or glyphs that should be placed separately each other. This parameter will be applied to document only when the value of FixedLayout attribute is true.
     @param "splitCssIntoPages" (bool) When multipage-mode selected(i.e &#39;SplitIntoPages&#39; is &#39;true&#39;), then this attribute defines whether should be created separate CSS-file for each result HTML page.
     @param "splitIntoPages" (bool) The flag that indicates whether each page of source document will be converted into it&#39;s own target HTML document, i.e whether result HTML will be splitted into several HTML-pages.
     @param "useZOrder" (bool) If attribute UseZORder set to true, graphics and text are added to resultant HTML document accordingly Z-order in original PDF document. If this attribute is false all graphics is put as single layer which may cause some unnecessary effects for overlapped objects.
     @param "antialiasingProcessing" (string) The parameter defines required antialiasing measures during conversion of compound background images from PDF to HTML.
     @param "cssClassNamesPrefix" (string) When PDFtoHTML converter generates result CSSs, CSS class names (something like \&quot;.stl_01 {}\&quot; ... \&quot;.stl_NN {}) are generated and used in result CSS. This property allows forcibly set class name prefix.
     @param "explicitListOfSavedPages" ([]int32) With this property You can explicitely define what pages of document should be converted. Pages in this list must have 1-based numbers. I.e. valid numbers of pages must be taken from range (1...[NumberOfPagesInConvertedDocument]) Order of appearing of pages in this list does not affect their order in result HTML page(s) - in result pages allways will go in order in which they are present in source PDF.
     @param "fontEncodingStrategy" (string) Defines encoding special rule to tune PDF decoding for current document.
     @param "fontSavingMode" (string) Defines font saving mode that will be used during saving of PDF to desirable format.
     @param "htmlMarkupGenerationMode" (string) Sometimes specific reqirments to generation of HTML markup are present. This parameter defines HTML preparing modes that can be used during conversion of PDF to HTML to match such specific requirments.
     @param "lettersPositioningMethod" (string) The mode of positioning of letters in words in result HTML.
     @param "pagesFlowTypeDependsOnViewersScreenSize" (bool) If attribute &#39;SplitOnPages&#x3D;false&#39;, than whole HTML representing all input PDF pages will be put into one big result HTML file. This flag defines whether result HTML will be generated in such way that flow of areas that represent PDF pages in result HTML will depend on screen resolution of viewer.
     @param "partsEmbeddingMode" (string) It defines whether referenced files (HTML, Fonts,Images, CSSes) will be embedded into main HTML file or will be generated as apart binary entities.
     @param "rasterImagesSavingMode" (string) Converted PDF can contain raster images This parameter defines how they should be handled during conversion of PDF to HTML.
     @param "removeEmptyAreasOnTopAndBottom" (bool) Defines whether in created HTML will be removed top and bottom empty area without any content (if any).
     @param "saveShadowedTextsAsTransparentTexts" (bool) Pdf can contain texts that are shadowed by another elements (f.e. by images) but can be selected to clipboard in Acrobat Reader (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML to mimic behaviour of Acrobat Reader (othervise such texts are usually saved as hidden, not available for copying to clipboard).
     @param "saveTransparentTexts" (bool) Pdf can contain transparent texts that can be selected to clipboard (usually it happen when document contains images and OCRed texts extracted from it). This settings tells to converter whether we need save such texts as transparent selectable texts in result HTML.
     @param "specialFolderForAllImages" (string) The path to directory to which must be saved any images if they are encountered during saving of document as HTML. If parameter is empty or null then image files(if any) wil be saved together with other files linked to HTML It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file.
     @param "specialFolderForSvgImages" (string) The path to directory to which must be saved only SVG-images if they are encountered during saving of document as HTML. If parameter is empty or null then SVG files(if any) wil be saved together with other image-files (near to output file) or in special folder for images (if it specified in SpecialImagesFolderIfAny option). It does not affect anything if CustomImageSavingStrategy property was successfully used to process relevant image file.
     @param "trySaveTextUnderliningAndStrikeoutingInCss" (bool) PDF itself does not contain underlining markers for texts. It emulated with line situated under text. This option allows converter try guess that this or that line is a text&#39;s underlining and put this info into CSS instead of drawing of underlining graphically.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToHtml(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/html"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["additionalMarginWidthInPoints"], "int32", "additionalMarginWidthInPoints"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["compressSvgGraphicsIfAny"], "bool", "compressSvgGraphicsIfAny"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["convertMarkedContentToLayers"], "bool", "convertMarkedContentToLayers"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["defaultFontName"], "string", "defaultFontName"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["documentType"], "string", "documentType"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fixedLayout"], "bool", "fixedLayout"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["imageResolution"], "int32", "imageResolution"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["minimalLineWidth"], "int32", "minimalLineWidth"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["preventGlyphsGrouping"], "bool", "preventGlyphsGrouping"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["splitCssIntoPages"], "bool", "splitCssIntoPages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["splitIntoPages"], "bool", "splitIntoPages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["useZOrder"], "bool", "useZOrder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["antialiasingProcessing"], "string", "antialiasingProcessing"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["cssClassNamesPrefix"], "string", "cssClassNamesPrefix"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fontEncodingStrategy"], "string", "fontEncodingStrategy"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["fontSavingMode"], "string", "fontSavingMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["htmlMarkupGenerationMode"], "string", "htmlMarkupGenerationMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["lettersPositioningMethod"], "string", "lettersPositioningMethod"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pagesFlowTypeDependsOnViewersScreenSize"], "bool", "pagesFlowTypeDependsOnViewersScreenSize"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["partsEmbeddingMode"], "string", "partsEmbeddingMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["rasterImagesSavingMode"], "string", "rasterImagesSavingMode"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["removeEmptyAreasOnTopAndBottom"], "bool", "removeEmptyAreasOnTopAndBottom"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["saveShadowedTextsAsTransparentTexts"], "bool", "saveShadowedTextsAsTransparentTexts"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["saveTransparentTexts"], "bool", "saveTransparentTexts"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["specialFolderForAllImages"], "string", "specialFolderForAllImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["specialFolderForSvgImages"], "string", "specialFolderForSvgImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["trySaveTextUnderliningAndStrikeoutingInCss"], "bool", "trySaveTextUnderliningAndStrikeoutingInCss"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["additionalMarginWidthInPoints"].(int32); localVarOk {
		localVarQueryParams.Add("additionalMarginWidthInPoints", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["compressSvgGraphicsIfAny"].(bool); localVarOk {
		localVarQueryParams.Add("compressSvgGraphicsIfAny", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["convertMarkedContentToLayers"].(bool); localVarOk {
		localVarQueryParams.Add("convertMarkedContentToLayers", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["defaultFontName"].(string); localVarOk {
		localVarQueryParams.Add("defaultFontName", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["documentType"].(string); localVarOk {
		localVarQueryParams.Add("documentType", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fixedLayout"].(bool); localVarOk {
		localVarQueryParams.Add("fixedLayout", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["imageResolution"].(int32); localVarOk {
		localVarQueryParams.Add("imageResolution", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["minimalLineWidth"].(int32); localVarOk {
		localVarQueryParams.Add("minimalLineWidth", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["preventGlyphsGrouping"].(bool); localVarOk {
		localVarQueryParams.Add("preventGlyphsGrouping", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["splitCssIntoPages"].(bool); localVarOk {
		localVarQueryParams.Add("splitCssIntoPages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["splitIntoPages"].(bool); localVarOk {
		localVarQueryParams.Add("splitIntoPages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["useZOrder"].(bool); localVarOk {
		localVarQueryParams.Add("useZOrder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["antialiasingProcessing"].(string); localVarOk {
		localVarQueryParams.Add("antialiasingProcessing", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["cssClassNamesPrefix"].(string); localVarOk {
		localVarQueryParams.Add("cssClassNamesPrefix", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["explicitListOfSavedPages"].([]int32); localVarOk {
		localVarQueryParams.Add("explicitListOfSavedPages", parameterToString(localVarTempParam, "multi"))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fontEncodingStrategy"].(string); localVarOk {
		localVarQueryParams.Add("fontEncodingStrategy", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["fontSavingMode"].(string); localVarOk {
		localVarQueryParams.Add("fontSavingMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["htmlMarkupGenerationMode"].(string); localVarOk {
		localVarQueryParams.Add("htmlMarkupGenerationMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["lettersPositioningMethod"].(string); localVarOk {
		localVarQueryParams.Add("lettersPositioningMethod", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pagesFlowTypeDependsOnViewersScreenSize"].(bool); localVarOk {
		localVarQueryParams.Add("pagesFlowTypeDependsOnViewersScreenSize", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["partsEmbeddingMode"].(string); localVarOk {
		localVarQueryParams.Add("partsEmbeddingMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["rasterImagesSavingMode"].(string); localVarOk {
		localVarQueryParams.Add("rasterImagesSavingMode", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["removeEmptyAreasOnTopAndBottom"].(bool); localVarOk {
		localVarQueryParams.Add("removeEmptyAreasOnTopAndBottom", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["saveShadowedTextsAsTransparentTexts"].(bool); localVarOk {
		localVarQueryParams.Add("saveShadowedTextsAsTransparentTexts", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["saveTransparentTexts"].(bool); localVarOk {
		localVarQueryParams.Add("saveTransparentTexts", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["specialFolderForAllImages"].(string); localVarOk {
		localVarQueryParams.Add("specialFolderForAllImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["specialFolderForSvgImages"].(string); localVarOk {
		localVarQueryParams.Add("specialFolderForSvgImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["trySaveTextUnderliningAndStrikeoutingInCss"].(bool); localVarOk {
		localVarQueryParams.Add("trySaveTextUnderliningAndStrikeoutingInCss", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to LaTeX format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.tex)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "pagesCount" (int32) Pages count.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToLaTeX(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/latex"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["pagesCount"], "int32", "pagesCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["pagesCount"].(int32); localVarOk {
		localVarQueryParams.Add("pagesCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to MOBIXML format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.mobixml)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToMobiXml(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/mobixml"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to PdfA format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.pdf)
 @param type_ Type of PdfA format.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToPdfA(name string, outPath string, type_ string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/pdfa"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	localVarQueryParams.Add("type", parameterToString(type_, ""))
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to PPTX format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.pptx)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "separateImages" (bool) Separate images.
     @param "slidesAsImages" (bool) Slides as images.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToPptx(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/pptx"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["separateImages"], "bool", "separateImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["slidesAsImages"], "bool", "slidesAsImages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["separateImages"].(bool); localVarOk {
		localVarQueryParams.Add("separateImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["slidesAsImages"].(bool); localVarOk {
		localVarQueryParams.Add("slidesAsImages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to SVG format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.svg)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToSvg(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/svg"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to TIFF format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.tiff)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "brightness" (float64) Image brightness.
     @param "compression" (string) Tiff compression. Possible values are: LZW, CCITT4, CCITT3, RLE, None.
     @param "colorDepth" (string) Image color depth. Possible valuse are: Default, Format8bpp, Format4bpp, Format1bpp.
     @param "leftMargin" (int32) Left image margin.
     @param "rightMargin" (int32) Right image margin.
     @param "topMargin" (int32) Top image margin.
     @param "bottomMargin" (int32) Bottom image margin.
     @param "orientation" (string) Image orientation. Possible values are: None, Landscape, Portait.
     @param "skipBlankPages" (bool) Skip blank pages flag.
     @param "width" (int32) Image width.
     @param "height" (int32) Image height.
     @param "xResolution" (int32) Horizontal resolution.
     @param "yResolution" (int32) Vertical resolution.
     @param "pageIndex" (int32) Start page to export.
     @param "pageCount" (int32) Number of pages to export.
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToTiff(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/tiff"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["brightness"], "float64", "brightness"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["compression"], "string", "compression"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["colorDepth"], "string", "colorDepth"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["leftMargin"], "int32", "leftMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["rightMargin"], "int32", "rightMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["topMargin"], "int32", "topMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["bottomMargin"], "int32", "bottomMargin"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["orientation"], "string", "orientation"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["skipBlankPages"], "bool", "skipBlankPages"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["width"], "int32", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "int32", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["xResolution"], "int32", "xResolution"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["yResolution"], "int32", "yResolution"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageIndex"], "int32", "pageIndex"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageCount"], "int32", "pageCount"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["brightness"].(float64); localVarOk {
		localVarQueryParams.Add("brightness", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["compression"].(string); localVarOk {
		localVarQueryParams.Add("compression", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["colorDepth"].(string); localVarOk {
		localVarQueryParams.Add("colorDepth", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["leftMargin"].(int32); localVarOk {
		localVarQueryParams.Add("leftMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["rightMargin"].(int32); localVarOk {
		localVarQueryParams.Add("rightMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["topMargin"].(int32); localVarOk {
		localVarQueryParams.Add("topMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["bottomMargin"].(int32); localVarOk {
		localVarQueryParams.Add("bottomMargin", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["orientation"].(string); localVarOk {
		localVarQueryParams.Add("orientation", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["skipBlankPages"].(bool); localVarOk {
		localVarQueryParams.Add("skipBlankPages", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["width"].(int32); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(int32); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["xResolution"].(int32); localVarOk {
		localVarQueryParams.Add("xResolution", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["yResolution"].(int32); localVarOk {
		localVarQueryParams.Add("yResolution", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageIndex"].(int32); localVarOk {
		localVarQueryParams.Add("pageIndex", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageCount"].(int32); localVarOk {
		localVarQueryParams.Add("pageCount", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to XLS format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.xls)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "insertBlankColumnAtFirst" (bool) Insert blank column at first
     @param "minimizeTheNumberOfWorksheets" (bool) Minimize the number of worksheets
     @param "scaleFactor" (float64) Scale factor
     @param "uniformWorksheets" (bool) Uniform worksheets
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToXls(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/xls"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["insertBlankColumnAtFirst"], "bool", "insertBlankColumnAtFirst"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["minimizeTheNumberOfWorksheets"], "bool", "minimizeTheNumberOfWorksheets"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["scaleFactor"], "float64", "scaleFactor"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["uniformWorksheets"], "bool", "uniformWorksheets"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["insertBlankColumnAtFirst"].(bool); localVarOk {
		localVarQueryParams.Add("insertBlankColumnAtFirst", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["minimizeTheNumberOfWorksheets"].(bool); localVarOk {
		localVarQueryParams.Add("minimizeTheNumberOfWorksheets", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["scaleFactor"].(float64); localVarOk {
		localVarQueryParams.Add("scaleFactor", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["uniformWorksheets"].(bool); localVarOk {
		localVarQueryParams.Add("uniformWorksheets", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to XML format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.xml)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToXml(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/xml"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document (located on storage) to XPS format and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.xps)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPdfInStorageToXps(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/xps"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Replace document polyline annotation
 @param name The document name.
 @param annotationId The annotation ID.
 @param annotation Annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return PolyLineAnnotationResponse*/
func (a *PdfApiService) PutPolyLineAnnotation(name string, annotationId string, annotation PolyLineAnnotation, localVarOptionals map[string]interface{}) (PolyLineAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolyLineAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/polyline/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotation
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Replace document polygon annotation
 @param name The document name.
 @param annotationId The annotation ID.
 @param annotation Annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return PolygonAnnotationResponse*/
func (a *PdfApiService) PutPolygonAnnotation(name string, annotationId string, annotation PolygonAnnotation, localVarOptionals map[string]interface{}) (PolygonAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  PolygonAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/polygon/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotation
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Update privilege document.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "privileges" (DocumentPrivilege) Document privileges. 
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return AsposeResponse*/
func (a *PdfApiService) PutPrivileges(name string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/privileges"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["privileges"].(DocumentPrivilege); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert PS file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.ps)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutPsInStorageToPdf(name string, srcPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/ps"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Replace document image.
 @param name The document name.
 @param imageId The image ID.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "imageFilePath" (string) Path to image file if specified. Request content is used otherwise.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "image" (*os.File) Image file.
 @return ImageResponse*/
func (a *PdfApiService) PutReplaceImage(name string, imageId string, localVarOptionals map[string]interface{}) (ImageResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  ImageResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/images/{imageId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"imageId"+"}", fmt.Sprintf("%v", imageId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["imageFilePath"], "string", "imageFilePath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["imageFilePath"].(string); localVarOk {
		localVarQueryParams.Add("imageFilePath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["image"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Create searchable PDF document. Generate OCR layer for images in input PDF document.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
     @param "lang" (string) language for OCR engine. Possible values: eng, ara, bel, ben, bul, ces, dan, deu, ell, fin, fra, heb, hin, ind, isl, ita, jpn, kor, nld, nor, pol, por, ron, rus, spa, swe, tha, tur, ukr, vie, chi_sim, chi_tra or thier combination e.g. eng,rus 
 @return AsposeResponse*/
func (a *PdfApiService) PutSearchableDocument(name string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/ocr"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["lang"], "string", "lang"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["lang"].(string); localVarOk {
		localVarQueryParams.Add("lang", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Add/update document property.
 @param name 
 @param propertyName 
 @param value 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) 
     @param "folder" (string) 
 @return DocumentPropertyResponse*/
func (a *PdfApiService) PutSetProperty(name string, propertyName string, value string, localVarOptionals map[string]interface{}) (DocumentPropertyResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  DocumentPropertyResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/documentproperties/{propertyName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"propertyName"+"}", fmt.Sprintf("%v", propertyName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("value", parameterToString(value, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Replace document square annotation
 @param name The document name.
 @param annotationId The annotation ID.
 @param annotation Annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return SquareAnnotationResponse*/
func (a *PdfApiService) PutSquareAnnotation(name string, annotationId string, annotation SquareAnnotation, localVarOptionals map[string]interface{}) (SquareAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  SquareAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/square/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotation
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert SVG file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.svg)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "adjustPageSize" (bool) Adjust page size
     @param "height" (float64) Page height
     @param "width" (float64) Page width
     @param "isLandscape" (bool) Is page landscaped
     @param "marginLeft" (float64) Page margin left
     @param "marginBottom" (float64) Page margin bottom
     @param "marginRight" (float64) Page margin right
     @param "marginTop" (float64) Page margin top
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutSvgInStorageToPdf(name string, srcPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/svg"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["adjustPageSize"], "bool", "adjustPageSize"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["height"], "float64", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["width"], "float64", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["isLandscape"], "bool", "isLandscape"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginLeft"], "float64", "marginLeft"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginBottom"], "float64", "marginBottom"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginRight"], "float64", "marginRight"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginTop"], "float64", "marginTop"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["adjustPageSize"].(bool); localVarOk {
		localVarQueryParams.Add("adjustPageSize", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["height"].(float64); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["width"].(float64); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["isLandscape"].(bool); localVarOk {
		localVarQueryParams.Add("isLandscape", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginLeft"].(float64); localVarOk {
		localVarQueryParams.Add("marginLeft", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginBottom"].(float64); localVarOk {
		localVarQueryParams.Add("marginBottom", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginRight"].(float64); localVarOk {
		localVarQueryParams.Add("marginRight", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginTop"].(float64); localVarOk {
		localVarQueryParams.Add("marginTop", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Replace document text annotation
 @param name The document name.
 @param annotationId The annotation ID.
 @param annotation Annotation.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return TextAnnotationResponse*/
func (a *PdfApiService) PutTextAnnotation(name string, annotationId string, annotation TextAnnotation, localVarOptionals map[string]interface{}) (TextAnnotationResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  TextAnnotationResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/annotations/text/{annotationId}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"annotationId"+"}", fmt.Sprintf("%v", annotationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &annotation
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Update field.
 @param name The document name.
 @param fieldName The name of a field to be updated.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "field" (Field) with the field data.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return FieldResponse*/
func (a *PdfApiService) PutUpdateField(name string, fieldName string, localVarOptionals map[string]interface{}) (FieldResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FieldResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/fields/{fieldName}"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"fieldName"+"}", fmt.Sprintf("%v", fieldName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["field"].(Field); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Update fields.
 @param name The document name.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "fields" (Fields) with the fields data.
     @param "storage" (string) The document storage.
     @param "folder" (string) The document folder.
 @return FieldsResponse*/
func (a *PdfApiService) PutUpdateFields(name string, localVarOptionals map[string]interface{}) (FieldsResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  FieldsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/fields"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarTempParam, localVarOk := localVarOptionals["fields"].(Fields); localVarOk {
		localVarPostBody = &localVarTempParam
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert web page to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param url Source url
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "height" (float64) Page height
     @param "width" (float64) Page width
     @param "isLandscape" (bool) Is page landscaped
     @param "marginLeft" (float64) Page margin left
     @param "marginBottom" (float64) Page margin bottom
     @param "marginRight" (float64) Page margin right
     @param "marginTop" (float64) Page margin top
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutWebInStorageToPdf(name string, url string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/web"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["height"], "float64", "height"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["width"], "float64", "width"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["isLandscape"], "bool", "isLandscape"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginLeft"], "float64", "marginLeft"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginBottom"], "float64", "marginBottom"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginRight"], "float64", "marginRight"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["marginTop"], "float64", "marginTop"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("url", parameterToString(url, ""))
	if localVarTempParam, localVarOk := localVarOptionals["height"].(float64); localVarOk {
		localVarQueryParams.Add("height", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["width"].(float64); localVarOk {
		localVarQueryParams.Add("width", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["isLandscape"].(bool); localVarOk {
		localVarQueryParams.Add("isLandscape", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginLeft"].(float64); localVarOk {
		localVarQueryParams.Add("marginLeft", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginBottom"].(float64); localVarOk {
		localVarQueryParams.Add("marginBottom", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginRight"].(float64); localVarOk {
		localVarQueryParams.Add("marginRight", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["marginTop"].(float64); localVarOk {
		localVarQueryParams.Add("marginTop", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document which contatins XFA form (in request content) to PDF with AcroForm and uploads resulting file to storage.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.pdf)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "storage" (string) The document storage.
     @param "file" (*os.File) A file to be converted.
 @return AsposeResponse*/
func (a *PdfApiService) PutXfaPdfInRequestToAcroForm(outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/convert/xfatoacroform"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "multipart/form-data",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var localVarFile (*os.File)
	if localVarTempParam, localVarOk := localVarOptionals["file"].(*os.File); localVarOk {
		localVarFile = localVarTempParam
	}
	if localVarFile != nil {
		fbs, _ := ioutil.ReadAll(localVarFile)
		localVarFileBytes = fbs
		localVarFileName = localVarFile.Name()
		localVarFile.Close()
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Converts PDF document which contatins XFA form (located on storage) to PDF with AcroForm and uploads resulting file to storage
 @param name The document name.
 @param outPath Full resulting filename (ex. /folder1/folder2/result.pdf)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "folder" (string) The document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutXfaPdfInStorageToAcroForm(name string, outPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/convert/xfatoacroform"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["folder"], "string", "folder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("outPath", parameterToString(outPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["folder"].(string); localVarOk {
		localVarQueryParams.Add("folder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert XML file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.xml)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "xslFilePath" (string) Full XSL source filename (ex. /folder1/folder2/template.xsl)
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutXmlInStorageToPdf(name string, srcPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/xml"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["xslFilePath"], "string", "xslFilePath"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["xslFilePath"].(string); localVarOk {
		localVarQueryParams.Add("xslFilePath", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert XPS file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.xps)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutXpsInStorageToPdf(name string, srcPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/xps"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* PdfApiService Convert XslFo file (located on storage) to PDF format and upload resulting file to storage. 
 @param name The document name.
 @param srcPath Full source filename (ex. /folder1/folder2/template.xpsfo)
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "dstFolder" (string) The destination document folder.
     @param "storage" (string) The document storage.
 @return AsposeResponse*/
func (a *PdfApiService) PutXslFoInStorageToPdf(name string, srcPath string, localVarOptionals map[string]interface{}) (AsposeResponse,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  AsposeResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/pdf/{name}/create/xslfo"
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", fmt.Sprintf("%v", name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _url.Values{}
	localVarFormParams := _url.Values{}

	if err := typeCheckParameter(localVarOptionals["dstFolder"], "string", "dstFolder"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["storage"], "string", "storage"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("srcPath", parameterToString(srcPath, ""))
	if localVarTempParam, localVarOk := localVarOptionals["dstFolder"].(string); localVarOk {
		localVarQueryParams.Add("dstFolder", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["storage"].(string); localVarOk {
		localVarQueryParams.Add("storage", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

