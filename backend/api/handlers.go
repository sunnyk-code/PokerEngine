package api

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "os"
    "github.com/joho/godotenv"
    "strconv"
    p "backend/pokerLogic"
)


func init() {
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }
}

// Middleware to check the API key
func ApiKeyMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {

        apiKey := os.Getenv("API_KEY")
        requestApiKey := r.Header.Get("API-KEY")

        if apiKey == "" || apiKey != requestApiKey {
            http.Error(w, "Invalid API key", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    }
}




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

    cardCountStr := r.FormValue("cardCount")
    cardCount, err := strconv.Atoi(cardCountStr)
    if err != nil {
        http.Error(w, "Failed to read card count", http.StatusBadRequest)
        return
    }

    response, err := invokeLambda(imageBytes1, imageBytes2)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to invoke Lambda: %v", err), http.StatusInternalServerError)
        return
    }

    communityCardsFromImage := response.Results1
    cardsSeen := make(map[string]bool)
    communityCardsAsStrings := make([]string, 0, cardCount)
    i := 0
    for len(communityCardsAsStrings) != cardCount && i < len(communityCardsFromImage) {
        card := communityCardsFromImage[i]
        if !cardsSeen[card] {
            communityCardsAsStrings = append(communityCardsAsStrings, card)
            cardsSeen[card] = true
        }
        i += 1
    }

    handAsStrings := response.Results2

    for i, card := range communityCardsAsStrings {
        communityCardsAsStrings[i] = ConvertCardFormat(card)
    }

    for i, card := range handAsStrings {
        handAsStrings[i] = ConvertCardFormat(card)
    }

    deck := p.NewDeck()

    deck.RemoveCard(p.NewCard(handAsStrings[0]))
    deck.RemoveCard(p.NewCard(handAsStrings[1]))

    hand := []p.Card{
		p.NewCard(handAsStrings[0]),
		p.NewCard(handAsStrings[1]),
	}

    communityCards := make([]p.Card, len(communityCardsAsStrings))
    for i, card := range communityCardsAsStrings {
        communityCards[i] = p.NewCard(card)
        deck.RemoveCard(p.NewCard(card))
    }

    winPercentage := p.MonteCarloSimulation(2, 100000, p.Hand(hand), communityCards, deck)

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(winPercentage)
    
}

