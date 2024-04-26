import './App.css';
import { useState } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Search from './Search.js';
import Navbar from './components/Navbar';
import About from './About';
import lightLogo from "./Logo/logo-light.png";
import darkLogo from "./Logo/logo-dark.png";

function App() {
  const [darkmode, setDarkMode] = useState(false); 
  const logo = darkmode ? darkLogo : lightLogo;

  return (
    <Router>
      <div className={`flex flex-col items-center justify-center h-max w-auto ${darkmode ? 'bg-black' : 'bg-white'}`}>
        <Navbar darkmode={darkmode} />
        <img className="w-auto h-96 top-20" src={logo} alt="Description of the image" />
        <div className="text-center mb-10">
          <h1 className={`font-sans font-bold text-xl ${darkmode ? 'text-white' : 'text-black'}`}>
            Made By DinastiPakLurah 
          </h1>
          <h2 className={`${darkmode ? 'text-white' : 'text-black'}`} >
            made with Happy(Tears) and Joy(Pain)
          </h2>
        </div>
        <Routes>
          <Route path="/about" target element={<About/>} />
        </Routes>
        <Search />
        <button onClick={() => setDarkMode(!darkmode)} className={`rounded-lg fixed top-4 right-4 ${darkmode ? 'bg-white hover:bg-gray-600' : 'bg-gray-300 hover:bg-gray-700 hover:text-white'}`}>Dark Mode</button>
      </div>
    </Router>
  );
}

export default App;
