import React, { useEffect } from "react";
import { useParams } from "react-router-dom";
import Content from "../components/Content";
import { Post } from "../Utils/Types/Models";
import CreatePost from "./CreatePost";

const PostPage = () => {
    const postID: number = parseInt(useParams()!.postID!);
    const [post, setPost] = React.useState<Post>({
        postid: postID,
        title: "",
        content: "",
        username: "",
    });

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
    }, []);

    return (
        <Content>
            <h1 style={{ width: "800px" }}>{post.title}</h1>
            <h4>
                <a href={"../users/" + post.username}>{post.username}</a>
            </h4>
            <p>{post.content}</p>
        </Content>
    );
};

export default PostPage;
