import Container from 'react-bootstrap/Container';
import Badge from 'react-bootstrap/Badge';
import Button from 'react-bootstrap/Button';
import Navbar from 'react-bootstrap/Navbar';
import logo from './img/logo.svg';

function Header() {
  return (
    <Navbar className="bg-body-tertiary" >
      <Container>
        <img
          src={logo}
          width="30"
          height="30"
          className="d-inline-block align-top"
          alt="React Bootstrap logo"
        />
        <Navbar.Brand href="#home">INDIGO</Navbar.Brand>
        <Navbar.Toggle />
        <Navbar.Collapse className="justify-content-end">
          <Button variant="info"> Notification <Badge bg="secondary">9</Badge>
            <span className="visually-hidden">unread messages</span>
          </Button>
          <div style={{ width:'20px'}}></div>
          <Navbar.Text>
            Signed in as: <a href="#login">Mark Otto</a>
          </Navbar.Text>
        </Navbar.Collapse>
      </Container>
    </Navbar>
  );
}

export default Header;