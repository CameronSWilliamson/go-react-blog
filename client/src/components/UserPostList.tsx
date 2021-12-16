import React, { useEffect } from "react";
import PostCard from "./PostCard";

interface Props {
    username: string;
}

const UserPostList = (props: Props) => {
    useEffect(() => {
        // fetch data from API
        const url = "/api/users/" + props.username + "/posts";
        fetch(url, {
            method: "GET",
        }).then((response) => {
            console.log(response);
        });
    });

    return <p>Hello world</p>;
};

export default UserPostList;
