import { useState, useEffect } from 'react';
import Container from 'react-bootstrap/Container';
import Badge from 'react-bootstrap/Badge';
import Button from 'react-bootstrap/Button';
import Navbar from 'react-bootstrap/Navbar';
import Tooltip from 'react-bootstrap/Tooltip';
import OverlayTrigger from 'react-bootstrap/OverlayTrigger';
import logo from './img/logo.svg';

// eslint-disable-next-line react/prop-types
function Header({websocketUrl}) {

  const [notifications, setNotifications] = useState(() => {
    const savedNotifications = sessionStorage.getItem('notifications');
    return savedNotifications ? JSON.parse(savedNotifications) : 0;
  });
  const [message, setMessage] = useState(() => {
    const savedMessage = sessionStorage.getItem('message');
    return savedMessage ? savedMessage : '';
  });

  useEffect(() => {
    const ws = new WebSocket(websocketUrl);

    ws.onmessage = (event) => {
      const receivedMessage = event.data;
      setNotifications((prev) => {
        const newCount = prev + 1;
        localStorage.setItem('notifications', JSON.stringify(newCount));
        return newCount;
      });
      setMessage(receivedMessage);
      localStorage.setItem('message', receivedMessage);
    };

    return () => {
      ws.close();
    };
  }, [websocketUrl]);

  const renderTooltip = (props) => (
    <Tooltip id="button-tooltip" {...props}>
      {message ? message : 'No new notifications'}
    </Tooltip>
  );

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
        <OverlayTrigger
            placement="bottom"
            delay={{ show: 250, hide: 400 }}
            overlay={renderTooltip}
          >
            <Button variant="info">
              Notification <Badge bg="secondary">{notifications}</Badge>
              <span className="visually-hidden">unread messages</span>
            </Button>
          </OverlayTrigger>
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