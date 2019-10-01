package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
)

type Docxmerge struct {
	baseUrl string
	apiKey  string
	tenant  string
}

type DocxmergeOptions struct {
	BaseUrl string
	ApiKey  string
	Tenant  string
}
type Data map[string]interface{}

func newMultipartFile(reader io.Reader, body *bytes.Buffer) (*multipart.Writer, error) {
	fileContents, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "file.docx")
	if err != nil {
		return nil, err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return writer, nil
}

// Creates a new file upload http request with optional extra params
func newMultipartData(reader io.Reader, data Data, body *bytes.Buffer, conversionType string) (*multipart.Writer, error) {
	fileContents, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", "file.docx")
	if err != nil {
		return nil, err
	}
	_, err = part.Write(fileContents)
	if err != nil {
		return nil, err
	}
	jsonBytes, err := json.Marshal(data)
	part, err = writer.CreateFormField("data")
	if err != nil {
		return nil, err
	}
	_, err = part.Write(jsonBytes)
	if err != nil {
		return nil, err
	}

	part, err = writer.CreateFormField("conversionType")
	if err != nil {
		return nil, err
	}
	_, err = part.Write([]byte(conversionType))
	if err != nil {
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	return writer, nil
}

func (d *Docxmerge) hidrateRequest(request *http.Request) {
	request.Header.Set("api-key", d.apiKey)
	request.Header.Set("x-tenant", d.tenant)
}
func (d *Docxmerge) RenderFile(reader io.Reader, data Data, conversionType string) (io.Reader, error) {
	uri := fmt.Sprintf("%s/api/v1/Admin/RenderFile", d.baseUrl)
	body := new(bytes.Buffer)
	w, err := newMultipartData(reader, data, body, conversionType)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}
	d.hidrateRequest(request)
	request.Header.Set("Content-Type", w.FormDataContentType())
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected status code %d", resp.StatusCode))
	}
	return resp.Body, nil
}
func (d *Docxmerge) RenderTemplate(templateName string, data Data, conversionType string, version string) (io.Reader, error) {
	uri := fmt.Sprintf("%s/api/v1/Admin/RenderTemplate", d.baseUrl)
	dataBody := map[string]interface{}{
		"data":           data,
		"conversionType": conversionType,
		"template":       templateName,
		"version":        version,
	}
	body, err := json.Marshal(dataBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	d.hidrateRequest(request)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected status code %d", resp.StatusCode))
	}
	return resp.Body, nil
}
func (d *Docxmerge) RenderUrl(url string, data Data, conversionType string) (io.Reader, error) {
	uri := fmt.Sprintf("%s/api/v1/Admin/RenderUrl", d.baseUrl)
	dataBody := map[string]interface{}{
		"data":           data,
		"conversionType": conversionType,
		"url":            url,
	}
	body, err := json.Marshal(dataBody)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	d.hidrateRequest(request)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf("Unexpected status code %d", resp.StatusCode))
	}
	return resp.Body, nil
}

func NewDocxmerge(options DocxmergeOptions) *Docxmerge {
	return &Docxmerge{
		baseUrl: options.BaseUrl,
		apiKey:  options.ApiKey,
		tenant:  options.Tenant,
	}
}

func main() {
	data := Data{
		"name": "David Viejo",
		"logo": "https://docxmerge.com/assets/android-chrome-512x512.png",
	}
	docxmerge := NewDocxmerge(DocxmergeOptions{
		BaseUrl: "https://api.docxmerge.com",
		ApiKey:  "API_KEY",
		Tenant:  "default",
	})
	url := "https://api.docxmerge.com/api/v1/File/GetContenido?id=cdb9842d-5e38-4149-a06b-e1079a208fc3&download=true"

	pdf, err := docxmerge.RenderUrl(url, data, "PDF")
	if err != nil {
		panic(err)
	}
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(pdf)
	if err != nil {
		panic(err)
	}
	output := "./hello_world_golang.pdf"
	err = ioutil.WriteFile(output, buf.Bytes(), 0644)
	if err != nil {
		panic(err)
	}
	log.Printf("Check %s", output)
}
