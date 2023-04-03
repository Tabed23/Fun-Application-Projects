const mongoose = require('mongoose');

const db = async () => {
    try {
        mongoose.set('strictQuery', false);
        await mongoose.connect(process.env.MONGO_URL)
        console.log('mongos connection established')
    }catch (error) {
        console.log('cannot connect to Mongo')

    }
}

module.exports = {db}
