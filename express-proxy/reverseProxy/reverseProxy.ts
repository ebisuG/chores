import express from "express"
import proxy from "express-http-proxy"
import dotenv from "dotenv"

dotenv.config()

const PROXY_PORT = process.env.PROXY_PORT
const API_PORT = process.env.API_PORT
const WEB_PORT = process.env.WEB_PORT
const app = express()

//Type http://localhost:50002/proxy to browser URL, get "hello from "+WEB_PORT
app.use('/proxy', proxy("localhost:"+WEB_PORT))
// app.use('/proxy', (req, res)=>{
// // app.use('/proxy', (req, res, next)=>{
//     console.log("hostname : ", req.hostname)
//     proxy("localhost:"+WEB_PORT)
//     // proxy(req.hostname+":"+WEB_PORT)
//     // next()
// })
console.log("proxy at ", PROXY_PORT)
console.log("proxy to : ", "localhost:"+WEB_PORT)
app.listen(PROXY_PORT)


