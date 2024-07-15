package api

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)


func WinningPercentageHandler(w http.ResponseWriter, r *http.Request) {
    imageData1, _, err := r.FormFile("image1")
    if err != nil {
        http.Error(w, "Failed to read image data", http.StatusBadRequest)
        return
    }
    defer imageData1.Close()

    imageData2, _, err := r.FormFile("image2")
    if err != nil {
        http.Error(w, "Failed to read image data", http.StatusBadRequest)
        return
    }
    defer imageData2.Close()

    imageBytes1, err := io.ReadAll(imageData1)
    if err != nil {
        http.Error(w, "Failed to read image data", http.StatusBadRequest)
        return
    }

    imageBytes2, err := io.ReadAll(imageData2)
    if err != nil {
        http.Error(w, "Failed to read image data", http.StatusBadRequest)
        return
    }

    response, err := invokeLambda(imageBytes1, imageBytes2)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to invoke Lambda: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)


	// Encode the image bytes to Base64
    // encodedImage1 := base64.StdEncoding.EncodeToString(imageBytes1)
    // encodedImage2 := base64.StdEncoding.EncodeToString(imageBytes2)

    // // Create a JSON response object
    // response := map[string]string{
    //     "image1": encodedImage1,
    //     "image2": encodedImage2,
    // }

    // // Set the Content-Type header
    // w.Header().Set("Content-Type", "application/json")

    // // Encode and write the JSON response
    // if err := json.NewEncoder(w).Encode(response); err != nil {
    //     http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    // }
}