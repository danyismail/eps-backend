# eps-backend

## backend api for eps

## locally testing prerequisite

- tunnel database through from backend server to local use ssh -L localhost:1433:192.168.26.28:1433 eps@36.92.58.82

- build image from Dockerfile

- run container from image already created with describe volume

- e.g. docker run -v /Users/daniismail/Documents/uploads:/app/uploads -v /Users/daniismail/Documents/backend-logs/app.log:/app/app.log -p 1717:1525 -d eps-backend-api
