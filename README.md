## Video of the running HTTP and Websocket Servers

## Clone the repository

```
git clone https://github.com/akshat2602/bharatx-task
```

## Dependencies

The python HTTP Server has the following dependencies -

```
fastAPI
uvicorn
websocket-client
```

The golang Websocket Server has the following dependencies -

```
nhooyr.io/websocket
```

## To run the HTTP server, follow the instructions listed below

Run the following command to install required dependencies

```
pip3 install -r app/requirements.txt
```

Run the http server with the command

```
python3 -m uvicorn app.main:app --reload
```

## To run the Websocket server, follow the instructions listed below

Run the following command to install required dependencies

```
go get .
```

Run the websocket server with the command

```
go build -v -o server

./server
```
