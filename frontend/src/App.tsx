import React from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Login from './components/Login';
import SpotifyCallback from './components/SpotifyCallback';
import Home from './components/Home';


function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/spotify/callback" element={<SpotifyCallback />} />
      </Routes>
    </Router>
  );
}

export default App;
