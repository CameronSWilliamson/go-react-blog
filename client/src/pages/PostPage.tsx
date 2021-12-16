import React, { useEffect } from "react";
import { Card } from "react-bootstrap";
import { useParams } from "react-router-dom";
import Content from "../components/Content";
import { Post, Comment } from "../Utils/Types/Models";

const PostPage = () => {
    const postID: number = parseInt(useParams()!.postID!);
    const [post, setPost] = React.useState<Post>({
        postid: postID,
        title: "",
        content: "",
        username: "",
    });
    const [comments, setComments] = React.useState<Array<Comment>>([]);

    useEffect(() => {
        fetch("/api/posts/" + postID, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        }).then((response) => {
            response.text().then((res) => {
                console.log(res);
                setPost(JSON.parse(res));
            });
        });

        fetch("/api/posts/" + postID + "/comments", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        }).then((response) => {
            response.text().then((res) => {
                console.log(res);
                setComments(JSON.parse(res));
            });
        });
    }, [postID]);

    return (
        <Content>
            <h1 style={{ width: "800px" }}>{post.title}</h1>
            <h4>
                <a href={"../users/" + post.username}>{post.username}</a>
            </h4>
            <p>{post.content}</p>
            <h4>Comments</h4>
            {comments.map((comment) => {
                return (
                    <Card>
                        <Card.Header>{comment.username}</Card.Header>
                        <Card.Body>
                            <Card.Text>{comment.content}</Card.Text>
                        </Card.Body>
                    </Card>
                );
            })}
        </Content>
    );
};

export default PostPage;
