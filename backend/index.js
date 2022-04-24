const express = require('express')
const app = express()
const cors = require('cors')
const fs = require('fs')
const path = require('path');

const PORT = process.env.PORT ?? 5000

app.use(express.json())
app.use(cors({
    origin: '*'
}))

fs.readdir('./routes', (err, files) => {
    if (err) return;
    files.forEach(file => {
        const nameSplitted = file.split('.')
        const fileNameWithoutExtension = nameSplitted[0]
        const fromPath = path.join('routes', file);

        fs.stat(fromPath, (error, stat) => {
            if (error) return;
            if (stat.isDirectory()) return;
            
            const extension = nameSplitted[nameSplitted.length - 1]
            if (extension !== 'js') return;

            const requiredModule = require('./' + fromPath);
            
            app.use(`/${fileNameWithoutExtension}`, requiredModule)
        })
    })
})

app.listen(PORT, () => {
    console.log(`server is listening on port ${PORT}`)
})