import React from "react";

interface Props {
    username: string;
    id: number;
}

const UserListItem = (props: Props) => {
    return (
        <li>
            <a href={"/users/" + props.username}>{props.username}</a>
        </li>
    );
};
