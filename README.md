# Poker Engine

## Overview

Poker Engine is a powerful and intuitive app designed to enhance your Texas Hold'em poker learning experience by providing real-time metrics on your chances of winning. Using Monte Carlo simulations, Poker Engine analyzes the state of your game from photos and delivers accurate winning probabilities, helping you learn and become a better player by verifying your own analysis in real time.

## How It Works

1. **Capture Game State**: Use your device's camera to take clear pictures of the current state of your Texas Hold'em game, including your cards and the community cards.
2. **Image Processing**: The app processes the images using a [Computer Vision model](https://universe.roboflow.com/augmented-startups/playing-cards-ow27d) to identify the cards and game state.
3. **Simulation**: Monte Carlo simulations are run based on the identified game state to calculate your chances of winning.
4. **Results Display**: View your winning probabilities directly within the app.

## Installation

To install Poker Engine, follow these steps:

1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/PokerEngine.git
    ```
2. Navigate to the frontend directory:
    ```bash
    cd PokerEngine/frontend
    ```
3. Install the necessary dependencies:
    ```bash
    npm install
    ```
4. Start the app:
    ```bash
    npm start
    ```

5. Open the app on your mobile device using the Expo Go mobile app
