// src/FlightStatus.js
import React from 'react';
import { Table, Container, Row, Col } from 'react-bootstrap';

const mockData = [
  { flight: 'AA123', status: 'On Time', gate: 'A1', remarks: 'Boarding' },
  { flight: 'UA456', status: 'Delayed', gate: 'B2', remarks: 'Delayed 30 mins' },
  { flight: 'DL789', status: 'Cancelled', gate: 'C3', remarks: 'Cancelled' },
  { flight: 'SW101', status: 'Gate Change', gate: 'D4', remarks: 'Gate changed to D4' }
];

const FlightStatus = () => {
  return (
    <Container>
      <Row className="my-4">
        <Col>
          <h1 className="text-center">Current Flight Status</h1>
          <Table striped bordered hover>
            <thead>
              <tr>
                <th>Flight</th>
                <th>Status</th>
                <th>Gate</th>
                <th>Remarks</th>
              </tr>
            </thead>
            <tbody>
              {mockData.map((flight, index) => (
                <tr key={index}>
                  <td>{flight.flight}</td>
                  <td>{flight.status}</td>
                  <td>{flight.gate}</td>
                  <td>{flight.remarks}</td>
                </tr>
              ))}
            </tbody>
          </Table>
        </Col>
      </Row>
    </Container>
  );
};

export default FlightStatus;
