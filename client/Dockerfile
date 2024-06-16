# Build stage
FROM node:alpine3.20 AS builder

WORKDIR /app

COPY package.json .

RUN npm install

COPY . .

# Generates a dist folder as defined in the package.json file
RUN npm run build

# Runtime
FROM nginx:alpine

# We have to copy the contents of the dist folder to the following path for ngnix to serve
# Copy the build output from the previous stage to the Nginx web server directory
COPY --from=builder /app/dist/ /usr/share/nginx/html

EXPOSE 80

CMD ["nginx",  "-g", "daemon off;"]