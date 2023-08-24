package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type InferResponse struct {
	Predictions string `json:"predictions"`
}

// 发送POST请求
func sendPostRequest(dataFile, token, url string) ([]byte, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, io.NopCloser(bytes.NewReader(data)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", fmt.Sprintf("token:%s", token))
	req.SetBasicAuth("token", token)
	fmt.Printf("token: %s\n", token)

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Body)
		err = fmt.Errorf("req not success:\n%s", string(body))
		return nil, err
	}

	return body, err
}

// 将base64字符串解码为图片
func decodeBase64ToImage(base64Str, outputFileName string) error {
	_ = os.WriteFile("base64str2.txt", []byte(base64Str), 0644)
	decodedData, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return err
	}
	return os.WriteFile(outputFileName, decodedData, 0644)
}

func processRequest(dataFile, url string) error {
	token := os.Getenv("DATABRICKS_TOKEN")

	responseBody, err := sendPostRequest(dataFile, token, url)
	if err != nil {
		return err
	}

	var result InferResponse
	err = json.Unmarshal(responseBody, &result)
	if err != nil {
		return err
	}

	err = decodeBase64ToImage(result.Predictions, "output_image.jpg")
	return err
}

func Infer(dataFile, url string) {
	//url := "https://dbc-a04074f7-325a.cloud.databricks.com/serving-endpoints/poctest1/invocation"
	//dataFile := "data.json"
	if err := processRequest(dataFile, url); err != nil {
		fmt.Println("Error:", err)
	}
}
