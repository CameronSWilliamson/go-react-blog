import React from "react";
import { useParams } from "react-router-dom";

const UserProfile = () => {
    const { username } = useParams();

    return <div>hi there {username}</div>;
};

export default UserProfile;
