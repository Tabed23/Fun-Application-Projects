const IncomeSchema = require('../model/incomeModel')

//Create a new income
exports.addIncome = async (req,res) => {
    const {title,amount, category,description, date} = req.body;
    const income = IncomeSchema({
            title, amount, category, description, date
    })

    try {
        if (!title || !amount || !category || !description || !date) {
          return res.status(404).json({"msg": "Invalid data"});
        }

        if (amount <=0 || amount  === 'number') {
            return res.status(404).json({"msg": "Invalid amount"});
        }

        await income.save()

        res.status(200).json({"msg":"income successfully added"});

    }catch (err) {
        res.status(500).json({"msg": "internal  server error," + err.message});
    }
 }

//Get all the Incomes
 exports.getIncome = async (req, res) => {
    try {

        const incomes = await IncomeSchema.find().sort({createdAt:-1});
        res.status(200).json({"incomes": incomes})

    }catch (err) {
        res.status(500).json({"msg": "internal server error, cannot get data, " + err.message});
    }
}
//Delete Income
exports.deleteIncome = async (req, res) => {
    const {id} = req.params;
    IncomeSchema.findByIdAndDelete(id).then((income) =>{
        res.status(200).json({message: 'Deleted income successfully'})
    }).catch((err) =>{
        res.status(500).json({message: 'invalid income, error: ' + err.message})
    })
}
