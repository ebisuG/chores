import express from "express"
import dotenv from "dotenv"

dotenv.config()

const API_PORT = process.env.API_PORT
const app = express()

app.all("/api/*", (req,res)=>{
    console.log("came to all")
    res.send("hello from api, at "+API_PORT)
})

console.log("api at ", API_PORT)
app.listen(API_PORT)

