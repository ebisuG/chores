import './App.css'
import { useEffect, useState } from 'react'

function FetchInServer() {
    const [timer, setTimer] = useState("init")
    const [valueFromApi, setValueFromApi] = useState("")

    useEffect(() => {
        const timeoutId = setTimeout(() => {
            setTimer("set!")
            console.log("setTimeout!")
        }, 1 * 5000)
        const fetchData = async()=>{
            const data = await waitApiResult()
            setValueFromApi(data)
        }
        fetchData().catch(console.error)
        return () => clearTimeout(timeoutId)
    }, [])

    async function waitApiResult(){
        const result = await fetch("http://localhost:3000/mock")
        console.log("called")
        console.log("result : ",result)
        return result
    }

    return (
        <>
        <div>
            Html element before timer<br></br>
            {timer}<br></br>
            Result from Api.
            {valueFromApi}
        </div>
        </>
    )
}

export default FetchInServer
