# Используем базовый образ nginx
FROM nginx:alpine

# Копируем простой HTML-файл в директорию nginx
COPY index.html /usr/share/nginx/html/index.html

# Создаем простой HTML-файл
RUN echo '<!DOCTYPE html><html><head><title>Empty Page</title></head><body><h1>Hello from Docker!</h1></body></html>' > /usr/share/nginx/html/index.html

# Открываем порт 80
EXPOSE 80

# Запускаем nginx
CMD ["nginx", "-g", "daemon off;"]

# # pull official base image
# FROM node:lts

# # set working directory
# WORKDIR /usr/src/app

# # copy package.json and install dependencies
# COPY ./package.json ./package-lock.json ./
# RUN npm install

# # copy react project
# COPY . .

# # Build with environment variables passed as build arguments
# ARG REACT_APP_API_HOST
# ARG REACT_APP_API_PORT
# ARG REACT_APP_API_SECURE

# # Print environment variables during build for verification
# RUN echo "Building with API_HOST=${REACT_APP_API_HOST}, API_PORT=${REACT_APP_API_PORT}, API_SECURE=${REACT_APP_API_SECURE}"

# # build react project
# RUN npm run build
# # serve the built project
# RUN npm install -g serve
# WORKDIR /usr/src/app/build
