const express = require('express')
const pool = require('../pool')
const router = express.Router()

router.get('/', async (req, res) => {
    try {
        const allCars = await pool.getInstance().query('SELECT * FROM pojazd')
        
        return res.status(200).json(allCars.rows)
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.get('/:id', async (req, res) => {
    const id = parseInt(req.params.id)
    
    if (typeof id !== 'number') return res.sendStatus(400)

    try {
        const specificCar = await pool.getInstance().query('SELECT * FROM pojazd WHERE id = $1', [id])

        if (specificCar.rows.length === 0) return res.sendStatus(404)
        
        return res.status(200).json(specificCar.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.post('/', async (req, res) => {
    const { numer_silnika, pojemnosc_silnika, marka_id } = req.body

    if (!numer_silnika || !pojemnosc_silnika || !marka_id) return res.sendStatus(400)

    try {
        const addedCar = await pool.getInstance().query('INSERT INTO pojazd (numer_silnika, pojemnosc_silnika, marka_id) VALUES ($1, $2, $3) RETURNING *', 
        [numer_silnika, pojemnosc_silnika, marka_id])
        
        return res.status(200).json(addedCar.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.delete('/:id', async (req, res) => {
    const id = parseInt(req.params.id)

    if (typeof id !== 'number') return res.sendStatus(400)

    try {
        const deletedCar = await pool.getInstance().query('DELETE FROM pojazd WHERE id = $1 RETURNING *', [id])

        if (deletedCar.rows.length === 0) return res.sendStatus(404)

        return res.status(200).json(deletedCar.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.put('/', async (req, res) => {
    const { id, numer_silnika, pojemnosc_silnika, marka_id } = req.body

    if (!id || !numer_silnika || !pojemnosc_silnika || !marka_id) return res.sendStatus(400)

    try {
        const specificCar = typeof id === 'number' ?
            await pool.getInstance().query('SELECT * FROM pojazd WHERE id = $1', [id]) :
            {rows: []}
        let putCar

        if (specificCar.rows.length === 0) {
            putCar = await pool.getInstance().query('INSERT INTO pojazd (numer_silnika, pojemnosc_silnika, marka_id) VALUES ($1, $2, $3) RETURNING *', 
            [numer_silnika, pojemnosc_silnika, marka_id])
        }
        else {
            putCar = await pool.getInstance().query('UPDATE pojazd SET numer_silnika = $1, pojemnosc_silnika = $2, marka_id = $3 WHERE id = $4 RETURNING *',
            [numer_silnika, pojemnosc_silnika, marka_id, id])
        }
        
        return res.status(200).json(putCar.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.patch('/', async (req, res) => {
    if (typeof req.body.id !== 'number') return res.sendStatus(400)

    try {
        const specificCar = await pool.getInstance().query('SELECT * FROM pojazd WHERE id = $1', [req.body.id])

        if (specificCar.rows.length === 0) return res.sendStatus(404)

        const { id, numer_silnika, pojemnosc_silnika, marka_id } = {...specificCar.rows[0], ...req.body}

        const updatedCar = await pool.getInstance().query('UPDATE pojazd SET numer_silnika = $1, pojemnosc_silnika = $2, marka_id = $3 WHERE id = $4 RETURNING *',
        [numer_silnika, pojemnosc_silnika, marka_id, id])

        return res.status(200).json(updatedCar.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

module.exports = router