import express from 'express'
import {Pool as pool} from '../pool.js'
const router = express.Router({mergeParams: true})

router.get('/', async (req, res) => {
    try {
        const allRoutes = await pool.getInstance().query('SELECT * FROM trasa')
        
        return res.status(200).json(allRoutes.rows)
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
        const specificRoute = await pool.getInstance().query('SELECT * FROM trasa WHERE id = $1', [id])

        if (specificRoute.rows.length === 0) return res.sendStatus(404)
        
        return res.status(200).json(specificRoute.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.post('/', async (req, res) => {
    const { miejscowosc_poczatkowa, miejscowosc_koncowa } = req.body

    if (!miejscowosc_poczatkowa || !miejscowosc_koncowa) return res.sendStatus(400)

    try {
        const addedRoute = await pool.getInstance().query('INSERT INTO trasa (miejscowosc_poczatkowa, miejscowosc_koncowa) VALUES ($1, $2) RETURNING *', 
        [miejscowosc_poczatkowa, miejscowosc_koncowa])
        
        return res.status(200).json(addedRoute.rows[0])
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
        const deletedRoute = await pool.getInstance().query('DELETE FROM trasa WHERE id = $1 RETURNING *', [id])

        if (deletedRoute.rows.length === 0) return res.sendStatus(404)

        return res.status(200).json(deletedRoute.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.put('/', async (req, res) => {
    const { id, miejscowosc_poczatkowa, miejscowosc_koncowa } = req.body

    if (!id || !miejscowosc_poczatkowa || !miejscowosc_koncowa) return res.sendStatus(400)

    try {
        const specificRoute = typeof id === 'number' ?
            await pool.getInstance().query('SELECT * FROM trasa WHERE id = $1', [id]) :
            {rows: []}
        let putRoute

        if (specificRoute.rows.length === 0) {
            putRoute = await pool.getInstance().query('INSERT INTO trasa (miejscowosc_poczatkowa, miejscowosc_koncowa) VALUES ($1) RETURNING *', 
            [miejscowosc_poczatkowa, miejscowosc_koncowa])
        }
        else {
            putRoute = await pool.getInstance().query('UPDATE trasa SET miejscowosc_poczatkowa = $1, miejscowosc_koncowa = $2 WHERE id = $3 RETURNING *',
            [miejscowosc_poczatkowa, miejscowosc_koncowa, id])
        }
        
        return res.status(200).json(putRoute.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

router.patch('/', async (req, res) => {
    if (typeof req.body.id !== 'number') return res.sendStatus(400)

    try {
        const specificRoute = await pool.getInstance().query('SELECT * FROM trasa WHERE id = $1', [req.body.id])

        if (specificRoute.rows.length === 0) return res.sendStatus(404)

        const { id, miejscowosc_poczatkowa, miejscowosc_koncowa } = {...specificRoute.rows[0], ...req.body}

        const updatedRoute = await pool.getInstance().query('UPDATE trasa SET miejscowosc_poczatkowa = $1, miejscowosc_koncowa = $2 WHERE id = $3 RETURNING *',
        [miejscowosc_poczatkowa, miejscowosc_koncowa, id])

        return res.status(200).json(updatedRoute.rows[0])
    }
    catch(err) {
        console.log(err)
        return res.sendStatus(500)
    }
})

export default router