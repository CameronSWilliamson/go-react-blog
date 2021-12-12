import React, { useEffect } from "react";
import "./stylesheets/App.css";
import AlbumList from "./components/AlbumList";
import Header from "./components/Header";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Loader from "./components/Loader";
import HomePage from "./pages/HomePage";
import UserProfile from "./pages/UserProfile";
import PostPage from "./pages/PostPage";
import UserSearch from "./pages/UserSearch";
import LoginPage from "./pages/LoginPage";
import { checkLogin } from "./Utils/Auth";

function App() {
    const [username, setUsername] = React.useState("");
    const [loggedIn, setLoggedIn] = React.useState(false);

    useEffect(() => {
        const loginStatus = checkLogin();
        setUsername(loginStatus.username);
        setLoggedIn(loginStatus.loggedIn);
    }, []);

    return (
        <Router>
            <div>
                <Header
                    loggedIn={loggedIn}
                    username={username}
                    setLoggedIn={setLoggedIn}
                    setUsername={setUsername}
                />
                <Routes>
                    <Route path="/" element={<HomePage />} />
                    <Route path="/loader" element={<Loader loading={true} />} />
                    <Route path="/user" element={<UserSearch />} />
                    <Route
                        path="/users/:username"
                        element={
                            <UserProfile
                                loggedIn={loggedIn}
                                username={username}
                                setLoggedIn={setLoggedIn}
                                setUsername={setUsername}
                            />
                        }
                    />
                    <Route path="/apitesting/" element={<AlbumList />} />
                    <Route path="/posts/:postID" element={<PostPage />} />
                    <Route path="/newPost" element={<PostPage />} />
                    <Route
                        path="/login"
                        element={
                            <LoginPage
                                loggedIn={loggedIn}
                                username={username}
                                setLoggedIn={setLoggedIn}
                                setUsername={setUsername}
                            />
                        }
                    />
                    <Route path="*" element={<div>Not Found</div>} />
                    {/* <Route
                        path="/login"
                        element={
                            <LoginPage
                                username={username}
                                setUsername={setUsername}
                                loggedIn={loggedIn}
                                setLoggedIn={setLoggedIn}
                            />
                        }
                    /> */}
                    "
                </Routes>
            </div>
        </Router>
    );
}

export default App;
