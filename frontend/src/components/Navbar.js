import React from 'react';
import { Link } from 'react-router-dom';
import IconMe from '../components/IconMe'
import {faGithub} from '@fortawesome/free-brands-svg-icons';
function Navbar({ darkmode }) {
  return (
    <nav className={`flex justify-between items-center py-4 w-screen ${darkmode ? 'bg-gray-300' : 'bg-gray-600'}`}>
      <ul className="flex mr-10">
        <li className="mx-4">
          <Link to="/" className="text-white hover:scale-110 transition-transform duration-300">Home</Link>
        </li>
        <li className="mx-4">
          <Link to="/about" className="text-white hover:text-gray-300">About</Link>
        </li>
        <li className="mx-4">
          <IconMe 
            url = {"https://github.com/MRafliRasyiidin/Tubes2_DinastiPakLurah"} 
            description={"Github"} 
            icon={faGithub}
            textColor={"text-white"}/>
        </li>
      </ul>
    </nav>
  );
}

export default Navbar;
