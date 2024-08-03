import './App.css'
import Header from './components/Header'
import FlightStatus from './components/FlightStatus'
import FlightSearch from './components/Hero'

function App() {

  const websocketUrl = import.meta.env.VITE_WEBSOCKET_URL;
  const flightStatus = import.meta.env.VITE_FLIGHT_STATUS_URL;

  return (
    <>
      <Header websocketUrl={websocketUrl} ></Header >
      <FlightSearch></FlightSearch>
      <FlightStatus flightStatusUrl={flightStatus} ></FlightStatus>
    </>
  )
}

export default App
