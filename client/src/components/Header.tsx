import React from "react";
import { Navbar, Container, Nav } from "react-bootstrap";
import { UsernameProp } from "../Utils/Types/Props";

const Header = (props: UsernameProp) => {
    return (
        <Navbar bg="dark" expand="lg">
            <Container>
                <Navbar.Brand href="/" className="text-white">
                    React-Bootstrap
                </Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse
                    id="basic-navbar-nav"
                    className="justify-content-end"
                >
                    <Nav className="me-auto">
                        <Nav.Link href="/" className="text-white">
                            Home
                        </Nav.Link>
                        <Nav.Link href="/user" className="text-white">
                            User
                        </Nav.Link>
                        <Nav.Link href="/newPost" className="text-white">
                            Create Post
                        </Nav.Link>
                    </Nav>
                    {!props.loggedIn ? (
                        <Navbar.Text className="text-white">
                            <a href="/login" className="text-white">
                                Click to Login
                            </a>
                        </Navbar.Text>
                    ) : (
                        <Navbar.Text className="text-white">
                            Signed in as:{" "}
                            <a
                                href={"/users/" + props.username}
                                className="text-white"
                            >
                                {props.username}
                            </a>
                        </Navbar.Text>
                    )}
                </Navbar.Collapse>
            </Container>
        </Navbar>
    );
};

export default Header;
