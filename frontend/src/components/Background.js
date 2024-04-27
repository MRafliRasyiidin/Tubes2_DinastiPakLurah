import { useEffect, useState } from "react";
import Particles, { initParticlesEngine } from "@tsparticles/react";
import type { Container } from "@tsparticles/engine";
import { loadSlim } from "@tsparticles/slim";

const ParticleApp = () => {
    const [init, setInit] = useState(false);
  
    useEffect(() => {
      initParticlesEngine(async (engine) => {
        await loadSlim(engine);
      }).then(() => {
        setInit(true);
      });
    }, []);
  
    const particlesLoaded = (container) => {
      console.log(container);
    };
  
    const particleOptions = {
      fpsLimit: 60,
      particles: {
        number: {
          value: 90,
          density: {
            enable: true,
            value_area: 1000,
          },
        },
        color: {
          value: "#82807a",
        },
        shape: {
          type: "triangle",
        },
        opacity: {
          value: 0.5,
          random: true,
        },
        size: {
          value: 5,
          random: true,
        },
        move: {
          enable: true,
          speed: 2,
          direction: "none",
          random: false,
          straight: false,
          out_mode: "out",
          bounce: false,
        },
        links: {
          enable: true,
          opacity: 0.5,
          color: "#82807a",
          distance: 200, 
          width: 2,
          blink: false,
        },
      },
    };
  
    if (init) {
      return (
        <Particles
          id="tsparticles"
          options={particleOptions}
          particlesLoaded={particlesLoaded}
        />
      );
    }
  
    return <></>;
  };
  
  export default ParticleApp;
  