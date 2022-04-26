import express from 'express'
import {Pool as pool} from '../pool.js'
const router = express.Router({mergeParams: true})

router.get('/', async (req, res) => {
    try {
        const allPositions = await pool.getInstance().query('SELECT * FROM stanowisko_administracyjne')
        
        return res.status(200).json(allPositions.rows)
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
        const specificPosition = await pool.getInstance().query('SELECT * FROM stanowisko_administracyjne WHERE id = $1', [id])

        if (specificPosition.rows.length === 0) return res.sendStatus(404)
        
        return res.status(200).json(specificPosition.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.post('/', async (req, res) => {
    const { typ } = req.body

    if (!typ) return res.sendStatus(400)

    try {
        const addedPosition = await pool.getInstance().query('INSERT INTO stanowisko_administracyjne (typ) VALUES ($1) RETURNING *', 
        [typ])
        
        return res.status(200).json(addedPosition.rows[0])
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
        const deletedPosition = await pool.getInstance().query('DELETE FROM stanowisko_administracyjne WHERE id = $1 RETURNING *', [id])

        if (deletedPosition.rows.length === 0) return res.sendStatus(404)

        return res.status(200).json(deletedPosition.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.put('/', async (req, res) => {
    const { id, typ } = req.body

    if (!id || !typ) return res.sendStatus(400)

    try {
        const specificPosition = typeof id === 'number' ?
            await pool.getInstance().query('SELECT * FROM stanowisko_administracyjne WHERE id = $1', [id]) :
            {rows: []}
        let putPosition

        if (specificPosition.rows.length === 0) {
            putPosition = await pool.getInstance().query('INSERT INTO stanowisko_administracyjne (typ) VALUES ($1) RETURNING *', 
            [typ])
        }
        else {
            putPosition = await pool.getInstance().query('UPDATE stanowisko_administracyjne SET typ = $1 WHERE id = $2 RETURNING *',
            [typ, id])
        }
        
        return res.status(200).json(putPosition.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.patch('/', async (req, res) => {
    if (typeof req.body.id !== 'number') return res.sendStatus(400)

    try {
        const specificPosition = await pool.getInstance().query('SELECT * FROM stanowisko_administracyjne WHERE id = $1', [req.body.id])

        if (specificPosition.rows.length === 0) return res.sendStatus(404)

        const { id, typ } = {...specificPosition.rows[0], ...req.body}

        const updatedPosition = await pool.getInstance().query('UPDATE stanowisko_administracyjne SET typ = $1 WHERE id = $2 RETURNING *',
        [typ, id])

        return res.status(200).json(updatedPosition.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

export default router