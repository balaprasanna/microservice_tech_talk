const express = require('express');
const fortune = require('fortune-teller')

const app = express();

app.use((req, resp) => {
	const wiseSaying = fortune.fortune();

	resp.status(200);
	resp.format({
		"text/plain": () => {
			resp.send(wiseSaying);
		},
		"text/html": () => {
			resp.send("<h2>" + wiseSaying + "<h2>");
		},
		"application/json": () => {
			resp.send({ fortune: wiseSaying });
		},
		"default": () => {
			res.status(406).send("Not Acceptable");
		}
	});
});

const port = process.argv[2] || process.env.APP_PORT || 3000;

app.listen(port, () => {
	console.info("Application started on port %d at %s"
			, port, new Date());
});
