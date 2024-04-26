import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import About from './About';
import Home from './Home'

function App() {

  return (
    <div className={`flex flex-col items-center justify-center h-max w-auto`}>
        <Router>
        <Routes>
          <Route path="/" target element={<Home/>} />
          <Route path="/about" target element={<About/>} />
        </Routes>
      </Router>
      </div>
  );
}

export default App;
