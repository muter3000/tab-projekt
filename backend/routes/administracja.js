const express = require('express')
const pool = require('../pool')
const router = express.Router()

router.get('/', async (req, res) => {
    try {
        const allAdministrations = await pool.getInstance().query('SELECT * FROM administracja')
        
        return res.status(200).json(allAdministrations.rows)
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.get('/:administracja_id', async (req, res) => {
    const administracja_id = parseInt(req.params.administracja_id)
    
    if (typeof administracja_id !== 'number') return res.sendStatus(400)

    try {
        const specificAdministration = await pool.getInstance().query('SELECT * FROM administracja WHERE administracja_id = $1', [administracja_id])

        if (specificAdministration.rows.length === 0) return res.sendStatus(404)
        
        return res.status(200).json(specificAdministration.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.post('/', async (req, res) => {
    const { pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id } = req.body

    if (!stanowisko_administracyjne_id || !login || !haslo || !pesel || !imie || !nazwisko || pesel.length !== 11) return res.sendStatus(400)

    try {
        const addedAdministration = await pool.getInstance().query('INSERT INTO administracja (pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *', 
        [pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id])
        
        return res.status(200).json(addedAdministration.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.delete('/:administracja_id', async (req, res) => {
    const administracja_id = parseInt(req.params.administracja_id)

    if (typeof administracja_id !== 'number') return res.sendStatus(400)

    try {
        const deletedAdministration = await pool.getInstance().query('DELETE FROM administracja WHERE administracja_id = $1 RETURNING *', [administracja_id])

        if (deletedAdministration.rows.length === 0) return res.sendStatus(404)

        return res.status(200).json(deletedAdministration.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.put('/', async (req, res) => {
    const { administracja_id, pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id } = req.body

    if (!administracja_id || !stanowisko_administracyjne_id || !login || !haslo || !pesel || !imie || !nazwisko || pesel.length !== 11) return res.sendStatus(400)

    try {
        const specificAdministration =  typeof administracja_id === 'number' ? 
            await pool.getInstance().query('SELECT * FROM administracja WHERE administracja_id = $1', [administracja_id]) :
            {rows: []}
        let putAdministration

        if (specificAdministration.rows.length === 0) {
            putAdministration = await pool.getInstance().query('INSERT INTO administracja (pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *', 
            [pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id])
        }
        else {
            putAdministration = await pool.getInstance().query('UPDATE administracja SET pesel = $1, imie = $2, nazwisko = $3, login = $4, haslo = $5, stanowisko_administracyjne_id = $6 WHERE administracja_id = $7 RETURNING *',
            [pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id, administracja_id])
        }
        
        return res.status(200).json(putAdministration.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

router.patch('/', async (req, res) => {
    if (typeof req.body.administracja_id !== 'number' || (!!req.body.pesel && req.body.pesel.length !== 11)) return res.sendStatus(400)

    try {
        const specificAdministration = await pool.getInstance().query('SELECT * FROM administracja WHERE administracja_id = $1', [req.body.administracja_id])

        if (specificAdministration.rows.length === 0) return res.sendStatus(404)

        const { administracja_id, pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id } = {...specificAdministration.rows[0], ...req.body}

        const updatedAdministration = await pool.getInstance().query('UPDATE administracja SET pesel = $1, imie = $2, nazwisko = $3, login = $4, haslo = $5, stanowisko_administracyjne_id = $6 WHERE administracja_id = $7 RETURNING *',
        [pesel, imie, nazwisko, login, haslo, stanowisko_administracyjne_id, administracja_id])

        return res.status(200).json(updatedAdministration.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.status(500)
    }
})

module.exports = router