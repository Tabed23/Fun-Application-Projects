const express = require('express');
const app = express();
const mongoose = require('mongoose');
const body_parser = require('body-parser');
const axios = require('axios');

require('./model')

const Order = mongoose.model('order');

mongoose.connect("mongodb://localhost:27017/orderservice");

app.use(body_parser.json());


app.get('/', (req, res) =>{
    res.send("health check")
});


// Post order
app.post('/order', (req, res) =>{

    cid = new mongoose.Types.ObjectId(req.body.customerId)
    bid = new mongoose.Types.ObjectId(req.body.bookId)
    var neworder  =  {

        customerId: cid,
        bookId: bid,
        initialDate : req.body.initialDate,
        deliveryDate: req.body.deliveryDate
    }

    var order =  new Order(neworder);

    order.save().then(()=>{
        res.send('new order created');
    }).catch(err =>{

        if (err){
            throw err;
        }
    });



});

// Get order
app.get('/orders', (req, res) =>{

    Order.find().then((order) => {

        if (order.length > 0) {
            res.json(order);
        }else{
            res.json({"data": [], "error": "orders not found"})
        }

    }).catch((err) =>{

        if (err){
            throw err;
        }

    })
});


// Get order by ID
app.get('/order/:id', (req, res) =>{
    Order.findById(req.params.id).then((order) =>{

        if(order){
            axios.get("http://localhost:8081/customer/" + order.customerId).then((response) =>{
                var orderObject = {customerName: response.data.name, bookTitle: '', publisher: '', author: ''}

                axios.get("http://localhost:8080/book/" + order.bookId).then((response)=> {
                    
                    orderObject.bookTitle = response.data.title
                    orderObject.publisher = response.data.publisher
                    orderObject.author = response.data.author

                    console.log(orderObject);

                    res.json(orderObject)
                })
            })
        }else{
            res.send("invalid order");
        }

    }).catch(err =>{

        if (err){

            throw err;
        }
    });
})


// Delete order
app.delete('/order/:id', (req, res) =>{
    Order.findByIdAndRemove(req.params.id).then((order) =>{
        res.json({"msg": "order has been deleted"})
    }).catch(err =>{

        if (err){

            throw err;
        }
    });
})

app.listen(8085, ()=> {
    console.log('order server is running.....!');
});
