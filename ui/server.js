const express = require('express');
const next = require('next');
const { createProxyMiddleware } = require('http-proxy-middleware');

const dev = process.env.NODE_ENV !== 'production';
const app = next({ dev });
const handle = app.getRequestHandler();

const API_SERVER_URL = process.env.ML_SERVER_HOST || "http://localhost:8000"; // Change this to your other server's URL

app.prepare().then(() => {
    const server = express();

    // Proxy any requests to /api/* to the API server
    server.use('/api', createProxyMiddleware({
        target: API_SERVER_URL,
        changeOrigin: true,
        pathRewrite: {
            '^/api': '/', // if your API server routes start from the root
        },
    }));

    server.all('*', (req, res) => {
        return handle(req, res);
    });

    server.listen(3000, (err) => {
        if (err) throw err;
        console.log('> Listening on port 3000');
    });
});
