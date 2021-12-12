import React from "react";
import { Form, Button } from "react-bootstrap";
import { useNavigate } from "react-router";
import Content from "../components/Content";
import { logIn } from "../Utils/Auth";
import { UsernameProp } from "../Utils/Types/Props";

const LoginPage = (props: UsernameProp) => {
    const [newAccount, setNewAccount] = React.useState(false);
    const [formData, setFormData] = React.useState({ username: "", email: "" });
    const navigate = useNavigate();

    const resetForm = () => {
        setFormData({ username: "", email: "" });
    };

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({
            ...formData,
            [event.target.name]: event.target.value,
        });
        console.log(formData);
    };

    const saveInfo = () => {
        localStorage.setItem("username", formData.username);
        console.log("loggedIn: " + props.loggedIn);
        localStorage.setItem("loggedIn", String(props.loggedIn));
        console.log("saving info username: " + props.username);
    };

    const handleSubmit = () => {
        const userName = formData.username;
        const userEmail = formData.email;
        if (newAccount) {
            fetch("/api/login", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify({
                    username: userName,
                    email: userEmail,
                }),
            }).then((response) => {
                console.log(response);
                if (response.status === 200) {
                    props.setUsername(userName);
                    props.setLoggedIn(true);
                    logIn(userName);
                    navigate("/", { replace: true });
                } else {
                    alert("Invalid username");
                }
            });
        } else {
            fetch("/api/login", {
                method: "GET",
                headers: {
                    "Content-Type": "application/json",
                    username: userName,
                },
            }).then((response) => {
                console.log(response.status);
                if (response.status === 200) {
                    props.setUsername(userName);
                    logIn(userName);
                    props.setLoggedIn(true);
                    saveInfo();
                    console.log("logged in");
                    navigate("/", { replace: true });
                } else {
                    props.setUsername("");
                    props.setLoggedIn(false);
                    alert("Invalid username");
                }
            });
        }
    };

    return (
        <Content>
            {newAccount ? (
                <div style={{ width: "500px" }}>
                    <h1>New Account</h1>
                    <Form>
                        <Form.Group className="mb-3" controlId="formBasicEmail">
                            <Form.Label>Email</Form.Label>
                            <Form.Control
                                type="email"
                                name="email"
                                placeholder="Enter email"
                                onChange={handleChange}
                            />
                            <Form.Label>Username</Form.Label>
                            <Form.Control
                                type="username"
                                name="username"
                                placeholder="Enter username"
                                onChange={handleChange}
                                value={formData.username}
                            />
                        </Form.Group>
                        <Button variant="primary" onClick={handleSubmit}>
                            Submit
                        </Button>
                        <Button
                            variant="primary"
                            style={{ marginLeft: "10px" }}
                            onClick={() => {
                                setNewAccount(false);
                                resetForm();
                            }}
                        >
                            Existing Account
                        </Button>
                    </Form>
                </div>
            ) : (
                <div style={{ width: "500px" }}>
                    <Form>
                        <Form.Group className="mb-3" controlId="formBasicEmail">
                            <Form.Label>Username</Form.Label>
                            <Form.Control
                                type="username"
                                name="username"
                                placeholder="Enter Username"
                                onChange={handleChange}
                            />
                        </Form.Group>
                        <Button variant="primary" onClick={handleSubmit}>
                            Submit
                        </Button>
                        <Button
                            variant="primary"
                            style={{ marginLeft: "10px" }}
                            onClick={() => {
                                setNewAccount(true);
                                resetForm();
                            }}
                        >
                            Create Account
                        </Button>
                    </Form>
                </div>
            )}
        </Content>
    );
};

export default LoginPage;
