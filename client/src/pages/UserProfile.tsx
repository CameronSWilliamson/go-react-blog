import React, { useEffect } from "react";
import { Button } from "react-bootstrap";
import { useParams } from "react-router-dom";
import Content from "../components/Content";
import PostCard from "../components/PostCard";
import UserPostList from "../components/UserPostList";
import { logOut } from "../Utils/Auth";
import { Post } from "../Utils/Types/Models";
import { UsernameProp } from "../Utils/Types/Props";

const UserProfile = (props: UsernameProp) => {
    const { username } = useParams();
    const [posts, setPosts] = React.useState<Array<Post>>([]);

    useEffect(() => {
        fetch("/api/users/" + username + "/posts/5", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        }).then((response) => {
            response.text().then((res) => {
                let response: Array<Post> = JSON.parse(res);
                console.log(response);
                setPosts(response);
            });
        });
    }, []);

    const loggingOut = () => {
        logOut();
        props.setLoggedIn(false);
        props.setUsername("");
    };

    return (
        <Content>
            <h1>{username}</h1>
            {posts.map((post) => {
                return (
                    <PostCard
                        author={post.username}
                        content={post.content}
                        title={post.title}
                        postID={post.postid}
                    />
                );
            })}
            {/* <PostCard
                title="first post"
                content="This is my first post this is so super cool omg"
                postID={100000}
                author={"kikard"}
            /> */}
            {/* <UserPostList username={username!} /> */}
            {props.loggedIn && props.username === username ? (
                <Button onClick={loggingOut}>logOut</Button>
            ) : (
                <p></p>
            )}
        </Content>
    );
};

export default UserProfile;
