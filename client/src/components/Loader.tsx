import * as React from "react";
import { Container, Spinner } from "react-bootstrap";

interface Props {
    loading: boolean;
    children?: React.ReactNode;
}

const Loader = (props: Props) => {
    return (
        <div>
            {props.loading ? (
                <Container
                    style={{
                        display: "flex",
                        justifyContent: "center",
                        alignItems: "center",
                    }}
                >
                    <Spinner animation="border" />
                </Container>
            ) : (
                props.children
            )}
        </div>
    );
};

export default Loader;
