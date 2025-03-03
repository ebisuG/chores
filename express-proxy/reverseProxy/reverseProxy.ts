import express from 'express';
import dotenv from "dotenv"
import { createProxyMiddleware } from 'http-proxy-middleware';
import type { Request, Response, NextFunction } from 'express';

dotenv.config()

const PROXY_PORT = process.env.PROXY_PORT
const API_PORT = process.env.API_PORT
const WEB_PORT = process.env.WEB_PORT
const app = express()

const proxyWeb = createProxyMiddleware<Request, Response>({
    target: `http://localhost:${WEB_PORT}/webserver/`,
    changeOrigin: true,
})
const proxyApi = createProxyMiddleware<Request, Response>({
    target: `http://localhost:${API_PORT}/api/`,
    changeOrigin: true,
  })

app.use('/proxy/web/', proxyWeb);
app.use('/proxy/api/', proxyApi);
console.log("proxy at ", PROXY_PORT)
console.log("proxy to : ", "localhost:"+WEB_PORT)
app.listen(PROXY_PORT)


