//Load require libs
const express = require('express');

//configure port
const PORT = process.argv[2] || 5000;
const APIVERSION = "v1"

//Create an instance of Express
app = express();

//Define routes
//GET /
app.get('/', (req, resp) => {
    resp.status(200).json({ "status": "ok" })
})

var router = express.Router()

    //route GET /api/v1/time
    router.get('/time', function (req, res) {
        res.status(200)
        res.type("application/json")
        res.json({ "time": new Date().toISOString().replace('T', ' ') })
    })

    //route GET /api/v1/source
    router.get('/source', (req, res) => {
        res.status(200)
        res.type("application/json")
        res.json({})
    })

// Registering /api/v1 as route for /time & /source
app.use(`/api/${APIVERSION}`, router)
//Start the server
app.listen(PORT, () => {
    console.log("Application started at %s on port %d", new Date(), PORT);
})