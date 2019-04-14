## Settings
default root object: index.html

## Origins

### Lambda
origin: api entrypoint
origin path: /Prod
protocol: HTTPS only

### Frontend
origin: select s3-bucket

## Behaviors
Allow all HTTP methods
All cookie
Header: Accept

api/*

Default(*): s3-web, compress, redirect HTTPS

#Errors pages

403: /index.html 200
404: /index.html 200