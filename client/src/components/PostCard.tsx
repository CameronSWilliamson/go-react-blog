import React from "react";
import { Button, Card } from "react-bootstrap";
import { useParams } from "react-router-dom";
import Content from "../components/Content";

interface Props {
    title: string;
    content?: string;
    postID: number;
}

const PostCard = (props: Props) => {
    return (
        <Card style={{ width: "18rem" }}>
            <Card.Img variant="top" src="https://via.placeholder.com/286x180" />
            <Card.Body>
                <Card.Title>{props.title}</Card.Title>
                <Card.Text>{props.content}</Card.Text>
                <Button variant="primary" href={`/posts/${props.postID}`}>
                    Go somewhere
                </Button>
            </Card.Body>
        </Card>
    );
};

export default PostCard;
