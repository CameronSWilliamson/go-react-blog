import React, { useEffect, useState } from "react";
import { Button, Form } from "react-bootstrap";
import { useNavigate } from "react-router-dom";
import Content from "../components/Content";
import { Category } from "../Utils/Types/Models";

interface Props {
    username: string;
    loggedIn: boolean;
}

const CreatePost = (props: Props) => {
    const [categories, setCategories] = useState<Array<Category>>([]);
    const [formData, setFormData] = useState({
        category_name: "",
        post: { title: "", content: "" },
    });
    const navigate = useNavigate();

    useEffect(() => {
        fetch("/api/categories/", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        }).then((response) => {
            response.text().then((res) => {
                console.log(res);
                let response: Array<Category> = JSON.parse(res);
                setCategories(response);
            });
        });
    }, []);

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setFormData({
            ...formData,
            [event.target.name]: event?.target.value,
        });
    };

    const submit = () => {
        fetch("/api/createpost/" + props.username, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(formData),
        }).then((response) => {
            if (response.status === 200) {
                navigate("/");
            }
        });
    };

    return (
        <Content>
            <Form style={{ width: "800px" }}>
                <Form.Group className="mb-3" controlId="title">
                    <Form.Label>Title</Form.Label>
                    <Form.Control
                        type="title"
                        name="title"
                        placeholder="Enter title"
                        onChange={handleChange}
                    />
                </Form.Group>
                <Form.Group className="mb-3" controlId="content">
                    <Form.Label>Content</Form.Label>
                    <Form.Control
                        as="textarea"
                        type="content"
                        name="content"
                        placeholder="content"
                        onChange={handleChange}
                        rows={12}
                    />
                </Form.Group>
                <Form.Group className="mb-3" controlId="category">
                    <Form.Select size="lg">
                        {categories.map((category) => {
                            return <option>{category.category_name}</option>;
                        })}
                    </Form.Select>
                </Form.Group>
                <Button onClick={submit} variant="primary">
                    Create Post
                </Button>
            </Form>
        </Content>
    );
};

export default CreatePost;
