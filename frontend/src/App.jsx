import { useState } from 'react'
import './App.css'
import Header from './components/Header'
import FlightStatus from './components/HomePage'

function App() {

  return (
    <>
      <Header></Header >
      <FlightStatus></FlightStatus>
    </>
  )
}

export default App
