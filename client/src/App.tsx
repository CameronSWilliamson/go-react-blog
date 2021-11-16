import React from "react";
import "./stylesheets/App.css";
import AlbumList from "./components/AlbumList";
import Header from "./components/Header";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Loader from "./components/Loader";
import HomePage from "./pages/HomePage";
import UserProfile from "./pages/UserProfile";
import PostPage from "./pages/PostPage";
import UserSearch from "./pages/UserSearch";

function App() {
    return (
        <Router>
            <div>
                <Header />
                <Routes>
                    <Route path="/" element={<HomePage />} />
                    <Route path="/loader" element={<Loader loading={true} />} />
                    <Route path="/user" element={<UserSearch />} />
                    <Route path="/users/:username" element={<UserProfile />} />
                    <Route path="/apitesting/" element={<AlbumList />} />
                    <Route path="/posts/:postID" element={<PostPage />} />
                </Routes>
            </div>
        </Router>
    );
}

export default App;
