const express = require('express')
const app = express()
const cors = require('cors')
const fs = require('fs')
const path = require('path');

const ROUTES_FOLDER_NAME = "routes"

const prepareRoutes = (baseFolderName, insideFolderName = "") => {
    if (baseFolderName === undefined) return;
    const calculatedPath = path.join(baseFolderName, insideFolderName)

    fs.readdir(`./${calculatedPath}`, (err, files) => {
        if (err) return;
        files.forEach(file => {
            const nameSplitted = file.split('.')
            const fileNameWithoutExtension = nameSplitted[0]
            const fromPath = path.join(calculatedPath, file);
    
            fs.stat(fromPath, (error, stat) => {
                if (error) return;
                if (stat.isDirectory()) {
                    prepareRoutes(calculatedPath, file)
                    return;
                }
                
                const extension = nameSplitted[nameSplitted.length - 1]
                if (extension !== 'js') return;
    
                const requiredModule = require('./' + fromPath);
                const pathWithoutBaseArgument = calculatedPath.replace('\\', '/').substring(ROUTES_FOLDER_NAME.length)
                app.use(`${pathWithoutBaseArgument}/${fileNameWithoutExtension}`, requiredModule)
            })
        })
    })
}

prepareRoutes(ROUTES_FOLDER_NAME)

const PORT = process.env.PORT ?? 5000

app.use(express.json())
app.use(cors({
    origin: '*'
}))

app.listen(PORT, () => {
    console.log(`server is listening on port ${PORT}`)
})