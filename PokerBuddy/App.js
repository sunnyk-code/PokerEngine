import React, { useState, useRef, useEffect } from 'react';
import { View, StyleSheet, Button, Image, Text, Alert, Dimensions } from 'react-native';
import { Camera } from 'expo-camera';

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
        if (!hasTakenFlopPicture) {
          setHasTakenFlopPicture(true);
        }
      } catch (error) {
        Alert.alert("Error", "Failed to take picture: " + error.message);
      }
    } else {
      Alert.alert("Error", "Camera is not ready yet.");
    }
  };

  const renderCamera = (cameraRef, photo, setPhoto, cameraId, isCameraReady, setIsCameraReady) => (
    <View style={styles.cameraContainer}>
      {photo ? (
        <View style={styles.imagePreview}>
          <Image source={{ uri: photo }} style={styles.image} />
          <Button title="Retake" onPress={() => setPhoto(null)} color="#007AFF" />
        </View>
      ) : (
        <Camera 
          style={styles.camera} 
          ref={cameraRef} 
          onCameraReady={() => setIsCameraReady(true)}
        >
          <View style={styles.buttonContainer}>
            <Button 
              title={cameraId === 1 && !hasTakenFlopPicture ? "Take a Picture of the Flop" : "Take a Picture of the Hand"} 
              onPress={() => takePicture(cameraRef, setPhoto, isCameraReady)} 
              color="#007AFF" 
            />
          </View>
        </Camera>
      )}
    </View>
  );

  if (hasPermission === null) {
    return <View />;
  }
  if (hasPermission === false) {
    return <Text>No access to camera</Text>;
  }

  return (
    <View style={styles.container}>
      {renderCamera(cameraRef1, photo1, setPhoto1, 1, isCameraReady1, setIsCameraReady1)}
      {hasTakenFlopPicture && renderCamera(cameraRef2, photo2, setPhoto2, 2, isCameraReady2, setIsCameraReady2)}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#000000', // Black background
    justifyContent: 'center',
    alignItems: 'center',
  },
  cameraContainer: {
    width: Dimensions.get('window').width * 0.9,
    height: Dimensions.get('window').height * 0.4,
    borderRadius: 10,
    overflow: 'hidden', // Needed to make the rounded corners work
    backgroundColor: '#FFFFFF', // White background for the camera container
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
    paddingBottom: 10,
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
});
