const express = require('express')
const pool = require('../pool')
const router = express.Router()

router.get('/', async (req, res) => {
    try {
        const allCategories = await pool.getInstance().query('SELECT * FROM kategoria_prawa_jazdy')
        
        return res.status(200).json(allCategories.rows)
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
        const specificCategory = await pool.getInstance().query('SELECT * FROM kategoria_prawa_jazdy WHERE id = $1', [id])

        if (specificCategory.rows.length === 0) return res.sendStatus(404)
        
        return res.status(200).json(specificCategory.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.post('/', async (req, res) => {
    const { kategoria } = req.body

    if (!typ) return res.sendStatus(400)

    try {
        const addedCategory = await pool.getInstance().query('INSERT INTO kategoria_prawa_jazdy (kategoria) VALUES ($1) RETURNING *', 
        [kategoria])
        
        return res.status(200).json(addedCategory.rows[0])
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
        const deletedCategory = await pool.getInstance().query('DELETE FROM kategoria_prawa_jazdy WHERE id = $1 RETURNING *', [id])

        if (deletedCategory.rows.length === 0) return res.sendStatus(404)

        return res.status(200).json(deletedCategory.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.put('/', async (req, res) => {
    const { id, kategoria } = req.body

    if (!id || !kategoria) return res.sendStatus(400)

    try {
        const specificCategory = typeof id === 'number' ?
            await pool.getInstance().query('SELECT * FROM kategoria_prawa_jazdy WHERE id = $1', [id]) :
            {rows: []}
        let putCategory

        if (specificCategory.rows.length === 0) {
            putCategory = await pool.getInstance().query('INSERT INTO kategoria_prawa_jazdy (kategoria) VALUES ($1) RETURNING *', 
            [kategoria])
        }
        else {
            putCategory = await pool.getInstance().query('UPDATE kategoria_prawa_jazdy SET kategoria = $1 WHERE id = $2 RETURNING *',
            [kategoria, id])
        }
        
        return res.status(200).json(putCategory.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.patch('/', async (req, res) => {
    if (typeof req.body.id !== 'number') return res.sendStatus(400)

    try {
        const specificCategory = await pool.getInstance().query('SELECT * FROM kategoria_prawa_jazdy WHERE id = $1', [req.body.id])

        if (specificCategory.rows.length === 0) return res.sendStatus(404)

        const { id, kategoria } = {...specificCategory.rows[0], ...req.body}

        const updatedCategory = await pool.getInstance().query('UPDATE kategoria_prawa_jazdy SET kategoria = $1 WHERE id = $2 RETURNING *',
        [kategoria, id])

        return res.status(200).json(updatedCategory.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

module.exports = router