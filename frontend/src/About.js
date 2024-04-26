import React from 'react';
import Navbar from './components/Navbar';

function About() {
  return (
    <div className="mx-auto h-screen">
      <Navbar/>
      <h1 className="text-3xl font-bold mt-10">About Us</h1>
      <p className="mt-4">This is the About page content.</p>
    </div>
  );
}

export default About;
