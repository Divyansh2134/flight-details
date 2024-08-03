import  { useState } from 'react';
import { Typeahead } from 'react-bootstrap-typeahead';
import 'react-bootstrap-typeahead/css/Typeahead.css';
import { Form, Button, Col, Row, Card } from 'react-bootstrap';
import 'bootstrap/dist/css/bootstrap.min.css';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';
import FlightCard from './FlightCard'; 

const FlightSearch = () => {
  const [from, setFrom] = useState('');
  const [to, setTo] = useState('');
  const [departureDate, setDepartureDate] = useState(null);
  const [returnDate, setReturnDate] = useState(null);
  const [flightClass, setFlightClass] = useState('Economy');
  const [flights, setFlights] = useState([]);

  const mockPlaces = ['New York', 'Los Angeles', 'Chicago', 'Houston'];

  const handleSearch = () => {
    // Mock search function that returns flights based on the search criteria
    const mockFlights = [
      { from: 'New York', to: 'Los Angeles', date: '2023-05-01', class: 'Economy', price: 300 },
      { from: 'Chicago', to: 'Houston', date: '2023-06-15', class: 'Business', price: 500 },
      // ... other mock flight data
    ];
    const results = mockFlights.filter(flight => 
      flight.from.includes(from) &&
      flight.to.includes(to) &&
      flight.class === flightClass
    );
    setFlights(results);
  };

  return (
    <div className="container mt-5">
      <Card className="mb-4">
        <Card.Body>
          <Card.Title className="text-center">Flight Search</Card.Title>
          <Form>
            <Row className="mb-3">
              <Col md={3}>
                <Typeahead
                  id="from-where"
                  onChange={setFrom}
                  options={mockPlaces}
                  placeholder="From where?"
                />
              </Col>
              <Col md={3}>
                <Typeahead
                  id="where-to"
                  onChange={setTo}
                  options={mockPlaces}
                  placeholder="Where to?"
                />
              </Col>
              <Col md={2}>
                <DatePicker
                  selected={departureDate}
                  onChange={date => setDepartureDate(date)}
                  selectsStart
                  startDate={departureDate}
                  endDate={returnDate}
                  minDate={new Date()}
                  placeholderText="Departure"
                  className="form-control"
                />
              </Col>
              <Col md={2}>
                <DatePicker
                  selected={returnDate}
                  onChange={date => setReturnDate(date)}
                  selectsEnd
                  startDate={departureDate}
                  endDate={returnDate}
                  minDate={departureDate}
                  placeholderText="Return"
                  className="form-control"
                />
              </Col>
              <Col md={1}>
                <Form.Control as="select" value={flightClass} onChange={e => setFlightClass(e.target.value)}>
                  <option>Economy</option>
                  <option>Business</option>
                  {/* Add more classes as needed */}
                </Form.Control>
              </Col>
              <Col md={1}>
                <Button variant="primary" onClick={handleSearch}>Search</Button>
              </Col>
            </Row>
          </Form>
        </Card.Body>
      </Card>

      {/* Displaying the mock flight results */}
      <Card>
        <Card.Body>
          <Card.Title className="text-center">Search Results</Card.Title>
          <div>
            {flights.length > 0 ? (
              flights.map((flight, index) => (
                <FlightCard key={index} flight={flight} />
              ))
            ) : (
              <p className="text-center">No flights found.</p>
            )}
          </div>
        </Card.Body>
      </Card>
    </div>
  );
};

export default FlightSearch;
