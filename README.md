# receipt-app

![Receipt App Screenshot](./docs/screenshot.png)

This is an example service developed to showcase the abilities of Common Fate Cloud.

Receipt App helps users store their receipts in the cloud. It's build with a cloud-native architecture and uses AWS S3 to store the receipt files. This app is intended to be deployed as a Docker container on AWS using a service such as AWS EKS.

## Development

### Tech stack

- Go 1.16
- NodeJS 14 (used to compile CSS styles using PostCSS)
- [Yarn v1](https://classic.yarnpkg.com/lang/en/)
- HTML using Go `template/html` standard HTML templating library
- [TailwindCSS](https://tailwindcss.com/) for styling

### Getting started

Before starting development, create an AWS S3 bucket. As shown below, when running the server you must pass the `S3_BUCKET_NAME` environment variable to point to the name of the bucket you wish to use.

The AWS role which you are using in your terminal session must have read/write access to the bucket.

1. Install NodeJS dependencies: `yarn install`
2. Build CSS styles: `yarn build`
3. Export the name of the S3 bucket: `export S3_BUCKET_NAME=___replace_with_your_bucket_name___`
4. Run the Go server: `go run main.go`

The app should be available on http://localhost:8080/

### Building the Docker container

To verify that the Docker container builts without errors, you can run

```
docker build -t receipt-app .
```
