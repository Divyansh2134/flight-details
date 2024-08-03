import { Card, Button } from 'react-bootstrap';
import PropTypes from 'prop-types';

const FlightCard = ({ flight }) => {
  return (
    <Card className="mb-3">
      <Card.Body>
        <Card.Title>Flight from {flight.from} to {flight.to}</Card.Title>
        <Card.Subtitle className="mb-2 text-muted">Date: {flight.date}</Card.Subtitle>
        <Card.Text>Class: {flight.class}</Card.Text>
        <Card.Text>Price: ${flight.price}</Card.Text>
        <Button variant="primary">Book Flight</Button>
      </Card.Body>
    </Card>
  );
};

FlightCard.propTypes = {
    flight: PropTypes.shape({
      from: PropTypes.string.isRequired,
      to: PropTypes.string.isRequired,
      date: PropTypes.string.isRequired,
      class: PropTypes.string.isRequired,
      price: PropTypes.number.isRequired
    }).isRequired
  };

export default FlightCard;
