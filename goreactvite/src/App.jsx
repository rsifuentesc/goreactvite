import React from 'react'

function App() {
  return (
<div>
    <h1>Hello world</h1>
    <button onClick={ async ()=>{
        const response = await fetch('http://localhost:8080/users')
        const data = await response.json()
        console.log(data)
    }}>Obterner datos</button>
</div>
  )
}

export default App
