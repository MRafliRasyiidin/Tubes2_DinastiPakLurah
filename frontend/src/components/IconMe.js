import React from 'react';
import  {FontAwesomeIcon}  from '@fortawesome/react-fontawesome';


function IconMe({url, description, icon}){
  return (
        <a
                href={url}
                className = "hover:scale-125"
                target="_blank"
                rel="noopener noreferrer"
                style={{ textDecoration: 'none', color: 'inherit' }} // Optional: style for link
        >
        <FontAwesomeIcon icon={icon} style={{ marginRight: '5px' }} />{description}
        </a>
  )
}

export default IconMe
