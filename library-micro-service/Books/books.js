const express = require('express');
const app = express();
const mongoose = require('mongoose');
const bodyparser = require('body-parser');


require('./model')

const books = mongoose.model('book');

mongoose.connect("mongodb://localhost:27017/bookservice");

app.use(bodyparser.json());


app.get('/', (req, res) =>{
    res.send("health check")
});


// Post Book
app.post('/book', (req, res) =>{

    var newbook  =  {
        title: req.body.title,
        author: req.body.author,
        numberPages :  req.body.   numberPages,
        publisher : req.body.publisher
    }

    var book =  new books(newbook);

    book.save();

    res.send('new book created');

});

// Get Books
app.get('/books', (req, res) =>{

    books.find().then((books) => {

        if (books.length > 0) {
            res.json(books);
        }else{
            res.json({"data": [], "error": "book not found"})
        }

    }).catch((err) =>{

        if (err){
            throw err;
        }

    })
});


// Get Book by ID
app.get('/book/:id', (req, res) =>{
    books.findById(req.params.id).then((book) =>{

        if(book){
            res.json(book);
        }else{
            res.sendStatus(404);
        }

    }).catch(err =>{

        if (err){

            throw err;
        }
    });
})


// Delete Book
app.delete('/book/:id', (req, res) =>{
    books.findByIdAndRemove(req.params.id).then((book) =>{
        res.json({"msg": "book has been deleted"})
    }).catch(err =>{

        if (err){

            throw err;
        }
    });
})

app.listen(8080, ()=> {
    console.log('Books server is running.....!');
});
