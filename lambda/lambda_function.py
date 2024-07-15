import json
from inference import get_model
import supervision as sv
import cv2
import numpy as np
import base64


def process_image(image_data):
    image_bytes = base64.b64decode(image_data)
    image_np = np.frombuffer(image_bytes, np.uint8)
    image = cv2.imdecode(image_np, cv2.IMREAD_COLOR)

    model = get_model(model_id="playing-cards-ow27d/4")
    results = model.infer(image)[0]
    results_class_names = []
    for prediction in results.predictions:
        results_class_names.append(prediction.class_name)


    return results_class_names

def lambda_handler(event, context):
    image_data1 = event['image_data1']
    image_data2 = event['image_data2']

    results1 = process_image(image_data1)
    results2 = process_image(image_data2)

    return {
        'statusCode': 200,
        'body': json.dumps({
            'results1': results1,
            'results2': results2
        })
    }

