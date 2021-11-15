import React from "react";
import { Album } from "../Utils/Types/Albums";
import axios from "axios";
import https from "https";

interface state {
    albums: Album[];
    isLoaded: boolean;
}

export default class AlbumList extends React.Component<{}, state> {
    constructor(props: {}) {
        super(props);
        this.state = { albums: [], isLoaded: false };
    }

    componentDidMount() {
        fetch("/albums")
            .then((res) => res.json())
            .then(
                (result) => {
                    console.log(result);
                    this.setState({
                        isLoaded: true,
                        albums: result,
                    });
                },
                // Note: it's important to handle errors here
                // instead of a catch() block so that we don't swallow
                // exceptions from actual bugs in components.
                (error) => {
                    this.setState({
                        isLoaded: true,
                    });
                }
            );
    }

    render() {
        return (
            <div>
                <h1>list</h1>
                <ul>
                    {this.state.albums.map((album, idx) => {
                        return (
                            <li>
                                {album.title} - {album.artist} - {album.price}
                            </li>
                        );
                    })}
                </ul>
            </div>
        );
    }
}
