package api

import (
    "encoding/base64"
    "encoding/json"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/lambda"
    "fmt"
)

type LambdaResponse struct {
    Results1 []string `json:"results1"`
    Results2 []string `json:"results2"`
}

func invokeLambda(imageData1, imageData2 []byte) (LambdaResponse, error) {
    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("us-east-1")},
    )
    if err != nil {
        return LambdaResponse{}, err
    }

    svc := lambda.New(sess)

    payload := map[string]interface{}{
        "image_data1": base64.StdEncoding.EncodeToString(imageData1),
        "image_data2": base64.StdEncoding.EncodeToString(imageData2),
    }
    payloadBytes, err := json.Marshal(payload)
    if err != nil {
        return LambdaResponse{}, err
    }

    result, err := svc.Invoke(&lambda.InvokeInput{
        FunctionName: aws.String("poker-lambda"),
        Payload:      payloadBytes,
    })
    if err != nil {
        return LambdaResponse{}, err
    }

    var response LambdaResponse
    err = json.Unmarshal(result.Payload, &response)
    if err != nil {
        return LambdaResponse{}, err
    }

    return response, nil
}

// Converts from {rank}{suit} to {rank}_{suit} with suit symbol
func ConvertCardFormat(cardStr string) string {
	if len(cardStr) < 2 {
		return "Invalid card format"
	}

	rank := string(cardStr[:len(cardStr)-1])
    suit := string(cardStr[len(cardStr)-1])

	suitMap := map[string]string{
		"S": "♤",
		"H": "♡",
		"D": "♦",
		"C": "♣",
	}

	newSuit, exists := suitMap[suit]
	if !exists {
		return "Invalid suit"
	}

	return fmt.Sprintf("%s_%s", rank, newSuit)
}