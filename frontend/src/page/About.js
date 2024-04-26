import React from 'react';
import Navbar from '../components/Navbar';
import RoundedImage from '../components/RoundedImage';
import logo from '../Logo/logo-light.png';
import profileImage1 from '../Logo/rafli.jpg'; 
import profileImage2 from '../Logo/indra.jpg'; 
import profileImage3 from '../Logo/hanief.png'; 


function About() {
  const imagePath = [
    profileImage1,
    profileImage2,
    profileImage3
  ];

  return (
    <div className="mx-auto h-screen">
      <Navbar/>
      <div className = "flex items-center justify-center">
        <img className="w-auto h-80" src={logo} alt="Description of the image" />
      </div>
        <div className = "flex flex-col items-center ">
          <h1 className="text-3xl font-bold mt-10">Dinasti Pak Lurah Crew</h1>
          <p className="mt-4 text-lg flex justify-center text-center">
          Ini adalah kami, sekumpulan orang yang bahagia dan beruntung dibalik pengerjaan tugas besar ini, kami mengerjakan
          <br/>
          tanpa adanya paksaan dari pihak manapun justru bahagia (bukan sarkas).
          </p>
        </div>

      <div >
        {/* Container for the three boxes */}
        <div className="mt-8 grid grid-cols-1 gap-4 sm:grid-cols-3">
          {/* First box */}
          <div className="flex flex-col items-center bg-gray-200 p-4 rounded-md w-4/5 mx-auto">
            <h2 className="text-lg font-semibold">Muhamad Rafli Rasyidin</h2>
            <RoundedImage imageUrl={imagePath[0]} altText="Profile Image 1" />
            <p className="mt-2 text-center"> Rafli, masuk IF, tapi masih bisa turu, hebat bukan?
            <br/>
            
            </p>
            <div className="flex flex-row mt-4 space-x-2">
              <a href="https://github.com/MRafliRasyiidin" target="_blank" rel="noopener noreferrer" className="text-gray-500 hover:underline">GitHub</a>
              <a href="/" className="text-gray-500 hover:underline">LinkedIn</a>
              <a href="/" className="text-gray-500 hover:underline">Instagram</a>
              {/* Add more links as needed */}
            </div>
          </div>

          {/* Second box */}
          <div className="flex flex-col items-center bg-gray-200 p-4 rounded-md w-4/5 mx-auto">
            <h2 className="text-lg font-semibold">M. Hanief Fatkhan Nashrullah</h2>
            <RoundedImage imageUrl={imagePath[2]} altText="Profile Image 2" />
            <p className="mt-2 text-center">Biasa dipanggil Hanief, orangnya oke
            <br/>ga juga sih, orang ini sering banget ngasih rickroll
            </p>
            <div className="flex flex-row mt-4 space-x-2">
              <a href="https://github.com/hannoobz" target="_blank" rel="noopener noreferrer" className="text-gray-500 hover:underline">GitHub</a>
              <a href="/" className="text-gray-500 hover:underline">LinkedIn</a>
              <a href="/" className="text-gray-500 hover:underline">Instagram</a>
              <a href="/" className="text-gray-500 hover:underline">Special</a>
              {/* Add more links as needed */}
            </div>
          </div>

          {/* Third box */}
          <div className="flex flex-col items-center bg-gray-200 p-4 rounded-md w-4/5 mx-auto">
            <h2 className="text-lg font-semibold">Indraswara Galih Jayanegara</h2>
            <RoundedImage imageUrl={imagePath[1]} altText="Profile Image 3" />
            <p className="mt-2 text-center">Foto ini diambil pada saat website ini dibuat, bahagia bukan?
            <br/>Aku biasanya dipanggil Indra
            <br/>anak yang masuk IF, tapi setelah masuk merasa...
            </p>
            <div className="flex flex-row mt-4 space-x-2">
              <a href="https://github.com/Indraswara" target="_blank" rel="noopener noreferrer" className="text-gray-500 hover:underline">GitHub</a>
              <a href="/" className="text-gray-500 hover:underline">LinkedIn</a>
              <a href="/" className="text-gray-500 hover:underline">Instagram</a>
              {/* Add more links as needed */}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default About;


