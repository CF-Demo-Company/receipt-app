################################################
# Build the frontend
################################################

FROM node:14.18.0-alpine as frontend_builder

WORKDIR /app

# install and cache app dependencies
COPY package.json ./
COPY yarn.lock ./
COPY postcss.config.js ./
COPY tailwind.config.js ./

RUN yarn install --frozen-lockfile

# bundle static source inside Docker image
COPY static/css/. ./static/css
RUN yarn run build

FROM golang:1.16-alpine AS server_builder

WORKDIR /app

COPY go.* ./

# install dependencies
RUN go mod download

# copy built frontend static files
COPY --from=frontend_builder /app/static/dist ./static/dist

# add all other folders required for the Go build
COPY . .

RUN go build -o bin/receipt-app main.go

FROM alpine:3.13.5

WORKDIR /app

COPY --from=server_builder /app/bin/receipt-app /app/receipt-app

# Web HTTP
EXPOSE 8080

CMD /app/receipt-app