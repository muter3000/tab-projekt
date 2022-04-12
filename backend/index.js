const express = require('express')
const app = express()
const cors = require('cors')

const PORT = process.env.PORT ?? 5000

app.use(express.json())
app.use(cors({
    origin: '*'
}))

const kierowcy = require('./routes/kierowcy')
const stanowisko_administracyjne = require('./routes/stanowisko_administracyjne')

app.use('/kierowcy', kierowcy);
app.use('/stanowisko_administracyjne', stanowisko_administracyjne)

app.listen(PORT, () => {
    console.log(`server is listening on port ${PORT}`)
})