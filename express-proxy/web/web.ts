import express from "express"
import dotenv from "dotenv"

dotenv.config()

const PROXY_PORT = process.env.PROXY_PORT
const WEB_PORT = process.env.WEB_PORT
const app = express()

app.all("/*", (req,res)=>{
    console.log("came to all")
    res.send("hello from "+WEB_PORT)
})

console.log("WEB at ", WEB_PORT)
app.listen(WEB_PORT)

