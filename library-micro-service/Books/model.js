const mongoose = require('mongoose');

mongoose.model('book', {
   title: {
        type: String,
     required: true
   },

   author : {
    type: String,
    required: true
   },

   numberPages : {
    type: Number,
    required: true
   },

   publisher : {
    type: String,
    required: true
   },

});

