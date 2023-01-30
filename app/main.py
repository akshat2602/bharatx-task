from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from pydantic import BaseModel
import websocket
import json


app = FastAPI()


origins = ["*"]
app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


class NormalRequest(BaseModel):
    value: int


class NormalResponse(BaseModel):
    c: int


def get_c_from_service2() -> NormalResponse:
    ws = websocket.WebSocket()
    ws.connect("ws://ws-server:8080/sum")
    resp = ws.recv()
    c = json.loads(resp)["value"]
    ws.close()
    return NormalResponse(c=c)


def handle_post_to_service2(request: NormalRequest, endpoint: str = "a"):
    ws = websocket.WebSocket()
    ws.connect("ws://ws-server:8080/send")
    ws.send(json.dumps({"value": request.value, "endpoint": endpoint}))
    ws.close()
    return


@app.get("/c", response_model=NormalResponse)
def c():
    return get_c_from_service2()


@app.post("/a")
def a(request: NormalRequest):
    return handle_post_to_service2(request=request, endpoint="a")


@app.post("/b")
def b(request: NormalRequest):
    return handle_post_to_service2(request=request, endpoint="b")
