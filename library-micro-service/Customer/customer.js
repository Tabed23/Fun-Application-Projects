const express = require('express');
const app = express();
const mongoose = require('mongoose');
const body_parser = require('body-parser');


require('./model')

const Customer = mongoose.model('customer');

mongoose.connect("mongodb://localhost:27017/customerservice");

app.use(body_parser.json());


app.get('/', (req, res) =>{
    res.send("health check")
});


// Post customer
app.post('/customer', (req, res) =>{

    var newcustomer  =  {
        name: req.body.name,
        age: req.body.age,
        address : req.body.address
    }

    var customer =  new Customer(newcustomer);

    customer.save().then(()=>{
        res.send('new customer created');
    }).catch(err =>{

        if (err){
            throw err;
        }
    });



});

// Get customer
app.get('/customers', (req, res) =>{

    Customer.find().then((customer) => {

        if (customer.length > 0) {
            res.json(customer);
        }else{
            res.json({"data": [], "error": "book not found"})
        }

    }).catch((err) =>{

        if (err){
            throw err;
        }

    })
});


// Get customer by ID
app.get('/customer/:id', (req, res) =>{
    Customer.findById(req.params.id).then((customer) =>{

        if(customer){
            res.json(customer);
        }else{
            res.sendStatus(404);
        }

    }).catch(err =>{

        if (err){

            throw err;
        }
    });
})


// Delete customer
app.delete('/customer/:id', (req, res) =>{
    Customer.findByIdAndRemove(req.params.id).then((customer) =>{
        res.json({"msg": "customer has been deleted"})
    }).catch(err =>{

        if (err){

            throw err;
        }
    });
})

app.listen(8081, ()=> {
    console.log('Customer server is running.....!');
});
