from typing import Union
from fastapi.middleware.cors import CORSMiddleware
from fastapi.background import BackgroundTasks
from fastapi import FastAPI
from redis_om import get_redis_connection, HashModel
from starlette.requests import Request
import requests
import time

app = FastAPI()

redis = get_redis_connection(
    host="localhost",
    port= 6379,
    decode_responses=True
)


class Order(HashModel):
    product_id: str
    product_name : str
    price: float
    fee: float
    total: float
    quantity: float
    status: str # pending , completed, refunded
    class Mets:
        database = redis


@app.post('/orders') # id and quantity
async def create(req: Request, background_task: BackgroundTasks):
    body = await req.json()
    r = requests.get('http://localhost:8000/product/%s' % body['id'])

    product = r.json()

    order = Order(
        product_id = body['id'],
        product_name = product['name'],
        price = product['price'],
        fee = 0.2 * product['price'],
        total = 1.2 * product['price'],
        quantity = body['quantity'],
        status = 'pending'
        )

    order.save()

    background_task.add_task(order_completed, order)

    return order



def order_completed(order):
    time.sleep(3)
    order.status = 'completed'
    order.save()
    redis.xadd("order_completed", order.dict(), '*')


@app.get('/orders/{pk}')
def get(pk):
    return Order.get(pk)
