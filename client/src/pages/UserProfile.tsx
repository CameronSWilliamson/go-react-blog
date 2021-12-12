import React from "react";
import { Button } from "react-bootstrap";
import { useParams } from "react-router-dom";
import Content from "../components/Content";
import PostCard from "../components/PostCard";
import UserPostList from "../components/UserPostList";
import { logOut } from "../Utils/Auth";
import { UsernameProp } from "../Utils/Types/Props";

const UserProfile = (props: UsernameProp) => {
    const { username } = useParams();

    const loggingOut = () => {
        logOut();
        props.setLoggedIn(false);
        props.setUsername("");
    };

    return (
        <Content>
            <h1>{username}</h1>
            <PostCard title="first post" content="" postID={100000} />
            <UserPostList username={username!} />
            {props.loggedIn && props.username === username ? (
                <Button onClick={loggingOut}>logOut</Button>
            ) : (
                <p></p>
            )}
        </Content>
    );
};

export default UserProfile;
