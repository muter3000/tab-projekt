import express from 'express'
const app = express()
import cors from 'cors'
import fs from 'fs'
import path from 'path'
import administracja from './routes/administracja.js'
import kierowcy from './routes/kierowcy.js'
import marka from './routes/marka.js'
import stanowisko_administracyjne from './routes/stanowisko_administracyjne.js'
import pojazd from './routes/pojazd.js'
import trasa from './routes/trasa.js'
import kategoria_prawa_jazdy from './routes/kategoria_prawa_jazdy.js'

const ROUTES_FOLDER_NAME = "routes"

const router = express.Router()
router.use('/kierowcy', kierowcy)
router.use('/pojazd',pojazd)


const PORT = process.env.PORT ?? 5000

app.use(router)

app.use(express.json())
app.use(cors({
    origin: '*'
}))

app.listen(PORT, () => {
    console.log(`server is listening on port ${PORT}`)
})