import React, { useState, useRef, useEffect } from 'react';
import { View, StyleSheet, TouchableOpacity, Image, Text, Alert, Dimensions, ScrollView, TextInput, Button } from 'react-native';
import { Camera } from 'expo-camera/legacy';
import axios from 'axios';
import { API_KEY } from '@env'
import Modal from 'react-native-modal';

const { width, height } = Dimensions.get('window');

export default function App() {
  const [hasPermission, setHasPermission] = useState(null);
  const [photo1, setPhoto1] = useState(null);
  const [photo2, setPhoto2] = useState(null);
  const [isCameraReady1, setIsCameraReady1] = useState(false);
  const [isCameraReady2, setIsCameraReady2] = useState(false);
  const cameraRef1 = useRef(null);
  const cameraRef2 = useRef(null);
  const [hasTakenFlopPicture, setHasTakenFlopPicture] = useState(false);
  const [isModalVisible, setModalVisible] = useState(false);
  const [cardCount, setCardCount] = useState('');
  const [cardCountInteger, setCardCountInteger] = useState(0);
  const [error, setError] = useState('');
  const [winningPercentage, setWinningPercentage] = useState('');

  
  useEffect(() => {
    (async () => {
      const { status } = await Camera.requestCameraPermissionsAsync();
      setHasPermission(status === 'granted');
    })();
  }, []);
  

  const handleInputChange = (value) => {
    setCardCount(value);
  };

  const takePicture = async (cameraRef, setPhoto, isCameraReady) => {
    if (cameraRef.current && isCameraReady) {
      try {
        let photo = await cameraRef.current.takePictureAsync();
        setPhoto(photo.uri);
        console.log(photo.uri)
      } catch (error) {
        Alert.alert("Error", "Failed to take picture: " + error.message);
      }
    } else {
      Alert.alert("Error", "Camera is not ready yet.");
    }
  };

  const mountError = () => {
    console.log('mount error')
  }

  const submitCommunityCards = () => {
    setModalVisible(true);
    console.log('cc run')
  }
  
  const handleSubmit = () => {
    const number = parseInt(cardCount, 10);
    if (isNaN(number) || number < 0 || number > 5) {
      setError('Please enter a valid number between 0 and 5.');
    } else {
      setCardCountInteger(number)
      setError('');
      setModalVisible(false);
      // Handle the valid input here
      console.log('Number of cards:', number);
      setHasTakenFlopPicture(true);
    }
  }

  const handleCancel = () => {
    setModalVisible(false)
    setCardCount('')
    setError('')
  }

  

  const submitFlop = async () => {
    console.log('submit flop')
    const formData = new FormData();

    formData.append('image1', {
      uri: photo1,
      name: 'image1.jpg',
      type: 'image/jpeg',
    });

    formData.append('image2', {
      uri: photo2,
      name: 'image2.jpg',
      type: 'image/jpeg',
    });

    formData.append('cardCount', cardCountInteger);

    try {
      const response = await axios.post('http://ec2-23-22-41-166.compute-1.amazonaws.com:80/winning-percentage', formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
          'API-KEY': API_KEY,
        },
      });
      console.log('Response:', response.data);
      setWinningPercentage(response.data)
    } catch (error) {
      console.error('Error uploading images:', error);
    }
  }

  const renderCamera = (cameraRef, photo, setPhoto, cameraId, isCameraReady, setIsCameraReady, submit) => (
    <View style={styles.cameraContainer}>
      {photo ? (
        <View style={styles.imagePreview}>
          <Image source={{ uri: photo }} style={styles.image} />
          <View style={styles.buttonContainer}>
            <TouchableOpacity 
              style={styles.button} 
              onPress={() => setPhoto(null)}>
              <Text style={styles.buttonText}>Retake</Text>
            </TouchableOpacity>
            <TouchableOpacity 
              style={styles.button} 
              onPress={submit}>
              <Text style={styles.buttonText}>Submit</Text>
            </TouchableOpacity>
          </View>
        </View>
      ) : (
        <Camera 
          style={styles.camera} 
          ref={cameraRef} 
          onCameraReady={() => setIsCameraReady(true)}
          onMountError={mountError}>
          <View style={styles.buttonContainer}>
            <TouchableOpacity 
              style={styles.button} 
              onPress={() => takePicture(cameraRef, setPhoto, isCameraReady)}>
              <Text style={styles.buttonText}>
                {cameraId === 1 && !hasTakenFlopPicture ? "Take a Picture of the Community Cards" : "Take a Picture of the Hand"}
              </Text>
            </TouchableOpacity>
          </View>
        </Camera>
      )}
    </View>
  );

  const renderWinningPercentage = () => {
    return (
      <View style={styles.winningPercentageContainer}>
        <Text style={styles.winningPercentageText}>{winningPercentage}</Text>
      </View>
    );
  }

  const renderResetButton = () => {
    return (
      <View style={styles.resetButtonContainer}>
        <TouchableOpacity style={styles.resetButton} onPress={reset}>
          <Text style={styles.resetButtonText}>Reset</Text>
        </TouchableOpacity>
      </View>
    );
  }

  const reset = () => {
    setPhoto1(null);
    setPhoto2(null);
    setHasTakenFlopPicture(false);
    setCardCount('');
    setCardCountInteger(0);
    setWinningPercentage('');
    setError('');
    setIsCameraReady1(false);
    setIsCameraReady2(false);
  }


  if (hasPermission === null) {
    return <View />;
  }
  if (hasPermission === false) {
    return <Text style={styles.permissionText}>No access to camera</Text>;
  }

  return (
    <ScrollView contentContainerStyle={styles.container}>
      <View style={styles.header}>
        <Text style={styles.headerText}>PokerBuddy</Text>
      </View>
      {!hasTakenFlopPicture && renderCamera(cameraRef1, photo1, setPhoto1, 1, isCameraReady1, setIsCameraReady1, submitCommunityCards)}
      {hasTakenFlopPicture && !winningPercentage && renderCamera(cameraRef2, photo2, setPhoto2, 2, isCameraReady2, setIsCameraReady2, submitFlop)}
      <Modal isVisible={isModalVisible}>
        <View style={styles.modalContent}>
          <Text>Enter the number of cards in the photo (0-5):</Text>
          <TextInput
            style={styles.input}
            keyboardType="numeric"
            value={cardCount}
            onChangeText={handleInputChange}
          />
          {error ? <Text style={styles.error}>{error}</Text> : null}
          <Button title="Submit" onPress={handleSubmit} />
          <Button title="Cancel" onPress={handleCancel} />
        </View>
      </Modal>
      {winningPercentage && renderWinningPercentage()}
      {winningPercentage && renderResetButton()}
      
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flexGrow: 1,
    backgroundColor: '#f0f0f0',
    alignItems: 'center',
  },
  header: {
    backgroundColor: '#007AFF',
    width: width,
    padding: 15,
    alignItems: 'center',
    justifyContent: 'flex-end', 
    paddingTop: 50,
  },
  headerText: {
    fontSize: 24,
    fontWeight: 'bold',
    color: '#FFFFFF',
    fontFamily: 'Arial',
  },
  cameraContainer: {
    width: width * 0.9,
    height: height * 0.75,
    borderRadius: 20,
    marginVertical: 10,
    backgroundColor: '#FFFFFF',
    shadowColor: "#000",
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.25,
    shadowRadius: 3.84,
    elevation: 5,
  },
  camera: {
    flex: 1,
    justifyContent: 'flex-end',
    alignItems: 'center',
  },
  buttonContainer: {
    backgroundColor: 'transparent',
    padding: 20,
    flexDirection: 'row'
  },
  button: {
    backgroundColor: '#007AFF',
    paddingTop: 10,
    paddingHorizontal: 20,
    paddingVertical: 10,
    borderRadius: 20,
    margin: 10,
  },
  buttonText: {
    color: '#FFFFFF',
    fontSize: 16,
    fontWeight: 'bold',
  },
  imagePreview: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  image: {
    width: '100%',
    height: '80%',
    borderRadius: 8,
  },
  permissionText: {
    fontSize: 18,
    color: '#000',
    marginTop: 20,
  },
  analyzeButton: {
    backgroundColor: '#007AFF',
    padding: 20,
    borderRadius: 10,
    marginTop: 20,
    marginBottom: 20,
    width: '100%',
  },
  analyzeButtonText: {
    color: '#007AFF',
    fontSize: 20,
    fontWeight: 'bold',
  },
  modalContent: {
    backgroundColor: 'white',
    padding: 20,
    borderRadius: 10,
    alignItems: 'center',
  },
  input: {
    borderWidth: 1,
    borderColor: '#ccc',
    padding: 10,
    marginVertical: 10,
    width: '80%',
    textAlign: 'center',
  },
  error: {
    color: 'red',
    marginBottom: 10,
  },
  winningPercentageContainer: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    height: height,
  },
  winningPercentageText: {
    fontSize: 24,
    fontWeight: 'bold',
    color: 'blue',
  },
  resetButtonContainer: {
    marginTop: 20,
    alignItems: 'center',
    justifyContent: 'center',
  },
  resetButton: {
    backgroundColor: 'red',
    padding: 15,
    borderRadius: 10,
  },
  resetButtonText: {
    color: 'white',
    fontSize: 18,
    fontWeight: 'bold',
  },
});
