import React from "react";
import { Col, Container } from "react-bootstrap";

interface Props {
    children?: React.ReactNode;
}

const Content = (props: Props) => {
    return (
        <Container>
            <Col></Col>
            <Col>{props.children}</Col>
            <Col></Col>
        </Container>
    );
};

export default Content;
