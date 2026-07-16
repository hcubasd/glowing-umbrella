FROM node:26
WORKDIR /root/app
COPY package*.json ./
RUN npm ci --omit=dev
COPY dist/ ./dist
