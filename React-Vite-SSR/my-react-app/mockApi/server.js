import express from "express";
import cors from "cors"

const app = express()

// app.use(cors);

const allowedOrigins = ['http://localhost:5137'];

const options= {
  origin: allowedOrigins
};

app.use(cors(options));

app.get("/mock", function(req, res){
    console.log("request came")
    let result = "none"
    for(let i = 0; i<=5000000000; i++){
        continue
    }
    result = "done!"
    console.log("result : ",result)
    res.send(result)
})

app.listen(3000)