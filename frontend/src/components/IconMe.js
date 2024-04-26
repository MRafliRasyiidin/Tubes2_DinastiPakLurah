import React from 'react';
import  {FontAwesomeIcon}  from '@fortawesome/react-fontawesome';


function IconMe({url, description, icon, textColor}){
  return (
        <a
                href={url}
                className = {`hover:scale-125`}
                target="_blank"
                rel="noopener noreferrer"
                style={{ textDecoration: 'none', color: 'inherit' }} // Optional: style for link
        >
        <FontAwesomeIcon className = {textColor} icon={icon} style={{ marginRight: '5px' }} />
        <span className={textColor}>{description}</span>
        </a>
  )
}

export default IconMe
