import '../App.css';
import { useState } from 'react';
import Search from '../Search';
import Navbar from '../components/Navbar';
import lightLogo from "../Logo/logo-light.png";
import darkLogo from "../Logo/logo-dark.png";
import ToggleSwitch from "../components/ToggleSwitch";

function App() {
  const [searchAlgorithm, setSearchAlgorithm] = useState('BFS');
  const [searchAll, setSearchAll] = useState('All');
  const [darkmode, setDarkMode] = useState(false); 
  const logo = darkmode ? darkLogo : lightLogo;

  const toggleAlgorithm = () => {
    setSearchAlgorithm(searchAlgorithm === 'BFS' ? 'IDS' : 'BFS');
    console.log(searchAlgorithm);
  };

  const toggleAll = () =>{
    setSearchAll(searchAll === 'All' ? 'One' : 'All');
    console.log(searchAll);
  }
  return (

    <div className={`flex flex-col items-center justify-center h-max w-auto`}>
        <Navbar darkmode={darkmode} />
        <img className="w-auto h-80 top-20" src={logo} alt="Description of the image" />
        <div className="text-center mb-10">
          <h1 className={`font-sans font-bold text-xl ${darkmode ? 'text-white' : 'text-black'}`}>
            Made By DinastiPakLurah 
          </h1>
          <h2 className={`${darkmode ? 'text-white' : 'text-black'}`} >
            made with Happy(Tears) and Joy(Pain)
          </h2>
        </div>
        <div className = "flex flex-col justify-center items-center">
          <ToggleSwitch
            checked={searchAll === 'All'} 
            onChange={toggleAll}
            leftInfo={'Yes'}
            rightInfo={'No'}
            info={'Search all'}
          />
          <ToggleSwitch
            checked={searchAlgorithm === 'IDS'}
            onChange={toggleAlgorithm}
            leftInfo={'IDS'}
            rightInfo={'BFS'}
            info={'Algorithm'}
          />
        </div>

        <Search 
          searchAlgorithm={searchAlgorithm}
        />
        <button onClick={() => setDarkMode(!darkmode)} className={`rounded-lg fixed top-4 right-4 ${darkmode ? 'bg-white hover:bg-gray-600' : 'bg-gray-300 hover:bg-gray-700 hover:text-white'}`}>Dark Mode</button>
      </div>
  );
}

export default App;
