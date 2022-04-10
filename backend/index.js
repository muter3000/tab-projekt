const express = require('express')
const app = express()
const cors = require('cors')

const PORT = process.env.PORT ?? 5000

app.use(express.json())
app.use(cors({
    origin: '*'
}))

app.listen(PORT, () => {
    console.log(`server is listening on port ${PORT}`)
})