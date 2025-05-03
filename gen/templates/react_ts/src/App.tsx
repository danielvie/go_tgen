import { useState } from 'react'
import './App.css'

function App() {
  const [count, set_ount] = useState(0)

  return (
    <>
      <div>
        <h1>my app here!!</h1>
        <div className='mt-6'>
          <button onClick={() => set_ount((c) => c + 1)}>
            count : {count}
          </button>
        </div>
      </div>
    </>
  )
}

export default App
