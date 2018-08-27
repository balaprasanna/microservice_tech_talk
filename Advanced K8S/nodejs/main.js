//Load require libs
const express = require('express');
var request = require('request');


//configure port
const PORT = process.argv[2] || 5000;
const APIVERSION = "v1"
const SOURCE_URL = "http://api.pnd.gs/v1/sources/"
//Create an instance of Express
app = express();

//Define routes
//GET /
app.get('/', (req, resp) => {
    resp.status(200).json({ 
        "status": "ok",
        "hostname": process.env.host
    })
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
        request(SOURCE_URL, (error, response, body) => {
            if (error) {
                console.log('error:', error)
                res.status(500).end()
                return
            }
            var data = {
                "latest": [],
                "popular": []
            }
            body = JSON.parse(body)
            for (let index = 0; index < body.length; index++) {
                const item = body[index];
                data["latest"].push( item["endpoints"]["latest"] )
                data["popular"].push( item["endpoints"]["popular"] )  
            }

            res.status(200)
            res.type("application/json")
            res.json( data )
        });

        
    })

// Registering /api/v1 as route for /time & /source
app.use(`/api/${APIVERSION}`, router)
//Start the server
app.listen(PORT, () => {
    console.log("Application started at %s on port %d", new Date(), PORT);
})