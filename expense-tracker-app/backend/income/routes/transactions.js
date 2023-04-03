const router = require('express').Router();
const {addIncome, getIncome, deleteIncome} = require('../controller/income')

router.get('/', (req, res) => {
    res.send('Hello, world!');
})

router.post('/add-income', addIncome)
router.get('/get-incomes', getIncome)
router.delete('/delete-income/:id', deleteIncome)

module.exports = router
