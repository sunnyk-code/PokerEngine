# Use an official AWS Lambda image as a parent image
FROM public.ecr.aws/lambda/python:3.10

# Install necessary system dependencies
RUN yum install -y \
    gcc \
    gcc-c++ \
    cmake \
    make \
    git \
    wget \
    unzip \
    python3-devel \
    epel-release \
    libjpeg-devel \
    libpng-devel \
    libtiff-devel \
    jasper-devel \
    openexr-devel \
    libwebp-devel \
    tbb-devel \
    libv4l-devel \
    v4l-utils \
    ffmpeg-devel \
    gstreamer-plugins-base-devel \
    gstreamer1-devel \
    gtk3-devel \
    mesa-libGL-devel \
    libdc1394-devel \
    tbb-devel \
    libopenblas-dev \
    liblapacke-dev \
    libavcodec-devel \
    libavformat-devel \
    libswscale-devel \
    libeigen3-devel \
    python3-opencv \
    libcanberra-gtk3* \
    openblas-devel \
    libxine2-devel \
    libgstreamer-plugins-base1.0-dev \
    libsm6 \
    libxext6 \
    libxrender1 \
    opencv

# Install Python dependencies
RUN pip install --upgrade pip
RUN pip install \
    opencv-python-headless \
    numpy \
    supervision \
    inference

# Copy the function code
COPY lambda_function.py ${LAMBDA_TASK_ROOT}

# Set the CMD to your handler
CMD ["lambda_function.lambda_handler"]
