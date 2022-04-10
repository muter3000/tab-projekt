const express = require('express')
const app = express()
const cors = require('cors')

const PORT = process.env.PORT ?? 5000

app.use(express.json())
app.use(cors({
    origin: '*'
}))

const kierowcy = require('./routes/kierowcy')

app.use('/kierowcy', kierowcy);

app.listen(PORT, () => {
    console.log(`server is listening on port ${PORT}`)
})