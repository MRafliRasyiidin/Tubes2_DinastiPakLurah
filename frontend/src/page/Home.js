import '../App.css';
import { useState } from 'react';
import Search from '../Search';
import Navbar from '../components/Navbar';
import lightLogo from "../Logo/logo-light.png";
import darkLogo from "../Logo/logo-dark.png";
import ToggleSwitch from "../components/ToggleSwitch";
import ParticleApp from "../components/Background";

function App() {
  const [searchAlgorithm, setSearchAlgorithm] = useState('BFS');
  const [searchAll, setSearchAll] = useState('No');
  const [darkmode, setDarkMode] = useState(false); 
  const logo = darkmode ? darkLogo : lightLogo;

  const toggleAlgorithm = () => {
    setSearchAlgorithm(searchAlgorithm === 'BFS' ? 'IDS' : 'BFS');
    // console.log(searchAlgorithm);
  };
  
  const toggleAll = () =>{
    setSearchAll(searchAll === 'No' ? 'Yes' : 'No');
    console.log(searchAll);
  }
  return (
    
    <div className={`flex flex-col items-center justify-center h-max w-auto`}>
      <ParticleApp/>
        <Navbar darkmode={darkmode} />
        <img className="w-auto h-80 top-20 z-10" src={logo} alt="Description of the image" />
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
            checked={searchAll === 'No'} 
            onChange={toggleAll}
            leftInfo={'Yes'}
            rightInfo={'No'}
            color={"#2E51A2"}
            scolor={"#D1D5DB"}
            info={'Search all'}
          />
          <ToggleSwitch
            checked={searchAlgorithm === 'BFS'}
            onChange={toggleAlgorithm}
            leftInfo={'IDS'}
            rightInfo={'BFS'}
            info={'Algorithm'}
            color={"#2E51A2"}
          />
        </div>

        <Search
          searchAlgorithm={searchAlgorithm} searchAll={searchAll}
        />
        <button onClick={() => setDarkMode(!darkmode)} className={`rounded-lg fixed top-4 right-4 ${darkmode ? 'bg-white hover:bg-gray-600' : 'bg-gray-300 hover:bg-gray-700 hover:text-white'}`}>Dark Mode</button>
      </div>
  );
}

export default App;
