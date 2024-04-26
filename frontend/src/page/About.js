import React from 'react';
import Navbar from '../components/Navbar';
import RoundedImage from '../components/RoundedImage';
import profileImage1 from '../Logo/logo-light.png'; 


function About() {
  const imagePath = [
    profileImage1,
    'https://example.com/profile2.jpg',
    'https://example.com/profile3.jpg'
  ];

  return (
    <div className="mx-auto h-screen">
      <Navbar/>
      <div className = "flex flex-col items-center">
        <h1 className="text-3xl font-bold mt-10">Dinasti Pak Lurah Crew </h1>
        <p className="mt-4">This is the About page content.</p>
      </div>


      {/* Container for the three boxes */}
      <div className="mt-8 grid grid-cols-1 gap-4 sm:grid-cols-3">
        {/* First box */}
        <div className="flex flex-col items-center bg-gray-200 p-4 rounded-md">
          <h2 className="text-lg font-semibold">Muhamad Rafli Rasyidin</h2>
          <RoundedImage imageUrl={imagePath[0]} altText="Profile Image 1" />
          <p className="mt-2">Content for box 1 goes here.</p>
        </div>

        {/* Second box */}
        <div className="flex flex-col items-center bg-gray-200 p-4 rounded-md">
          <h2 className="text-lg font-semibold">M. Hanief Fatkhan Nashrullah</h2>
          <RoundedImage imageUrl={imagePath[0]} altText="Profile Image 2" />
          <p className="mt-2">Content for box 2 goes here.</p>
        </div>

        {/* Third box */}
        <div className="flex flex-col items-center bg-gray-200 p-4 rounded-md">
          <h2 className="text-lg font-semibold">Indraswara Galih Jayanegara</h2>
          <RoundedImage imageUrl={imagePath[0]} altText="Profile Image 3" />
          <p className="mt-2">Content for box 3 goes here.</p>
        </div>
      </div>
    </div>
  );
}

export default About;


