# Use the official Node.js 18 base image
FROM node:18

# Set the working directory inside the container to /app
WORKDIR /app

# Copy application source code from the host machine's source_code/ folder into the container
COPY source_code/ .

# (Optional) Install project dependencies if needed
# RUN npm install

# Expose port 3000 so it can be accessed outside the container
EXPOSE 3000

# Define the command to run your app
# Build number shown on the home page (set by CI with --build-arg)
ARG APP_VERSION=dev
ENV APP_VERSION=$APP_VERSION

CMD ["node", "index.js"]
