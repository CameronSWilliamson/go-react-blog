import React, { useEffect, useState } from "react";

interface Data {
    ID: number;
    Status: string;
}

const AlbumList = () => {
    const [loading, setLoading] = useState(false);
    const [status, setStatus] = useState<string>();
    const [ID, setID] = useState<number>();

    const fetchApi = () => {
        fetch("api/test", {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        }).then((response) => {
            response.text().then((res) => {
                console.log("Getting response");
                let response: Data = JSON.parse(res);
                setID(response.ID);
                setStatus(response.Status);
                console.log(ID);
            });
            // response.json();
        });
    };

    useEffect(() => fetchApi(), []);

    return (
        <div>
            Your data is here: {status} {ID}
        </div>
    );
};

export default AlbumList;
