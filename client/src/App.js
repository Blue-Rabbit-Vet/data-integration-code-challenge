import React, { useEffect, useState } from "react";
import { Button, Form, Table, Row, Col, Container } from "react-bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";

const App = (props) => {
  const [fetchedData, setFetchedData] = useState([]);
  const [payload, setPayload] = useState({
    petName: "",
    ownerEmail: "",
    procedure: "",
    cost: "",
    isPaid: false,
  });

  //recieves pet records from db and sets to state
  const getData = async () => {
    const response = await fetch("/api");
    if (response.ok) {
      const body = await response.json();
      console.log("Body: ", body.records);
      setFetchedData(body.records);
    } else {
      setFetchedData("Unsuccessful Fetch");
    }
  };

  useEffect(() => {
    getData();
  }, []);

  //sends form data to backend where it's persisted to sql db and kafka
  const postData = async () => {
    try {
      const response = await fetch("/api", {
        method: "POST",
        body: JSON.stringify(payload),
        headers: new Headers({
          "Content-Type": "application/json",
        }),
      });
      const body = await response.json();
      setFetchedData([...fetchedData, body.record]);
    } catch (error) {
      console.log("Eror: ", error);
    }
  };

  const onSubmit = (event) => {
    event.preventDefault();
    postData();
    setPayload({
      petName: "",
      ownerEmail: "",
      procedure: "",
      cost: "",
      isPaid: false,
    });
  };

  const onInputChange = (event) => {
    setPayload({
      ...payload,
      [event.currentTarget.name]: event.currentTarget.value,
    });
  };

  const setIsPaid = () => {
    payload.isPaid = !payload.isPaid;
    console.log("Has paid: ", payload.isPaid);
  };

  const recordTDs = fetchedData.map((row) => {
    return (
      <tr key={row.id}>
        <td>{row.petName}</td>
        <td>{row.ownerEmail}</td>
        <td>{row.procedure}</td>
        <td>{row.cost}</td>
      </tr>
    );
  });

  return (
    <Container className="appContainer">
      <h2>Blue Rabbit Code Challenge</h2>
      <Row>
        <Col>
          <Form onSubmit={onSubmit} className="formContainer">
            <Form.Group className="mb-3" controlId="formBasicPetName">
              <Form.Label>Pet Name</Form.Label>
              <Form.Control
                name="petName"
                value={payload.petName}
                onChange={onInputChange}
                type="text"
              />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicEmail">
              <Form.Label>Pet Owner's Email</Form.Label>
              <Form.Control
                name="ownerEmail"
                value={payload.ownerEmail}
                onChange={onInputChange}
                type="text"
              />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicProcedure">
              <Form.Label>Procedure</Form.Label>
              <Form.Control
                name="procedure"
                value={payload.procedure}
                onChange={onInputChange}
                type="text"
              />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicCost">
              <Form.Label>Cost</Form.Label>
              <Form.Control
                name="cost"
                value={payload.cost}
                onChange={onInputChange}
                type="text"
              />
            </Form.Group>

            <Form.Group className="mb-3" controlId="formBasicCheckbox">
              <Form.Check
                onChange={setIsPaid}
                type="checkbox"
                label="Procedure Paid For"
              />
            </Form.Group>

            <Button variant="primary" type="submit">
              Submit
            </Button>
          </Form>
        </Col>

        <Col>
          <Table className="tableContainer" striped bordered hover>
            <thead>
              <tr>
                <th>Pet Name</th>
                <th>Owner Email</th>
                <th>Procedure</th>
                <th>Cost</th>
              </tr>
            </thead>
            <tbody>{recordTDs}</tbody>
          </Table>
        </Col>
      </Row>
    </Container>
  );
};

export default App;
