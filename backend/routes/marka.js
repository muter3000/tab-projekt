import express from 'express'
import {Pool as pool} from '../pool.js'
const router = express.Router({mergeParams: true})

router.get('/', async (req, res) => {
    try {
        const allBrands = await pool.getInstance().query('SELECT * FROM marka')
        
        return res.status(200).json(allBrands.rows)
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
        const specificBrand = await pool.getInstance().query('SELECT * FROM marka WHERE id = $1', [id])

        if (specificBrand.rows.length === 0) return res.sendStatus(404)
        
        return res.status(200).json(specificBrand.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.post('/', async (req, res) => {
    const { nazwa } = req.body

    if (!nazwa) return res.sendStatus(400)

    try {
        const addedBrand = await pool.getInstance().query('INSERT INTO marka (nazwa) VALUES ($1) RETURNING *', 
        [nazwa])
        
        return res.status(200).json(addedBrand.rows[0])
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
        const deletedBrand = await pool.getInstance().query('DELETE FROM marka WHERE id = $1 RETURNING *', [id])

        if (deletedBrand.rows.length === 0) return res.sendStatus(404)

        return res.status(200).json(deletedBrand.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.put('/', async (req, res) => {
    const { id, nazwa } = req.body

    if (!id || !nazwa) return res.sendStatus(400)

    try {
        const specificBrand = typeof id === 'number' ?
            await pool.getInstance().query('SELECT * FROM marka WHERE id = $1', [id]) :
            {rows: []}
        let putBrand

        if (specificBrand.rows.length === 0) {
            putBrand = await pool.getInstance().query('INSERT INTO marka (nazwa) VALUES ($1) RETURNING *', 
            [nazwa])
        }
        else {
            putBrand = await pool.getInstance().query('UPDATE marka SET nazwa = $1 WHERE id = $2 RETURNING *',
            [nazwa, id])
        }
        
        return res.status(200).json(putBrand.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.patch('/', async (req, res) => {
    if (typeof req.body.id !== 'number') return res.sendStatus(400)

    try {
        const specificBrand = await pool.getInstance().query('SELECT * FROM marka WHERE id = $1', [req.body.id])

        if (specificBrand.rows.length === 0) return res.sendStatus(404)

        const { id, nazwa } = {...specificBrand.rows[0], ...req.body}

        const updatedBrand = await pool.getInstance().query('UPDATE marka SET nazwa = $1 WHERE id = $2 RETURNING *',
        [nazwa, id])

        return res.status(200).json(updatedBrand.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

export default router