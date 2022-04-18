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

router.get('/:kierowcy_id', async (req, res) => {
    const id = parseInt(req.params.id)
    
    if (typeof id !== 'number') return res.sendStatus(400)

    try {
        const specificDriver = await pool.getInstance().query('SELECT * FROM kierowcy WHERE kierowcy_id = $1', [kierowcy_id])

        if (specificDriver.rows.length === 0) return res.sendStatus(404)
        
        return res.status(200).json(specificDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.post('/', async (req, res) => {
    const { pesel, imie, nazwisko, login, haslo } = req.body

    if (!login || !haslo || !pesel || !imie || !nazwisko || pesel.length !== 11) return res.sendStatus(400)

    try {
        const addedDriver = await pool.getInstance().query('INSERT INTO kierowcy (pesel, imie, nazwisko, login, haslo) VALUES ($1, $2, $3, $4, $5) RETURNING *', 
        [pesel, imie, nazwisko, login, haslo])
        
        return res.status(200).json(addedDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.delete('/:kierowcy_id', async (req, res) => {
    const kierowcy_id = parseInt(req.params.kierowcy_id)

    if (typeof kierowcy_id !== 'number') return res.sendStatus(400)

    try {
        const deletedDriver = await pool.getInstance().query('DELETE FROM kierowcy WHERE kierowcy_id = $1 RETURNING *', [kierowcy_id])

        if (deletedDriver.rows.length === 0) return res.sendStatus(404)

        return res.status(200).json(deletedDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.put('/', async (req, res) => {
    const { kierowcy_id, pesel, imie, nazwisko, login, haslo } = req.body

    if (!login || !haslo || !pesel || !imie || !nazwisko || pesel.length !== 11) return res.sendStatus(400)

    try {
        const specificDriver = await pool.getInstance().query('SELECT * FROM kierowcy WHERE kierowcy_id = $1', [kierowcy_id])
        let putDriver

        if (specificDriver.rows.length === 0) {
            putDriver = await pool.getInstance().query('INSERT INTO kierowcy (pesel, imie, nazwisko, login, haslo) VALUES ($1, $2, $3, $4, $5) RETURNING *', 
            [pesel, imie, nazwisko, login, haslo])
        }
        else {
            putDriver = await pool.getInstance().query('UPDATE kierowcy SET pesel = $1, imie = $2, nazwisko = $3, login = $4, haslo = $5 WHERE kierowcy_id = $6 RETURNING *',
            [pesel, imie, nazwisko, login, haslo, kierowcy_id])
        }
        
        return res.status(200).json(putDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.patch('/', async (req, res) => {
    if (typeof req.body.kierowcy_id !== 'number' || (!!req.body.pesel && req.body.pesel.length !== 11)) return res.sendStatus(400)

    try {
        const specificDriver = await pool.getInstance().query('SELECT * FROM kierowcy WHERE kierowcy_id = $1', [req.body.kierowcy_id])

        if (specificDriver.rows.length === 0) return res.sendStatus(404)

        const { kierowcy_id, pesel, imie, nazwisko, login, haslo } = {...specificDriver.rows[0], ...req.body}

        const updatedDriver = await pool.getInstance().query('UPDATE kierowcy SET pesel = $1, imie = $2, nazwisko = $3, login = $4, haslo = $5 WHERE kierowcy_id = $6 RETURNING *',
        [pesel, imie, nazwisko, login, haslo, kierowcy_id])

        return res.status(200).json(updatedDriver.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

module.exports = router