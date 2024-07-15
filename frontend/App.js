import React, { useState, useRef, useEffect } from 'react';
import { View, StyleSheet, TouchableOpacity, Image, Text, Alert, Dimensions, ScrollView } from 'react-native';
import { Camera } from 'expo-camera/legacy';
import axios from 'axios';

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

  useEffect(() => {
    (async () => {
      const { status } = await Camera.requestCameraPermissionsAsync();
      setHasPermission(status === 'granted');
    })();
  }, []);

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
    setHasTakenFlopPicture(true);
    console.log('cc run')
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

  try {
    const response = await axios.post('http://10.0.0.91:8080/winning-percentage', formData, {
      headers: {
        'Content-Type': 'multipart/form-data',
      },
    });
    console.log('Response:', response.data);
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
      {hasTakenFlopPicture && renderCamera(cameraRef2, photo2, setPhoto2, 2, isCameraReady2, setIsCameraReady2, submitFlop)}
      
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
});
