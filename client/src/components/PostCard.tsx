import React from "react";
import { Card } from "react-bootstrap";

interface Props {
    title: string;
    author: string;
    content?: string;
    postID: number;
}

const PostCard = (props: Props) => {
    return (
        <Card style={{ width: "18rem" }}>
            <Card.Body>
                <Card.Title>{props.title}</Card.Title>
                <Card.Subtitle className="mb-2 text-muted">
                    {props.author}
                </Card.Subtitle>
                <Card.Text>{props.content}</Card.Text>
                <Card.Link href={"../posts/" + props.postID}>
                    Full Post
                </Card.Link>
            </Card.Body>
        </Card>
    );
};

export default PostCard;
