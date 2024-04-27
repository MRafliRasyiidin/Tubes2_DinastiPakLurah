import React from 'react';

function RoundedImage({ imageUrl, altText }) {
  return (
    <div 
        className="flex items-center justify-center rounded-full overflow-hidden w-32 h-32 z-10"
        style={{ width: '150px', height: '150px' }}>
      <img
        src={imageUrl}
        alt={altText}
        className="object-cover w-full h-full"
      />
    </div>
  );
}

export default RoundedImage;
