from typing import Union
from fastapi.middleware.cors import CORSMiddleware
from fastapi import FastAPI
from redis_om import get_redis_connection, HashModel

app = FastAPI()

redis = get_redis_connection(
    host="localhost",
    port= 6379,
    decode_responses=True
)

class Product(HashModel):
    name: str
    price: float
    quantity: int

    app.add_middleware(
        CORSMiddleware,
        allow_methods=['*'],
        allow_headers=['*']
    )
    class Meta:
        database = redis



@app.get('/product')
def all():
    return [format(pk) for pk in Product.all_pks()]
def format(pk):
    product = Product.get(pk)
    return {
        'id': product.pk,
        'name': product.name,
        'price': product.price,
        'quantity': product.quantity
    }

@app.post('/product')
def create(product: Product):
    return product.save()


@app.get('/product/{pk}')
def get(pk):
    return Product.get(pk)

@app.delete('/product/{pk}')
def delete(pk):
    return Product.delete(pk)
