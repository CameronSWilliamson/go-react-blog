import React from "react";
import { useParams } from "react-router-dom";
import Content from "../components/Content";
import PostCard from "../components/PostCard";

const UserProfile = () => {
    const { username } = useParams();

    return (
        <Content>
            <h1>{username}</h1>
            <PostCard title="first post" content="" postID={100000} />
        </Content>
    );
};

export default UserProfile;
