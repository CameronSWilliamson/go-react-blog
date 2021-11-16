import React from "react";
import { useParams } from "react-router-dom";
import Content from "../components/Content";

const PostPage = () => {
    const { postID } = useParams();

    return <Content>{postID}</Content>;
};

export default PostPage;
