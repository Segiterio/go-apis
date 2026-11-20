/**
 * Get started with your virtual shopping cart.
 *   1. POST /cart/items   add items to buy
 *   2. GET  /cart         see what's in the cart
 *   3. POST /checkout     place the order
 * 
 * Click Start to build your cart and checkout.
 * 
 * Add more routes and annotate using @endpoint <method> /<path>.
 * You can also select existing requests and examples to mock.
 * 
 * Deploying this mock gives you a url which can be shared, used in CI and more.
 */
const http = require("http");
const PORT = process.env.PORT || 4500;
const JSON_HEADER = { "Content-Type": "application/json" };

const server = http.createServer(async (req, res) => {
    const { method, url } = req;

    // @endpoint GET /health
    if (method === "GET" && url === "/health") {
        res.writeHead(200, JSON_HEADER);
        return res.end(JSON.stringify({ status: "ok" }));
    }

    // @endpoint GET /listings
    if (method === "GET" && url === "/listings") {
        res.writeHead(200, JSON_HEADER);
        return res.end(JSON.stringify({ status: "ok" }));
    }

    // @endpoint POST /listings
    if (method === "POST" && url === "/listings") {
        res.writeHead(200, JSON_HEADER);
        return res.end(JSON.stringify({ status: "ok" }));
    }

    // @endpoint GET /listings/:id
    if (method === "GET" && url === "/listings/:id") {
        res.writeHead(200, JSON_HEADER);
        return res.end(JSON.stringify({ status: "ok" }));
    }

    // @endpoint DELETE /listings/:id
    if (method === "GET" && url === "/listings/:id") {
        res.writeHead(200, JSON_HEADER);
        return res.end(JSON.stringify({ status: "ok" }));
    }

    // Select requests and examples to mock
    if (pm.mock.matchRequest("pm-stable:req:GET /health", req)) {
        return pm.mock.sendExample("pm-stable:ex:GET /health#200#Success", res);
    }

    res.writeHead(404, JSON_HEADER);
    res.end(JSON.stringify({
        error: "Endpoint not defined",
        message: `Please add ${url} endpoint to this mock.`
    }));
});

server.listen(PORT, () => console.log(`Mock server on port ${PORT}`));