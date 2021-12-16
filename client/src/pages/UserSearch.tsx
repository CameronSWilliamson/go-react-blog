import React, { useEffect } from "react";
import Content from "../components/Content";

const UserSearch = () => {
    useEffect(() => {
        fetch("/api/users", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        }).then((response) => {
            response.json().then((data) => {
                console.log(data);
            });
        });
    });

    return (
        <Content>
            <ul>
                <li>
                    <a href="/users/therealkey">Therealkey</a>
                </li>
            </ul>
        </Content>
    );
};

export default UserSearch;
