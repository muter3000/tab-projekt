const express = require('express')
const pool = require('../pool')
const router = express.Router()

router.get('/', async (req, res) => {
    try {
        const allDrivers = await pool.getInstance().query('SELECT * FROM kierowcy')
        
        return res.status(200).json(allDrivers.rows)
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.get('/:id', async (req, res) => {
    const id = parseInt(req.params.id)
    
    if (typeof id !== 'number') return res.sendStatus(400)

    try {
        const specificDriver = await pool.getInstance().query('SELECT * FROM kierowcy WHERE id = $1', [id])

        if (specificDriver.rows.length === 0) return res.sendStatus(404)
        
        return res.status(200).json(specificDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.post('/', async (req, res) => {
    const { pesel, imie, nazwisko } = req.body

    if (!pesel || !imie || !nazwisko || pesel.length !== 11) return res.sendStatus(400)

    try {
        const addedDriver = await pool.getInstance().query('INSERT INTO kierowcy (pesel, imie, nazwisko) VALUES ($1, $2, $3) RETURNING *', 
        [pesel, imie, nazwisko])
        
        return res.status(200).json(addedDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.delete('/:id', async (req, res) => {
    const id = parseInt(req.params.id)

    if (typeof id !== 'number') return res.sendStatus(400)

    try {
        const deletedDriver = await pool.getInstance().query('DELETE FROM kierowcy WHERE id = $1 RETURNING *', [id])

        if (deletedDriver.rows.length === 0) return res.sendStatus(404)

        return res.status(200).json(deletedDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.put('/', async (req, res) => {
    const { id, pesel, imie, nazwisko } = req.body

    if (!id || !pesel || !imie || !nazwisko || pesel.length !== 11) return res.sendStatus(400)

    try {
        const specificDriver = await pool.getInstance().query('SELECT * FROM kierowcy WHERE id = $1', [id])
        let putDriver

        if (specificDriver.rows.length === 0) {
            putDriver = await pool.getInstance().query('INSERT INTO kierowcy (pesel, imie, nazwisko) VALUES ($1, $2, $3) RETURNING *', 
            [pesel, imie, nazwisko])
        }
        else {
            putDriver = await pool.getInstance().query('UPDATE kierowcy SET pesel = $1, imie = $2, nazwisko = $3 WHERE id = $4 RETURNING *',
            [pesel, imie, nazwisko, id])
        }
        
        return res.status(200).json(putDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.patch('/', async (req, res) => {
    if (typeof req.body.id !== 'number' || !!req.body.pesel && req.body.pesel.length !== 11) return res.sendStatus(400)

    try {
        const specificDriver = await pool.getInstance().query('SELECT * FROM kierowcy WHERE id = $1', [req.body.id])

        if (specificDriver.rows.length === 0) return res.sendStatus(404)

        const { id, pesel, imie, nazwisko } = {...specificDriver.rows[0], ...req.body}

        const updatedDriver = await pool.getInstance().query('UPDATE kierowcy SET pesel = $1, imie = $2, nazwisko = $3 WHERE id = $4 RETURNING *',
        [pesel, imie, nazwisko, id])

        return res.status(200).json(updatedDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

module.exports = router