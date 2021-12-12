import React from "react";
import { Col, Container } from "react-bootstrap";

interface Props {
    children?: React.ReactNode;
}

const Content = (props: Props) => {
    return (
        <div>
            <Container style={{ display: "flex" }}>
                <Col></Col>
                <Col md="auto">{props.children}</Col>
                <Col></Col>
            </Container>
        </div>
    );
};

export default Content;
