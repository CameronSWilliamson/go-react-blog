import React, { useEffect } from "react";
import Content from "../components/Content";
import { User } from "../Utils/Types/Models";

const UserSearch = () => {
    const [users, setUsers] = React.useState<Array<User>>([]);

    useEffect(() => {
        fetch("/api/users", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        }).then((response) => {
            response.text().then((res) => {
                let response: Array<User> = JSON.parse(res);
                console.log(response);
                setUsers(response);
            });
        });
    }, []);

    return (
        <Content>
            <ul>
                {users.map((user) => {
                    return (
                        <li>
                            <a href={"../users/" + user.username}>
                                {user.username}
                            </a>
                        </li>
                    );
                })}
            </ul>
        </Content>
    );
};

export default UserSearch;
